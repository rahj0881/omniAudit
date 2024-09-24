package module

import (
	"github.com/omni-network/omni/halo/attest/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

// noopMigration doesn't perform any store migrations.
var noopMigration = func(_ sdk.Context) error { return nil }

func registerMigrations(cfg module.Configurator) {
	migrations := []struct {
		FromVersion uint64
		Handler     module.MigrationHandler
	}{
		{
			// 1_uluwatu doesn't include any store migrations.
			// It only includes constructor and logic changes.
			FromVersion: 1,
			Handler:     noopMigration,
		},
	}

	for _, m := range migrations {
		err := cfg.RegisterMigration(types.ModuleName, m.FromVersion, m.Handler)
		if err != nil {
			panic(err)
		}
	}
}
