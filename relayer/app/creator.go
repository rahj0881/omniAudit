package relayer

import (
	"github.com/omni-network/omni/lib/xchain"
)

// CreateSubmissions creates submissions from the given stream update by destination chain id.
// It creates a merkle tree from the block and uses it to create a multi-proof for the given messages.
func CreateSubmissions(streamUpdate StreamUpdate) ([]xchain.Submission, error) {
	// todo(lazar): in future this will receive receipts as well
	tree, err := xchain.NewBlockTree(xchain.Block{
		BlockHeader: streamUpdate.AggAttestation.BlockHeader,
		Msgs:        streamUpdate.Msgs,
	})
	if err != nil {
		return nil, err
	}

	groupedMsgs := map[uint64][]xchain.Msg{}
	for _, msg := range streamUpdate.Msgs {
		groupedMsgs[msg.StreamID.DestChainID] = append(groupedMsgs[msg.StreamID.DestChainID], msg)
	}

	submissions := make([]xchain.Submission, len(groupedMsgs))

	for i, msgs := range groupedMsgs {
		multi, err := tree.Proof(streamUpdate.AggAttestation.BlockHeader, msgs)
		if err != nil {
			return nil, err
		}

		submissions[i] = xchain.Submission{
			AttestationRoot: tree.Root(),
			BlockHeader:     streamUpdate.AggAttestation.BlockHeader,
			Msgs:            msgs,
			Proof:           multi.Proof,
			ProofFlags:      multi.ProofFlags,
			Signatures:      streamUpdate.AggAttestation.Signatures,
		}
	}

	return submissions, nil
}
