package relayer

import (
	"context"
	"net/http"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// startMonitoring starts the monitoring goroutines.
func startMonitoring(ctx context.Context, network netconf.Network, xprovider xchain.Provider,
	addr common.Address, rpcClients map[uint64]*ethclient.Client) {
	for _, srcChain := range network.Chains {
		go monitorAccountForever(ctx, addr, srcChain.Name, rpcClients[srcChain.ID])
		go monitorHeadsForever(ctx, srcChain.Name, rpcClients[srcChain.ID])

		for _, dstChain := range network.Chains {
			if srcChain.ID == dstChain.ID {
				continue
			}

			go monitorOffsetsForever(ctx, srcChain.ID, dstChain.ID, srcChain.Name, dstChain.Name, xprovider)
		}
	}
}

// monitorHeadsForever blocks and periodically monitors the heads of the given chain.
func monitorHeadsForever(ctx context.Context, chainName string, client *ethclient.Client) {
	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			monitorHeadsOnce(ctx, chainName, client)
		}
	}
}

// monitorHeadsOnce monitors the heads of the given chain.
func monitorHeadsOnce(ctx context.Context, chainName string, client *ethclient.Client) {
	for _, typ := range []string{"latest", "safe", "finalized"} {
		head, err := getHead(ctx, client, typ)
		if err != nil {
			// Not all chains support all types, so just swallow the errors, this is best effort monitoring.
			continue
		}
		headHeight.WithLabelValues(chainName, typ).Set(float64(head))
	}
}

// monitorAccountsForever blocks and periodically monitors the relayer accounts
// for the given chain.
func monitorAccountForever(ctx context.Context, addr common.Address, chainName string, client *ethclient.Client) {
	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := monitorAccountOnce(ctx, addr, chainName, client)
			if ctx.Err() != nil {
				return
			} else if err != nil {
				log.Error(ctx, "Monitoring account failed (will retry)", err,
					"chain", chainName)

				continue
			}
		}
	}
}

// monitorAccountOnce monitors the relayer account for the given chain.
func monitorAccountOnce(ctx context.Context, addr common.Address, chainName string, client *ethclient.Client) error {
	balance, err := client.BalanceAt(ctx, addr, nil)
	if err != nil {
		return errors.Wrap(err, "balance at")
	}

	nonce, err := client.NonceAt(ctx, addr, nil)
	if err != nil {
		return errors.Wrap(err, "nonce at")
	}

	bf, _ := balance.Float64()
	bf /= params.Ether
	accountBalance.WithLabelValues(chainName).Set(bf)
	accountNonce.WithLabelValues(chainName).Set(float64(nonce))

	return nil
}

// monitorOffsetsForever blocks and periodically monitors the emitted and submitted
// offsets for a given source and destination chain.
func monitorOffsetsForever(ctx context.Context, src, dst uint64, srcChain, dstChain string,
	xprovider xchain.Provider) {
	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := monitorOffsetsOnce(ctx, src, dst, srcChain, dstChain, xprovider)
			if ctx.Err() != nil {
				return
			} else if err != nil {
				log.Error(ctx, "Monitoring stream offsets failed (will retry)", err,
					"src_chain", srcChain, "dst_chain", dstChain)

				continue
			}
		}
	}
}

// monitorOffsetsOnce monitors the emitted and submitted offsets for a given source and
// destination chain.
func monitorOffsetsOnce(ctx context.Context, src, dst uint64, srcChain, dstChain string,
	xprovider xchain.Provider) error {
	emitted, ok, err := xprovider.GetEmittedCursor(ctx, src, dst)
	if err != nil {
		return err
	} else if !ok {
		return nil
	}

	submitted, _, err := xprovider.GetSubmittedCursor(ctx, dst, src)
	if err != nil {
		return err
	}

	emitCursor.WithLabelValues(srcChain, dstChain).Set(float64(emitted.Offset))
	submitCursor.WithLabelValues(srcChain, dstChain).Set(float64(submitted.Offset))

	return nil
}

// serveMonitoring starts a goroutine that serves the monitoring API. It
// returns a channel that will receive an error if the server fails to start.
func serveMonitoring(address string) <-chan error {
	errChan := make(chan error)
	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.Handler())

		srv := &http.Server{
			Addr:              address,
			ReadHeaderTimeout: 5 * time.Second,
			IdleTimeout:       5 * time.Second,
			WriteTimeout:      5 * time.Second,
			Handler:           mux,
		}
		errChan <- errors.Wrap(srv.ListenAndServe(), "serve monitoring")
	}()

	return errChan
}

// getHead returns the head of the chain for the given type.
func getHead(ctx context.Context, rpcClient *ethclient.Client, typ string) (uint64, error) {
	var header *types.Header
	err := rpcClient.Client().CallContext(
		ctx,
		&header,
		"eth_getBlockByNumber",
		typ,
		false,
	)
	if err != nil {
		return 0, errors.Wrap(err, "could not get block")
	}

	return header.Number.Uint64(), nil
}
