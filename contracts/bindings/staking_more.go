package bindings

import (
	_ "embed"
)

const (
	StakingDeployedBytecode = "0x6080604052600436106100345760003560e01c80635c19a95c14610039578063a5a470ad1461004e578063e1e158a514610061575b600080fd5b61004c610047366004610272565b610090565b005b61004c61005c3660046102a2565b610180565b34801561006d57600080fd5b5061007e68056bc75e2d6310000081565b60405190815260200160405180910390f35b600034116100e55760405162461bcd60e51b815260206004820152601d60248201527f5374616b696e673a20696e73756666696369656e74206465706f73697400000060448201526064015b60405180910390fd5b336001600160a01b0382161461013d5760405162461bcd60e51b815260206004820152601d60248201527f5374616b696e673a206f6e6c792073656c662064656c65676174696f6e00000060448201526064016100dc565b6040513481526001600160a01b0382169033907f510b11bb3f3c799b11307c01ab7db0d335683ef5b2da98f7697de744f465eacc9060200160405180910390a350565b602181146101d05760405162461bcd60e51b815260206004820152601e60248201527f5374616b696e673a20696e76616c6964207075626b6579206c656e677468000060448201526064016100dc565b68056bc75e2d631000003410156102295760405162461bcd60e51b815260206004820152601d60248201527f5374616b696e673a20696e73756666696369656e74206465706f73697400000060448201526064016100dc565b336001600160a01b03167fc7abef7b73f049da6a9bc2349ba5066a39e316eabc9f671b6f9406aa9490a45383833460405161026693929190610314565b60405180910390a25050565b60006020828403121561028457600080fd5b81356001600160a01b038116811461029b57600080fd5b9392505050565b600080602083850312156102b557600080fd5b823567ffffffffffffffff808211156102cd57600080fd5b818501915085601f8301126102e157600080fd5b8135818111156102f057600080fd5b86602082850101111561030257600080fd5b60209290920196919550909350505050565b604081528260408201528284606083013760006060848301015260006060601f19601f860116830101905082602083015294935050505056fea26469706673582212209e42a613e2b43c0c3cace889e784f24b0709198a43fcf64b7e1ce2169241b5d464736f6c63430008180033"
)

//go:embed staking_storage_layout.json
var stakingStorageLayoutJSON []byte

var StakingStorageLayout = mustGetStorageLayout(stakingStorageLayoutJSON)
