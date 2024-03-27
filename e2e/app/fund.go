package app

import (
	"context"

	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"

	"cosmossdk.io/math"
)

// noAnvilDev returns a list of accounts that are not dev anvil accounts.
func noAnvilDev(accounts []common.Address) []common.Address {
	var nonDevAccounts []common.Address
	for _, account := range accounts {
		if !anvil.IsDevAccount(account) {
			nonDevAccounts = append(nonDevAccounts, account)
		}
	}

	return nonDevAccounts
}

// accountsToFund returns a list of accounts to fund on anvil chains, based on the network.
func accountsToFund(network netconf.ID) []common.Address {
	switch network {
	case netconf.Staging:
		return []common.Address{
			contracts.StagingCreate3Deployer(),
			contracts.StagingDeployer(),
			contracts.StagingProxyAdminOwner(),
			contracts.StagingPortalAdmin(),
			contracts.StagingAVSAdmin(),
		}
	case netconf.Devnet:
		return []common.Address{
			contracts.DevnetCreate3Deployer(),
			contracts.DevnetDeployer(),
			contracts.DevnetProxyAdminOwner(),
			contracts.DevnetPortalAdmin(),
			contracts.DevnetAVSAdmin(),
		}
	default:
		return []common.Address{}
	}
}

// fundAccounts funds the EOAs that need funding (just on anvil chains, for now).
func fundAccounts(ctx context.Context, def Definition) error {
	accounts := accountsToFund(def.Testnet.Network)
	eth1M := math.NewInt(1_000_000).MulRaw(params.Ether) // 1_000_000 ETH
	for _, chain := range def.Testnet.AnvilChains {
		if err := anvil.FundAccounts(ctx, chain.ExternalRPC, eth1M.BigInt(), noAnvilDev(accounts)...); err != nil {
			return errors.Wrap(err, "fund anvil account")
		}
	}

	return nil
}
