package app

import (
	"context"
	"fmt"
	"net"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/omni-network/omni/e2e/app/agent"
	"github.com/omni-network/omni/e2e/app/key"
	"github.com/omni-network/omni/e2e/docker"
	"github.com/omni-network/omni/e2e/netman"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/e2e/vmcompose"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/fireblocks"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/xchain"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
	"github.com/cometbft/cometbft/test/e2e/pkg/exec"

	"github.com/ethereum/go-ethereum/p2p/enode"

	"github.com/BurntSushi/toml"
)

// DefinitionConfig is the configuration required to create a full Definition.
type DefinitionConfig struct {
	AgentSecrets agent.Secrets

	ManifestFile  string
	InfraProvider string

	// Secrets (not required for devnet)
	DeployKeyFile string
	FireAPIKey    string
	FireKeyPath   string
	RPCOverrides  map[string]string // map[chainName]rpcURL1,rpcURL2,...

	InfraDataFile string // Not required for docker provider
	OmniImgTag    string // OmniImgTag is the docker image tag used for halo and relayer.

	ExplorerDBConn string // ExplorerDBConn is the connection string for the explorer database.
	GraphQLURL     string // GraphQLURL is the host for the GraphQL server.

	TracingEndpoint string
	TracingHeaders  string
}

// DefaultDefinitionConfig returns a default configuration for a Definition.
func DefaultDefinitionConfig(ctx context.Context) DefinitionConfig {
	defaultTag := "main"
	if out, err := exec.CommandOutput(ctx, "git", "rev-parse", "--short=7", "HEAD"); err == nil {
		defaultTag = strings.TrimSpace(string(out))
	}

	return DefinitionConfig{
		AgentSecrets:   agent.Secrets{}, // empty agent.Secrets by default
		InfraProvider:  docker.ProviderName,
		OmniImgTag:     defaultTag,
		ExplorerDBConn: "postgres://omni:password@explorer_db:5432/omni_db",
		GraphQLURL:     "http://explorer_graphql:8080/query",
	}
}

// Definition defines a e2e network. All (sub)commands of the e2e cli requires a definition operate.
// Armed with a definition, a e2e network can be deployed, started, tested, stopped, etc.
type Definition struct {
	Manifest    types.Manifest
	Testnet     types.Testnet // Note that testnet is the cometBFT term.
	Infra       types.InfraProvider
	Cfg         DefinitionConfig // Original config used to construct the Definition.
	lazyNetwork *lazyNetwork     // lazyNetwork does lazy setup of backends and netman (only if required).
}

// InitLazyNetwork initializes the lazy network, which is the backends and netman.
func (d Definition) InitLazyNetwork() error {
	return d.lazyNetwork.Init()
}

// Backends returns the backends.
func (d Definition) Backends() ethbackend.Backends {
	return d.lazyNetwork.MustBackends()
}

// Netman returns the netman.
func (d Definition) Netman() netman.Manager {
	return d.lazyNetwork.MustNetman()
}

// DeployInfos returns the deploy information of the OmniPortal and OmniAVS contracts.
func (d Definition) DeployInfos() types.DeployInfos {
	resp := make(types.DeployInfos)

	for chainID, info := range d.Netman().DeployInfo() {
		resp.Set(chainID, types.ContractPortal, info.PortalAddress, info.DeployHeight)
	}

	return resp
}

func MakeDefinition(ctx context.Context, cfg DefinitionConfig, commandName string) (Definition, error) {
	if strings.TrimSpace(cfg.ManifestFile) == "" {
		return Definition{}, errors.New("manifest not specified, use --manifest-file or -f")
	}

	manifest, err := LoadManifest(cfg.ManifestFile)
	if err != nil {
		return Definition{}, errors.Wrap(err, "loading manifest")
	}

	var infd types.InfrastructureData
	switch cfg.InfraProvider {
	case docker.ProviderName:
		infd, err = docker.NewInfraData(manifest)
	case vmcompose.ProviderName:
		infd, err = vmcompose.LoadData(cfg.InfraDataFile)
	default:
		return Definition{}, errors.New("unknown infra provider", "provider", cfg.InfraProvider)
	}
	if err != nil {
		return Definition{}, errors.Wrap(err, "loading infrastructure data")
	}

	testnet, err := TestnetFromManifest(ctx, manifest, infd, cfg)
	if err != nil {
		return Definition{}, errors.Wrap(err, "loading testnet")
	}

	// Setup lazy network, this is only executed by command that require networking.
	lazy := func() (ethbackend.Backends, netman.Manager, error) {
		backends, err := newBackends(ctx, cfg, testnet, commandName)
		if err != nil {
			return ethbackend.Backends{}, nil, errors.Wrap(err, "new backends")
		}

		netman, err := netman.NewManager(testnet, backends)
		if err != nil {
			return ethbackend.Backends{}, nil, errors.Wrap(err, "get network")
		}

		return backends, netman, nil
	}

	var infp types.InfraProvider
	switch cfg.InfraProvider {
	case docker.ProviderName:
		infp = docker.NewProvider(testnet, infd, cfg.OmniImgTag, cfg.GraphQLURL)
	case vmcompose.ProviderName:
		infp = vmcompose.NewProvider(testnet, infd, cfg.OmniImgTag, cfg.GraphQLURL)
	default:
		return Definition{}, errors.New("unknown infra provider", "provider", cfg.InfraProvider)
	}

	return Definition{
		Manifest:    manifest,
		Testnet:     testnet,
		Infra:       infp,
		lazyNetwork: &lazyNetwork{initFunc: lazy},
		Cfg:         cfg,
	}, nil
}

func newBackends(ctx context.Context, cfg DefinitionConfig, testnet types.Testnet, commandName string) (ethbackend.Backends, error) {
	// Skip backends if only deploying monitor, since there are no EVM to connect to.
	if testnet.OnlyMonitor {
		return ethbackend.Backends{}, nil
	}

	// If no fireblocks API key, use in-memory keys.
	if cfg.FireAPIKey == "" {
		return ethbackend.NewBackends(testnet, cfg.DeployKeyFile)
	}

	key, err := fireblocks.LoadKey(cfg.FireKeyPath)
	if err != nil {
		return ethbackend.Backends{}, errors.Wrap(err, "load fireblocks key")
	}

	opts := []fireblocks.Option{
		fireblocks.WithSignNote(fmt.Sprintf("omni e2e %s %s", commandName, testnet.Network)),
		fireblocks.WithQueryInterval(5 * time.Second), // If we retry too often we get rate limited.
	}
	fireCl, err := fireblocks.New(testnet.Network, cfg.FireAPIKey, key, opts...)
	if err != nil {
		return ethbackend.Backends{}, errors.Wrap(err, "new fireblocks")
	}

	// TODO(corver): Fireblocks keys need to be funded on private/internal chains we deploy.

	return ethbackend.NewFireBackends(ctx, testnet, fireCl)
}

// adaptCometTestnet adapts the default comet testnet for omni specific changes and custom config.
func adaptCometTestnet(ctx context.Context, manifest types.Manifest, testnet *e2e.Testnet, imgTag string) (*e2e.Testnet, error) {
	testnet.Dir = runsDir(testnet.File)
	testnet.VoteExtensionsEnableHeight = 1
	testnet.UpgradeVersion = "omniops/halo:" + imgTag

	for i := range testnet.Nodes {
		var err error
		testnet.Nodes[i], err = adaptNode(ctx, manifest, testnet, testnet.Nodes[i], imgTag)
		if err != nil {
			return nil, err
		}
	}

	return testnet, nil
}

// adaptNode adapts the default comet node for omni specific changes and custom config.
func adaptNode(ctx context.Context, manifest types.Manifest, testnet *e2e.Testnet, node *e2e.Node, tag string) (*e2e.Node, error) {
	valKey, err := getOrGenKey(ctx, manifest, node.Name, key.Validator)
	if err != nil {
		return nil, err
	}
	nodeKey, err := getOrGenKey(ctx, manifest, node.Name, key.P2PConsensus)
	if err != nil {
		return nil, err
	}

	node.Version = "omniops/halo:" + tag
	node.PrivvalKey = valKey.PrivKey
	node.NodeKey = nodeKey.PrivKey

	// Add seeds (cometBFT only adds seeds defined explicitly per node, we auto-add all seeds).
	seeds := manifest.Seeds()
	for seed := range seeds {
		if seed == node.Name {
			continue // Skip self
		}
		node.Seeds = append(node.Seeds, testnet.LookupNode(seed))
	}
	// Remove seeds from persisted peers (cometBFT adds all nodes as peers by default).
	var persisted []*e2e.Node
	for _, peer := range node.PersistentPeers {
		if seeds[peer.Name] {
			continue
		}
		persisted = append(persisted, peer)
	}
	node.PersistentPeers = persisted

	return node, nil
}

// runsDir returns the runs directory for a given manifest file.
// E.g. /path/to/manifests/manifest.toml > /path/to/runs/manifest.
func runsDir(manifestFile string) string {
	resp := strings.TrimSuffix(manifestFile, filepath.Ext(manifestFile))
	return strings.Replace(resp, "manifests", "runs", 1)
}

// LoadManifest loads a manifest from disk.
func LoadManifest(path string) (types.Manifest, error) {
	manifest := types.Manifest{}
	_, err := toml.DecodeFile(path, &manifest)
	if err != nil {
		return manifest, errors.Wrap(err, "decode manifest")
	}

	return manifest, nil
}

func NoNodesTestnet(manifest types.Manifest, infd types.InfrastructureData, cfg DefinitionConfig) (types.Testnet, error) {
	publics, err := publicChains(manifest, cfg)
	if err != nil {
		return types.Testnet{}, err
	}

	cmtTestnet, err := noNodesTestnet(manifest.Manifest, cfg.ManifestFile, infd.InfrastructureData)
	if err != nil {
		return types.Testnet{}, errors.Wrap(err, "testnet from manifest")
	}

	return types.Testnet{
		Network:      manifest.Network,
		Testnet:      cmtTestnet,
		PublicChains: publics,
		OnlyMonitor:  manifest.OnlyMonitor,
	}, nil
}

// noNodesTestnet returns a bare minimum instance of *e2e.Omega. It doesn't have any nodes or chain details setup.
func noNodesTestnet(manifest e2e.Manifest, file string, ifd e2e.InfrastructureData) (*e2e.Testnet, error) {
	dir := strings.TrimSuffix(file, filepath.Ext(file))

	_, ipNet, err := net.ParseCIDR(ifd.Network)
	if err != nil {
		return nil, errors.Wrap(err, "parse network ip", "network", ifd.Network)
	}

	testnet := &e2e.Testnet{
		Name:         filepath.Base(dir),
		File:         file,
		Dir:          runsDir(file),
		IP:           ipNet,
		InitialState: manifest.InitialState,
		Prometheus:   manifest.Prometheus,
	}

	return testnet, nil
}

//nolint:nosprintfhostport // Not an issue for non-critical e2e test code.
func TestnetFromManifest(ctx context.Context, manifest types.Manifest, infd types.InfrastructureData, cfg DefinitionConfig) (types.Testnet, error) {
	if manifest.OnlyMonitor || len(manifest.Nodes) == 0 {
		// Create a bare minimum comet testnet only with test di, prometheus and ipnet.
		// Otherwise e2e.NewTestnetFromManifest panics because there are no nodes set
		// in the only_monitor manifest.
		return NoNodesTestnet(manifest, infd, cfg)
	}

	cmtTestnet, err := e2e.NewTestnetFromManifest(manifest.Manifest, cfg.ManifestFile, infd.InfrastructureData)
	if err != nil {
		return types.Testnet{}, errors.Wrap(err, "testnet from manifest")
	}
	cmtTestnet, err = adaptCometTestnet(ctx, manifest, cmtTestnet, cfg.OmniImgTag)
	if err != nil {
		return types.Testnet{}, errors.Wrap(err, "adapt comet testnet")
	}

	var omniEVMS []types.OmniEVM
	for name, isArchive := range manifest.OmniEVMs() {
		inst, ok := infd.Instances[name]
		if !ok {
			return types.Testnet{}, errors.New("omni evm instance not found in infrastructure data")
		}

		pk, err := getOrGenKey(ctx, manifest, name, key.P2PExecution)
		if err != nil {
			return types.Testnet{}, errors.Wrap(err, "execution node key")
		}
		nodeKey, err := pk.ECDSA()
		if err != nil {
			return types.Testnet{}, err
		}

		en := enode.NewV4(&nodeKey.PublicKey, inst.IPAddress, 30303, 30303)

		internalIP := inst.IPAddress.String()
		advertisedIP := inst.ExtIPAddress // EVM P2P NAT advertised address.
		if infd.Provider == docker.ProviderName {
			internalIP = name             // For docker, we use container names
			advertisedIP = inst.IPAddress // For docker, we use container IPs for evm p2p networking, not localhost.
		}

		omniEVMS = append(omniEVMS, types.OmniEVM{
			Chain:        types.OmniEVMByNetwork(manifest.Network),
			InstanceName: name,
			AdvertisedIP: advertisedIP,
			ProxyPort:    inst.Port,
			InternalRPC:  fmt.Sprintf("http://%s:8545", internalIP),
			ExternalRPC:  fmt.Sprintf("http://%s:%d", inst.ExtIPAddress.String(), inst.Port),
			NodeKey:      nodeKey,
			Enode:        en,
			IsArchive:    isArchive,
			JWTSecret:    tutil.RandomHash().Hex(),
		})
	}

	// Second pass to mesh the bootnodes
	for i := range omniEVMS {
		var bootnodes []*enode.Node
		for j, bootEVM := range omniEVMS {
			if i == j {
				continue // Skip self
			}
			bootnodes = append(bootnodes, bootEVM.Enode)
		}
		omniEVMS[i].Peers = bootnodes
	}

	anvilEVMs, err := types.AnvilChainsByNames(manifest.AnvilChains)
	if err != nil {
		return types.Testnet{}, err
	}

	var anvils []types.AnvilChain
	for _, chain := range anvilEVMs {
		inst, ok := infd.Instances[chain.Name]
		if !ok {
			return types.Testnet{}, errors.New("anvil chain instance not found in infrastructure data")
		}

		internalIP := inst.IPAddress.String()
		if infd.Provider == docker.ProviderName {
			internalIP = chain.Name // For docker, we use container names
		}

		anvils = append(anvils, types.AnvilChain{
			Chain:       chain,
			InternalIP:  inst.IPAddress,
			ProxyPort:   inst.Port,
			LoadState:   "./anvil/state.json",
			InternalRPC: fmt.Sprintf("http://%s:8545", internalIP),
			ExternalRPC: fmt.Sprintf("http://%s:%d", inst.ExtIPAddress.String(), inst.Port),
		})
	}

	publics, err := publicChains(manifest, cfg)
	if err != nil {
		return types.Testnet{}, err
	}

	return types.Testnet{
		Network:      manifest.Network,
		Testnet:      cmtTestnet,
		OmniEVMs:     omniEVMS,
		AnvilChains:  anvils,
		PublicChains: publics,
		Explorer:     manifest.Explorer,
		Perturb:      manifest.Perturb,
	}, nil
}

// getOrGenKey gets (based on manifest) or creates a private key for the given node and type.
func getOrGenKey(ctx context.Context, manifest types.Manifest, nodeName string, typ key.Type) (key.Key, error) {
	addr, ok := manifest.Keys[nodeName][typ]
	if !ok { // No key in manifest
		// Generate an insecure deterministic key for devnet
		if manifest.Network == netconf.Devnet {
			return key.GenerateInsecureDeterministic(manifest.Network, typ, nodeName), nil
		}

		// Otherwise generate a proper key
		return key.Generate(typ), nil
	}

	// Address configured in manifest, download from GCP
	return key.Download(ctx, manifest.Network, nodeName, typ, addr)
}

func publicChains(manifest types.Manifest, cfg DefinitionConfig) ([]types.PublicChain, error) {
	var publics []types.PublicChain
	for _, name := range manifest.PublicChains {
		chain, err := types.PublicChainByName(name)
		if err != nil {
			return nil, errors.Wrap(err, "get public chain")
		}

		addr, ok := cfg.RPCOverrides[name]
		if !ok {
			addr = types.PublicRPCByName(name)
		}

		publics = append(publics, types.NewPublicChain(chain, strings.Split(addr, ",")))
	}

	return publics, nil
}

// externalEndpoints returns the evm rpc endpoints for access from inside the
// docker network.
func internalEndpoints(def Definition, nodePrefix string) xchain.RPCEndpoints {
	endpoints := make(xchain.RPCEndpoints)

	// Add all public chains
	for _, public := range def.Testnet.PublicChains {
		endpoints[public.Chain().Name] = public.NextRPCAddress()
	}

	// In monitor only mode, there is only public chains, so skip omni and anvil chains.
	if def.Testnet.OnlyMonitor {
		return endpoints
	}

	omniEVM := omniEVMByPrefix(def.Testnet, nodePrefix)
	endpoints[omniEVM.Chain.Name] = omniEVM.InternalRPC

	node := nodeByPrefix(def.Testnet, nodePrefix)
	endpoints[def.Testnet.Network.Static().OmniConsensusChain().Name] = node.AddressRPC()

	// Add all anvil chains
	for _, anvil := range def.Testnet.AnvilChains {
		endpoints[anvil.Chain.Name] = anvil.InternalRPC
	}

	return endpoints
}

// externalEndpoints returns the evm rpc endpoints for access from outside the
// docker network.
func externalEndpoints(def Definition) xchain.RPCEndpoints {
	endpoints := make(xchain.RPCEndpoints)

	// Add all public chains
	for _, public := range def.Testnet.PublicChains {
		endpoints[public.Chain().Name] = public.NextRPCAddress()
	}

	// In monitor only mode, there is only public chains, so skip omni and anvil chains.
	if def.Testnet.OnlyMonitor {
		return endpoints
	}

	// Connect to a proper omni_evm that isn't unavailable
	omniEVM := def.Testnet.BroadcastOmniEVM()
	endpoints[omniEVM.Chain.Name] = omniEVM.ExternalRPC

	// Add omni consensus chain
	endpoints[def.Testnet.Network.Static().OmniConsensusChain().Name] = def.Testnet.BroadcastNode().AddressRPC()

	// Add all anvil chains
	for _, anvil := range def.Testnet.AnvilChains {
		endpoints[anvil.Chain.Name] = anvil.ExternalRPC
	}

	return endpoints
}

// networkFromDef returns the network configuration from the definition.
func networkFromDef(def Definition) netconf.Network {
	var chains []netconf.Chain

	newChain := func(chain types.EVMChain) netconf.Chain {
		depInfo := def.DeployInfos()[chain.ChainID]
		return netconf.Chain{
			ID:             chain.ChainID,
			Name:           chain.Name,
			BlockPeriod:    chain.BlockPeriod,
			Shards:         chain.Shards,
			AttestInterval: chain.AttestInterval(def.Testnet.Network),
			PortalAddress:  depInfo[types.ContractPortal].Address,
			DeployHeight:   depInfo[types.ContractPortal].Height,
		}
	}

	// Add all public chains
	for _, public := range def.Testnet.PublicChains {
		chains = append(chains, newChain(public.Chain()))
	}

	// In monitor only mode, there is only public chains, so skip omni and anvil chains.
	if def.Testnet.OnlyMonitor {
		return netconf.Network{
			ID:     def.Testnet.Network,
			Chains: chains,
		}
	}

	// Connect to a proper omni_evm that isn't unavailable
	omniEVM := def.Testnet.BroadcastOmniEVM()
	chains = append(chains, newChain(omniEVM.Chain))

	// Add omni consensus chain
	chains = append(chains, def.Testnet.Network.Static().OmniConsensusChain())

	// Add all anvil chains
	for _, anvil := range def.Testnet.AnvilChains {
		chains = append(chains, newChain(anvil.Chain))
	}

	return netconf.Network{
		ID:     def.Testnet.Network,
		Chains: chains,
	}
}

// omniEVMByPrefix returns a omniEVM from the testnet with the given prefix.
// Or broadcast omniEVM if prefix is empty.
// Or the only omniEVM if there is only one.
func omniEVMByPrefix(testnet types.Testnet, prefix string) types.OmniEVM {
	if prefix == "" {
		return testnet.BroadcastOmniEVM()
	} else if len(testnet.OmniEVMs) == 1 {
		return testnet.OmniEVMs[0]
	}

	for _, evm := range testnet.OmniEVMs {
		if strings.HasPrefix(evm.InstanceName, prefix) {
			return evm
		}
	}

	panic("evm not found")
}

// nodeByPrefix returns a halo node from the testnet with the given prefix.
// Or a random node if prefix is empty.
// Or the only node if there is only one.
func nodeByPrefix(testnet types.Testnet, prefix string) *e2e.Node {
	if prefix == "" {
		return random(testnet.Nodes)
	} else if len(testnet.Nodes) == 1 {
		return testnet.Nodes[0]
	}

	for _, node := range testnet.Nodes {
		if strings.HasPrefix(node.Name, prefix) {
			return node
		}
	}

	panic("node not found")
}

// random returns a random item from a slice.
func random[T any](items []T) T {
	var zero T
	if len(items) == 0 {
		return zero
	}

	return items[int(time.Now().UnixNano())%len(items)]
}

// lazyNetwork is a lazy network setup that initializes the backends and netman only if required.
// Some e2e commands do not require networking, so this mitigates the need for special networking flags in that case.
type lazyNetwork struct {
	once     sync.Once
	initFunc func() (ethbackend.Backends, netman.Manager, error)
	backends ethbackend.Backends
	netman   netman.Manager
}

func (l *lazyNetwork) Init() error {
	var err error
	l.once.Do(func() {
		l.backends, l.netman, err = l.initFunc()
	})

	return err
}

func (l *lazyNetwork) mustInit() {
	if err := l.Init(); err != nil {
		panic(err)
	}
}

func (l *lazyNetwork) MustBackends() ethbackend.Backends {
	l.mustInit()
	return l.backends
}

func (l *lazyNetwork) MustNetman() netman.Manager {
	l.mustInit()
	return l.netman
}
