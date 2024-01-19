package consensus

import (
	"sync"

	"github.com/omni-network/omni/halo/attest"
	"github.com/omni-network/omni/lib/engine"
	"github.com/omni-network/omni/lib/xchain"

	abci "github.com/cometbft/cometbft/abci/types"
)

// Core of the Halo consensus client with the following responsibilities:
// - Implements the server side of the ABCI++ interface, see abci.go.
// - Maintains the consensus chain state.
type Core struct {
	// Immutable fields (configured at construction)
	ethCl            engine.API
	state            *State
	attestSvc        attest.Service
	snapshots        *SnapshotStore
	snapshotInterval uint64

	// Mutable restore snapshot fields
	restore struct { //nolint: revive // Nested struct use to isolate mutable fields
		sync.Mutex
		Snapshot *abci.Snapshot
		Chunks   [][]byte
	}
}

// NewCore returns a new Core instance.
func NewCore(ethCl engine.API, attestSvc attest.Service, state *State, snapshots *SnapshotStore,
	snapshotInterval uint64,
) *Core {
	return &Core{
		ethCl:            ethCl,
		attestSvc:        attestSvc,
		state:            state,
		snapshots:        snapshots,
		snapshotInterval: snapshotInterval,
	}
}

// ApprovedAggregates returns a copy of the latest state's approved aggregates.
// For testing purposes only.
func (c *Core) ApprovedAggregates() []xchain.AggAttestation {
	return c.state.ApprovedAggregates()
}

// ApprovedFrom returns a sequential range of approved aggregates from the provided chain ID and height.
// It returns at most max aggregates. Their block heights are sequentially increasing.
// For testing purposes only.
func (c *Core) ApprovedFrom(chainID uint64, height uint64, max uint64) []xchain.AggAttestation {
	return c.state.ApprovedFrom(chainID, height, max)
}

var _ abci.Application = (*Core)(nil)
