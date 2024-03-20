package agent

import (
	"context"
	"testing"

	"github.com/omni-network/omni/e2e/tutil"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/netconf"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func TestPromGen(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		network      string
		nodes        []string
		newNodes     []string
		newRelayer   bool
		newMonitor   bool
		hostname     string
		agentSecrets bool
	}{
		{
			name:         "manifest1",
			network:      netconf.Devnet,
			nodes:        []string{"validator01", "validator02"},
			hostname:     "localhost",
			newNodes:     []string{"validator01"},
			newRelayer:   false,
			newMonitor:   false,
			agentSecrets: false,
		},
		{
			name:         "manifest2",
			network:      netconf.Staging,
			nodes:        []string{"validator01", "validator02", "fullnode03"},
			hostname:     "vm",
			newNodes:     []string{"fullnode04"},
			newRelayer:   true,
			newMonitor:   false,
			agentSecrets: true,
		},
		{
			name:         "manifest3",
			network:      netconf.Devnet,
			nodes:        []string{"validator01", "validator02"},
			hostname:     "localhost",
			newNodes:     []string{"validator01"},
			newMonitor:   true,
			newRelayer:   false,
			agentSecrets: false,
		},
		{
			name:         "manifest4",
			network:      netconf.Staging,
			nodes:        []string{"validator01", "validator02", "fullnode03"},
			hostname:     "vm",
			newNodes:     []string{"fullnode04"},
			newMonitor:   true,
			newRelayer:   true,
			agentSecrets: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()

			var nodes []*e2e.Node
			for _, name := range test.nodes {
				nodes = append(nodes, &e2e.Node{Name: name})
			}

			testnet := types.Testnet{
				Network: test.network,
				Testnet: &e2e.Testnet{
					Name:  test.name,
					Nodes: nodes,
				},
			}

			var agentSecrets Secrets
			if test.agentSecrets {
				agentSecrets = Secrets{
					URL:  "https://grafana.com",
					User: "admin",
					Pass: "password",
				}
			}

			cfg1, err := genPromConfig(ctx, testnet, agentSecrets, test.hostname)
			require.NoError(t, err)

			cfg2 := ConfigForHost(cfg1, test.hostname+"-2", test.newNodes, test.newRelayer, test.newMonitor)

			t.Run("gen", func(t *testing.T) {
				t.Parallel()
				tutil.RequireGoldenBytes(t, cfg1)
			})

			t.Run("update", func(t *testing.T) {
				t.Parallel()
				tutil.RequireGoldenBytes(t, cfg2)
			})
		})
	}
}
