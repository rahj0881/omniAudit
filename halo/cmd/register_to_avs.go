package cmd

import (
	"context"
	"math/big"
	"os"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/cometbft/cometbft/privval"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/spf13/cobra"
)

func RegisterOperatorToOmniAVS(cfg *OperatorConfig) *cobra.Command {
	registerToAVSCmd := &cobra.Command{
		Use:   "register",
		Short: "register validator to omni avs",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg.HaloConfig.HomeDir = cfg.HomeDir

			// read the comet config based on the home directory
			cometCfg, err := parseCometConfig(cmd.Context(), cfg.HaloConfig.HomeDir)
			if err != nil {
				return err
			}
			cfg.CometConfig = cometCfg

			return register(cmd.Context(), cfg)
		},
	}

	bindOperatorFlags(registerToAVSCmd.Flags(), cfg)

	return registerToAVSCmd
}

func register(ctx context.Context, cfg *OperatorConfig) error {
	// check for home directory where the config files exist
	if !directoryExists(cfg.HaloConfig.HomeDir) {
		log.Info(ctx, "Make sure to run \"init\" command before running operator commands")
		err := errors.New("directory does not exists", "home", cfg.HaloConfig.HomeDir)

		return err
	}

	// load network config for the layer1 chain
	chain, err := getChainConfig(cfg)
	if err != nil {
		return err
	}

	// load a private validator key and state from disk (this hard exits on any error).
	privVal := privval.LoadFilePVEmptyState(cfg.CometConfig.PrivValidatorKeyFile(), cfg.CometConfig.PrivValidatorStateFile())

	// connect to the rpc endpoint
	client, err := ethclient.Dial(chain.Name, chain.RPCURL)
	if err != nil {
		return err
	}

	// load contract bindings
	omniAvs, err := bindings.NewOmniAVS(common.HexToAddress(cfg.AVSDirectoryAddr), client)
	if err != nil {
		return err
	}
	avsDirectory, err := bindings.NewAVSDirectory(common.HexToAddress(cfg.OmniAVSAddr), client)
	if err != nil {
		return err
	}

	operPK, err := crypto.ToECDSA(privVal.Key.PrivKey.Bytes())
	if err != nil {
		log.Info(ctx, "Could not convert private keys", err)

		return errors.Wrap(err, "could not convert pk to ecdsa")
	}
	operAddr := common.HexToAddress(privVal.GetAddress().String())

	// calculate operator signature and digest hash
	blockNumber, err := client.BlockNumber(ctx)
	if err != nil {
		return errors.Wrap(err, "getting blockNumber ")
	}
	block, err := client.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		return errors.Wrap(err, "getting blockByNumber ")
	}

	operatorSignatureWithSaltAndExpiry := bindings.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: []byte{0},
		Salt:      crypto.Keccak256Hash(operAddr.Bytes()),
		Expiry:    big.NewInt(int64(block.Time()) + int64(24*time.Hour)),
	}
	digestHash, err := avsDirectory.CalculateOperatorAVSRegistrationDigestHash(&bind.CallOpts{},
		common.HexToAddress(privVal.GetAddress().String()),
		common.HexToAddress(cfg.OmniAVSAddr),
		operatorSignatureWithSaltAndExpiry.Salt,
		operatorSignatureWithSaltAndExpiry.Expiry)
	if err != nil {
		return err
	}

	operatorSignatureWithSaltAndExpiry.Signature, err = crypto.Sign(digestHash[:32], operPK)
	if err != nil {
		return errors.Wrap(err, "error signing)")
	}

	if len(operatorSignatureWithSaltAndExpiry.Signature) != 65 {
		return errors.New("invalid signature length")
	}
	operatorSignatureWithSaltAndExpiry.Signature[64] += 27
	txOpts, err := bind.NewKeyedTransactorWithChainID(operPK, big.NewInt(int64(chain.ID)))
	if err != nil {
		return errors.Wrap(err, "error getting txopts")
	}

	txOpts.Context = ctx
	tx, err := omniAvs.RegisterOperatorToAVS(txOpts, operAddr, operatorSignatureWithSaltAndExpiry)
	if err != nil {
		return err
	}
	log.Info(ctx, "Submitted registration to AVS")

	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		return errors.Wrap(err, "error waiting for mining tx")
	}
	log.Info(ctx, "Operator registered with AVS", "address", operAddr.String(), "txHash", receipt.TxHash)

	return nil
}

func getChainConfig(cfg *OperatorConfig) (*netconf.Chain, error) {
	network, err := netconf.Load(cfg.HaloConfig.NetworkFile())
	if err != nil {
		return nil, errors.Wrap(err, "load network config")
	} else if err := network.Validate(); err != nil {
		return nil, errors.Wrap(err, "validate network config")
	}

	for _, c := range network.Chains {
		if c.Name == cfg.L1ChainName {
			return &c, nil
		}
	}

	return nil, errors.New("chain not found")
}

func directoryExists(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}

	return true
}
