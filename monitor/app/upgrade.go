package monitor

import (
	"context"
	"strconv"
	"time"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/log"

	utypes "cosmossdk.io/x/upgrade/types"
)

func monitorUpgradesForever(ctx context.Context, cprov cchain.Provider) {
	ticker := time.NewTicker(time.Second * 15)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			plan, ok, err := cprov.CurrentPlannedPlan(ctx)
			if err != nil {
				log.Warn(ctx, "Failed fetching planned upgrade (will retry)", err)
				continue
			} else if !ok {
				plan = utypes.Plan{
					Name:   "none",
					Height: 0,
				}
			}

			plannedUpgradeGauge.Reset()
			plannedUpgradeGauge.WithLabelValues(plan.Name, strconv.FormatInt(plan.Height, 10)).Set(1)
		}
	}
}
