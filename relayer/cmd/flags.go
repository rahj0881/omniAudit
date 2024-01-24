package cmd

import (
	relayer "github.com/omni-network/omni/relayer/app"

	"github.com/spf13/pflag"
)

func bindRunFlags(flags *pflag.FlagSet, cfg *relayer.Config) {
	flags.StringVar(&cfg.PrivateKey, "private-key", cfg.PrivateKey, "The path to the private key e.g path/private.key")
	flags.StringVar(&cfg.HaloURL, "halo-url", cfg.HaloURL, "The URL of the halo node e.g http://localhost:26657")
}
