package relayer

import (
	"testing"

	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/stretchr/testify/require"
)

func TestGasEstimator(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		network netconf.ID
		msgs    []xchain.Msg
		gas     uint64
	}{
		{
			name: "no msgs",
			gas:  subGasBase,
		},
		{
			name:    "devnet consensus",
			network: netconf.Devnet,
			msgs: []xchain.Msg{
				{
					MsgID: xchain.MsgID{
						StreamID: xchain.StreamID{
							SourceChainID: netconf.Devnet.Static().OmniConsensusChainIDUint64(),
						},
					},
				},
			},
			gas: subEphemeralConsensusGas,
		},
		{
			name:    "mainnet consensus",
			network: netconf.Mainnet,
			msgs: []xchain.Msg{
				{
					MsgID: xchain.MsgID{
						StreamID: xchain.StreamID{
							SourceChainID: netconf.Mainnet.Static().OmniConsensusChainIDUint64(),
						},
					},
				},
			},
			gas: properGasEstimation,
		},
		{
			name:    "arb destination",
			network: netconf.Mainnet,
			msgs: []xchain.Msg{
				{
					MsgID: xchain.MsgID{
						StreamID: xchain.StreamID{
							DestChainID: evmchain.IDArbSepolia,
						},
					},
				},
			},
			gas: properGasEstimation,
		},
		{
			name:    "naive gas model",
			network: netconf.Mainnet,
			msgs: []xchain.Msg{
				{
					MsgID: xchain.MsgID{
						StreamID: xchain.StreamID{
							DestChainID: evmchain.IDOpSepolia,
						},
					},
					DestGasLimit: 99,
				},
			},
			gas: subGasBase + subGasXmsgOverhead + 99,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			estimator := newGasEstimator(test.network)
			gas := estimator(test.msgs)
			require.Equal(t, test.gas, gas)
		})
	}
}
