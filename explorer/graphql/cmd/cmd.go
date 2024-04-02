// Package cmd provides the cli for running the api.
package cmd

import (
	"github.com/omni-network/omni/explorer/graphql/app"
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// New returns a new root cobra command that runs the graphql server.
func New() *cobra.Command {
	cmd := libcmd.NewRootCmd("graphql", "Explorer GraphQL Server")

	cfg := app.DefaultConfig()
	logCfg := log.DefaultConfig()

	bindGraphQLFlags(cmd.Flags(), &cfg)
	log.BindFlags(cmd.Flags(), &logCfg)

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		ctx, err := log.Init(cmd.Context(), logCfg)
		if err != nil {
			return err
		}

		if err := libcmd.LogFlags(ctx, cmd.Flags()); err != nil {
			return err
		}

		return app.Run(ctx, cfg)
	}

	return cmd
}

func bindGraphQLFlags(flags *pflag.FlagSet, cfg *app.Config) {
	flags.StringVar(&cfg.DBUrl, "db-url", cfg.DBUrl, "URL to the database")
	flags.StringVar(&cfg.ListenAddress, "listen-addr", cfg.ListenAddress, "Address to listen on")
}
