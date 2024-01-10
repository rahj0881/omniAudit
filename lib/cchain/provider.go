package cchain

import (
	"context"

	"github.com/omni-network/omni/lib/xchain"
)

// ProviderCallback is the callback function signature that will be called with all approved attestation per
// consensus block.
type ProviderCallback func(ctx context.Context, height uint64, approved []xchain.AggAttestation) error

// Provider abstracts connecting to the omni consensus chain and streaming approved
// aggregate attestations for each block from a specific height.
//
// It provides exactly once-delivery guarantees for the callback function.
// It will exponentially backoff and retry forever while the callback function returns an error.
type Provider interface {
	// Subscribe registers a callback function that will be called with all approved aggregate
	// attestations (as they become available per block) on the consensus chain from the provided height (inclusive).
	Subscribe(ctx context.Context, height uint64, callback ProviderCallback)
}
