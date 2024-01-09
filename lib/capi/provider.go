package capi

import (
	"context"

	"github.com/omni-network/omni/lib/xchain"
)

// ProviderCallback is the callback function signature that will be called with all approved attestation per
// consensus block.
type ProviderCallback func(ctx context.Context, height uint64, aggs []xchain.AggAttestation)

// Provider abstracts connecting to the omni consensus chainand streaming approved
// aggregate attestations from a specific height.
//
// It provides exactly once-delivery guarantees for the callback function.
// It will exponentially backoff and retry forever while the callback function returns an error.
type Provider interface {
	Subscribe(ctx context.Context, height uint64, callback ProviderCallback)
}
