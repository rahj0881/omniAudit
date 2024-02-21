package txsender

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"
	"github.com/omni-network/omni/test/e2e/netman"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type TxSenderManager struct {
	txSender map[uint64]Sender
	abi      *abi.ABI
}

func Deploy(
	ctx context.Context,
	portals map[uint64]netman.Portal,
	privateKey *ecdsa.PrivateKey,
) (TxSenderManager, error) {
	// create ABI interface
	parsedAbi, err := abi.JSON(strings.NewReader(bindings.OmniPortalMetaData.ABI))
	if err != nil {
		return TxSenderManager{}, errors.Wrap(err, "parse abi error")
	}

	// create tx sender manager
	manager := TxSenderManager{
		txSender: make(map[uint64]Sender),
		abi:      &parsedAbi,
	}

	// deploy tx sender for each portal
	for _, portal := range portals {
		err := manager.deployTx(ctx, portal, privateKey)
		if err != nil {
			return TxSenderManager{}, errors.Wrap(err, "deploy tx sender")
		}
	}

	return manager, nil
}

func (s TxSenderManager) deployTx(
	ctx context.Context,
	portal netman.Portal,
	privateKey *ecdsa.PrivateKey,
) error {
	chain := portal.Chain

	if _, ok := s.txSender[chain.ID]; ok {
		return errors.New("tx sender already exists", "chain", chain.ID)
	}

	txSender, err := DeployTxSender(
		ctx,
		portal.RPCURL,
		chain.BlockPeriod,
		chain.ID,
		portal.DeployInfo.PortalAddress,
		portal.Client,
		privateKey,
		chain.Name,
		*s.abi,
	)
	if err != nil {
		return errors.Wrap(err, "deploy tx sender", "chain", chain.ID)
	}

	s.txSender[chain.ID] = txSender

	return nil
}

func (s TxSenderManager) SendXCallTransaction(
	ctx context.Context,
	msg xchain.Msg,
	value *big.Int,
	sourceChainID uint64,
) error {
	txSender := s.txSender[sourceChainID]
	bytes, err := s.XCallBytes(MsgToBindings(msg))
	if err != nil {
		return errors.Wrap(err, "get xsubmit bytes")
	}

	return txSender.sendTransaction(ctx, msg.DestChainID, bytes, value)
}

// getXCallBytes returns the byte representation of the xcall function call.
func (s TxSenderManager) XCallBytes(
	sub bindings.XTypesMsg,
) ([]byte, error) {
	bytes, err := s.abi.Pack("xcall", sub)
	if err != nil {
		return nil, errors.Wrap(err, "pack xcall")
	}

	return bytes, nil
}
