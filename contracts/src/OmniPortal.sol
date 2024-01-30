// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { IOmniPortal } from "./interfaces/IOmniPortal.sol";
import { XBlockMerkleProof } from "./libraries/XBlockMerkleProof.sol";
import { XTypes } from "./libraries/XTypes.sol";

contract OmniPortal is IOmniPortal {
    /// @inheritdoc IOmniPortal
    uint64 public constant XMSG_DEFAULT_GAS_LIMIT = 200_000;

    /// @inheritdoc IOmniPortal
    uint64 public constant XMSG_MAX_GAS_LIMIT = 5_000_000;

    /// @inheritdoc IOmniPortal
    uint64 public constant XMSG_MIN_GAS_LIMIT = 21_000;

    /// @inheritdoc IOmniPortal
    uint64 public immutable chainId;

    /// @inheritdoc IOmniPortal
    mapping(uint64 => uint64) public outXStreamOffset;

    /// @inheritdoc IOmniPortal
    mapping(uint64 => uint64) public inXStreamOffset;

    /// @inheritdoc IOmniPortal
    mapping(uint64 => uint64) public inXStreamBlockHeight;

    /// @dev The current XMsg being executed, exposed via xmsg() getter
    ///      Private state + public getter preferred over public state with default getter,
    ///      so that we can use the XMsg struct type in the interface.
    XTypes.Msg private _currentXmsg;

    constructor() {
        chainId = uint64(block.chainid);
    }

    /// @inheritdoc IOmniPortal
    function xmsg() external view returns (XTypes.Msg memory) {
        return _currentXmsg;
    }

    /// @inheritdoc IOmniPortal
    function isXCall() external view returns (bool) {
        return _currentXmsg.sourceChainId != 0;
    }

    /// @inheritdoc IOmniPortal
    function xcall(uint64 destChainId, address to, bytes calldata data) external payable {
        _xcall(destChainId, msg.sender, to, data, XMSG_DEFAULT_GAS_LIMIT);
    }

    /// @inheritdoc IOmniPortal
    function xcall(uint64 destChainId, address to, bytes calldata data, uint64 gasLimit) external payable {
        _xcall(destChainId, msg.sender, to, data, gasLimit);
    }

    /// @inheritdoc IOmniPortal
    function xsubmit(XTypes.Submission calldata xsub) external {
        // TODO: verify a quorum of validators have signed off on the attestation root.

        require(
            XBlockMerkleProof.verify(xsub.attestationRoot, xsub.blockHeader, xsub.msgs, xsub.proof, xsub.proofFlags),
            "OmniPortal: invalid proof"
        );

        inXStreamBlockHeight[xsub.blockHeader.sourceChainId] = xsub.blockHeader.blockHeight;

        for (uint256 i = 0; i < xsub.msgs.length; i++) {
            _exec(xsub.msgs[i]);
        }
    }

    /// @dev Emit an XMsg event, increment dest chain outXStreamOffset
    function _xcall(uint64 destChainId, address sender, address to, bytes calldata data, uint64 gasLimit) private {
        require(gasLimit <= XMSG_MAX_GAS_LIMIT, "OmniPortal: gasLimit too high");
        require(gasLimit >= XMSG_MIN_GAS_LIMIT, "OmniPortal: gasLimit too low");
        require(destChainId != chainId, "OmniPortal: no same-chain xcall");

        outXStreamOffset[destChainId] += 1;

        emit XMsg(destChainId, outXStreamOffset[destChainId], sender, to, data, gasLimit);
    }

    /// @dev Verify an XMsg is next in its XStream, execute it, increment inXStreamOffset, emit an XReceipt
    function _exec(XTypes.Msg calldata _xmsg) internal {
        require(_xmsg.destChainId == chainId, "OmniPortal: wrong destChainId");
        require(_xmsg.streamOffset == inXStreamOffset[_xmsg.sourceChainId] + 1, "OmniPortal: wrong streamOffset");

        // set xmsg to the one we're executing
        _currentXmsg = _xmsg;

        // increment offset before executing xcall, to avoid reentrancy loop
        inXStreamOffset[_xmsg.sourceChainId] += 1;

        // we enforce a maximum on xcall, but we trim to max here just in case
        uint256 gasLimit = _xmsg.gasLimit > XMSG_MAX_GAS_LIMIT ? XMSG_MAX_GAS_LIMIT : _xmsg.gasLimit;

        // execute xmsg, tracking gas used
        uint256 gasUsed = gasleft();
        (bool success,) = _xmsg.to.call{ gas: gasLimit }(_xmsg.data);
        gasUsed = gasUsed - gasleft();

        // reset xmsg to zero
        _currentXmsg = XTypes.zeroMsg();

        emit XReceipt(_xmsg.sourceChainId, _xmsg.streamOffset, gasUsed, msg.sender, success);
    }
}
