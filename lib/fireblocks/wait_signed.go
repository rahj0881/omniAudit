package fireblocks

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

func (c Client) WaitSigned(ctx context.Context, opts TransactionRequestOptions) (*TransactionResponse, error) {
	createTransactionRequest, err := NewTransactionRequest(opts)
	if err != nil {
		return nil, err
	}

	resp, err := c.CreateTransaction(ctx, createTransactionRequest)
	if err != nil {
		return nil, err
	}

	transactionID := resp.ID
	attempt := 1
	queryTicker := time.NewTicker(c.cfg.QueryInterval)

	defer queryTicker.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil, errors.Wrap(ctx.Err(), "context canceled")
		case <-queryTicker.C:
			resp, err := c.GetTransactionByID(ctx, transactionID)
			if err != nil {
				return nil, err
			}
			if resp == nil {
				return nil, errors.New("failed to fetch transaction by id")
			}

			pending, err := evaluateTransactionStatus(*resp)
			if err != nil {
				return resp, err
			}
			if attempt%c.cfg.LogFreqFactor == 0 {
				log.Warn(ctx, "Transaction not signed yet", nil, "attempt", attempt)
			}
			if !pending {
				return resp, nil
			}
			attempt++
		}
	}
}

// EvaluateTransactionStatus checks the status of a transaction and returns a boolean indicating whether the transaction is still pending or not.
func evaluateTransactionStatus(resp TransactionResponse) (bool, error) {
	switch resp.Status {
	case "COMPLETED":
		return false, nil
	case "CANCELED", "BLOCKED_BY_POLICY", "REJECTED", "FAILED":
		return false, errors.New("transaction failed", "status", resp.Status)
	default:
		return true, nil
	}
}
