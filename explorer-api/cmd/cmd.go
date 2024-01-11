// Package cmd provides the cli for running the explorer-api.
package cmd

import (
	"github.com/omni-network/omni/explorer-api/app"
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/log"

	"github.com/spf13/cobra"
)

// New returns a new root cobra command that handles our command line tool.
func New() *cobra.Command {
	return libcmd.NewRootCmd(
		"explorer-api",
		"Explorer API is a service that serves as the intermediary between our Explorer and our Omni Blocks DB while generating appropriate response objects and coalesncing data for the explorer to show/visualize",
		newRunCmd(),
	)
}

// newRunCmd returns a new cobra command that runs the explorer-api.
func newRunCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "Runs the explorer-api",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			log.Info(ctx, "ExplorerAPI started")
			conf := app.Config{}

			err := app.Run(ctx, conf)
			if err != nil {
				log.Error(ctx, "failed to start Explorer API", err)
				<-ctx.Done()
				return err
			}

			log.Info(ctx, "Press Ctrl+C to stop")
			<-ctx.Done()
			log.Info(ctx, "explorer-api stopped")

			return nil
		},
	}
}
