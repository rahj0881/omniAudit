// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { OmniPortal } from "src/protocol/OmniPortal.sol";
import { XTypes } from "src/libraries/XTypes.sol";
import { Validators } from "src/libraries/Validators.sol";

/**
 * @title PortalHarness
 * @dev A test contract that exposes OmniPortal internal functions, and allows state manipulation.
 */
contract PortalHarness is OmniPortal {
    function exec(XTypes.Msg calldata xmsg) external {
        _exec(xmsg);
    }
}
