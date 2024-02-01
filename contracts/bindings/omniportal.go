// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// XTypesBlockHeader is an auto generated low-level Go binding around an user-defined struct.
type XTypesBlockHeader struct {
	SourceChainId uint64
	BlockHeight   uint64
	BlockHash     [32]byte
}

// XTypesMsg is an auto generated low-level Go binding around an user-defined struct.
type XTypesMsg struct {
	SourceChainId uint64
	DestChainId   uint64
	StreamOffset  uint64
	Sender        common.Address
	To            common.Address
	Data          []byte
	GasLimit      uint64
}

// XTypesSigTuple is an auto generated low-level Go binding around an user-defined struct.
type XTypesSigTuple struct {
	ValidatorPubKey []byte
	Signature       []byte
}

// XTypesSubmission is an auto generated low-level Go binding around an user-defined struct.
type XTypesSubmission struct {
	AttestationRoot [32]byte
	BlockHeader     XTypesBlockHeader
	Msgs            []XTypesMsg
	Proof           [][32]byte
	ProofFlags      []bool
	Signatures      []XTypesSigTuple
}

// OmniPortalMetaData contains all meta data concerning the OmniPortal contract.
var OmniPortalMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeOracle_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"XMSG_DEFAULT_GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"XMSG_MAX_GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"XMSG_MIN_GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"chainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collectFees\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeFor\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeFor\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeOracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inXStreamBlockHeight\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inXStreamOffset\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isXCall\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"outXStreamOffset\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeOracle\",\"inputs\":[{\"name\":\"feeOracle_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"xcall\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"xcall\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"xmsg\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXTypes.Msg\",\"components\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"streamOffset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xsubmit\",\"inputs\":[{\"name\":\"xsub\",\"type\":\"tuple\",\"internalType\":\"structXTypes.Submission\",\"components\":[{\"name\":\"attestationRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockHeader\",\"type\":\"tuple\",\"internalType\":\"structXTypes.BlockHeader\",\"components\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"msgs\",\"type\":\"tuple[]\",\"internalType\":\"structXTypes.Msg[]\",\"components\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"streamOffset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"proof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"proofFlags\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"},{\"name\":\"signatures\",\"type\":\"tuple[]\",\"internalType\":\"structXTypes.SigTuple[]\",\"components\":[{\"name\":\"validatorPubKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FeeOracleChanged\",\"inputs\":[{\"name\":\"oldFeeOracle\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newFeeOracle\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeesCollected\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XMsg\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"streamOffset\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XReceipt\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"streamOffset\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"gasUsed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"relayer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"MerkleProofInvalidMultiproof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a06040523480156200001157600080fd5b50604051620021bf380380620021bf8339810160408190526200003491620001a8565b816001600160a01b0381166200006557604051631e4fbdf760e01b8152600060048201526024015b60405180910390fd5b620000708162000091565b506001600160401b0346166080526200008981620000e1565b5050620001e0565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b038116620001395760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e69506f7274616c3a206e6f207a65726f206665654f7261636c6500000060448201526064016200005c565b600180546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f2819896846ed9ab612eb19218fd845f8328f084c8706b9ec2c47eabd479037a290600090a35050565b80516001600160a01b0381168114620001a357600080fd5b919050565b60008060408385031215620001bc57600080fd5b620001c7836200018b565b9150620001d7602084016200018b565b90509250929050565b608051611fb56200020a600039600081816102da01528181610a0f0152610f1c0152611fb56000f3fe60806040526004361061011f5760003560e01c806390ab417c116100a0578063a480ca7911610064578063a480ca7914610360578063a8a9896214610380578063b58e964f146103a0578063f2fde38b146103d6578063fa590d14146103f657600080fd5b806390ab417c1461027a5780639a8a0592146102c85780639c346d99146102fc5780639dad9aae14610332578063a2cc111b1461034957600080fd5b806355e2448e116100e757806355e2448e146101ea57806370e8b56a14610214578063715018a6146102275780638da5cb5b1461023c5780638dd9523c1461025a57600080fd5b806306f9f174146101245780632f32700e146101465780634115ab7914610171578063500b19e71461019f57806350e646dd146101d7575b600080fd5b34801561013057600080fd5b5061014461013f366004611538565b61040c565b005b34801561015257600080fd5b5061015b610556565b6040516101689190611573565b60405180910390f35b34801561017d57600080fd5b5061019161018c36600461169c565b610695565b604051908152602001610168565b3480156101ab57600080fd5b506001546101bf906001600160a01b031681565b6040516001600160a01b039091168152602001610168565b6101446101e5366004611710565b61071a565b3480156101f657600080fd5b506005546040516001600160401b0390911615158152602001610168565b610144610222366004611774565b610797565b34801561023357600080fd5b5061014461080f565b34801561024857600080fd5b506000546001600160a01b03166101bf565b34801561026657600080fd5b506101916102753660046117f0565b610823565b34801561028657600080fd5b506102b0610295366004611857565b6002602052600090815260409020546001600160401b031681565b6040516001600160401b039091168152602001610168565b3480156102d457600080fd5b506102b07f000000000000000000000000000000000000000000000000000000000000000081565b34801561030857600080fd5b506102b0610317366004611857565b6004602052600090815260409020546001600160401b031681565b34801561033e57600080fd5b506102b062030d4081565b34801561035557600080fd5b506102b0624c4b4081565b34801561036c57600080fd5b5061014461037b366004611874565b6108a4565b34801561038c57600080fd5b5061014461039b366004611874565b61092c565b3480156103ac57600080fd5b506102b06103bb366004611857565b6003602052600090815260409020546001600160401b031681565b3480156103e257600080fd5b506101446103f1366004611874565b610940565b34801561040257600080fd5b506102b061520881565b6104418135602083016104226080850185611891565b61042f60a0870187611891565b61043c60c0890189611891565b61097b565b6104925760405162461bcd60e51b815260206004820152601960248201527f4f6d6e69506f7274616c3a20696e76616c69642070726f6f660000000000000060448201526064015b60405180910390fd5b6104a26060820160408301611857565b600460006104b66040850160208601611857565b6001600160401b03166001600160401b0316815260200190815260200160002060006101000a8154816001600160401b0302191690836001600160401b0316021790555060005b61050a6080830183611891565b90508110156105525761054a6105236080840184611891565b83818110610533576105336118da565b905060200281019061054591906118f0565b610a05565b6001016104fd565b5050565b6040805160e08101825260008082526020820181905291810182905260608082018390526080820183905260a082015260c08101919091526040805160e081018252600580546001600160401b038082168452600160401b820481166020850152600160801b90910416928201929092526006546001600160a01b0390811660608301526007541660808201526008805491929160a0840191906105f990611910565b80601f016020809104026020016040519081016040528092919081815260200182805461062590611910565b80156106725780601f1061064757610100808354040283529160200191610672565b820191906000526020600020905b81548152906001019060200180831161065557829003601f168201915b5050509183525050600491909101546001600160401b0316602090910152919050565b600154604051632376548f60e21b81526000916001600160a01b031690638dd9523c906106cf9087908790879062030d4090600401611973565b602060405180830381865afa1580156106ec573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061071091906119ab565b90505b9392505050565b83828262030d4061072d84848484610823565b34101561077c5760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e69506f7274616c3a20696e73756666696369656e7420666565000000006044820152606401610489565b61078d883389898962030d40610e63565b5050505050505050565b848383836107a784848484610823565b3410156107f65760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e69506f7274616c3a20696e73756666696369656e7420666565000000006044820152606401610489565b61080489338a8a8a8a610e63565b505050505050505050565b610817611048565b6108216000611075565b565b600154604051632376548f60e21b81526000916001600160a01b031690638dd9523c9061085a908890889088908890600401611973565b602060405180830381865afa158015610877573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061089b91906119ab565b95945050505050565b6108ac611048565b60405147906001600160a01b0383169082156108fc029083906000818181858888f193505050501580156108e4573d6000803e3d6000fd5b50816001600160a01b03167f9dc46f23cfb5ddcad0ae7ea2be38d47fec07bb9382ec7e564efc69e036dd66ce8260405161092091815260200190565b60405180910390a25050565b610934611048565b61093d816110c5565b50565b610948611048565b6001600160a01b03811661097257604051631e4fbdf760e01b815260006004820152602401610489565b61093d81611075565b60006109f8858580806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250506040805160208089028281018201909352888252909350889250879182918501908490808284376000920191909152508e92506109f391508d90508c8c61116d565b611280565b9998505050505050505050565b6001600160401b037f000000000000000000000000000000000000000000000000000000000000000016610a3f6040830160208401611857565b6001600160401b031614610a955760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e69506f7274616c3a2077726f6e672064657374436861696e49640000006044820152606401610489565b60036000610aa66020840184611857565b6001600160401b039081168252602082019290925260400160002054610ace911660016119da565b6001600160401b0316610ae76060830160408401611857565b6001600160401b031614610b3d5760405162461bcd60e51b815260206004820152601e60248201527f4f6d6e69506f7274616c3a2077726f6e672073747265616d4f666673657400006044820152606401610489565b806005610b4a8282611ba8565b506001905060036000610b606020850185611857565b6001600160401b0390811682526020820192909252604001600090812080549092610b8d918591166119da565b92506101000a8154816001600160401b0302191690836001600160401b031602179055506000624c4b406001600160401b03168260c0016020810190610bd39190611857565b6001600160401b031611610bf657610bf160e0830160c08401611857565b610bfb565b624c4b405b6001600160401b0316905060005a90506000610c1d60a0850160808601611874565b6001600160a01b031683610c3460a0870187611a3b565b604051610c42929190611cce565b60006040518083038160008787f1925050503d8060008114610c80576040519150601f19603f3d011682016040523d82523d6000602084013e610c85565b606091505b505090505a610c949083611cde565b9150610d1b6040805160e08101825260008082526020820181905291810182905260608082018390526080820183905260a082015260c0810191909152506040805160e081018252600080825260208083018290528284018290526060830182905260808301829052835190810190935280835260a082019290925260c081019190915290565b805160058054602084015160408501516001600160401b03908116600160801b0267ffffffffffffffff60801b19928216600160401b026fffffffffffffffffffffffffffffffff19909416919095161791909117169190911781556060820151600680546001600160a01b039283166001600160a01b03199182161790915560808401516007805491909316911617905560a0820151600890610dbf9082611cf1565b5060c091909101516004909101805467ffffffffffffffff19166001600160401b03909216919091179055610dfa6060850160408601611857565b6001600160401b0316610e106020860186611857565b604080518581523360208201528415158183015290516001600160401b0392909216917f34515b4105a7bb34f3af3cd490137ab292bb2ff14efb800df5c7d59e28944f259181900360600190a350505050565b624c4b406001600160401b0382161115610ebf5760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e69506f7274616c3a206761734c696d697420746f6f20686967680000006044820152606401610489565b6152086001600160401b0382161015610f1a5760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e69506f7274616c3a206761734c696d697420746f6f206c6f77000000006044820152606401610489565b7f00000000000000000000000000000000000000000000000000000000000000006001600160401b0316866001600160401b031603610f9b5760405162461bcd60e51b815260206004820152601f60248201527f4f6d6e69506f7274616c3a206e6f2073616d652d636861696e207863616c6c006044820152606401610489565b6001600160401b0380871660009081526002602052604081208054600193919291610fc8918591166119da565b82546101009290920a6001600160401b038181021990931691831602179091558781166000818152600260205260409081902054905192169250907fac3afbbff5be7c4af1610721cf4793840bd167251fd6f184ee708f752a731283906110389089908990899089908990611db0565b60405180910390a3505050505050565b6000546001600160a01b031633146108215760405163118cdaa760e01b8152336004820152602401610489565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811661111b5760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e69506f7274616c3a206e6f207a65726f206665654f7261636c650000006044820152606401610489565b600180546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f2819896846ed9ab612eb19218fd845f8328f084c8706b9ec2c47eabd479037a290600090a35050565b6060600061117c836001611df8565b6001600160401b0381111561119357611193611a81565b6040519080825280602002602001820160405280156111bc578160200160208202803683370190505b5090506111e7856040516020016111d39190611e0b565b604051602081830303815290604052611298565b816000815181106111fa576111fa6118da565b60200260200101818152505060005b8381101561127757611248858583818110611226576112266118da565b905060200281019061123891906118f0565b6040516020016111d39190611e92565b82611254836001611df8565b81518110611264576112646118da565b6020908102919091010152600101611209565b50949350505050565b60008261128e8686856112d1565b1495945050505050565b600081805190602001206040516020016112b491815260200190565b604051602081830303815290604052805190602001209050919050565b80518351835160009291906112e7816001611df8565b6112f18385611df8565b1461130f57604051631a8a024960e11b815260040160405180910390fd5b6000816001600160401b0381111561132957611329611a81565b604051908082528060200260200182016040528015611352578160200160208202803683370190505b5090506000806000805b8581101561148657600088851061139757858461137881611f66565b95508151811061138a5761138a6118da565b60200260200101516113bd565b8a856113a281611f66565b9650815181106113b4576113b46118da565b60200260200101515b905060008c83815181106113d3576113d36118da565b6020026020010151611409578d846113ea81611f66565b9550815181106113fc576113fc6118da565b6020026020010151611453565b89861061142d57868561141b81611f66565b9650815181106113fc576113fc6118da565b8b8661143881611f66565b97508151811061144a5761144a6118da565b60200260200101515b905061145f8282611504565b878481518110611471576114716118da565b6020908102919091010152505060010161135c565b5084156114d8578581146114ad57604051631a8a024960e11b815260040160405180910390fd5b8360018603815181106114c2576114c26118da565b6020026020010151975050505050505050610713565b86156114f157886000815181106114c2576114c26118da565b8a6000815181106114c2576114c26118da565b600081831061152057600082815260208490526040902061152f565b60008381526020839052604090205b90505b92915050565b60006020828403121561154a57600080fd5b81356001600160401b0381111561156057600080fd5b8201610100818503121561071357600080fd5b600060208083526101006001600160401b038086511683860152808387015116604086015280604087015116606086015250606085015160018060a01b0380821660808701528060808801511660a0870152505060a085015160e060c08601528051808387015260005b818110156115fa57828101850151878201850186015284016115dd565b506000868201840185015260c08701516001600160401b03811660e08801529150601f01601f19169490940101019392505050565b6001600160401b038116811461093d57600080fd5b803561164f8161162f565b919050565b60008083601f84011261166657600080fd5b5081356001600160401b0381111561167d57600080fd5b60208301915083602082850101111561169557600080fd5b9250929050565b6000806000604084860312156116b157600080fd5b83356116bc8161162f565b925060208401356001600160401b038111156116d757600080fd5b6116e386828701611654565b9497909650939450505050565b6001600160a01b038116811461093d57600080fd5b803561164f816116f0565b6000806000806060858703121561172657600080fd5b84356117318161162f565b93506020850135611741816116f0565b925060408501356001600160401b0381111561175c57600080fd5b61176887828801611654565b95989497509550505050565b60008060008060006080868803121561178c57600080fd5b85356117978161162f565b945060208601356117a7816116f0565b935060408601356001600160401b038111156117c257600080fd5b6117ce88828901611654565b90945092505060608601356117e28161162f565b809150509295509295909350565b6000806000806060858703121561180657600080fd5b84356118118161162f565b935060208501356001600160401b0381111561182c57600080fd5b61183887828801611654565b909450925050604085013561184c8161162f565b939692955090935050565b60006020828403121561186957600080fd5b81356107138161162f565b60006020828403121561188657600080fd5b8135610713816116f0565b6000808335601e198436030181126118a857600080fd5b8301803591506001600160401b038211156118c257600080fd5b6020019150600581901b360382131561169557600080fd5b634e487b7160e01b600052603260045260246000fd5b6000823560de1983360301811261190657600080fd5b9190910192915050565b600181811c9082168061192457607f821691505b60208210810361194457634e487b7160e01b600052602260045260246000fd5b50919050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60006001600160401b0380871683526060602084015261199760608401868861194a565b915080841660408401525095945050505050565b6000602082840312156119bd57600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b6001600160401b038181168382160190808211156119fa576119fa6119c4565b5092915050565b600081356115328161162f565b60008135611532816116f0565b80546001600160a01b0319166001600160a01b0392909216919091179055565b6000808335601e19843603018112611a5257600080fd5b8301803591506001600160401b03821115611a6c57600080fd5b60200191503681900382131561169557600080fd5b634e487b7160e01b600052604160045260246000fd5b601f821115611ae3576000816000526020600020601f850160051c81016020861015611ac05750805b601f850160051c820191505b81811015611adf57828155600101611acc565b5050505b505050565b6001600160401b03831115611aff57611aff611a81565b611b1383611b0d8354611910565b83611a97565b6000601f841160018114611b475760008515611b2f5750838201355b600019600387901b1c1916600186901b178355611ba1565b600083815260209020601f19861690835b82811015611b785786850135825560209485019460019092019101611b58565b5086821015611b955760001960f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b8135611bb38161162f565b815467ffffffffffffffff19166001600160401b038216178255506020820135611bdc8161162f565b81546fffffffffffffffff0000000000000000604092831b166fffffffffffffffff00000000000000001982168117845591840135611c1a8161162f565b77ffffffffffffffffffffffffffffffff0000000000000000199190911690911760809190911b67ffffffffffffffff60801b16178155611c69611c6060608401611a0e565b60018301611a1b565b611c81611c7860808401611a0e565b60028301611a1b565b611c8e60a0830183611a3b565b611c9c818360038601611ae8565b5050610552611cad60c08401611a01565b600483016001600160401b0382166001600160401b03198254161781555050565b8183823760009101908152919050565b81810381811115611532576115326119c4565b81516001600160401b03811115611d0a57611d0a611a81565b611d1e81611d188454611910565b84611a97565b602080601f831160018114611d535760008415611d3b5750858301515b600019600386901b1c1916600185901b178555611adf565b600085815260208120601f198616915b82811015611d8257888601518255948401946001909101908401611d63565b5085821015611da05787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6001600160a01b03868116825285166020820152608060408201819052600090611ddd908301858761194a565b90506001600160401b03831660608301529695505050505050565b80820180821115611532576115326119c4565b606081018235611e1a8161162f565b6001600160401b039081168352602084013590611e368261162f565b166020830152604092830135929091019190915290565b6000808335601e19843603018112611e6457600080fd5b83016020810192503590506001600160401b03811115611e8357600080fd5b80360382131561169557600080fd5b6020815260008235611ea38161162f565b6001600160401b03808216602085015260208501359150611ec38261162f565b808216604085015260408501359150611edb8261162f565b16606083810191909152830135611ef1816116f0565b6001600160a01b038116608084015250611f0d60808401611705565b6001600160a01b03811660a084015250611f2a60a0840184611e4d565b60e060c0850152611f406101008501828461194a565b915050611f4f60c08501611644565b6001600160401b03811660e0850152509392505050565b600060018201611f7857611f786119c4565b506001019056fea2646970667358221220daf9832f1d31b36edcef7233523940c11a11f3074d3ab245423cac8c5a60a88464736f6c63430008170033",
}

// OmniPortalABI is the input ABI used to generate the binding from.
// Deprecated: Use OmniPortalMetaData.ABI instead.
var OmniPortalABI = OmniPortalMetaData.ABI

// OmniPortalBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OmniPortalMetaData.Bin instead.
var OmniPortalBin = OmniPortalMetaData.Bin

// DeployOmniPortal deploys a new Ethereum contract, binding an instance of OmniPortal to it.
func DeployOmniPortal(auth *bind.TransactOpts, backend bind.ContractBackend, owner_ common.Address, feeOracle_ common.Address) (common.Address, *types.Transaction, *OmniPortal, error) {
	parsed, err := OmniPortalMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OmniPortalBin), backend, owner_, feeOracle_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OmniPortal{OmniPortalCaller: OmniPortalCaller{contract: contract}, OmniPortalTransactor: OmniPortalTransactor{contract: contract}, OmniPortalFilterer: OmniPortalFilterer{contract: contract}}, nil
}

// OmniPortal is an auto generated Go binding around an Ethereum contract.
type OmniPortal struct {
	OmniPortalCaller     // Read-only binding to the contract
	OmniPortalTransactor // Write-only binding to the contract
	OmniPortalFilterer   // Log filterer for contract events
}

// OmniPortalCaller is an auto generated read-only Go binding around an Ethereum contract.
type OmniPortalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniPortalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OmniPortalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniPortalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OmniPortalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniPortalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OmniPortalSession struct {
	Contract     *OmniPortal       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OmniPortalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OmniPortalCallerSession struct {
	Contract *OmniPortalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// OmniPortalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OmniPortalTransactorSession struct {
	Contract     *OmniPortalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OmniPortalRaw is an auto generated low-level Go binding around an Ethereum contract.
type OmniPortalRaw struct {
	Contract *OmniPortal // Generic contract binding to access the raw methods on
}

// OmniPortalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OmniPortalCallerRaw struct {
	Contract *OmniPortalCaller // Generic read-only contract binding to access the raw methods on
}

// OmniPortalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OmniPortalTransactorRaw struct {
	Contract *OmniPortalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOmniPortal creates a new instance of OmniPortal, bound to a specific deployed contract.
func NewOmniPortal(address common.Address, backend bind.ContractBackend) (*OmniPortal, error) {
	contract, err := bindOmniPortal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OmniPortal{OmniPortalCaller: OmniPortalCaller{contract: contract}, OmniPortalTransactor: OmniPortalTransactor{contract: contract}, OmniPortalFilterer: OmniPortalFilterer{contract: contract}}, nil
}

// NewOmniPortalCaller creates a new read-only instance of OmniPortal, bound to a specific deployed contract.
func NewOmniPortalCaller(address common.Address, caller bind.ContractCaller) (*OmniPortalCaller, error) {
	contract, err := bindOmniPortal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OmniPortalCaller{contract: contract}, nil
}

// NewOmniPortalTransactor creates a new write-only instance of OmniPortal, bound to a specific deployed contract.
func NewOmniPortalTransactor(address common.Address, transactor bind.ContractTransactor) (*OmniPortalTransactor, error) {
	contract, err := bindOmniPortal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OmniPortalTransactor{contract: contract}, nil
}

// NewOmniPortalFilterer creates a new log filterer instance of OmniPortal, bound to a specific deployed contract.
func NewOmniPortalFilterer(address common.Address, filterer bind.ContractFilterer) (*OmniPortalFilterer, error) {
	contract, err := bindOmniPortal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OmniPortalFilterer{contract: contract}, nil
}

// bindOmniPortal binds a generic wrapper to an already deployed contract.
func bindOmniPortal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OmniPortalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OmniPortal *OmniPortalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OmniPortal.Contract.OmniPortalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OmniPortal *OmniPortalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.Contract.OmniPortalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OmniPortal *OmniPortalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OmniPortal.Contract.OmniPortalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OmniPortal *OmniPortalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OmniPortal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OmniPortal *OmniPortalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OmniPortal *OmniPortalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OmniPortal.Contract.contract.Transact(opts, method, params...)
}

// XMSGDEFAULTGASLIMIT is a free data retrieval call binding the contract method 0x9dad9aae.
//
// Solidity: function XMSG_DEFAULT_GAS_LIMIT() view returns(uint64)
func (_OmniPortal *OmniPortalCaller) XMSGDEFAULTGASLIMIT(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "XMSG_DEFAULT_GAS_LIMIT")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// XMSGDEFAULTGASLIMIT is a free data retrieval call binding the contract method 0x9dad9aae.
//
// Solidity: function XMSG_DEFAULT_GAS_LIMIT() view returns(uint64)
func (_OmniPortal *OmniPortalSession) XMSGDEFAULTGASLIMIT() (uint64, error) {
	return _OmniPortal.Contract.XMSGDEFAULTGASLIMIT(&_OmniPortal.CallOpts)
}

// XMSGDEFAULTGASLIMIT is a free data retrieval call binding the contract method 0x9dad9aae.
//
// Solidity: function XMSG_DEFAULT_GAS_LIMIT() view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) XMSGDEFAULTGASLIMIT() (uint64, error) {
	return _OmniPortal.Contract.XMSGDEFAULTGASLIMIT(&_OmniPortal.CallOpts)
}

// XMSGMAXGASLIMIT is a free data retrieval call binding the contract method 0xa2cc111b.
//
// Solidity: function XMSG_MAX_GAS_LIMIT() view returns(uint64)
func (_OmniPortal *OmniPortalCaller) XMSGMAXGASLIMIT(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "XMSG_MAX_GAS_LIMIT")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// XMSGMAXGASLIMIT is a free data retrieval call binding the contract method 0xa2cc111b.
//
// Solidity: function XMSG_MAX_GAS_LIMIT() view returns(uint64)
func (_OmniPortal *OmniPortalSession) XMSGMAXGASLIMIT() (uint64, error) {
	return _OmniPortal.Contract.XMSGMAXGASLIMIT(&_OmniPortal.CallOpts)
}

// XMSGMAXGASLIMIT is a free data retrieval call binding the contract method 0xa2cc111b.
//
// Solidity: function XMSG_MAX_GAS_LIMIT() view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) XMSGMAXGASLIMIT() (uint64, error) {
	return _OmniPortal.Contract.XMSGMAXGASLIMIT(&_OmniPortal.CallOpts)
}

// XMSGMINGASLIMIT is a free data retrieval call binding the contract method 0xfa590d14.
//
// Solidity: function XMSG_MIN_GAS_LIMIT() view returns(uint64)
func (_OmniPortal *OmniPortalCaller) XMSGMINGASLIMIT(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "XMSG_MIN_GAS_LIMIT")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// XMSGMINGASLIMIT is a free data retrieval call binding the contract method 0xfa590d14.
//
// Solidity: function XMSG_MIN_GAS_LIMIT() view returns(uint64)
func (_OmniPortal *OmniPortalSession) XMSGMINGASLIMIT() (uint64, error) {
	return _OmniPortal.Contract.XMSGMINGASLIMIT(&_OmniPortal.CallOpts)
}

// XMSGMINGASLIMIT is a free data retrieval call binding the contract method 0xfa590d14.
//
// Solidity: function XMSG_MIN_GAS_LIMIT() view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) XMSGMINGASLIMIT() (uint64, error) {
	return _OmniPortal.Contract.XMSGMINGASLIMIT(&_OmniPortal.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_OmniPortal *OmniPortalCaller) ChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_OmniPortal *OmniPortalSession) ChainId() (uint64, error) {
	return _OmniPortal.Contract.ChainId(&_OmniPortal.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) ChainId() (uint64, error) {
	return _OmniPortal.Contract.ChainId(&_OmniPortal.CallOpts)
}

// FeeFor is a free data retrieval call binding the contract method 0x4115ab79.
//
// Solidity: function feeFor(uint64 destChainId, bytes data) view returns(uint256)
func (_OmniPortal *OmniPortalCaller) FeeFor(opts *bind.CallOpts, destChainId uint64, data []byte) (*big.Int, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "feeFor", destChainId, data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeFor is a free data retrieval call binding the contract method 0x4115ab79.
//
// Solidity: function feeFor(uint64 destChainId, bytes data) view returns(uint256)
func (_OmniPortal *OmniPortalSession) FeeFor(destChainId uint64, data []byte) (*big.Int, error) {
	return _OmniPortal.Contract.FeeFor(&_OmniPortal.CallOpts, destChainId, data)
}

// FeeFor is a free data retrieval call binding the contract method 0x4115ab79.
//
// Solidity: function feeFor(uint64 destChainId, bytes data) view returns(uint256)
func (_OmniPortal *OmniPortalCallerSession) FeeFor(destChainId uint64, data []byte) (*big.Int, error) {
	return _OmniPortal.Contract.FeeFor(&_OmniPortal.CallOpts, destChainId, data)
}

// FeeFor0 is a free data retrieval call binding the contract method 0x8dd9523c.
//
// Solidity: function feeFor(uint64 destChainId, bytes data, uint64 gasLimit) view returns(uint256)
func (_OmniPortal *OmniPortalCaller) FeeFor0(opts *bind.CallOpts, destChainId uint64, data []byte, gasLimit uint64) (*big.Int, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "feeFor0", destChainId, data, gasLimit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeFor0 is a free data retrieval call binding the contract method 0x8dd9523c.
//
// Solidity: function feeFor(uint64 destChainId, bytes data, uint64 gasLimit) view returns(uint256)
func (_OmniPortal *OmniPortalSession) FeeFor0(destChainId uint64, data []byte, gasLimit uint64) (*big.Int, error) {
	return _OmniPortal.Contract.FeeFor0(&_OmniPortal.CallOpts, destChainId, data, gasLimit)
}

// FeeFor0 is a free data retrieval call binding the contract method 0x8dd9523c.
//
// Solidity: function feeFor(uint64 destChainId, bytes data, uint64 gasLimit) view returns(uint256)
func (_OmniPortal *OmniPortalCallerSession) FeeFor0(destChainId uint64, data []byte, gasLimit uint64) (*big.Int, error) {
	return _OmniPortal.Contract.FeeFor0(&_OmniPortal.CallOpts, destChainId, data, gasLimit)
}

// FeeOracle is a free data retrieval call binding the contract method 0x500b19e7.
//
// Solidity: function feeOracle() view returns(address)
func (_OmniPortal *OmniPortalCaller) FeeOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "feeOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeOracle is a free data retrieval call binding the contract method 0x500b19e7.
//
// Solidity: function feeOracle() view returns(address)
func (_OmniPortal *OmniPortalSession) FeeOracle() (common.Address, error) {
	return _OmniPortal.Contract.FeeOracle(&_OmniPortal.CallOpts)
}

// FeeOracle is a free data retrieval call binding the contract method 0x500b19e7.
//
// Solidity: function feeOracle() view returns(address)
func (_OmniPortal *OmniPortalCallerSession) FeeOracle() (common.Address, error) {
	return _OmniPortal.Contract.FeeOracle(&_OmniPortal.CallOpts)
}

// InXStreamBlockHeight is a free data retrieval call binding the contract method 0x9c346d99.
//
// Solidity: function inXStreamBlockHeight(uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCaller) InXStreamBlockHeight(opts *bind.CallOpts, arg0 uint64) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "inXStreamBlockHeight", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// InXStreamBlockHeight is a free data retrieval call binding the contract method 0x9c346d99.
//
// Solidity: function inXStreamBlockHeight(uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalSession) InXStreamBlockHeight(arg0 uint64) (uint64, error) {
	return _OmniPortal.Contract.InXStreamBlockHeight(&_OmniPortal.CallOpts, arg0)
}

// InXStreamBlockHeight is a free data retrieval call binding the contract method 0x9c346d99.
//
// Solidity: function inXStreamBlockHeight(uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) InXStreamBlockHeight(arg0 uint64) (uint64, error) {
	return _OmniPortal.Contract.InXStreamBlockHeight(&_OmniPortal.CallOpts, arg0)
}

// InXStreamOffset is a free data retrieval call binding the contract method 0xb58e964f.
//
// Solidity: function inXStreamOffset(uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCaller) InXStreamOffset(opts *bind.CallOpts, arg0 uint64) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "inXStreamOffset", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// InXStreamOffset is a free data retrieval call binding the contract method 0xb58e964f.
//
// Solidity: function inXStreamOffset(uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalSession) InXStreamOffset(arg0 uint64) (uint64, error) {
	return _OmniPortal.Contract.InXStreamOffset(&_OmniPortal.CallOpts, arg0)
}

// InXStreamOffset is a free data retrieval call binding the contract method 0xb58e964f.
//
// Solidity: function inXStreamOffset(uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) InXStreamOffset(arg0 uint64) (uint64, error) {
	return _OmniPortal.Contract.InXStreamOffset(&_OmniPortal.CallOpts, arg0)
}

// IsXCall is a free data retrieval call binding the contract method 0x55e2448e.
//
// Solidity: function isXCall() view returns(bool)
func (_OmniPortal *OmniPortalCaller) IsXCall(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "isXCall")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsXCall is a free data retrieval call binding the contract method 0x55e2448e.
//
// Solidity: function isXCall() view returns(bool)
func (_OmniPortal *OmniPortalSession) IsXCall() (bool, error) {
	return _OmniPortal.Contract.IsXCall(&_OmniPortal.CallOpts)
}

// IsXCall is a free data retrieval call binding the contract method 0x55e2448e.
//
// Solidity: function isXCall() view returns(bool)
func (_OmniPortal *OmniPortalCallerSession) IsXCall() (bool, error) {
	return _OmniPortal.Contract.IsXCall(&_OmniPortal.CallOpts)
}

// OutXStreamOffset is a free data retrieval call binding the contract method 0x90ab417c.
//
// Solidity: function outXStreamOffset(uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCaller) OutXStreamOffset(opts *bind.CallOpts, arg0 uint64) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "outXStreamOffset", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// OutXStreamOffset is a free data retrieval call binding the contract method 0x90ab417c.
//
// Solidity: function outXStreamOffset(uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalSession) OutXStreamOffset(arg0 uint64) (uint64, error) {
	return _OmniPortal.Contract.OutXStreamOffset(&_OmniPortal.CallOpts, arg0)
}

// OutXStreamOffset is a free data retrieval call binding the contract method 0x90ab417c.
//
// Solidity: function outXStreamOffset(uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) OutXStreamOffset(arg0 uint64) (uint64, error) {
	return _OmniPortal.Contract.OutXStreamOffset(&_OmniPortal.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniPortal *OmniPortalCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniPortal *OmniPortalSession) Owner() (common.Address, error) {
	return _OmniPortal.Contract.Owner(&_OmniPortal.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniPortal *OmniPortalCallerSession) Owner() (common.Address, error) {
	return _OmniPortal.Contract.Owner(&_OmniPortal.CallOpts)
}

// Xmsg is a free data retrieval call binding the contract method 0x2f32700e.
//
// Solidity: function xmsg() view returns((uint64,uint64,uint64,address,address,bytes,uint64))
func (_OmniPortal *OmniPortalCaller) Xmsg(opts *bind.CallOpts) (XTypesMsg, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "xmsg")

	if err != nil {
		return *new(XTypesMsg), err
	}

	out0 := *abi.ConvertType(out[0], new(XTypesMsg)).(*XTypesMsg)

	return out0, err

}

// Xmsg is a free data retrieval call binding the contract method 0x2f32700e.
//
// Solidity: function xmsg() view returns((uint64,uint64,uint64,address,address,bytes,uint64))
func (_OmniPortal *OmniPortalSession) Xmsg() (XTypesMsg, error) {
	return _OmniPortal.Contract.Xmsg(&_OmniPortal.CallOpts)
}

// Xmsg is a free data retrieval call binding the contract method 0x2f32700e.
//
// Solidity: function xmsg() view returns((uint64,uint64,uint64,address,address,bytes,uint64))
func (_OmniPortal *OmniPortalCallerSession) Xmsg() (XTypesMsg, error) {
	return _OmniPortal.Contract.Xmsg(&_OmniPortal.CallOpts)
}

// CollectFees is a paid mutator transaction binding the contract method 0xa480ca79.
//
// Solidity: function collectFees(address to) returns()
func (_OmniPortal *OmniPortalTransactor) CollectFees(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "collectFees", to)
}

// CollectFees is a paid mutator transaction binding the contract method 0xa480ca79.
//
// Solidity: function collectFees(address to) returns()
func (_OmniPortal *OmniPortalSession) CollectFees(to common.Address) (*types.Transaction, error) {
	return _OmniPortal.Contract.CollectFees(&_OmniPortal.TransactOpts, to)
}

// CollectFees is a paid mutator transaction binding the contract method 0xa480ca79.
//
// Solidity: function collectFees(address to) returns()
func (_OmniPortal *OmniPortalTransactorSession) CollectFees(to common.Address) (*types.Transaction, error) {
	return _OmniPortal.Contract.CollectFees(&_OmniPortal.TransactOpts, to)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniPortal *OmniPortalTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniPortal *OmniPortalSession) RenounceOwnership() (*types.Transaction, error) {
	return _OmniPortal.Contract.RenounceOwnership(&_OmniPortal.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniPortal *OmniPortalTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OmniPortal.Contract.RenounceOwnership(&_OmniPortal.TransactOpts)
}

// SetFeeOracle is a paid mutator transaction binding the contract method 0xa8a98962.
//
// Solidity: function setFeeOracle(address feeOracle_) returns()
func (_OmniPortal *OmniPortalTransactor) SetFeeOracle(opts *bind.TransactOpts, feeOracle_ common.Address) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "setFeeOracle", feeOracle_)
}

// SetFeeOracle is a paid mutator transaction binding the contract method 0xa8a98962.
//
// Solidity: function setFeeOracle(address feeOracle_) returns()
func (_OmniPortal *OmniPortalSession) SetFeeOracle(feeOracle_ common.Address) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetFeeOracle(&_OmniPortal.TransactOpts, feeOracle_)
}

// SetFeeOracle is a paid mutator transaction binding the contract method 0xa8a98962.
//
// Solidity: function setFeeOracle(address feeOracle_) returns()
func (_OmniPortal *OmniPortalTransactorSession) SetFeeOracle(feeOracle_ common.Address) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetFeeOracle(&_OmniPortal.TransactOpts, feeOracle_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniPortal *OmniPortalTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniPortal *OmniPortalSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OmniPortal.Contract.TransferOwnership(&_OmniPortal.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniPortal *OmniPortalTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OmniPortal.Contract.TransferOwnership(&_OmniPortal.TransactOpts, newOwner)
}

// Xcall is a paid mutator transaction binding the contract method 0x50e646dd.
//
// Solidity: function xcall(uint64 destChainId, address to, bytes data) payable returns()
func (_OmniPortal *OmniPortalTransactor) Xcall(opts *bind.TransactOpts, destChainId uint64, to common.Address, data []byte) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "xcall", destChainId, to, data)
}

// Xcall is a paid mutator transaction binding the contract method 0x50e646dd.
//
// Solidity: function xcall(uint64 destChainId, address to, bytes data) payable returns()
func (_OmniPortal *OmniPortalSession) Xcall(destChainId uint64, to common.Address, data []byte) (*types.Transaction, error) {
	return _OmniPortal.Contract.Xcall(&_OmniPortal.TransactOpts, destChainId, to, data)
}

// Xcall is a paid mutator transaction binding the contract method 0x50e646dd.
//
// Solidity: function xcall(uint64 destChainId, address to, bytes data) payable returns()
func (_OmniPortal *OmniPortalTransactorSession) Xcall(destChainId uint64, to common.Address, data []byte) (*types.Transaction, error) {
	return _OmniPortal.Contract.Xcall(&_OmniPortal.TransactOpts, destChainId, to, data)
}

// Xcall0 is a paid mutator transaction binding the contract method 0x70e8b56a.
//
// Solidity: function xcall(uint64 destChainId, address to, bytes data, uint64 gasLimit) payable returns()
func (_OmniPortal *OmniPortalTransactor) Xcall0(opts *bind.TransactOpts, destChainId uint64, to common.Address, data []byte, gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "xcall0", destChainId, to, data, gasLimit)
}

// Xcall0 is a paid mutator transaction binding the contract method 0x70e8b56a.
//
// Solidity: function xcall(uint64 destChainId, address to, bytes data, uint64 gasLimit) payable returns()
func (_OmniPortal *OmniPortalSession) Xcall0(destChainId uint64, to common.Address, data []byte, gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.Xcall0(&_OmniPortal.TransactOpts, destChainId, to, data, gasLimit)
}

// Xcall0 is a paid mutator transaction binding the contract method 0x70e8b56a.
//
// Solidity: function xcall(uint64 destChainId, address to, bytes data, uint64 gasLimit) payable returns()
func (_OmniPortal *OmniPortalTransactorSession) Xcall0(destChainId uint64, to common.Address, data []byte, gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.Xcall0(&_OmniPortal.TransactOpts, destChainId, to, data, gasLimit)
}

// Xsubmit is a paid mutator transaction binding the contract method 0x06f9f174.
//
// Solidity: function xsubmit((bytes32,(uint64,uint64,bytes32),(uint64,uint64,uint64,address,address,bytes,uint64)[],bytes32[],bool[],(bytes,bytes)[]) xsub) returns()
func (_OmniPortal *OmniPortalTransactor) Xsubmit(opts *bind.TransactOpts, xsub XTypesSubmission) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "xsubmit", xsub)
}

// Xsubmit is a paid mutator transaction binding the contract method 0x06f9f174.
//
// Solidity: function xsubmit((bytes32,(uint64,uint64,bytes32),(uint64,uint64,uint64,address,address,bytes,uint64)[],bytes32[],bool[],(bytes,bytes)[]) xsub) returns()
func (_OmniPortal *OmniPortalSession) Xsubmit(xsub XTypesSubmission) (*types.Transaction, error) {
	return _OmniPortal.Contract.Xsubmit(&_OmniPortal.TransactOpts, xsub)
}

// Xsubmit is a paid mutator transaction binding the contract method 0x06f9f174.
//
// Solidity: function xsubmit((bytes32,(uint64,uint64,bytes32),(uint64,uint64,uint64,address,address,bytes,uint64)[],bytes32[],bool[],(bytes,bytes)[]) xsub) returns()
func (_OmniPortal *OmniPortalTransactorSession) Xsubmit(xsub XTypesSubmission) (*types.Transaction, error) {
	return _OmniPortal.Contract.Xsubmit(&_OmniPortal.TransactOpts, xsub)
}

// OmniPortalFeeOracleChangedIterator is returned from FilterFeeOracleChanged and is used to iterate over the raw logs and unpacked data for FeeOracleChanged events raised by the OmniPortal contract.
type OmniPortalFeeOracleChangedIterator struct {
	Event *OmniPortalFeeOracleChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OmniPortalFeeOracleChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalFeeOracleChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OmniPortalFeeOracleChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OmniPortalFeeOracleChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalFeeOracleChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalFeeOracleChanged represents a FeeOracleChanged event raised by the OmniPortal contract.
type OmniPortalFeeOracleChanged struct {
	OldFeeOracle common.Address
	NewFeeOracle common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeOracleChanged is a free log retrieval operation binding the contract event 0x2819896846ed9ab612eb19218fd845f8328f084c8706b9ec2c47eabd479037a2.
//
// Solidity: event FeeOracleChanged(address indexed oldFeeOracle, address indexed newFeeOracle)
func (_OmniPortal *OmniPortalFilterer) FilterFeeOracleChanged(opts *bind.FilterOpts, oldFeeOracle []common.Address, newFeeOracle []common.Address) (*OmniPortalFeeOracleChangedIterator, error) {

	var oldFeeOracleRule []interface{}
	for _, oldFeeOracleItem := range oldFeeOracle {
		oldFeeOracleRule = append(oldFeeOracleRule, oldFeeOracleItem)
	}
	var newFeeOracleRule []interface{}
	for _, newFeeOracleItem := range newFeeOracle {
		newFeeOracleRule = append(newFeeOracleRule, newFeeOracleItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "FeeOracleChanged", oldFeeOracleRule, newFeeOracleRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalFeeOracleChangedIterator{contract: _OmniPortal.contract, event: "FeeOracleChanged", logs: logs, sub: sub}, nil
}

// WatchFeeOracleChanged is a free log subscription operation binding the contract event 0x2819896846ed9ab612eb19218fd845f8328f084c8706b9ec2c47eabd479037a2.
//
// Solidity: event FeeOracleChanged(address indexed oldFeeOracle, address indexed newFeeOracle)
func (_OmniPortal *OmniPortalFilterer) WatchFeeOracleChanged(opts *bind.WatchOpts, sink chan<- *OmniPortalFeeOracleChanged, oldFeeOracle []common.Address, newFeeOracle []common.Address) (event.Subscription, error) {

	var oldFeeOracleRule []interface{}
	for _, oldFeeOracleItem := range oldFeeOracle {
		oldFeeOracleRule = append(oldFeeOracleRule, oldFeeOracleItem)
	}
	var newFeeOracleRule []interface{}
	for _, newFeeOracleItem := range newFeeOracle {
		newFeeOracleRule = append(newFeeOracleRule, newFeeOracleItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "FeeOracleChanged", oldFeeOracleRule, newFeeOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalFeeOracleChanged)
				if err := _OmniPortal.contract.UnpackLog(event, "FeeOracleChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeeOracleChanged is a log parse operation binding the contract event 0x2819896846ed9ab612eb19218fd845f8328f084c8706b9ec2c47eabd479037a2.
//
// Solidity: event FeeOracleChanged(address indexed oldFeeOracle, address indexed newFeeOracle)
func (_OmniPortal *OmniPortalFilterer) ParseFeeOracleChanged(log types.Log) (*OmniPortalFeeOracleChanged, error) {
	event := new(OmniPortalFeeOracleChanged)
	if err := _OmniPortal.contract.UnpackLog(event, "FeeOracleChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalFeesCollectedIterator is returned from FilterFeesCollected and is used to iterate over the raw logs and unpacked data for FeesCollected events raised by the OmniPortal contract.
type OmniPortalFeesCollectedIterator struct {
	Event *OmniPortalFeesCollected // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OmniPortalFeesCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalFeesCollected)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OmniPortalFeesCollected)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OmniPortalFeesCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalFeesCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalFeesCollected represents a FeesCollected event raised by the OmniPortal contract.
type OmniPortalFeesCollected struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeesCollected is a free log retrieval operation binding the contract event 0x9dc46f23cfb5ddcad0ae7ea2be38d47fec07bb9382ec7e564efc69e036dd66ce.
//
// Solidity: event FeesCollected(address indexed to, uint256 amount)
func (_OmniPortal *OmniPortalFilterer) FilterFeesCollected(opts *bind.FilterOpts, to []common.Address) (*OmniPortalFeesCollectedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "FeesCollected", toRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalFeesCollectedIterator{contract: _OmniPortal.contract, event: "FeesCollected", logs: logs, sub: sub}, nil
}

// WatchFeesCollected is a free log subscription operation binding the contract event 0x9dc46f23cfb5ddcad0ae7ea2be38d47fec07bb9382ec7e564efc69e036dd66ce.
//
// Solidity: event FeesCollected(address indexed to, uint256 amount)
func (_OmniPortal *OmniPortalFilterer) WatchFeesCollected(opts *bind.WatchOpts, sink chan<- *OmniPortalFeesCollected, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "FeesCollected", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalFeesCollected)
				if err := _OmniPortal.contract.UnpackLog(event, "FeesCollected", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeesCollected is a log parse operation binding the contract event 0x9dc46f23cfb5ddcad0ae7ea2be38d47fec07bb9382ec7e564efc69e036dd66ce.
//
// Solidity: event FeesCollected(address indexed to, uint256 amount)
func (_OmniPortal *OmniPortalFilterer) ParseFeesCollected(log types.Log) (*OmniPortalFeesCollected, error) {
	event := new(OmniPortalFeesCollected)
	if err := _OmniPortal.contract.UnpackLog(event, "FeesCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OmniPortal contract.
type OmniPortalOwnershipTransferredIterator struct {
	Event *OmniPortalOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OmniPortalOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OmniPortalOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OmniPortalOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalOwnershipTransferred represents a OwnershipTransferred event raised by the OmniPortal contract.
type OmniPortalOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OmniPortal *OmniPortalFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OmniPortalOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalOwnershipTransferredIterator{contract: _OmniPortal.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OmniPortal *OmniPortalFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OmniPortalOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalOwnershipTransferred)
				if err := _OmniPortal.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OmniPortal *OmniPortalFilterer) ParseOwnershipTransferred(log types.Log) (*OmniPortalOwnershipTransferred, error) {
	event := new(OmniPortalOwnershipTransferred)
	if err := _OmniPortal.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXMsgIterator is returned from FilterXMsg and is used to iterate over the raw logs and unpacked data for XMsg events raised by the OmniPortal contract.
type OmniPortalXMsgIterator struct {
	Event *OmniPortalXMsg // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OmniPortalXMsgIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXMsg)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OmniPortalXMsg)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OmniPortalXMsgIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXMsgIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXMsg represents a XMsg event raised by the OmniPortal contract.
type OmniPortalXMsg struct {
	DestChainId  uint64
	StreamOffset uint64
	Sender       common.Address
	To           common.Address
	Data         []byte
	GasLimit     uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterXMsg is a free log retrieval operation binding the contract event 0xac3afbbff5be7c4af1610721cf4793840bd167251fd6f184ee708f752a731283.
//
// Solidity: event XMsg(uint64 indexed destChainId, uint64 indexed streamOffset, address sender, address to, bytes data, uint64 gasLimit)
func (_OmniPortal *OmniPortalFilterer) FilterXMsg(opts *bind.FilterOpts, destChainId []uint64, streamOffset []uint64) (*OmniPortalXMsgIterator, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}
	var streamOffsetRule []interface{}
	for _, streamOffsetItem := range streamOffset {
		streamOffsetRule = append(streamOffsetRule, streamOffsetItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XMsg", destChainIdRule, streamOffsetRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalXMsgIterator{contract: _OmniPortal.contract, event: "XMsg", logs: logs, sub: sub}, nil
}

// WatchXMsg is a free log subscription operation binding the contract event 0xac3afbbff5be7c4af1610721cf4793840bd167251fd6f184ee708f752a731283.
//
// Solidity: event XMsg(uint64 indexed destChainId, uint64 indexed streamOffset, address sender, address to, bytes data, uint64 gasLimit)
func (_OmniPortal *OmniPortalFilterer) WatchXMsg(opts *bind.WatchOpts, sink chan<- *OmniPortalXMsg, destChainId []uint64, streamOffset []uint64) (event.Subscription, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}
	var streamOffsetRule []interface{}
	for _, streamOffsetItem := range streamOffset {
		streamOffsetRule = append(streamOffsetRule, streamOffsetItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XMsg", destChainIdRule, streamOffsetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXMsg)
				if err := _OmniPortal.contract.UnpackLog(event, "XMsg", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseXMsg is a log parse operation binding the contract event 0xac3afbbff5be7c4af1610721cf4793840bd167251fd6f184ee708f752a731283.
//
// Solidity: event XMsg(uint64 indexed destChainId, uint64 indexed streamOffset, address sender, address to, bytes data, uint64 gasLimit)
func (_OmniPortal *OmniPortalFilterer) ParseXMsg(log types.Log) (*OmniPortalXMsg, error) {
	event := new(OmniPortalXMsg)
	if err := _OmniPortal.contract.UnpackLog(event, "XMsg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXReceiptIterator is returned from FilterXReceipt and is used to iterate over the raw logs and unpacked data for XReceipt events raised by the OmniPortal contract.
type OmniPortalXReceiptIterator struct {
	Event *OmniPortalXReceipt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OmniPortalXReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXReceipt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OmniPortalXReceipt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OmniPortalXReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXReceipt represents a XReceipt event raised by the OmniPortal contract.
type OmniPortalXReceipt struct {
	SourceChainId uint64
	StreamOffset  uint64
	GasUsed       *big.Int
	Relayer       common.Address
	Success       bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterXReceipt is a free log retrieval operation binding the contract event 0x34515b4105a7bb34f3af3cd490137ab292bb2ff14efb800df5c7d59e28944f25.
//
// Solidity: event XReceipt(uint64 indexed sourceChainId, uint64 indexed streamOffset, uint256 gasUsed, address relayer, bool success)
func (_OmniPortal *OmniPortalFilterer) FilterXReceipt(opts *bind.FilterOpts, sourceChainId []uint64, streamOffset []uint64) (*OmniPortalXReceiptIterator, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var streamOffsetRule []interface{}
	for _, streamOffsetItem := range streamOffset {
		streamOffsetRule = append(streamOffsetRule, streamOffsetItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XReceipt", sourceChainIdRule, streamOffsetRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalXReceiptIterator{contract: _OmniPortal.contract, event: "XReceipt", logs: logs, sub: sub}, nil
}

// WatchXReceipt is a free log subscription operation binding the contract event 0x34515b4105a7bb34f3af3cd490137ab292bb2ff14efb800df5c7d59e28944f25.
//
// Solidity: event XReceipt(uint64 indexed sourceChainId, uint64 indexed streamOffset, uint256 gasUsed, address relayer, bool success)
func (_OmniPortal *OmniPortalFilterer) WatchXReceipt(opts *bind.WatchOpts, sink chan<- *OmniPortalXReceipt, sourceChainId []uint64, streamOffset []uint64) (event.Subscription, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var streamOffsetRule []interface{}
	for _, streamOffsetItem := range streamOffset {
		streamOffsetRule = append(streamOffsetRule, streamOffsetItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XReceipt", sourceChainIdRule, streamOffsetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXReceipt)
				if err := _OmniPortal.contract.UnpackLog(event, "XReceipt", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseXReceipt is a log parse operation binding the contract event 0x34515b4105a7bb34f3af3cd490137ab292bb2ff14efb800df5c7d59e28944f25.
//
// Solidity: event XReceipt(uint64 indexed sourceChainId, uint64 indexed streamOffset, uint256 gasUsed, address relayer, bool success)
func (_OmniPortal *OmniPortalFilterer) ParseXReceipt(log types.Log) (*OmniPortalXReceipt, error) {
	event := new(OmniPortalXReceipt)
	if err := _OmniPortal.contract.UnpackLog(event, "XReceipt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
