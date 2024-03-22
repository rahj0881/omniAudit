package fireblocks

import (
	"crypto/rsa"
	"sync"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
)

const (
	endpointTransactions = "/v1/transactions"
	endpointAssets       = "/v1/supported_assets"
	endpointVaults       = "/v1/vault/accounts_paged"
	endpointPubkeyTmpl   = "/v1/vault/accounts/{{.VaultAccountID}}/{{.AssetID}}/0/0/public_key_info?compressed"

	assetHolesky = "ETH_TEST6"
	assetSepolia = "ETH_TEST5"
	assetMainnet = "ETH"

	hostProd    = "https://api.fireblocks.io"
	hostSandbox = "https://sandbox-api.fireblocks.io"
)

// Client is a JSON HTTP client for the FireBlocks API.
type Client struct {
	opts       options
	apiKey     string
	privateKey *rsa.PrivateKey
	jsonHTTP   jsonHTTP
	cache      *accountCache
}

// New creates a new FireBlocks client.
func New(apiKey string, privateKey *rsa.PrivateKey, opts ...func(*options)) (Client, error) {
	if apiKey == "" {
		return Client{}, errors.New("apiKey is required")
	}
	if privateKey == nil {
		return Client{}, errors.New("privateKey is required")
	}

	o := defaultOptions()
	for _, opt := range opts {
		opt(&o)
	}
	if err := o.check(); err != nil {
		return Client{}, errors.Wrap(err, "options check")
	}

	return Client{
		apiKey:     apiKey,
		privateKey: privateKey,
		jsonHTTP:   newJSONHTTP(o.host(), apiKey),
		opts:       o,
		cache:      newAccountCache(o.TestAccounts),
	}, nil
}

// authHeaders returns the authentication headers for the FireBlocks API.
func (c Client) authHeaders(endpoint string, request any) (map[string]string, error) {
	token, err := c.token(endpoint, request)
	if err != nil {
		return nil, errors.Wrap(err, "generating token")
	}

	return map[string]string{
		"X-API-KEY":     c.apiKey,
		"Authorization": "Bearer " + token,
	}, nil
}

func (c Client) getAssetID() string {
	switch c.opts.Network {
	case TestNet:
		return assetHolesky
	case MainNet:
		return assetMainnet
	default:
		return assetMainnet
	}
}

func newAccountCache(init map[common.Address]uint64) *accountCache {
	return &accountCache{
		accountsByAddress: init,
	}
}

type accountCache struct {
	sync.Mutex
	accountsByAddress map[common.Address]uint64
}

func (c *accountCache) Get(addr common.Address) (uint64, bool) {
	c.Lock()
	defer c.Unlock()

	acc, ok := c.accountsByAddress[addr]

	return acc, ok
}

func (c *accountCache) Set(addr common.Address, id uint64) {
	c.Lock()
	defer c.Unlock()

	c.accountsByAddress[addr] = id
}
