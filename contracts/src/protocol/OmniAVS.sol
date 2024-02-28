// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { OwnableUpgradeable } from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import { PausableUpgradeable } from "@openzeppelin-upgrades/contracts/security/PausableUpgradeable.sol";

import { IAVSDirectory } from "eigenlayer-contracts/src/contracts/interfaces/IAVSDirectory.sol";
import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import { ISignatureUtils } from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";
import { IServiceManager } from "eigenlayer-middleware/src/interfaces/IServiceManager.sol";

import { OmniPredeploys } from "../libraries/OmniPredeploys.sol";
import { IDelegationManager } from "../interfaces/IDelegationManager.sol";
import { IOmniEthRestaking } from "../interfaces/IOmniEthRestaking.sol";
import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { IOmniAVS } from "../interfaces/IOmniAVS.sol";
import { IOmniAVSAdmin } from "../interfaces/IOmniAVSAdmin.sol";

import { OmniAVSStorage } from "./OmniAVSStorage.sol";

contract OmniAVS is
    IOmniAVS,
    IOmniAVSAdmin,
    IServiceManager,
    OwnableUpgradeable,
    PausableUpgradeable,
    OmniAVSStorage
{
    /// @notice Constant used as a divisor in calculating weights
    uint256 internal constant STRATEGY_WEIGHTING_DIVISOR = 1e18;

    /// @notice EigenLayer core DelegationManager
    IDelegationManager internal immutable _delegationManager;

    /// @notice EigenLayer core AVSDirectory
    IAVSDirectory internal immutable _avsDirectory;

    constructor(IDelegationManager delegationManager_, IAVSDirectory avsDirectory_) {
        _delegationManager = delegationManager_;
        _avsDirectory = avsDirectory_;
        _disableInitializers();
    }

    /// @inheritdoc IOmniAVSAdmin
    function initialize(
        address owner_,
        IOmniPortal omni_,
        uint64 omniChainId_,
        uint96 minOperatorStake_,
        uint32 maxOperatorCount_,
        address[] calldata allowlist_,
        StrategyParams[] calldata strategyParams_
    ) external initializer {
        omni = omni_;
        omniChainId = omniChainId_;
        minOperatorStake = minOperatorStake_;
        maxOperatorCount = maxOperatorCount_;

        _transferOwnership(owner_);
        _setStrategyParams(strategyParams_);

        for (uint256 i = 0; i < allowlist_.length; i++) {
            _allowlist[allowlist_[i]] = true;
        }
    }

    /**
     * Operator registration
     */

    /// @inheritdoc IServiceManager
    function registerOperatorToAVS(
        address operator,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external whenNotPaused {
        require(msg.sender == operator, "OmniAVS: only operator");
        require(_allowlist[operator], "OmniAVS: not allowed");
        require(!_isOperator(operator), "OmniAVS: already an operator"); // we could let _avsDirectory.regsiterOperatorToAVS handle this, they do check
        require(_operators.length < maxOperatorCount, "OmniAVS: max operators reached");
        require(_getSelfDelegations(operator) >= minOperatorStake, "OmniAVS: min stake not met"); // TODO: should this be _getTotalDelegations?

        _avsDirectory.registerOperatorToAVS(operator, operatorSignature);
        _addOperator(operator);

        emit OperatorAdded(operator);
    }

    /// @inheritdoc IServiceManager
    function deregisterOperatorFromAVS(address operator) external whenNotPaused {
        require(msg.sender == operator || msg.sender == owner(), "OmniAVS: only operator or owner");
        require(_isOperator(operator), "OmniAVS: not an operator");

        _avsDirectory.deregisterOperatorFromAVS(operator);
        _removeOperator(operator);

        emit OperatorRemoved(operator);
    }

    /// @dev Adds an operator to the list of operators, does not check if operator already exists
    function _addOperator(address operator) private {
        _operators.push(operator);
    }

    /// @dev Removes an operator from the list of operators
    function _removeOperator(address operator) private {
        for (uint256 i = 0; i < _operators.length; i++) {
            if (_operators[i] == operator) {
                _operators[i] = _operators[_operators.length - 1];
                _operators.pop();
                break;
            }
        }
    }

    /// @dev Returns true if the operator is in the list of operators
    function _isOperator(address operator) private view returns (bool) {
        for (uint256 i = 0; i < _operators.length; i++) {
            if (_operators[i] == operator) {
                return true;
            }
        }
        return false;
    }

    /**
     * Omni sync
     */

    /// @inheritdoc IOmniAVS
    function feeForSync() external view returns (uint256) {
        Validator[] memory vals = _getValidators();
        return omni.feeFor(
            omniChainId, abi.encodeWithSelector(IOmniEthRestaking.sync.selector, vals), _xcallGasLimitFor(vals.length)
        );
    }

    /// @inheritdoc IOmniAVS
    function syncWithOmni() external payable whenNotPaused {
        Validator[] memory vals = _getValidators();
        omni.xcall{ value: msg.value }(
            omniChainId,
            OmniPredeploys.OMNI_ETH_RESTAKING,
            abi.encodeWithSelector(IOmniEthRestaking.sync.selector, vals),
            _xcallGasLimitFor(vals.length)
        );
    }

    /// @dev Returns the gas limit for OmniEthRestaking.sync xcall for some number of validators
    function _xcallGasLimitFor(uint256 numValidators) internal view returns (uint64) {
        return uint64(numValidators * xcallGasLimitPerValidator + xcallBaseGasLimit);
    }

    /**
     * View functions
     */

    /// @inheritdoc IOmniAVS
    function strategyParams() external view returns (StrategyParams[] memory) {
        return _strategyParams;
    }

    /// @inheritdoc IOmniAVS
    function getValidators() external view returns (Validator[] memory) {
        return _getValidators();
    }

    /// @inheritdoc IServiceManager
    function getRestakeableStrategies() external view returns (address[] memory) {
        address[] memory strategies = new address[](_strategyParams.length);
        for (uint256 j = 0; j < _strategyParams.length; j++) {
            strategies[j] = address(_strategyParams[j].strategy);
        }
        return strategies;
    }

    /// @inheritdoc IServiceManager
    function getOperatorRestakedStrategies(address operator) external view returns (address[] memory) {
        address[] memory strategies = new address[](_strategyParams.length);
        for (uint256 j = 0; j < _strategyParams.length; j++) {
            address strat = address(_strategyParams[j].strategy);
            if (_delegationManager.operatorShares(operator, IStrategy(strat)) > 0) {
                strategies[j] = strat;
            }
        }
        return strategies;
    }

    /// @inheritdoc IServiceManager
    function avsDirectory() external view returns (address) {
        return address(_avsDirectory);
    }

    /// @inheritdoc IOmniAVSAdmin
    function isInAllowlist(address operator) external view returns (bool) {
        return _allowlist[operator];
    }

    /// @dev Return current list of Validators, including their personal stake and delegated stake
    function _getValidators() internal view returns (Validator[] memory) {
        Validator[] memory vals = new Validator[](_operators.length);

        for (uint256 i = 0; i < vals.length; i++) {
            address operator = _operators[i];

            uint96 total = _getTotalDelegations(operator);
            uint96 staked = _getSelfDelegations(operator);

            // this should never happen, but just in case
            uint96 delegated = total > staked ? total - staked : 0;

            vals[i] = Validator(operator, delegated, staked);
        }

        return vals;
    }

    /// @dev Returns total delegations to the operator, including self delegations
    function _getTotalDelegations(address operator) internal view returns (uint96) {
        uint96 total;
        StrategyParams memory params;

        for (uint256 j = 0; j < _strategyParams.length; j++) {
            params = _strategyParams[j];
            uint256 shares = _delegationManager.operatorShares(operator, params.strategy);

            // TODO: should we convert shares to underlying?
            // uint256 amt = IStrategy(params.strategy).sharesToUnderlying(shares);
            // This would convert "shares in the stETH strategy" to "stETH tokens"
            // Shares do not map 1:1 to underlying for rebalancing tokens

            total += _weight(shares, params.multiplier);
        }

        return total;
    }

    /// @dev Returns the operator's self-delegations
    function _getSelfDelegations(address operator) internal view returns (uint96) {
        (IStrategy[] memory strategies, uint256[] memory shares) = _delegationManager.getDelegatableShares(operator);

        uint96 staked;
        for (uint256 i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];

            // find the strategy params for the strategy
            StrategyParams memory params;
            for (uint256 j = 0; j < _strategyParams.length; j++) {
                if (address(_strategyParams[j].strategy) == address(strat)) {
                    params = _strategyParams[j];
                    break;
                }
            }

            // if strategy is not found, do not consider it in stake
            if (address(params.strategy) == address(0)) continue;

            // TODO: should we convert shares to underlying?
            // uint256 amt = IStrategy(params.strategy).sharesToUnderlying(shares[i]);
            // This would convert "shares in the stETH strategy" to "stETH tokens"
            // Shares do not map 1:1 to underlying for rebalancing tokens

            staked += _weight(shares[i], params.multiplier);
        }

        return staked;
    }

    /// @dev Returns the weighted stake for shares with specified multiplier
    function _weight(uint256 shares, uint96 multiplier) internal pure returns (uint96) {
        return uint96(shares * multiplier / STRATEGY_WEIGHTING_DIVISOR);
    }

    /**
     * Admin functions
     */

    /// @inheritdoc IServiceManager
    function setMetadataURI(string memory metadataURI) external onlyOwner {
        _avsDirectory.updateAVSMetadataURI(metadataURI);
    }

    /// @inheritdoc IOmniAVSAdmin
    function setOmniPortal(IOmniPortal portal) external onlyOwner {
        omni = portal;
    }

    /// @inheritdoc IOmniAVSAdmin
    function setOmniChainId(uint64 chainId) external onlyOwner {
        omniChainId = chainId;
    }

    /// @inheritdoc IOmniAVSAdmin
    function setStrategyParams(StrategyParams[] calldata params) external onlyOwner {
        _setStrategyParams(params);
    }

    /// @inheritdoc IOmniAVSAdmin
    function setMinOperatorStake(uint96 stake) external onlyOwner {
        minOperatorStake = stake;
    }

    /// @inheritdoc IOmniAVSAdmin
    function setMaxOperatorCount(uint32 count) external onlyOwner {
        maxOperatorCount = count;
    }

    /// @inheritdoc IOmniAVSAdmin
    function setXcallGasLimits(uint64 base, uint64 perValidator) external onlyOwner {
        xcallBaseGasLimit = base;
        xcallGasLimitPerValidator = perValidator;
    }

    /// @inheritdoc IOmniAVSAdmin
    function addToAllowlist(address operator) external onlyOwner {
        require(operator != address(0), "OmniAVS: zero address");
        require(!_allowlist[operator], "OmniAVS: already in allowlist");
        _allowlist[operator] = true;
        emit OperatorAllowed(operator);
    }

    /// @inheritdoc IOmniAVSAdmin
    function removeFromAllowlist(address operator) external onlyOwner {
        require(_allowlist[operator], "OmniAVS: not in allowlist");
        _allowlist[operator] = false;
        emit OperatorDisallowed(operator);
    }

    /// @inheritdoc IOmniAVSAdmin
    function pause() external onlyOwner {
        _pause();
    }

    /// @inheritdoc IOmniAVSAdmin
    function unpause() external onlyOwner {
        _unpause();
    }

    /// @dev Set the strategy parameters
    function _setStrategyParams(StrategyParams[] calldata params) internal {
        delete _strategyParams;

        for (uint256 i = 0; i < params.length; i++) {
            // TODO: add zero addr and duplicate strat tests
            require(address(params[i].strategy) != address(0), "OmniAVS: zero strategy");

            // ensure no duplicates
            for (uint256 j = i + 1; j < params.length; j++) {
                require(address(params[i].strategy) != address(params[j].strategy), "OmniAVS: duplicate strategy");
            }

            _strategyParams.push(params[i]);
        }
    }
}
