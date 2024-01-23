// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { CommonTest } from "test/common/CommonTest.sol";
import { XChain } from "src/libraries/XChain.sol";
import { Vm } from "forge-std/Vm.sol";

/**
 * @title OmniPortal_exec_Test
 * @dev Test of OmniPortal._exec, an internal function made public for testing
 */
contract OmniPortal_exec_Test is CommonTest {
    /// @dev Test that exec of a valid XMsg succeeds, and emits the correct XReceipt
    function test_exec_xmsg_succeeds() public {
        XChain.Msg memory xmsg = _inbound_increment();

        uint256 count = counter.count();
        uint64 offset = portal.inXStreamOffset(xmsg.sourceChainId);

        vm.prank(xrelayer);
        vm.expectCall(xmsg.to, xmsg.data);
        vm.recordLogs();
        portal.exec(xmsg);

        assertEq(counter.count(), count + 1);
        assertEq(portal.inXStreamOffset(xmsg.sourceChainId), offset + 1);

        Vm.Log[] memory logs = vm.getRecordedLogs();

        _assertReceiptEmitted(
            logs[0],
            xmsg.sourceChainId,
            offset,
            xrelayer,
            true // success
        );
    }

    /// @dev Test that exec of an XMsg that reverts succeeds, and emits the correct XReceipt
    function test_exec_xmsgRevert_succeeds() public {
        XChain.Msg memory xmsg = _inbound_revert();

        uint256 count = counter.count();
        uint64 offset = portal.inXStreamOffset(xmsg.sourceChainId);

        vm.prank(xrelayer);
        vm.expectCall(xmsg.to, xmsg.data);
        vm.recordLogs();
        portal.exec(xmsg);

        assertEq(counter.count(), count);
        assertEq(portal.inXStreamOffset(xmsg.sourceChainId), offset + 1);

        Vm.Log[] memory logs = vm.getRecordedLogs();

        _assertReceiptEmitted(
            logs[0],
            xmsg.sourceChainId,
            offset,
            xrelayer,
            false // failure
        );
    }

    /// @dev Test that exec of an XMsg with the wrong destChainId reverts
    function test_exec_wrongChainId_reverts() public {
        XChain.Msg memory xmsg = _inbound_increment();

        xmsg.destChainId = xmsg.destChainId + 1; // intentionally wrong chainId

        vm.expectRevert("OmniPortal: wrong destChainId");
        portal.exec(xmsg);
    }

    /// @dev Test that exec of an XMsg ahead of the current offset reverts
    function test_exec_aheadOffset_reverts() public {
        XChain.Msg memory xmsg = _inbound_increment();

        xmsg.streamOffset = xmsg.streamOffset + 1; // intentionally ahead of offset

        vm.expectRevert("OmniPortal: wrong streamOffset");
        portal.exec(xmsg);
    }

    /// @dev Test that exec of an XMsg behind the current offset reverts
    function test_exec_behindOffset_reverts() public {
        XChain.Msg memory xmsg = _inbound_increment();

        portal.exec(xmsg); // execute, to increment offset

        vm.expectRevert("OmniPortal: wrong streamOffset");
        portal.exec(xmsg);
    }
}
