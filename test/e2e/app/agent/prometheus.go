package agent

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/test/e2e/types"

	_ "embed"
)

type Secrets struct {
	URL  string
	User string
	Pass string
}

const promPort = 26660 // Default metrics port for all omni apps (from cometBFT)

//go:embed prometheus.yml.tmpl
var promConfigTmpl []byte

func WriteConfig(ctx context.Context, testnet types.Testnet, secrets Secrets) error {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	bz, err := genPromConfig(ctx, testnet, secrets, hostname)
	if err != nil {
		return errors.Wrap(err, "generating prometheus config")
	}

	promFile := filepath.Join(testnet.Dir, "prometheus", "prometheus.yml")
	if err := os.MkdirAll(filepath.Dir(promFile), 0755); err != nil {
		return errors.Wrap(err, "creating prometheus dir")
	}

	if err := os.WriteFile(promFile, bz, 0644); err != nil {
		return errors.Wrap(err, "writing prometheus config")
	}

	return nil
}

func genPromConfig(ctx context.Context, testnet types.Testnet, secrets Secrets, hostname string) ([]byte, error) {
	var nodeTargets []string
	for _, node := range testnet.Nodes {
		// Prometheus is always inside the same docker-compose, so use service names.
		nodeTargets = append(nodeTargets, fmt.Sprintf("%s:%d", node.Name, promPort))
	}

	network := testnet.Network
	if network == netconf.Devnet {
		network = fmt.Sprintf("%s-%s", testnet.Name, hostname)
	}

	if secrets.URL == "" {
		log.Warn(ctx, "Prometheus remote URL not set, metrics not being pushed to Grafana cloud", nil)
	} else {
		log.Info(ctx, "Prometheus metrics pushed to Grafana cloud", "network", network)
	}

	data := promTmplData{
		Network:        network,
		Host:           hostname,
		RemoteURL:      secrets.URL,
		RemoteUsername: secrets.User,
		RemotePassword: secrets.Pass,
		ScrapeConfigs: []promScrapConfig{
			{
				JobName:     "relayer",
				MetricsPath: "/metrics",
				targets:     []string{fmt.Sprintf("relayer:%d", promPort)},
			},
			{
				JobName:     "halo",
				MetricsPath: "/metrics",
				targets:     nodeTargets,
			},
			{
				JobName:     "monitor",
				MetricsPath: "/metrics",
				targets:     []string{fmt.Sprintf("monitor:%d", promPort)},
			},
		},
	}

	t, err := template.New("").Parse(string(promConfigTmpl))
	if err != nil {
		return nil, errors.Wrap(err, "parsing template")
	}

	var bz bytes.Buffer
	if err := t.Execute(&bz, data); err != nil {
		return nil, errors.Wrap(err, "executing template")
	}

	return bz.Bytes(), nil
}

type promTmplData struct {
	Network        string            // Used a "network" label to all metrics
	Host           string            // Hostname of the docker host machine
	RemoteURL      string            // URL to the Grafana cloud server
	RemoteUsername string            // Username to the Grafana cloud server
	RemotePassword string            // Password to the Grafana cloud server
	ScrapeConfigs  []promScrapConfig // List of scrape configs
}

type promScrapConfig struct {
	JobName     string
	MetricsPath string
	targets     []string
}

func (c promScrapConfig) Targets() string {
	return strings.Join(c.targets, ",")
}

// ConfigForHost returns a new prometheus agent config with the given host and halo targets.
//
//	It removes the relayer targets if not enabled.
//	It replaces the halo targets with provided.
//	It replaces the host label.
func ConfigForHost(bz []byte, newHost string, halos []string, relayer bool, monitor bool) []byte {
	if !relayer {
		// Remove relayer target if not needed.
		bz = regexp.MustCompile(`(?m)\[.*\] # relayer targets$`).
			ReplaceAll(bz, []byte(`[] # relayer targets`))
	}

	if !monitor {
		// Remove monitor target if not needed.
		bz = regexp.MustCompile(`(?m)\[.*\] # monitor targets$`).
			ReplaceAll(bz, []byte(`[] # monitor targets`))
	}

	var haloTargets []string
	for _, halo := range halos {
		haloTargets = append(haloTargets, fmt.Sprintf(`"%s:%d"`, halo, promPort))
	}
	replace := fmt.Sprintf(`[%s] # halo targets`, strings.Join(haloTargets, ","))
	bz = regexp.MustCompile(`(?m)\[.*\] # halo targets$`).
		ReplaceAll(bz, []byte(replace))

	bz = regexp.MustCompile(`(?m)host: '.*'$`).
		ReplaceAll(bz, []byte(fmt.Sprintf(`host: '%s'`, newHost)))

	return bz
}
