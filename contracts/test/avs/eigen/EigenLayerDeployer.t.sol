// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { ERC20PresetFixedSupply } from "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import { IStrategyManager } from "eigenlayer-contracts/src/contracts/interfaces/IStrategyManager.sol";
import { StrategyManager } from "eigenlayer-contracts/src/contracts/core/StrategyManager.sol";
import { DelegationManager } from "eigenlayer-contracts/src/contracts/core/DelegationManager.sol";
import { AVSDirectory } from "eigenlayer-contracts/src/contracts/core/AVSDirectory.sol";
import { StrategyBase } from "eigenlayer-contracts/src/contracts/strategies/StrategyBase.sol";

import { EigenLayerGoerli } from "./deploy/EigenLayerGoerli.sol";
import { EigenLayerLocal } from "./deploy/EigenLayerLocal.sol";
import { IEigenDeployer } from "./deploy/IEigenDeployer.sol";

import { console } from "forge-std/console.sol";
import { Test } from "forge-std/Test.sol";

/**
 * @dev Repurposed from eignlayer-contracts src/test/EigenLayerDeployer.t.sol
 *      Unused storage variables and functions were removed
 * @custom:attribution https://github.com/Layr-Labs/eigenlayer-contracts/blob/m2-mainnet-fixes/src/test/EigenLayerDeployer.t.sol
 */
contract EigenLayerDeployer is Test {
    // eigen deployments
    DelegationManager delegation;
    AVSDirectory avsDirectory;
    StrategyManager strategyManager;

    // stragies
    address[] strategies;

    // unsupported strategy (always excluded from OmniAVS strategy params)
    StrategyBase unsupportedStrat;

    function isGoerli() public view returns (bool) {
        return block.chainid == 5;
    }

    function setUp() public virtual {
        IEigenDeployer deployer;

        if (isGoerli()) {
            deployer = new EigenLayerGoerli();
        } else {
            deployer = new EigenLayerLocal();
        }

        IEigenDeployer.Deployments memory deployments = deployer.deploy();

        // we always deploy unsupported strategy
        unsupportedStrat = StrategyBase(
            _deployUnsupportedStrategy(deployments.strategyManager, deployments.proxyAdmin, deployments.pauserRegistry)
        );

        delegation = DelegationManager(deployments.delegationManager);
        avsDirectory = AVSDirectory(deployments.avsDirectory);
        strategyManager = StrategyManager(deployments.strategyManager);
        strategies = deployments.strategies;
    }

    function _deployUnsupportedStrategy(address strategyManager_, address proxyAdmin_, address pauserRegistry_)
        internal
        returns (address)
    {
        uint256 totalSupply = 1000e18;
        IERC20 unsupportedToken = new ERC20PresetFixedSupply("unsupported", "UNSUPPORTED", totalSupply, address(this));
        StrategyBase impl = new StrategyBase(IStrategyManager(strategyManager_));
        return address(
            new TransparentUpgradeableProxy(
                address(impl),
                address(proxyAdmin_),
                abi.encodeWithSelector(StrategyBase.initialize.selector, unsupportedToken, pauserRegistry_)
            )
        );
    }
}
