package types

import (
	"context"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/ethereum/go-ethereum/common"

	"cosmossdk.io/depinject"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// VoteExtensionProvider abstracts logic that provides consensus payload messages
// from the last commits vote extensions.
//
// EVMEngine calls this during PreparePayload to collect all vote extensions msgs to include in
// the consensus block.
type VoteExtensionProvider interface {
	PrepareVotes(ctx context.Context, commit abci.ExtendedCommitInfo) ([]sdk.Msg, error)
}

// EvmEventProcessor abstracts logic that processes EVM log events of the
// previous execution payload (current head) identified by
// the provided block hash.
//
// EVMEngine calls this during PreparePayload to collect all EVM-log-events to include in
// the consensus block. It is also called during ProcessPayload to verify the proposed EVM events.
type EvmEventProcessor interface {
	Name() string
	Prepare(ctx context.Context, blockHash common.Hash) ([]*EVMEvent, error)
	Addresses() []common.Address
	Deliver(ctx context.Context, blockHash common.Hash, log *EVMEvent) error
}

var _ depinject.ManyPerContainerType = InjectedEventProc{}

// InjectedEventProc wraps an EvmEventProcessor such that
// many instances can be injected during app wiring.
type InjectedEventProc struct {
	EvmEventProcessor
}

// InjectEventProc returns an InjectedEventProc that wraps the given EvmEventProcessor.
func InjectEventProc(proc EvmEventProcessor) InjectedEventProc {
	return InjectedEventProc{EvmEventProcessor: proc}
}

func (InjectedEventProc) IsManyPerContainerType() {}
