// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { IFeeOracle } from "../interfaces/IFeeOracle.sol";
import { IFeeOracleV1 } from "../interfaces/IFeeOracleV1.sol";

/**
 * @title FeeOracleV1
 * @notice A simple fee oracle with a fixed fee, controlled by an admin account
 *         Used by OmniPortal to calculate xmsg fees
 */
contract FeeOracleV1 is IFeeOracle, IFeeOracleV1, OwnableUpgradeable {
    /**
     * @notice Base gas limit for each xmsg.
     */
    uint256 public baseGasLimit;

    /**
     * @notice Base protocol fee for each xmsg.
     */
    uint256 public protocolFee;

    /**
     * @notice Gas price per destination chain, in wei, of the chains native token.
     */
    mapping(uint64 => uint256) public gasPriceOn;

    /**
     * @notice Conversion rate from a destination chain's native token to this chain's native token.
     *         Rate is numerator over CONVERSION_RATE_DENOM.
     */
    mapping(uint64 => uint256) public toNativeRate;

    /**
     * @notice Denominator for conversion rate calculations.
     */
    uint256 public constant CONVERSION_RATE_DENOM = 1e6;

    constructor() {
        _disableInitializers();
    }

    function initialize(address owner_, uint256 baseGasLimit_, uint256 protocolFee_, ChainFeeParams[] calldata params)
        public
        initializer
    {
        _transferOwnership(owner_);
        _setBaseGasLimit(baseGasLimit_);
        _setProtocolFee(protocolFee_);
        _bulkSetFeeParams(params);
    }

    /// @inheritdoc IFeeOracle
    function feeFor(uint64 destChainId, bytes calldata, uint64 gasLimit) external view returns (uint256) {
        uint256 gasPrice = gasPriceOn[destChainId] * toNativeRate[destChainId] / CONVERSION_RATE_DENOM;
        return protocolFee + (baseGasLimit * gasPrice) + (gasLimit * gasPrice);
    }

    /**
     * @notice Set the fee parameters for a list of destination chains.
     */
    function bulkSetFeeParams(ChainFeeParams[] calldata params) external onlyOwner {
        _bulkSetFeeParams(params);
    }

    /**
     * @notice Set the gas price for a destination chain.
     */
    function setGasPrice(uint64 chainId, uint256 gasPrice) external onlyOwner {
        _setGasPrice(chainId, gasPrice);
    }

    /**
     * @notice Set the to native conversion rate for a destination chain.
     */
    function setToNativeRate(uint64 chainId, uint256 toNativeRate) external onlyOwner {
        _setToNativeRate(chainId, toNativeRate);
    }

    /**
     * @notice Set the base gas limit for each xmsg.
     */
    function setBaseGasLimit(uint256 baseGasLimit_) external onlyOwner {
        _setBaseGasLimit(baseGasLimit_);
    }

    /**
     * @notice Set the base protocol fee for each xmsg.
     */
    function setProtocolFee(uint256 protocolFee_) external onlyOwner {
        _setProtocolFee(protocolFee_);
    }

    /**
     * @notice Set the fee parameters for a list of destination chains.
     */
    function _bulkSetFeeParams(ChainFeeParams[] calldata params) internal {
        for (uint256 i = 0; i < params.length; i++) {
            ChainFeeParams memory p = params[i];
            _setGasPrice(p.chainId, p.gasPrice);
            _setToNativeRate(p.chainId, p.toNativeRate);
        }
    }

    /**
     * @notice Set the gas price for a destination chain.
     */
    function _setGasPrice(uint64 chainId, uint256 gasPrice) internal {
        require(gasPrice > 0, "FeeOracleV1: no zero gas price");
        require(chainId != 0, "FeeOracleV1: no zero chain id");

        gasPriceOn[chainId] = gasPrice;
        emit GasPriceSet(chainId, gasPrice);
    }

    /**
     * @notice Set the to-native conversion rate for a destination chain.
     */
    function _setToNativeRate(uint64 chainId, uint256 rate) internal {
        require(rate > 0, "FeeOracleV1: no zero rate");
        require(chainId != 0, "FeeOracleV1: no zero chain id");

        toNativeRate[chainId] = rate;
        emit ToNativeRateSet(chainId, rate);
    }

    /**
     * @notice Set the base gas limit for each xmsg.
     */
    function _setBaseGasLimit(uint256 baseGasLimit_) internal {
        baseGasLimit = baseGasLimit_;
        emit BaseGasLimitSet(baseGasLimit);
    }

    /**
     * @notice Set the base protocol fee for each xmsg.
     */
    function _setProtocolFee(uint256 protocolFee_) internal {
        protocolFee = protocolFee_;
        emit ProtocolFeeSet(protocolFee);
    }
}
