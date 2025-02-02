// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package challengeRPC

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ArbMachineABI is the input ABI used to generate the binding from.
const ArbMachineABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"instructionStackHash\",\"type\":\"bytes32\"},{\"name\":\"dataStackHash\",\"type\":\"bytes32\"},{\"name\":\"auxStackHash\",\"type\":\"bytes32\"},{\"name\":\"registerHash\",\"type\":\"bytes32\"},{\"name\":\"staticHash\",\"type\":\"bytes32\"},{\"name\":\"errHandlerHash\",\"type\":\"bytes32\"}],\"name\":\"machineHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArbMachineFuncSigs maps the 4-byte function signature to its string representation.
var ArbMachineFuncSigs = map[string]string{
	"c1355b59": "machineHash(bytes32,bytes32,bytes32,bytes32,bytes32,bytes32)",
}

// ArbMachineBin is the compiled bytecode used for deploying new contracts.
var ArbMachineBin = "0x6101d6610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c8063c1355b591461003a575b600080fd5b610075600480360360c081101561005057600080fd5b5080359060208101359060408101359060608101359060808101359060a00135610087565b60408051918252519081900360200190f35b604080516101008101825260e081018881528152815160208181018452888252808301919091528251808201845287815282840152825180820184528681526060830152825180820184528581526080830152825190810190925282825260a0810191909152600060c08201819052906101009061010b565b979650505050505050565b600060028260c0015114156101225750600061019c565b60018260c0015114156101375750600161019c565b508051516020808301515160408085015151606080870151516080808901515160a0808b0151518751808b019b909b528a8801989098529389019490945287015285015260c0808501929092528051808503909201825260e090930190925281519101205b91905056fea265627a7a723058202de57ff579deb2614671ecb57b8a6ba3e1da8cd1cc762df772d54742cf26db1964736f6c634300050a0032"

// DeployArbMachine deploys a new Ethereum contract, binding an instance of ArbMachine to it.
func DeployArbMachine(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArbMachine, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbMachineABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbMachineBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbMachine{ArbMachineCaller: ArbMachineCaller{contract: contract}, ArbMachineTransactor: ArbMachineTransactor{contract: contract}, ArbMachineFilterer: ArbMachineFilterer{contract: contract}}, nil
}

// ArbMachine is an auto generated Go binding around an Ethereum contract.
type ArbMachine struct {
	ArbMachineCaller     // Read-only binding to the contract
	ArbMachineTransactor // Write-only binding to the contract
	ArbMachineFilterer   // Log filterer for contract events
}

// ArbMachineCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbMachineCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbMachineTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbMachineTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbMachineFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbMachineFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbMachineSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbMachineSession struct {
	Contract     *ArbMachine       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbMachineCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbMachineCallerSession struct {
	Contract *ArbMachineCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ArbMachineTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbMachineTransactorSession struct {
	Contract     *ArbMachineTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ArbMachineRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbMachineRaw struct {
	Contract *ArbMachine // Generic contract binding to access the raw methods on
}

// ArbMachineCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbMachineCallerRaw struct {
	Contract *ArbMachineCaller // Generic read-only contract binding to access the raw methods on
}

// ArbMachineTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbMachineTransactorRaw struct {
	Contract *ArbMachineTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbMachine creates a new instance of ArbMachine, bound to a specific deployed contract.
func NewArbMachine(address common.Address, backend bind.ContractBackend) (*ArbMachine, error) {
	contract, err := bindArbMachine(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbMachine{ArbMachineCaller: ArbMachineCaller{contract: contract}, ArbMachineTransactor: ArbMachineTransactor{contract: contract}, ArbMachineFilterer: ArbMachineFilterer{contract: contract}}, nil
}

// NewArbMachineCaller creates a new read-only instance of ArbMachine, bound to a specific deployed contract.
func NewArbMachineCaller(address common.Address, caller bind.ContractCaller) (*ArbMachineCaller, error) {
	contract, err := bindArbMachine(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbMachineCaller{contract: contract}, nil
}

// NewArbMachineTransactor creates a new write-only instance of ArbMachine, bound to a specific deployed contract.
func NewArbMachineTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbMachineTransactor, error) {
	contract, err := bindArbMachine(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbMachineTransactor{contract: contract}, nil
}

// NewArbMachineFilterer creates a new log filterer instance of ArbMachine, bound to a specific deployed contract.
func NewArbMachineFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbMachineFilterer, error) {
	contract, err := bindArbMachine(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbMachineFilterer{contract: contract}, nil
}

// bindArbMachine binds a generic wrapper to an already deployed contract.
func bindArbMachine(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbMachineABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbMachine *ArbMachineRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbMachine.Contract.ArbMachineCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbMachine *ArbMachineRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbMachine.Contract.ArbMachineTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbMachine *ArbMachineRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbMachine.Contract.ArbMachineTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbMachine *ArbMachineCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbMachine.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbMachine *ArbMachineTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbMachine.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbMachine *ArbMachineTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbMachine.Contract.contract.Transact(opts, method, params...)
}

// MachineHash is a free data retrieval call binding the contract method 0xc1355b59.
//
// Solidity: function machineHash(bytes32 instructionStackHash, bytes32 dataStackHash, bytes32 auxStackHash, bytes32 registerHash, bytes32 staticHash, bytes32 errHandlerHash) constant returns(bytes32)
func (_ArbMachine *ArbMachineCaller) MachineHash(opts *bind.CallOpts, instructionStackHash [32]byte, dataStackHash [32]byte, auxStackHash [32]byte, registerHash [32]byte, staticHash [32]byte, errHandlerHash [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbMachine.contract.Call(opts, out, "machineHash", instructionStackHash, dataStackHash, auxStackHash, registerHash, staticHash, errHandlerHash)
	return *ret0, err
}

// MachineHash is a free data retrieval call binding the contract method 0xc1355b59.
//
// Solidity: function machineHash(bytes32 instructionStackHash, bytes32 dataStackHash, bytes32 auxStackHash, bytes32 registerHash, bytes32 staticHash, bytes32 errHandlerHash) constant returns(bytes32)
func (_ArbMachine *ArbMachineSession) MachineHash(instructionStackHash [32]byte, dataStackHash [32]byte, auxStackHash [32]byte, registerHash [32]byte, staticHash [32]byte, errHandlerHash [32]byte) ([32]byte, error) {
	return _ArbMachine.Contract.MachineHash(&_ArbMachine.CallOpts, instructionStackHash, dataStackHash, auxStackHash, registerHash, staticHash, errHandlerHash)
}

// MachineHash is a free data retrieval call binding the contract method 0xc1355b59.
//
// Solidity: function machineHash(bytes32 instructionStackHash, bytes32 dataStackHash, bytes32 auxStackHash, bytes32 registerHash, bytes32 staticHash, bytes32 errHandlerHash) constant returns(bytes32)
func (_ArbMachine *ArbMachineCallerSession) MachineHash(instructionStackHash [32]byte, dataStackHash [32]byte, auxStackHash [32]byte, registerHash [32]byte, staticHash [32]byte, errHandlerHash [32]byte) ([32]byte, error) {
	return _ArbMachine.Contract.MachineHash(&_ArbMachine.CallOpts, instructionStackHash, dataStackHash, auxStackHash, registerHash, staticHash, errHandlerHash)
}

// ArbProtocolABI is the input ABI used to generate the binding from.
const ArbProtocolABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_fields\",\"type\":\"bytes32[5]\"},{\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageData\",\"type\":\"bytes\"},{\"name\":\"_messageTokenNum\",\"type\":\"uint16[]\"},{\"name\":\"_messageAmount\",\"type\":\"uint256[]\"},{\"name\":\"_messageDestination\",\"type\":\"bytes32[]\"}],\"name\":\"unanimousAssertHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageTokenNums\",\"type\":\"uint16[]\"},{\"name\":\"_messageAmounts\",\"type\":\"uint256[]\"}],\"name\":\"calculateBeforeValues\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"name\":\"_firstMessageHash\",\"type\":\"bytes32\"},{\"name\":\"_lastMessageHash\",\"type\":\"bytes32\"},{\"name\":\"_firstLogHash\",\"type\":\"bytes32\"},{\"name\":\"_lastLogHash\",\"type\":\"bytes32\"},{\"name\":\"_totalMessageValueAmounts\",\"type\":\"uint256[]\"}],\"name\":\"generateAssertionHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_data\",\"type\":\"bytes\"},{\"name\":\"_tokenNums\",\"type\":\"uint16[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_destinations\",\"type\":\"bytes32[]\"}],\"name\":\"generateLastMessageHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_dest\",\"type\":\"bytes32\"},{\"name\":\"_data\",\"type\":\"bytes32\"},{\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_sender\",\"type\":\"bytes32\"}],\"name\":\"generateSentMessageHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_beforeBalances\",\"type\":\"uint256[]\"}],\"name\":\"generatePreconditionHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_beforeBalances\",\"type\":\"uint256[]\"}],\"name\":\"beforeBalancesValid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_dataHashes\",\"type\":\"bytes32[]\"},{\"name\":\"_tokenNums\",\"type\":\"uint16[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_destinations\",\"type\":\"bytes32[]\"}],\"name\":\"generateLastMessageHashStub\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes32\"},{\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_destination\",\"type\":\"bytes32\"}],\"name\":\"generateMessageStubHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pendingMessages\",\"type\":\"bytes32\"},{\"name\":\"_newMessage\",\"type\":\"bytes32\"}],\"name\":\"appendInboxPendingMessage\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_inboxHash\",\"type\":\"bytes32\"},{\"name\":\"_pendingMessages\",\"type\":\"bytes32\"}],\"name\":\"appendInboxMessages\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArbProtocolFuncSigs maps the 4-byte function signature to its string representation.
var ArbProtocolFuncSigs = map[string]string{
	"f11fcc26": "appendInboxMessages(bytes32,bytes32)",
	"d78d18ea": "appendInboxPendingMessage(bytes32,bytes32)",
	"af17d922": "beforeBalancesValid(bytes21[],uint256[])",
	"0f89fbff": "calculateBeforeValues(bytes21[],uint16[],uint256[])",
	"20903721": "generateAssertionHash(bytes32,uint32,bytes32,bytes32,bytes32,bytes32,uint256[])",
	"25200160": "generateLastMessageHash(bytes21[],bytes,uint16[],uint256[],bytes32[])",
	"b3277495": "generateLastMessageHashStub(bytes21[],bytes32[],uint16[],uint256[],bytes32[])",
	"ccf69dd7": "generateMessageStubHash(bytes32,bytes21,uint256,bytes32)",
	"3e285598": "generatePreconditionHash(bytes32,uint64[2],bytes32,bytes21[],uint256[])",
	"2a0500d8": "generateSentMessageHash(bytes32,bytes32,bytes21,uint256,bytes32)",
	"014bba5b": "unanimousAssertHash(bytes32[5],uint64[2],bytes21[],bytes,uint16[],uint256[],bytes32[])",
}

// ArbProtocolBin is the compiled bytecode used for deploying new contracts.
var ArbProtocolBin = "0x6121a6610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100b35760003560e01c80633e2855981161007b5780633e2855981461097e578063af17d92214610adc578063b327749514610c13578063ccf69dd714610ebc578063d78d18ea14610ef5578063f11fcc2614610f18576100b3565b8063014bba5b146100b85780630f89fbff146103d357806320903721146105c857806325200160146106935780632a0500d81461093f575b600080fd5b6103c160048036036101808110156100cf57600080fd5b810190808060a001906005806020026040519081016040528092919082600560200280828437600092019190915250506040805180820182529295949381810193925090600290839083908082843760009201919091525091949392602081019250359050600160201b81111561014557600080fd5b82018360208201111561015757600080fd5b803590602001918460208302840111600160201b8311171561017857600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156101c757600080fd5b8201836020820111156101d957600080fd5b803590602001918460018302840111600160201b831117156101fa57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561024c57600080fd5b82018360208201111561025e57600080fd5b803590602001918460208302840111600160201b8311171561027f57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156102ce57600080fd5b8201836020820111156102e057600080fd5b803590602001918460208302840111600160201b8311171561030157600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561035057600080fd5b82018360208201111561036257600080fd5b803590602001918460208302840111600160201b8311171561038357600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610f3b945050505050565b60408051918252519081900360200190f35b610578600480360360608110156103e957600080fd5b810190602081018135600160201b81111561040357600080fd5b82018360208201111561041557600080fd5b803590602001918460208302840111600160201b8311171561043657600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561048557600080fd5b82018360208201111561049757600080fd5b803590602001918460208302840111600160201b831117156104b857600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561050757600080fd5b82018360208201111561051957600080fd5b803590602001918460208302840111600160201b8311171561053a57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506110d4945050505050565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156105b457818101518382015260200161059c565b505050509050019250505060405180910390f35b6103c1600480360360e08110156105de57600080fd5b81359163ffffffff6020820135169160408201359160608101359160808201359160a08101359181019060e0810160c0820135600160201b81111561062257600080fd5b82018360208201111561063457600080fd5b803590602001918460208302840111600160201b8311171561065557600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506112b9945050505050565b6103c1600480360360a08110156106a957600080fd5b810190602081018135600160201b8111156106c357600080fd5b8201836020820111156106d557600080fd5b803590602001918460208302840111600160201b831117156106f657600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561074557600080fd5b82018360208201111561075757600080fd5b803590602001918460018302840111600160201b8311171561077857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156107ca57600080fd5b8201836020820111156107dc57600080fd5b803590602001918460208302840111600160201b831117156107fd57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561084c57600080fd5b82018360208201111561085e57600080fd5b803590602001918460208302840111600160201b8311171561087f57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156108ce57600080fd5b8201836020820111156108e057600080fd5b803590602001918460208302840111600160201b8311171561090157600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611324945050505050565b6103c1600480360360a081101561095557600080fd5b508035906020810135906001600160581b03196040820135169060608101359060800135611567565b6103c1600480360360c081101561099457600080fd5b6040805180820182528335939283019291606083019190602084019060029083908390808284376000920191909152509194833594909390925060408101915060200135600160201b8111156109e957600080fd5b8201836020820111156109fb57600080fd5b803590602001918460208302840111600160201b83111715610a1c57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610a6b57600080fd5b820183602082011115610a7d57600080fd5b803590602001918460208302840111600160201b83111715610a9e57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611745945050505050565b610bff60048036036040811015610af257600080fd5b810190602081018135600160201b811115610b0c57600080fd5b820183602082011115610b1e57600080fd5b803590602001918460208302840111600160201b83111715610b3f57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610b8e57600080fd5b820183602082011115610ba057600080fd5b803590602001918460208302840111600160201b83111715610bc157600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611830945050505050565b604080519115158252519081900360200190f35b6103c1600480360360a0811015610c2957600080fd5b810190602081018135600160201b811115610c4357600080fd5b820183602082011115610c5557600080fd5b803590602001918460208302840111600160201b83111715610c7657600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610cc557600080fd5b820183602082011115610cd757600080fd5b803590602001918460208302840111600160201b83111715610cf857600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610d4757600080fd5b820183602082011115610d5957600080fd5b803590602001918460208302840111600160201b83111715610d7a57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610dc957600080fd5b820183602082011115610ddb57600080fd5b803590602001918460208302840111600160201b83111715610dfc57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610e4b57600080fd5b820183602082011115610e5d57600080fd5b803590602001918460208302840111600160201b83111715610e7e57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611a35945050505050565b6103c160048036036080811015610ed257600080fd5b508035906001600160581b03196020820135169060408101359060600135611bd3565b6103c160048036036040811015610f0b57600080fd5b5080359060200135611cb1565b6103c160048036036040811015610f2e57600080fd5b5080359060200135611cf5565b6000878787878787876040516020018088600560200280838360005b83811015610f6f578181015183820152602001610f57565b5050505090500187600260200280838360005b83811015610f9a578181015183820152602001610f82565b50505050905001868051906020019060200280838360005b83811015610fca578181015183820152602001610fb2565b5050505090500185805190602001908083835b60208310610ffc5780518252601f199092019160209182019101610fdd565b51815160209384036101000a60001901801990921691161790528751919093019287810192500280838360005b83811015611041578181015183820152602001611029565b50505050905001838051906020019060200280838360005b83811015611071578181015183820152602001611059565b50505050905001828051906020019060200280838360005b838110156110a1578181015183820152602001611089565b50505050905001975050505050505050604051602081830303815290604052805190602001209050979650505050505050565b606060008351905060608551604051908082528060200260200182016040528015611109578160200160208202803883390190505b50905060005b828110156112af57600086828151811061112557fe5b60200260200101519050878161ffff168151811061113f57fe5b602002602001015160146015811061115357fe5b1a60f81b6001600160f81b0319166111a05785828151811061117157fe5b6020026020010151838261ffff168151811061118957fe5b6020026020010181815101915081815250506112a6565b828161ffff16815181106111b057fe5b602002602001015160001461120c576040805162461bcd60e51b815260206004820152601d60248201527f43616e277420696e636c756465204e465420746f6b656e207477696365000000604482015290519081900360640190fd5b85828151811061121857fe5b602002602001015160001415611275576040805162461bcd60e51b815260206004820152601f60248201527f4e465420746f6b656e206d7573742068617665206e6f6e2d7a65726f20696400604482015290519081900360640190fd5b85828151811061128157fe5b6020026020010151838261ffff168151811061129957fe5b6020026020010181815250505b5060010161110f565b5095945050505050565b600087878787878787604051602001808881526020018763ffffffff1663ffffffff1660e01b815260040186815260200185815260200184815260200183815260200182805190602001906020028083836000838110156110a1578181015183820152602001611089565b60008151835114611372576040805162461bcd60e51b8152602060048201526013602482015272092dce0eae840e6d2f4ca40dad2e6dac2e8c6d606b1b604482015290519081900360640190fd5b83518351146113be576040805162461bcd60e51b8152602060048201526013602482015272092dce0eae840e6d2f4ca40dad2e6dac2e8c6d606b1b604482015290519081900360640190fd5b825160009081908190815b818110156115595773__$0d86abb4a722a612872fb80f4c7e7e95bd$__6389df40da8b866040518363ffffffff1660e01b81526004018080602001838152602001828103825284818151815260200191508051906020019080838360005b8381101561143f578181015183820152602001611427565b50505050905090810190601f16801561146c5780820380516001836020036101000a031916815260200191505b509350505050604080518083038186803b15801561148957600080fd5b505af415801561149d573d6000803e3d6000fd5b505050506040513d60408110156114b357600080fd5b5080516020909101518a5191955093506115229084908d908c90859081106114d757fe5b602002602001015161ffff16815181106114ed57fe5b60200260200101518a848151811061150157fe5b60200260200101518a858151811061151557fe5b6020026020010151611bd3565b60408051602080820198909852808201839052815180820383018152606090910190915280519601959095209492506001016113c9565b505050505095945050505050565b60408051602080820188905281830187905260608083018690526001600160581b03198716608084015283518084036075018152609584018086528151919093012060048084526101358501909552600094909391929160b5015b6115ca61213b565b8152602001906001900390816115c25790505090506115e887611d0f565b816000815181106115f557fe5b602002602001018190525061160942611d6b565b8160018151811061161657fe5b602002602001018190525061162a43611d6b565b8160028151811061163757fe5b602090810291909101015261164b82611d6b565b8160038151811061165857fe5b602090810291909101015260408051600480825260a08201909252606091816020015b61168361213b565b81526020019060019003908161167b5790505090506116a182611dc5565b816000815181106116ae57fe5b60209081029190910101526116c285611d6b565b816001815181106116cf57fe5b60200260200101819052506116e386611d6b565b816002815181106116f057fe5b602090810291909101015261170e6001600160581b03198816611d6b565b8160038151811061171b57fe5b602002602001018190525061173761173282611dc5565b611e4d565b519998505050505050505050565b600085858260200201518660016020020151868686604051602001808781526020018667ffffffffffffffff1667ffffffffffffffff1660c01b81526008018567ffffffffffffffff1667ffffffffffffffff1660c01b8152600801848152602001838051906020019060200280838360005b838110156117d05781810151838201526020016117b8565b50505050905001828051906020019060200280838360005b838110156118005781810151838201526020016117e8565b50505050905001965050505050505060405160208183030381529060405280519060200120905095945050505050565b81516000908015806118425750806001145b15611851576001915050611a2f565b60005b600182038110156119e257600085828151811061186d57fe5b602002602001015160146015811061188157fe5b1a60f81b90506001600160f81b031981166118eb578582815181106118a257fe5b60200260200101516001600160581b0319168683600101815181106118c357fe5b60200260200101516001600160581b031916116118e65760009350505050611a2f565b6119d9565b600160f81b6001600160f81b0319821614156119cd5785828151811061190d57fe5b60200260200101516001600160581b03191686836001018151811061192e57fe5b60200260200101516001600160581b03191610806119bc575085828151811061195357fe5b60200260200101516001600160581b03191686836001018151811061197457fe5b60200260200101516001600160581b0319161480156119bc575084828151811061199a57fe5b60200260200101518583600101815181106119b157fe5b602002602001015111155b156118e65760009350505050611a2f565b60009350505050611a2f565b50600101611854565b50600160f81b8460018303815181106119f757fe5b6020026020010151601460158110611a0b57fe5b1a60f81b6001600160f81b0319161115611a29576000915050611a2f565b60019150505b92915050565b60008351855114611a83576040805162461bcd60e51b8152602060048201526013602482015272092dce0eae840e6d2f4ca40dad2e6dac2e8c6d606b1b604482015290519081900360640190fd5b8251855114611acf576040805162461bcd60e51b8152602060048201526013602482015272092dce0eae840e6d2f4ca40dad2e6dac2e8c6d606b1b604482015290519081900360640190fd5b8151855114611b1b576040805162461bcd60e51b8152602060048201526013602482015272092dce0eae840e6d2f4ca40dad2e6dac2e8c6d606b1b604482015290519081900360640190fd5b84516000908190815b81811015611bc557611b8e898281518110611b3b57fe5b60200260200101518b8a8481518110611b5057fe5b602002602001015161ffff1681518110611b6657fe5b6020026020010151898481518110611b7a57fe5b602002602001015189858151811061151557fe5b6040805160208082019790975280820183905281518082038301815260609091019091528051950194909420939250600101611b24565b509198975050505050505050565b60408051600480825260a0820190925260009160609190816020015b611bf761213b565b815260200190600190039081611bef579050509050611c1586611d0f565b81600081518110611c2257fe5b6020908102919091010152611c3683611d6b565b81600181518110611c4357fe5b6020026020010181905250611c5784611d6b565b81600281518110611c6457fe5b6020908102919091010152611c826001600160581b03198616611d6b565b81600381518110611c8f57fe5b6020026020010181905250611ca661173282611dc5565b519695505050505050565b6000611cee6040518060600160405280611ccb6000611d6b565b8152602001611cd986611d0f565b8152602001611ce785611d0f565b9052611f3c565b9392505050565b6000611cee6040518060600160405280611ccb6001611d6b565b611d1761213b565b604080516060810182528381528151600080825260208281019094529192830191611d58565b611d4561213b565b815260200190600190039081611d3d5790505b508152600260209091015290505b919050565b611d7361213b565b604080516060810182528381528151600080825260208281019094529192830191611db4565b611da161213b565b815260200190600190039081611d995790505b508152600060209091015292915050565b611dcd61213b565b611dd78251611fc4565b611e28576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b50604080516060810182526000815260208101839052915160030160ff169082015290565b611e5561215f565b6040820151600c60ff90911610611ea7576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b604082015160ff16611ed4576040518060200160405280611ecb8460000151611fcb565b90529050611d66565b604082015160ff1660021415611ef95750604080516020810190915281518152611d66565b600360ff16826040015160ff1610158015611f1d57506040820151600c60ff909116105b15611f3a576040518060200160405280611ecb8460200151611fef565bfe5b6040805160038082526080820190925260009160609190816020015b611f6061213b565b815260200190600190039081611f58575050805190915060005b81811015611fb257848160038110611f8e57fe5b6020020151838281518110611f9f57fe5b6020908102919091010152600101611f7a565b50611fbc82611fef565b949350505050565b6008101590565b60408051602080820193909352815180820384018152908201909152805191012090565b600060088251111561203f576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6060825160405190808252806020026020018201604052801561206c578160200160208202803883390190505b50805190915060005b818110156120c85761208561215f565b6120a186838151811061209457fe5b6020026020010151611e4d565b905080600001518483815181106120b457fe5b602090810291909101015250600101612075565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b838110156121115781810151838201526020016120f9565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b60405180606001604052806000815260200160608152602001600060ff1681525090565b6040805160208101909152600081529056fea265627a7a72305820395f6b6386945f9dea787f3c2fabcab2cab0869e553caa2fa6ab0c6b416a70fb64736f6c634300050a0032"

// DeployArbProtocol deploys a new Ethereum contract, binding an instance of ArbProtocol to it.
func DeployArbProtocol(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArbProtocol, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbProtocolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	arbValueAddr, _, _, _ := DeployArbValue(auth, backend)
	ArbProtocolBin = strings.Replace(ArbProtocolBin, "__$0d86abb4a722a612872fb80f4c7e7e95bd$__", arbValueAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbProtocolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbProtocol{ArbProtocolCaller: ArbProtocolCaller{contract: contract}, ArbProtocolTransactor: ArbProtocolTransactor{contract: contract}, ArbProtocolFilterer: ArbProtocolFilterer{contract: contract}}, nil
}

// ArbProtocol is an auto generated Go binding around an Ethereum contract.
type ArbProtocol struct {
	ArbProtocolCaller     // Read-only binding to the contract
	ArbProtocolTransactor // Write-only binding to the contract
	ArbProtocolFilterer   // Log filterer for contract events
}

// ArbProtocolCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbProtocolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbProtocolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbProtocolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbProtocolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbProtocolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbProtocolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbProtocolSession struct {
	Contract     *ArbProtocol      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbProtocolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbProtocolCallerSession struct {
	Contract *ArbProtocolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ArbProtocolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbProtocolTransactorSession struct {
	Contract     *ArbProtocolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ArbProtocolRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbProtocolRaw struct {
	Contract *ArbProtocol // Generic contract binding to access the raw methods on
}

// ArbProtocolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbProtocolCallerRaw struct {
	Contract *ArbProtocolCaller // Generic read-only contract binding to access the raw methods on
}

// ArbProtocolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbProtocolTransactorRaw struct {
	Contract *ArbProtocolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbProtocol creates a new instance of ArbProtocol, bound to a specific deployed contract.
func NewArbProtocol(address common.Address, backend bind.ContractBackend) (*ArbProtocol, error) {
	contract, err := bindArbProtocol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbProtocol{ArbProtocolCaller: ArbProtocolCaller{contract: contract}, ArbProtocolTransactor: ArbProtocolTransactor{contract: contract}, ArbProtocolFilterer: ArbProtocolFilterer{contract: contract}}, nil
}

// NewArbProtocolCaller creates a new read-only instance of ArbProtocol, bound to a specific deployed contract.
func NewArbProtocolCaller(address common.Address, caller bind.ContractCaller) (*ArbProtocolCaller, error) {
	contract, err := bindArbProtocol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbProtocolCaller{contract: contract}, nil
}

// NewArbProtocolTransactor creates a new write-only instance of ArbProtocol, bound to a specific deployed contract.
func NewArbProtocolTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbProtocolTransactor, error) {
	contract, err := bindArbProtocol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbProtocolTransactor{contract: contract}, nil
}

// NewArbProtocolFilterer creates a new log filterer instance of ArbProtocol, bound to a specific deployed contract.
func NewArbProtocolFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbProtocolFilterer, error) {
	contract, err := bindArbProtocol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbProtocolFilterer{contract: contract}, nil
}

// bindArbProtocol binds a generic wrapper to an already deployed contract.
func bindArbProtocol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbProtocolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbProtocol *ArbProtocolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbProtocol.Contract.ArbProtocolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbProtocol *ArbProtocolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbProtocol.Contract.ArbProtocolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbProtocol *ArbProtocolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbProtocol.Contract.ArbProtocolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbProtocol *ArbProtocolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbProtocol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbProtocol *ArbProtocolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbProtocol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbProtocol *ArbProtocolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbProtocol.Contract.contract.Transact(opts, method, params...)
}

// AppendInboxMessages is a free data retrieval call binding the contract method 0xf11fcc26.
//
// Solidity: function appendInboxMessages(bytes32 _inboxHash, bytes32 _pendingMessages) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) AppendInboxMessages(opts *bind.CallOpts, _inboxHash [32]byte, _pendingMessages [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "appendInboxMessages", _inboxHash, _pendingMessages)
	return *ret0, err
}

// AppendInboxMessages is a free data retrieval call binding the contract method 0xf11fcc26.
//
// Solidity: function appendInboxMessages(bytes32 _inboxHash, bytes32 _pendingMessages) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) AppendInboxMessages(_inboxHash [32]byte, _pendingMessages [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.AppendInboxMessages(&_ArbProtocol.CallOpts, _inboxHash, _pendingMessages)
}

// AppendInboxMessages is a free data retrieval call binding the contract method 0xf11fcc26.
//
// Solidity: function appendInboxMessages(bytes32 _inboxHash, bytes32 _pendingMessages) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) AppendInboxMessages(_inboxHash [32]byte, _pendingMessages [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.AppendInboxMessages(&_ArbProtocol.CallOpts, _inboxHash, _pendingMessages)
}

// AppendInboxPendingMessage is a free data retrieval call binding the contract method 0xd78d18ea.
//
// Solidity: function appendInboxPendingMessage(bytes32 _pendingMessages, bytes32 _newMessage) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) AppendInboxPendingMessage(opts *bind.CallOpts, _pendingMessages [32]byte, _newMessage [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "appendInboxPendingMessage", _pendingMessages, _newMessage)
	return *ret0, err
}

// AppendInboxPendingMessage is a free data retrieval call binding the contract method 0xd78d18ea.
//
// Solidity: function appendInboxPendingMessage(bytes32 _pendingMessages, bytes32 _newMessage) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) AppendInboxPendingMessage(_pendingMessages [32]byte, _newMessage [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.AppendInboxPendingMessage(&_ArbProtocol.CallOpts, _pendingMessages, _newMessage)
}

// AppendInboxPendingMessage is a free data retrieval call binding the contract method 0xd78d18ea.
//
// Solidity: function appendInboxPendingMessage(bytes32 _pendingMessages, bytes32 _newMessage) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) AppendInboxPendingMessage(_pendingMessages [32]byte, _newMessage [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.AppendInboxPendingMessage(&_ArbProtocol.CallOpts, _pendingMessages, _newMessage)
}

// BeforeBalancesValid is a free data retrieval call binding the contract method 0xaf17d922.
//
// Solidity: function beforeBalancesValid(bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bool)
func (_ArbProtocol *ArbProtocolCaller) BeforeBalancesValid(opts *bind.CallOpts, _tokenTypes [][21]byte, _beforeBalances []*big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "beforeBalancesValid", _tokenTypes, _beforeBalances)
	return *ret0, err
}

// BeforeBalancesValid is a free data retrieval call binding the contract method 0xaf17d922.
//
// Solidity: function beforeBalancesValid(bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bool)
func (_ArbProtocol *ArbProtocolSession) BeforeBalancesValid(_tokenTypes [][21]byte, _beforeBalances []*big.Int) (bool, error) {
	return _ArbProtocol.Contract.BeforeBalancesValid(&_ArbProtocol.CallOpts, _tokenTypes, _beforeBalances)
}

// BeforeBalancesValid is a free data retrieval call binding the contract method 0xaf17d922.
//
// Solidity: function beforeBalancesValid(bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bool)
func (_ArbProtocol *ArbProtocolCallerSession) BeforeBalancesValid(_tokenTypes [][21]byte, _beforeBalances []*big.Int) (bool, error) {
	return _ArbProtocol.Contract.BeforeBalancesValid(&_ArbProtocol.CallOpts, _tokenTypes, _beforeBalances)
}

// CalculateBeforeValues is a free data retrieval call binding the contract method 0x0f89fbff.
//
// Solidity: function calculateBeforeValues(bytes21[] _tokenTypes, uint16[] _messageTokenNums, uint256[] _messageAmounts) constant returns(uint256[])
func (_ArbProtocol *ArbProtocolCaller) CalculateBeforeValues(opts *bind.CallOpts, _tokenTypes [][21]byte, _messageTokenNums []uint16, _messageAmounts []*big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "calculateBeforeValues", _tokenTypes, _messageTokenNums, _messageAmounts)
	return *ret0, err
}

// CalculateBeforeValues is a free data retrieval call binding the contract method 0x0f89fbff.
//
// Solidity: function calculateBeforeValues(bytes21[] _tokenTypes, uint16[] _messageTokenNums, uint256[] _messageAmounts) constant returns(uint256[])
func (_ArbProtocol *ArbProtocolSession) CalculateBeforeValues(_tokenTypes [][21]byte, _messageTokenNums []uint16, _messageAmounts []*big.Int) ([]*big.Int, error) {
	return _ArbProtocol.Contract.CalculateBeforeValues(&_ArbProtocol.CallOpts, _tokenTypes, _messageTokenNums, _messageAmounts)
}

// CalculateBeforeValues is a free data retrieval call binding the contract method 0x0f89fbff.
//
// Solidity: function calculateBeforeValues(bytes21[] _tokenTypes, uint16[] _messageTokenNums, uint256[] _messageAmounts) constant returns(uint256[])
func (_ArbProtocol *ArbProtocolCallerSession) CalculateBeforeValues(_tokenTypes [][21]byte, _messageTokenNums []uint16, _messageAmounts []*big.Int) ([]*big.Int, error) {
	return _ArbProtocol.Contract.CalculateBeforeValues(&_ArbProtocol.CallOpts, _tokenTypes, _messageTokenNums, _messageAmounts)
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x20903721.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash, uint256[] _totalMessageValueAmounts) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateAssertionHash(opts *bind.CallOpts, _afterHash [32]byte, _numSteps uint32, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte, _totalMessageValueAmounts []*big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateAssertionHash", _afterHash, _numSteps, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash, _totalMessageValueAmounts)
	return *ret0, err
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x20903721.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash, uint256[] _totalMessageValueAmounts) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateAssertionHash(_afterHash [32]byte, _numSteps uint32, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte, _totalMessageValueAmounts []*big.Int) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateAssertionHash(&_ArbProtocol.CallOpts, _afterHash, _numSteps, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash, _totalMessageValueAmounts)
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x20903721.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash, uint256[] _totalMessageValueAmounts) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateAssertionHash(_afterHash [32]byte, _numSteps uint32, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte, _totalMessageValueAmounts []*big.Int) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateAssertionHash(&_ArbProtocol.CallOpts, _afterHash, _numSteps, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash, _totalMessageValueAmounts)
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0x25200160.
//
// Solidity: function generateLastMessageHash(bytes21[] _tokenTypes, bytes _data, uint16[] _tokenNums, uint256[] _amounts, bytes32[] _destinations) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateLastMessageHash(opts *bind.CallOpts, _tokenTypes [][21]byte, _data []byte, _tokenNums []uint16, _amounts []*big.Int, _destinations [][32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateLastMessageHash", _tokenTypes, _data, _tokenNums, _amounts, _destinations)
	return *ret0, err
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0x25200160.
//
// Solidity: function generateLastMessageHash(bytes21[] _tokenTypes, bytes _data, uint16[] _tokenNums, uint256[] _amounts, bytes32[] _destinations) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateLastMessageHash(_tokenTypes [][21]byte, _data []byte, _tokenNums []uint16, _amounts []*big.Int, _destinations [][32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateLastMessageHash(&_ArbProtocol.CallOpts, _tokenTypes, _data, _tokenNums, _amounts, _destinations)
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0x25200160.
//
// Solidity: function generateLastMessageHash(bytes21[] _tokenTypes, bytes _data, uint16[] _tokenNums, uint256[] _amounts, bytes32[] _destinations) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateLastMessageHash(_tokenTypes [][21]byte, _data []byte, _tokenNums []uint16, _amounts []*big.Int, _destinations [][32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateLastMessageHash(&_ArbProtocol.CallOpts, _tokenTypes, _data, _tokenNums, _amounts, _destinations)
}

// GenerateLastMessageHashStub is a free data retrieval call binding the contract method 0xb3277495.
//
// Solidity: function generateLastMessageHashStub(bytes21[] _tokenTypes, bytes32[] _dataHashes, uint16[] _tokenNums, uint256[] _amounts, bytes32[] _destinations) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateLastMessageHashStub(opts *bind.CallOpts, _tokenTypes [][21]byte, _dataHashes [][32]byte, _tokenNums []uint16, _amounts []*big.Int, _destinations [][32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateLastMessageHashStub", _tokenTypes, _dataHashes, _tokenNums, _amounts, _destinations)
	return *ret0, err
}

// GenerateLastMessageHashStub is a free data retrieval call binding the contract method 0xb3277495.
//
// Solidity: function generateLastMessageHashStub(bytes21[] _tokenTypes, bytes32[] _dataHashes, uint16[] _tokenNums, uint256[] _amounts, bytes32[] _destinations) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateLastMessageHashStub(_tokenTypes [][21]byte, _dataHashes [][32]byte, _tokenNums []uint16, _amounts []*big.Int, _destinations [][32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateLastMessageHashStub(&_ArbProtocol.CallOpts, _tokenTypes, _dataHashes, _tokenNums, _amounts, _destinations)
}

// GenerateLastMessageHashStub is a free data retrieval call binding the contract method 0xb3277495.
//
// Solidity: function generateLastMessageHashStub(bytes21[] _tokenTypes, bytes32[] _dataHashes, uint16[] _tokenNums, uint256[] _amounts, bytes32[] _destinations) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateLastMessageHashStub(_tokenTypes [][21]byte, _dataHashes [][32]byte, _tokenNums []uint16, _amounts []*big.Int, _destinations [][32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateLastMessageHashStub(&_ArbProtocol.CallOpts, _tokenTypes, _dataHashes, _tokenNums, _amounts, _destinations)
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0xccf69dd7.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, bytes32 _destination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateMessageStubHash(opts *bind.CallOpts, _data [32]byte, _tokenType [21]byte, _value *big.Int, _destination [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateMessageStubHash", _data, _tokenType, _value, _destination)
	return *ret0, err
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0xccf69dd7.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, bytes32 _destination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateMessageStubHash(_data [32]byte, _tokenType [21]byte, _value *big.Int, _destination [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateMessageStubHash(&_ArbProtocol.CallOpts, _data, _tokenType, _value, _destination)
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0xccf69dd7.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, bytes32 _destination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateMessageStubHash(_data [32]byte, _tokenType [21]byte, _value *big.Int, _destination [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateMessageStubHash(&_ArbProtocol.CallOpts, _data, _tokenType, _value, _destination)
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x3e285598.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox, bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GeneratePreconditionHash(opts *bind.CallOpts, _beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte, _tokenTypes [][21]byte, _beforeBalances []*big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generatePreconditionHash", _beforeHash, _timeBounds, _beforeInbox, _tokenTypes, _beforeBalances)
	return *ret0, err
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x3e285598.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox, bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GeneratePreconditionHash(_beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte, _tokenTypes [][21]byte, _beforeBalances []*big.Int) ([32]byte, error) {
	return _ArbProtocol.Contract.GeneratePreconditionHash(&_ArbProtocol.CallOpts, _beforeHash, _timeBounds, _beforeInbox, _tokenTypes, _beforeBalances)
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x3e285598.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox, bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GeneratePreconditionHash(_beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte, _tokenTypes [][21]byte, _beforeBalances []*big.Int) ([32]byte, error) {
	return _ArbProtocol.Contract.GeneratePreconditionHash(&_ArbProtocol.CallOpts, _beforeHash, _timeBounds, _beforeInbox, _tokenTypes, _beforeBalances)
}

// GenerateSentMessageHash is a free data retrieval call binding the contract method 0x2a0500d8.
//
// Solidity: function generateSentMessageHash(bytes32 _dest, bytes32 _data, bytes21 _tokenType, uint256 _value, bytes32 _sender) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateSentMessageHash(opts *bind.CallOpts, _dest [32]byte, _data [32]byte, _tokenType [21]byte, _value *big.Int, _sender [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateSentMessageHash", _dest, _data, _tokenType, _value, _sender)
	return *ret0, err
}

// GenerateSentMessageHash is a free data retrieval call binding the contract method 0x2a0500d8.
//
// Solidity: function generateSentMessageHash(bytes32 _dest, bytes32 _data, bytes21 _tokenType, uint256 _value, bytes32 _sender) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateSentMessageHash(_dest [32]byte, _data [32]byte, _tokenType [21]byte, _value *big.Int, _sender [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateSentMessageHash(&_ArbProtocol.CallOpts, _dest, _data, _tokenType, _value, _sender)
}

// GenerateSentMessageHash is a free data retrieval call binding the contract method 0x2a0500d8.
//
// Solidity: function generateSentMessageHash(bytes32 _dest, bytes32 _data, bytes21 _tokenType, uint256 _value, bytes32 _sender) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateSentMessageHash(_dest [32]byte, _data [32]byte, _tokenType [21]byte, _value *big.Int, _sender [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateSentMessageHash(&_ArbProtocol.CallOpts, _dest, _data, _tokenType, _value, _sender)
}

// UnanimousAssertHash is a free data retrieval call binding the contract method 0x014bba5b.
//
// Solidity: function unanimousAssertHash(bytes32[5] _fields, uint64[2] _timeBounds, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes32[] _messageDestination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) UnanimousAssertHash(opts *bind.CallOpts, _fields [5][32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination [][32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "unanimousAssertHash", _fields, _timeBounds, _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination)
	return *ret0, err
}

// UnanimousAssertHash is a free data retrieval call binding the contract method 0x014bba5b.
//
// Solidity: function unanimousAssertHash(bytes32[5] _fields, uint64[2] _timeBounds, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes32[] _messageDestination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) UnanimousAssertHash(_fields [5][32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination [][32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.UnanimousAssertHash(&_ArbProtocol.CallOpts, _fields, _timeBounds, _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination)
}

// UnanimousAssertHash is a free data retrieval call binding the contract method 0x014bba5b.
//
// Solidity: function unanimousAssertHash(bytes32[5] _fields, uint64[2] _timeBounds, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes32[] _messageDestination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) UnanimousAssertHash(_fields [5][32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination [][32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.UnanimousAssertHash(&_ArbProtocol.CallOpts, _fields, _timeBounds, _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination)
}

// ArbValueABI is the input ABI used to generate the binding from.
const ArbValueABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"hashIntValue\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"getNextValidValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"opcode\",\"type\":\"uint8\"},{\"name\":\"immediateVal\",\"type\":\"bytes32\"},{\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePointImmediateValue\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hashEmptyTuple\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"opcode\",\"type\":\"uint8\"},{\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePointBasicValue\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"deserializeValidValueHash\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deserializeValueHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"isValidTupleSize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArbValueFuncSigs maps the 4-byte function signature to its string representation.
var ArbValueFuncSigs = map[string]string{
	"89df40da": "deserializeValidValueHash(bytes,uint256)",
	"8f346036": "deserializeValueHash(bytes)",
	"1f3d4d4e": "getNextValidValue(bytes,uint256)",
	"53409fab": "hashCodePointBasicValue(uint8,bytes32)",
	"264f384b": "hashCodePointImmediateValue(uint8,bytes32,bytes32)",
	"364df277": "hashEmptyTuple()",
	"1667b411": "hashIntValue(uint256)",
	"b2b9dc62": "isValidTupleSize(uint256)",
}

// ArbValueBin is the compiled bytecode used for deploying new contracts.
var ArbValueBin = "0x610d71610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100925760003560e01c806353409fab1161006557806353409fab1461022157806389df40da146102475780638f34603614610308578063b2b9dc62146103ae57610092565b80631667b411146100975780631f3d4d4e146100c6578063264f384b146101ed578063364df27714610219575b600080fd5b6100b4600480360360208110156100ad57600080fd5b50356103df565b60408051918252519081900360200190f35b61016e600480360360408110156100dc57600080fd5b8101906020810181356401000000008111156100f757600080fd5b82018360208201111561010957600080fd5b8035906020019184600183028401116401000000008311171561012b57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610405915050565b6040518083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156101b1578181015183820152602001610199565b50505050905090810190601f1680156101de5780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b6100b46004803603606081101561020357600080fd5b5060ff813516906020810135906040013561049a565b6100b46104ec565b6100b46004803603604081101561023757600080fd5b5060ff813516906020013561055f565b6102ef6004803603604081101561025d57600080fd5b81019060208101813564010000000081111561027857600080fd5b82018360208201111561028a57600080fd5b803590602001918460018302840111640100000000831117156102ac57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506105a6915050565b6040805192835260208301919091528051918290030190f35b6100b46004803603602081101561031e57600080fd5b81019060208101813564010000000081111561033957600080fd5b82018360208201111561034b57600080fd5b8035906020019184600183028401116401000000008311171561036d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610631945050505050565b6103cb600480360360208110156103c457600080fd5b50356106b5565b604080519115158252519081900360200190f35b60408051602080820184905282518083038201815291830190925280519101205b919050565b60006060600080610414610d06565b61041e87876106bc565b919450925090508215610478576040805162461bcd60e51b815260206004820152601e60248201527f4d61727368616c6c65642076616c7565206d7573742062652076616c69640000604482015290519081900360640190fd5b8161048c888880840363ffffffff61081116565b945094505050509250929050565b60408051600160f81b60208083019190915260f89590951b6001600160f81b03191660218201526022810193909352604280840192909252805180840390920182526062909201909152805191012090565b6040805160008082526020808301808552600360f81b948401948552835192946003938593919260418501929091028083838a5b83811015610538578181015183820152602001610520565b50505050905001925050506040516020818303038152906040528051906020012091505090565b60408051600160f81b60208083019190915260f89490941b6001600160f81b0319166021820152602280820193909352815180820390930183526042019052805191012090565b6000806000806105b4610d06565b6105be87876106bc565b919450925090508215610618576040805162461bcd60e51b815260206004820152601e60248201527f4d61727368616c6c65642076616c7565206d7573742062652076616c69640000604482015290519081900360640190fd5b8161062282610891565b51909890975095505050505050565b6000808061063d610d06565b6106488560006106bc565b9194509250905082156106a2576040805162461bcd60e51b815260206004820152601e60248201527f4d61727368616c6c65642076616c7565206d7573742062652076616c69640000604482015290519081900360640190fd5b6106ab81610891565b5195945050505050565b6008101590565b6000806106c7610d06565b8451841061071c576040805162461bcd60e51b815260206004820152601960248201527f44617461206f6666736574206f7574206f6620626f756e647300000000000000604482015290519081900360640190fd5b6000849050600086828151811061072f57fe5b016020015160019092019160f81c9050600081610771576107508884610980565b9093509050600083610761836109a7565b9197509550935061080a92505050565b60ff821660021415610798576107878884610980565b909350905060008361076183610a01565b600360ff8316108015906107af5750600c60ff8316105b156107eb576002198201606060006107c8838c88610a5b565b9097509250905080866107da84610b16565b98509850985050505050505061080a565b8160ff166127100160006107ff60006109a7565b919750955093505050505b9250925092565b60608183018451101561082357600080fd5b60608215801561083e57604051915060208201604052610888565b6040519150601f8416801560200281840101858101878315602002848b0101015b8183101561087757805183526020928301920161085f565b5050858452601f01601f1916604052505b50949350505050565b610899610d2a565b6040820151600c60ff909116106108eb576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b604082015160ff1661091857604051806020016040528061090f84600001516103df565b90529050610400565b604082015160ff166002141561093d5750604080516020810190915281518152610400565b600360ff16826040015160ff161015801561096157506040820151600c60ff909116105b1561097e57604051806020016040528061090f8460200151610b9e565bfe5b6000808281610995868363ffffffff610cea16565b60209290920196919550909350505050565b6109af610d06565b6040805160608101825283815281516000808252602082810190945291928301916109f0565b6109dd610d06565b8152602001906001900390816109d55790505b508152600060209091015292915050565b610a09610d06565b604080516060810182528381528151600080825260208281019094529192830191610a4a565b610a37610d06565b815260200190600190039081610a2f5790505b508152600260209091015292915050565b60008060606000849050600060608860ff16604051908082528060200260200182016040528015610aa657816020015b610a93610d06565b815260200190600190039081610a8b5790505b50905060005b8960ff168160ff161015610b0057610ac489856106bc565b8451859060ff8616908110610ad557fe5b6020908102919091010152945092508215610af857509094509092509050610b0d565b600101610aac565b5060009550919350909150505b93509350939050565b610b1e610d06565b610b2882516106b5565b610b79576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b50604080516060810182526000815260208101839052915160030160ff169082015290565b6000600882511115610bee576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015610c1b578160200160208202803883390190505b50805190915060005b81811015610c7757610c34610d2a565b610c50868381518110610c4357fe5b6020026020010151610891565b90508060000151848381518110610c6357fe5b602090810291909101015250600101610c24565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015610cc0578181015183820152602001610ca8565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b60008160200183511015610cfd57600080fd5b50016020015190565b60405180606001604052806000815260200160608152602001600060ff1681525090565b6040805160208101909152600081529056fea265627a7a72305820e580d18e866df844b2ca541adbb02d9e0adb42ede410d1908ac1a96124e1d94b64736f6c634300050a0032"

// DeployArbValue deploys a new Ethereum contract, binding an instance of ArbValue to it.
func DeployArbValue(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArbValue, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbValueABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbValueBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbValue{ArbValueCaller: ArbValueCaller{contract: contract}, ArbValueTransactor: ArbValueTransactor{contract: contract}, ArbValueFilterer: ArbValueFilterer{contract: contract}}, nil
}

// ArbValue is an auto generated Go binding around an Ethereum contract.
type ArbValue struct {
	ArbValueCaller     // Read-only binding to the contract
	ArbValueTransactor // Write-only binding to the contract
	ArbValueFilterer   // Log filterer for contract events
}

// ArbValueCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbValueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbValueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbValueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbValueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbValueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbValueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbValueSession struct {
	Contract     *ArbValue         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbValueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbValueCallerSession struct {
	Contract *ArbValueCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ArbValueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbValueTransactorSession struct {
	Contract     *ArbValueTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ArbValueRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbValueRaw struct {
	Contract *ArbValue // Generic contract binding to access the raw methods on
}

// ArbValueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbValueCallerRaw struct {
	Contract *ArbValueCaller // Generic read-only contract binding to access the raw methods on
}

// ArbValueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbValueTransactorRaw struct {
	Contract *ArbValueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbValue creates a new instance of ArbValue, bound to a specific deployed contract.
func NewArbValue(address common.Address, backend bind.ContractBackend) (*ArbValue, error) {
	contract, err := bindArbValue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbValue{ArbValueCaller: ArbValueCaller{contract: contract}, ArbValueTransactor: ArbValueTransactor{contract: contract}, ArbValueFilterer: ArbValueFilterer{contract: contract}}, nil
}

// NewArbValueCaller creates a new read-only instance of ArbValue, bound to a specific deployed contract.
func NewArbValueCaller(address common.Address, caller bind.ContractCaller) (*ArbValueCaller, error) {
	contract, err := bindArbValue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbValueCaller{contract: contract}, nil
}

// NewArbValueTransactor creates a new write-only instance of ArbValue, bound to a specific deployed contract.
func NewArbValueTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbValueTransactor, error) {
	contract, err := bindArbValue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbValueTransactor{contract: contract}, nil
}

// NewArbValueFilterer creates a new log filterer instance of ArbValue, bound to a specific deployed contract.
func NewArbValueFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbValueFilterer, error) {
	contract, err := bindArbValue(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbValueFilterer{contract: contract}, nil
}

// bindArbValue binds a generic wrapper to an already deployed contract.
func bindArbValue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbValueABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbValue *ArbValueRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbValue.Contract.ArbValueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbValue *ArbValueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbValue.Contract.ArbValueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbValue *ArbValueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbValue.Contract.ArbValueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbValue *ArbValueCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbValue.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbValue *ArbValueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbValue.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbValue *ArbValueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbValue.Contract.contract.Transact(opts, method, params...)
}

// DeserializeValidValueHash is a free data retrieval call binding the contract method 0x89df40da.
//
// Solidity: function deserializeValidValueHash(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_ArbValue *ArbValueCaller) DeserializeValidValueHash(opts *bind.CallOpts, data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ArbValue.contract.Call(opts, out, "deserializeValidValueHash", data, offset)
	return *ret0, *ret1, err
}

// DeserializeValidValueHash is a free data retrieval call binding the contract method 0x89df40da.
//
// Solidity: function deserializeValidValueHash(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_ArbValue *ArbValueSession) DeserializeValidValueHash(data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	return _ArbValue.Contract.DeserializeValidValueHash(&_ArbValue.CallOpts, data, offset)
}

// DeserializeValidValueHash is a free data retrieval call binding the contract method 0x89df40da.
//
// Solidity: function deserializeValidValueHash(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_ArbValue *ArbValueCallerSession) DeserializeValidValueHash(data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	return _ArbValue.Contract.DeserializeValidValueHash(&_ArbValue.CallOpts, data, offset)
}

// DeserializeValueHash is a free data retrieval call binding the contract method 0x8f346036.
//
// Solidity: function deserializeValueHash(bytes data) constant returns(bytes32)
func (_ArbValue *ArbValueCaller) DeserializeValueHash(opts *bind.CallOpts, data []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbValue.contract.Call(opts, out, "deserializeValueHash", data)
	return *ret0, err
}

// DeserializeValueHash is a free data retrieval call binding the contract method 0x8f346036.
//
// Solidity: function deserializeValueHash(bytes data) constant returns(bytes32)
func (_ArbValue *ArbValueSession) DeserializeValueHash(data []byte) ([32]byte, error) {
	return _ArbValue.Contract.DeserializeValueHash(&_ArbValue.CallOpts, data)
}

// DeserializeValueHash is a free data retrieval call binding the contract method 0x8f346036.
//
// Solidity: function deserializeValueHash(bytes data) constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) DeserializeValueHash(data []byte) ([32]byte, error) {
	return _ArbValue.Contract.DeserializeValueHash(&_ArbValue.CallOpts, data)
}

// GetNextValidValue is a free data retrieval call binding the contract method 0x1f3d4d4e.
//
// Solidity: function getNextValidValue(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_ArbValue *ArbValueCaller) GetNextValidValue(opts *bind.CallOpts, data []byte, offset *big.Int) (*big.Int, []byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ArbValue.contract.Call(opts, out, "getNextValidValue", data, offset)
	return *ret0, *ret1, err
}

// GetNextValidValue is a free data retrieval call binding the contract method 0x1f3d4d4e.
//
// Solidity: function getNextValidValue(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_ArbValue *ArbValueSession) GetNextValidValue(data []byte, offset *big.Int) (*big.Int, []byte, error) {
	return _ArbValue.Contract.GetNextValidValue(&_ArbValue.CallOpts, data, offset)
}

// GetNextValidValue is a free data retrieval call binding the contract method 0x1f3d4d4e.
//
// Solidity: function getNextValidValue(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_ArbValue *ArbValueCallerSession) GetNextValidValue(data []byte, offset *big.Int) (*big.Int, []byte, error) {
	return _ArbValue.Contract.GetNextValidValue(&_ArbValue.CallOpts, data, offset)
}

// HashCodePointBasicValue is a free data retrieval call binding the contract method 0x53409fab.
//
// Solidity: function hashCodePointBasicValue(uint8 opcode, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueCaller) HashCodePointBasicValue(opts *bind.CallOpts, opcode uint8, nextCodePoint [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbValue.contract.Call(opts, out, "hashCodePointBasicValue", opcode, nextCodePoint)
	return *ret0, err
}

// HashCodePointBasicValue is a free data retrieval call binding the contract method 0x53409fab.
//
// Solidity: function hashCodePointBasicValue(uint8 opcode, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueSession) HashCodePointBasicValue(opcode uint8, nextCodePoint [32]byte) ([32]byte, error) {
	return _ArbValue.Contract.HashCodePointBasicValue(&_ArbValue.CallOpts, opcode, nextCodePoint)
}

// HashCodePointBasicValue is a free data retrieval call binding the contract method 0x53409fab.
//
// Solidity: function hashCodePointBasicValue(uint8 opcode, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) HashCodePointBasicValue(opcode uint8, nextCodePoint [32]byte) ([32]byte, error) {
	return _ArbValue.Contract.HashCodePointBasicValue(&_ArbValue.CallOpts, opcode, nextCodePoint)
}

// HashCodePointImmediateValue is a free data retrieval call binding the contract method 0x264f384b.
//
// Solidity: function hashCodePointImmediateValue(uint8 opcode, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueCaller) HashCodePointImmediateValue(opts *bind.CallOpts, opcode uint8, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbValue.contract.Call(opts, out, "hashCodePointImmediateValue", opcode, immediateVal, nextCodePoint)
	return *ret0, err
}

// HashCodePointImmediateValue is a free data retrieval call binding the contract method 0x264f384b.
//
// Solidity: function hashCodePointImmediateValue(uint8 opcode, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueSession) HashCodePointImmediateValue(opcode uint8, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	return _ArbValue.Contract.HashCodePointImmediateValue(&_ArbValue.CallOpts, opcode, immediateVal, nextCodePoint)
}

// HashCodePointImmediateValue is a free data retrieval call binding the contract method 0x264f384b.
//
// Solidity: function hashCodePointImmediateValue(uint8 opcode, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) HashCodePointImmediateValue(opcode uint8, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	return _ArbValue.Contract.HashCodePointImmediateValue(&_ArbValue.CallOpts, opcode, immediateVal, nextCodePoint)
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() constant returns(bytes32)
func (_ArbValue *ArbValueCaller) HashEmptyTuple(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbValue.contract.Call(opts, out, "hashEmptyTuple")
	return *ret0, err
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() constant returns(bytes32)
func (_ArbValue *ArbValueSession) HashEmptyTuple() ([32]byte, error) {
	return _ArbValue.Contract.HashEmptyTuple(&_ArbValue.CallOpts)
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) HashEmptyTuple() ([32]byte, error) {
	return _ArbValue.Contract.HashEmptyTuple(&_ArbValue.CallOpts)
}

// HashIntValue is a free data retrieval call binding the contract method 0x1667b411.
//
// Solidity: function hashIntValue(uint256 val) constant returns(bytes32)
func (_ArbValue *ArbValueCaller) HashIntValue(opts *bind.CallOpts, val *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbValue.contract.Call(opts, out, "hashIntValue", val)
	return *ret0, err
}

// HashIntValue is a free data retrieval call binding the contract method 0x1667b411.
//
// Solidity: function hashIntValue(uint256 val) constant returns(bytes32)
func (_ArbValue *ArbValueSession) HashIntValue(val *big.Int) ([32]byte, error) {
	return _ArbValue.Contract.HashIntValue(&_ArbValue.CallOpts, val)
}

// HashIntValue is a free data retrieval call binding the contract method 0x1667b411.
//
// Solidity: function hashIntValue(uint256 val) constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) HashIntValue(val *big.Int) ([32]byte, error) {
	return _ArbValue.Contract.HashIntValue(&_ArbValue.CallOpts, val)
}

// IsValidTupleSize is a free data retrieval call binding the contract method 0xb2b9dc62.
//
// Solidity: function isValidTupleSize(uint256 size) constant returns(bool)
func (_ArbValue *ArbValueCaller) IsValidTupleSize(opts *bind.CallOpts, size *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ArbValue.contract.Call(opts, out, "isValidTupleSize", size)
	return *ret0, err
}

// IsValidTupleSize is a free data retrieval call binding the contract method 0xb2b9dc62.
//
// Solidity: function isValidTupleSize(uint256 size) constant returns(bool)
func (_ArbValue *ArbValueSession) IsValidTupleSize(size *big.Int) (bool, error) {
	return _ArbValue.Contract.IsValidTupleSize(&_ArbValue.CallOpts, size)
}

// IsValidTupleSize is a free data retrieval call binding the contract method 0xb2b9dc62.
//
// Solidity: function isValidTupleSize(uint256 size) constant returns(bool)
func (_ArbValue *ArbValueCallerSession) IsValidTupleSize(size *big.Int) (bool, error) {
	return _ArbValue.Contract.IsValidTupleSize(&_ArbValue.CallOpts, size)
}

// BisectionABI is the input ABI used to generate the binding from.
const BisectionABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"assertionIndex\",\"type\":\"uint256\"}],\"name\":\"ContinuedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"bisecter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"afterHashAndMessageAndLogsBisections\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"name\":\"totalSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"totalMessageAmounts\",\"type\":\"uint256[]\"}],\"name\":\"BisectedAssertion\",\"type\":\"event\"}]"

// BisectionFuncSigs maps the 4-byte function signature to its string representation.
var BisectionFuncSigs = map[string]string{
	"1c81699c": "bisectAssertion(Bisection.Challenge storage,bytes32[2],bytes32[],uint256[],uint32,uint64[2],bytes21[],uint256[])",
	"04862117": "continueChallenge(Bisection.Challenge storage,uint256,bytes,bytes32,bytes32)",
}

// BisectionBin is the compiled bytecode used for deploying new contracts.
var BisectionBin = "0x61144d610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100405760003560e01c806304862117146100455780631c81699c14610109575b600080fd5b81801561005157600080fd5b50610107600480360360a081101561006857600080fd5b813591602081013591810190606081016040820135600160201b81111561008e57600080fd5b8201836020820111156100a057600080fd5b803590602001918460018302840111600160201b831117156100c157600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505050602001356103aa565b005b81801561011557600080fd5b50610107600480360361014081101561012d57600080fd5b60408051808201825283359392830192916060830191906020840190600290839083908082843760009201919091525091949392602081019250359050600160201b81111561017b57600080fd5b82018360208201111561018d57600080fd5b803590602001918460208302840111600160201b831117156101ae57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156101fd57600080fd5b82018360208201111561020f57600080fd5b803590602001918460208302840111600160201b8311171561023057600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080518082018252939663ffffffff8635169690959094606082019450925060200190600290839083908082843760009201919091525091949392602081019250359050600160201b8111156102b757600080fd5b8201836020820111156102c957600080fd5b803590602001918460208302840111600160201b831117156102ea57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561033957600080fd5b82018360208201111561034b57600080fd5b803590602001918460208302840111600160201b8311171561036c57600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061067f945050505050565b846001015482146103ec5760405162461bcd60e51b815260040180806020018281038252602b8152602001806113bf602b913960400191505060405180910390fd5b600585015467ffffffffffffffff1643111561044f576040805162461bcd60e51b815260206004820152601a60248201527f4368616c6c656e676520646561646c696e652065787069726564000000000000604482015290519081900360640190fd5b60038501600101546001600160a01b0316331461049d5760405162461bcd60e51b815260040180806020018281038252602f8152602001806113ea602f913960400191505060405180910390fd5b73__$a50780cb42d41d2927e39f529dc62d6697$__63b792d767848484886001016040518563ffffffff1660e01b81526004018080602001858152602001848152602001838152602001828103825286818151815260200191508051906020019080838360005b8381101561051c578181015183820152602001610504565b50505050905090810190601f1680156105495780820380516001836020036101000a031916815260200191505b509550505050505060206040518083038186803b15801561056957600080fd5b505af415801561057d573d6000803e3d6000fd5b505050506040513d602081101561059357600080fd5b50516105e6576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c696420617373657274696f6e2073656c6563746564000000000000604482015290519081900360640190fd5b60058501805467ffffffffffffffff4363ffffffff600160401b60ff60601b19909416600160601b1793840416011667ffffffffffffffff19919091161790556001850181905584546004860154604080516001600160a01b0390921682526020820187905280517ff5a6d1468c6ce7a03663fa2fd47dad9f6693a9289ebdbcf3309941caa75c7cf09281900390910190a25050505050565b60016005890154600160601b900460ff16600281111561069b57fe5b146106d75760405162461bcd60e51b815260040180806020018281038252603481526020018061138b6034913960400191505060405180910390fd5b815115806106ee575081518551816106eb57fe5b06155b610738576040805162461bcd60e51b8152602060048201526016602482015275092dcc6dee4e4cac6e840d2dce0eae840d8cadccee8d60531b604482015290519081900360640190fd5b8151158061075e5750815185518161074c57fe5b04600360018851038161075b57fe5b04145b6107a8576040805162461bcd60e51b8152602060048201526016602482015275092dcc6dee4e4cac6e840d2dce0eae840d8cadccee8d60531b604482015290519081900360640190fd5b80518251146107f7576040805162461bcd60e51b8152602060048201526016602482015275092dcc6dee4e4cac6e840d2dce0eae840d8cadccee8d60531b604482015290519081900360640190fd5b600588015467ffffffffffffffff1643111561085a576040805162461bcd60e51b815260206004820152601a60248201527f4368616c6c656e676520646561646c696e652065787069726564000000000000604482015290519081900360640190fd5b60038801600001546001600160a01b031633146108be576040805162461bcd60e51b815260206004820181905260248201527f4f6e6c79206f7269676e616c2061737365727465722063616e20626973656374604482015290519081900360640190fd5b60006060610945604051806101200160405280600360018c5103816108df57fe5b0463ffffffff1681526020018a81526020018981526020018863ffffffff1681526020018b60006002811061091057fe5b602002015181526020018781526020018681526020018581526020018b60016002811061093957fe5b60200201519052610b90565b60018c0154919350915082146109a2576040805162461bcd60e51b815260206004820152601960248201527f446f6573206e6f74206d61746368207072657620737461746500000000000000604482015290519081900360640190fd5b60058a01805467ffffffffffffffff4363ffffffff600160401b60ff60601b19909416600160611b1793840416011667ffffffffffffffff19919091161790556040516309898dc160e41b815260206004820181815283516024840152835173__$a50780cb42d41d2927e39f529dc62d6697$__93639898dc1093869392839260440191858101910280838360005b83811015610a49578181015183820152602001610a31565b505050509050019250505060206040518083038186803b158015610a6c57600080fd5b505af4158015610a80573d6000803e3d6000fd5b505050506040513d6020811015610a9657600080fd5b505160018b0155895460038b0154604080516001600160a01b0390921680835263ffffffff8a1691830191909152608060208084018281528d51928501929092528c517fa4f8cbbd195d8e69d66b332eb24e05c24a35ee00de81da472b89ecb42a70ef71948e938d938f93606084019160a0850191818901910280838360005b83811015610b2e578181015183820152602001610b16565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015610b6d578181015183820152602001610b55565b50505050905001965050505050505060405180910390a250505050505050505050565b60006060610b9c611356565b836000015163ffffffff16604051908082528060200260200182016040528015610bd0578160200160208202803883390190505b508160600181905250836000015163ffffffff16846060015163ffffffff1681610bf657fe5b0460010163ffffffff1660a0820152608084015181526000805b855163ffffffff1681101561133f57856000015163ffffffff16866060015163ffffffff1681610c3c57fe5b0663ffffffff16811415610c5d5760a0830180516000190163ffffffff1690525b8560c0015151604051908082528060200260200182016040528015610c8c578160200160208202803883390190505b506080840152600091505b8560c0015151821015610cf1578560400151828760c001515183020181518110610cbd57fe5b602002602001015183608001518381518110610cd557fe5b6020908102919091010180519091019052600190910190610c97565b73__$6b4cc75dad3e0abd6ad83b3d907747c608$__633e28559884600001518860a001518961010001518a60c001518b60e001516040518663ffffffff1660e01b81526004018086815260200185600260200280838360005b83811015610d62578181015183820152602001610d4a565b505050509050018481526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b83811015610daf578181015183820152602001610d97565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015610dee578181015183820152602001610dd6565b5050505090500197505050505050505060206040518083038186803b158015610e1657600080fd5b505af4158015610e2a573d6000803e3d6000fd5b505050506040513d6020811015610e4057600080fd5b50516020840152600091505b8560c0015151821015610e9e5782608001518281518110610e6957fe5b60200260200101518660e001518381518110610e8157fe5b602090810291909101018051919091039052600190910190610e4c565b826020015173__$6b4cc75dad3e0abd6ad83b3d907747c608$__632090372188602001518481518110610ecd57fe5b60200260200101518660a001518a60200151868c6000015163ffffffff160181518110610ef657fe5b60200260200101518b60200151878d6000015163ffffffff160160010181518110610f1d57fe5b60200260200101518c60200151888e6000015160020260010163ffffffff160181518110610f4757fe5b60200260200101518d60200151898f6000015160020260020163ffffffff160181518110610f7157fe5b60200260200101518b608001516040518863ffffffff1660e01b8152600401808881526020018763ffffffff1663ffffffff16815260200186815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019060200280838360005b83811015610ffd578181015183820152602001610fe5565b505050509050019850505050505050505060206040518083038186803b15801561102657600080fd5b505af415801561103a573d6000803e3d6000fd5b505050506040513d602081101561105057600080fd5b50516040805160208181019490945280820192909252805180830382018152606092830190915280519201919091209084015180518390811061108f57fe5b602002602001018181525050856020015181815181106110ab57fe5b6020908102919091010151835280611337578560c00151516040519080825280602002602001820160405280156110ec578160200160208202803883390190505b506080840152600091505b856040015151821015611156578560400151828151811061111457fe5b602002602001015183608001518760c0015151848161112f57fe5b068151811061113a57fe5b60209081029190910101805190910190526001909101906110f7565b826020015173__$6b4cc75dad3e0abd6ad83b3d907747c608$__6320903721886020015160018a600001510363ffffffff168151811061119257fe5b602002602001015189606001518a602001518b6000015163ffffffff16815181106111b957fe5b60200260200101518b602001518c6000015160020263ffffffff16815181106111de57fe5b60200260200101518c602001518d6000015160020260010163ffffffff168151811061120657fe5b60200260200101518d6020015160018f6020015151038151811061122657fe5b60200260200101518b608001516040518863ffffffff1660e01b8152600401808881526020018763ffffffff1663ffffffff16815260200186815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019060200280838360005b838110156112b257818101518382015260200161129a565b505050509050019850505050505050505060206040518083038186803b1580156112db57600080fd5b505af41580156112ef573d6000803e3d6000fd5b505050506040513d602081101561130557600080fd5b505160408051602081810194909452808201929092528051808303820181526060909201815281519190920120908401525b600101610c10565b505060408101516060909101519092509050915091565b6040805160c0810182526000808252602082018190529181018290526060808201819052608082015260a08101919091529056fe43616e206f6e6c792062697365637420617373657274696f6e20696e20726573706f6e736520746f2061206368616c6c656e6765636f6e74696e75654368616c6c656e67653a20496e636f72726563742070726576696f75732073746174654f6e6c79206f726967696e616c206368616c6c656e6765722063616e20636f6e74696e7565206368616c6c656e6765a265627a7a723058203a2dad7d607c5e1753de54b205235bd1f56dd9bc8c43d35369103da8a68c919664736f6c634300050a0032"

// DeployBisection deploys a new Ethereum contract, binding an instance of Bisection to it.
func DeployBisection(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bisection, error) {
	parsed, err := abi.JSON(strings.NewReader(BisectionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	arbProtocolAddr, _, _, _ := DeployArbProtocol(auth, backend)
	BisectionBin = strings.Replace(BisectionBin, "__$6b4cc75dad3e0abd6ad83b3d907747c608$__", arbProtocolAddr.String()[2:], -1)

	merkleLibAddr, _, _, _ := DeployMerkleLib(auth, backend)
	BisectionBin = strings.Replace(BisectionBin, "__$a50780cb42d41d2927e39f529dc62d6697$__", merkleLibAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BisectionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bisection{BisectionCaller: BisectionCaller{contract: contract}, BisectionTransactor: BisectionTransactor{contract: contract}, BisectionFilterer: BisectionFilterer{contract: contract}}, nil
}

// Bisection is an auto generated Go binding around an Ethereum contract.
type Bisection struct {
	BisectionCaller     // Read-only binding to the contract
	BisectionTransactor // Write-only binding to the contract
	BisectionFilterer   // Log filterer for contract events
}

// BisectionCaller is an auto generated read-only Go binding around an Ethereum contract.
type BisectionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BisectionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BisectionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BisectionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BisectionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BisectionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BisectionSession struct {
	Contract     *Bisection        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BisectionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BisectionCallerSession struct {
	Contract *BisectionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BisectionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BisectionTransactorSession struct {
	Contract     *BisectionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BisectionRaw is an auto generated low-level Go binding around an Ethereum contract.
type BisectionRaw struct {
	Contract *Bisection // Generic contract binding to access the raw methods on
}

// BisectionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BisectionCallerRaw struct {
	Contract *BisectionCaller // Generic read-only contract binding to access the raw methods on
}

// BisectionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BisectionTransactorRaw struct {
	Contract *BisectionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBisection creates a new instance of Bisection, bound to a specific deployed contract.
func NewBisection(address common.Address, backend bind.ContractBackend) (*Bisection, error) {
	contract, err := bindBisection(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bisection{BisectionCaller: BisectionCaller{contract: contract}, BisectionTransactor: BisectionTransactor{contract: contract}, BisectionFilterer: BisectionFilterer{contract: contract}}, nil
}

// NewBisectionCaller creates a new read-only instance of Bisection, bound to a specific deployed contract.
func NewBisectionCaller(address common.Address, caller bind.ContractCaller) (*BisectionCaller, error) {
	contract, err := bindBisection(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BisectionCaller{contract: contract}, nil
}

// NewBisectionTransactor creates a new write-only instance of Bisection, bound to a specific deployed contract.
func NewBisectionTransactor(address common.Address, transactor bind.ContractTransactor) (*BisectionTransactor, error) {
	contract, err := bindBisection(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BisectionTransactor{contract: contract}, nil
}

// NewBisectionFilterer creates a new log filterer instance of Bisection, bound to a specific deployed contract.
func NewBisectionFilterer(address common.Address, filterer bind.ContractFilterer) (*BisectionFilterer, error) {
	contract, err := bindBisection(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BisectionFilterer{contract: contract}, nil
}

// bindBisection binds a generic wrapper to an already deployed contract.
func bindBisection(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BisectionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bisection *BisectionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bisection.Contract.BisectionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bisection *BisectionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bisection.Contract.BisectionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bisection *BisectionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bisection.Contract.BisectionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bisection *BisectionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bisection.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bisection *BisectionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bisection.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bisection *BisectionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bisection.Contract.contract.Transact(opts, method, params...)
}

// BisectionBisectedAssertionIterator is returned from FilterBisectedAssertion and is used to iterate over the raw logs and unpacked data for BisectedAssertion events raised by the Bisection contract.
type BisectionBisectedAssertionIterator struct {
	Event *BisectionBisectedAssertion // Event containing the contract specifics and raw log

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
func (it *BisectionBisectedAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BisectionBisectedAssertion)
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
		it.Event = new(BisectionBisectedAssertion)
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
func (it *BisectionBisectedAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BisectionBisectedAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BisectionBisectedAssertion represents a BisectedAssertion event raised by the Bisection contract.
type BisectionBisectedAssertion struct {
	VmId                                 [32]byte
	Bisecter                             common.Address
	AfterHashAndMessageAndLogsBisections [][32]byte
	TotalSteps                           uint32
	TotalMessageAmounts                  []*big.Int
	Raw                                  types.Log // Blockchain specific contextual infos
}

// FilterBisectedAssertion is a free log retrieval operation binding the contract event 0xa4f8cbbd195d8e69d66b332eb24e05c24a35ee00de81da472b89ecb42a70ef71.
//
// Solidity: event BisectedAssertion(bytes32 indexed vmId, address bisecter, bytes32[] afterHashAndMessageAndLogsBisections, uint32 totalSteps, uint256[] totalMessageAmounts)
func (_Bisection *BisectionFilterer) FilterBisectedAssertion(opts *bind.FilterOpts, vmId [][32]byte) (*BisectionBisectedAssertionIterator, error) {

	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _Bisection.contract.FilterLogs(opts, "BisectedAssertion", vmIdRule)
	if err != nil {
		return nil, err
	}
	return &BisectionBisectedAssertionIterator{contract: _Bisection.contract, event: "BisectedAssertion", logs: logs, sub: sub}, nil
}

// WatchBisectedAssertion is a free log subscription operation binding the contract event 0xa4f8cbbd195d8e69d66b332eb24e05c24a35ee00de81da472b89ecb42a70ef71.
//
// Solidity: event BisectedAssertion(bytes32 indexed vmId, address bisecter, bytes32[] afterHashAndMessageAndLogsBisections, uint32 totalSteps, uint256[] totalMessageAmounts)
func (_Bisection *BisectionFilterer) WatchBisectedAssertion(opts *bind.WatchOpts, sink chan<- *BisectionBisectedAssertion, vmId [][32]byte) (event.Subscription, error) {

	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _Bisection.contract.WatchLogs(opts, "BisectedAssertion", vmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BisectionBisectedAssertion)
				if err := _Bisection.contract.UnpackLog(event, "BisectedAssertion", log); err != nil {
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

// ParseBisectedAssertion is a log parse operation binding the contract event 0xa4f8cbbd195d8e69d66b332eb24e05c24a35ee00de81da472b89ecb42a70ef71.
//
// Solidity: event BisectedAssertion(bytes32 indexed vmId, address bisecter, bytes32[] afterHashAndMessageAndLogsBisections, uint32 totalSteps, uint256[] totalMessageAmounts)
func (_Bisection *BisectionFilterer) ParseBisectedAssertion(log types.Log) (*BisectionBisectedAssertion, error) {
	event := new(BisectionBisectedAssertion)
	if err := _Bisection.contract.UnpackLog(event, "BisectedAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BisectionContinuedChallengeIterator is returned from FilterContinuedChallenge and is used to iterate over the raw logs and unpacked data for ContinuedChallenge events raised by the Bisection contract.
type BisectionContinuedChallengeIterator struct {
	Event *BisectionContinuedChallenge // Event containing the contract specifics and raw log

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
func (it *BisectionContinuedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BisectionContinuedChallenge)
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
		it.Event = new(BisectionContinuedChallenge)
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
func (it *BisectionContinuedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BisectionContinuedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BisectionContinuedChallenge represents a ContinuedChallenge event raised by the Bisection contract.
type BisectionContinuedChallenge struct {
	VmId           [32]byte
	Challenger     common.Address
	AssertionIndex *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterContinuedChallenge is a free log retrieval operation binding the contract event 0xf5a6d1468c6ce7a03663fa2fd47dad9f6693a9289ebdbcf3309941caa75c7cf0.
//
// Solidity: event ContinuedChallenge(bytes32 indexed vmId, address challenger, uint256 assertionIndex)
func (_Bisection *BisectionFilterer) FilterContinuedChallenge(opts *bind.FilterOpts, vmId [][32]byte) (*BisectionContinuedChallengeIterator, error) {

	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _Bisection.contract.FilterLogs(opts, "ContinuedChallenge", vmIdRule)
	if err != nil {
		return nil, err
	}
	return &BisectionContinuedChallengeIterator{contract: _Bisection.contract, event: "ContinuedChallenge", logs: logs, sub: sub}, nil
}

// WatchContinuedChallenge is a free log subscription operation binding the contract event 0xf5a6d1468c6ce7a03663fa2fd47dad9f6693a9289ebdbcf3309941caa75c7cf0.
//
// Solidity: event ContinuedChallenge(bytes32 indexed vmId, address challenger, uint256 assertionIndex)
func (_Bisection *BisectionFilterer) WatchContinuedChallenge(opts *bind.WatchOpts, sink chan<- *BisectionContinuedChallenge, vmId [][32]byte) (event.Subscription, error) {

	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _Bisection.contract.WatchLogs(opts, "ContinuedChallenge", vmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BisectionContinuedChallenge)
				if err := _Bisection.contract.UnpackLog(event, "ContinuedChallenge", log); err != nil {
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

// ParseContinuedChallenge is a log parse operation binding the contract event 0xf5a6d1468c6ce7a03663fa2fd47dad9f6693a9289ebdbcf3309941caa75c7cf0.
//
// Solidity: event ContinuedChallenge(bytes32 indexed vmId, address challenger, uint256 assertionIndex)
func (_Bisection *BisectionFilterer) ParseContinuedChallenge(log types.Log) (*BisectionContinuedChallenge, error) {
	event := new(BisectionContinuedChallenge)
	if err := _Bisection.contract.UnpackLog(event, "ContinuedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BytesLibABI is the input ABI used to generate the binding from.
const BytesLibABI = "[]"

// BytesLibBin is the compiled bytecode used for deploying new contracts.
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723058202a4b33a7382bd41fc01a7b408d0a3491e17cc0d680148cb3a692e8a4446bc59d64736f6c634300050a0032"

// DeployBytesLib deploys a new Ethereum contract, binding an instance of BytesLib to it.
func DeployBytesLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BytesLib, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BytesLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// BytesLib is an auto generated Go binding around an Ethereum contract.
type BytesLib struct {
	BytesLibCaller     // Read-only binding to the contract
	BytesLibTransactor // Write-only binding to the contract
	BytesLibFilterer   // Log filterer for contract events
}

// BytesLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type BytesLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BytesLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BytesLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BytesLibSession struct {
	Contract     *BytesLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BytesLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BytesLibCallerSession struct {
	Contract *BytesLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BytesLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BytesLibTransactorSession struct {
	Contract     *BytesLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BytesLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type BytesLibRaw struct {
	Contract *BytesLib // Generic contract binding to access the raw methods on
}

// BytesLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BytesLibCallerRaw struct {
	Contract *BytesLibCaller // Generic read-only contract binding to access the raw methods on
}

// BytesLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BytesLibTransactorRaw struct {
	Contract *BytesLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBytesLib creates a new instance of BytesLib, bound to a specific deployed contract.
func NewBytesLib(address common.Address, backend bind.ContractBackend) (*BytesLib, error) {
	contract, err := bindBytesLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// NewBytesLibCaller creates a new read-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibCaller(address common.Address, caller bind.ContractCaller) (*BytesLibCaller, error) {
	contract, err := bindBytesLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibCaller{contract: contract}, nil
}

// NewBytesLibTransactor creates a new write-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibTransactor(address common.Address, transactor bind.ContractTransactor) (*BytesLibTransactor, error) {
	contract, err := bindBytesLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibTransactor{contract: contract}, nil
}

// NewBytesLibFilterer creates a new log filterer instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibFilterer(address common.Address, filterer bind.ContractFilterer) (*BytesLibFilterer, error) {
	contract, err := bindBytesLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BytesLibFilterer{contract: contract}, nil
}

// bindBytesLib binds a generic wrapper to an already deployed contract.
func bindBytesLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.BytesLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transact(opts, method, params...)
}

// ChallengeManagerABI is the input ABI used to generate the binding from.
const ChallengeManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"bytes32\"},{\"name\":\"_players\",\"type\":\"address[2]\"},{\"name\":\"_escrows\",\"type\":\"uint128[2]\"},{\"name\":\"_challengePeriod\",\"type\":\"uint32\"},{\"name\":\"_challengeRoot\",\"type\":\"bytes32\"}],\"name\":\"initiateChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"bytes32\"}],\"name\":\"challengerTimedOut\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_challengeId\",\"type\":\"bytes32\"},{\"name\":\"_fields\",\"type\":\"bytes32[2]\"},{\"name\":\"_afterHashAndMessageAndLogsBisections\",\"type\":\"bytes32[]\"},{\"name\":\"_totalMessageAmounts\",\"type\":\"uint256[]\"},{\"name\":\"_totalSteps\",\"type\":\"uint32\"},{\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_beforeBalances\",\"type\":\"uint256[]\"}],\"name\":\"bisectAssertion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"bytes32\"},{\"name\":\"_assertionToChallenge\",\"type\":\"uint256\"},{\"name\":\"_proof\",\"type\":\"bytes\"},{\"name\":\"_bisectionRoot\",\"type\":\"bytes32\"},{\"name\":\"_bisectionHash\",\"type\":\"bytes32\"}],\"name\":\"continueChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"bytes32\"},{\"name\":\"_beforeHashAndInbox\",\"type\":\"bytes32[2]\"},{\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_beforeBalances\",\"type\":\"uint256[]\"},{\"name\":\"_afterHashAndMessages\",\"type\":\"bytes32[5]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"oneStepProof\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"bytes32\"}],\"name\":\"asserterTimedOut\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_vmTracker\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"assertionIndex\",\"type\":\"uint256\"}],\"name\":\"ContinuedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"bisecter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"afterHashAndMessageAndLogsBisections\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"name\":\"totalSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"totalMessageAmounts\",\"type\":\"uint256[]\"}],\"name\":\"BisectedAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"OneStepProofCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"challengerWrong\",\"type\":\"bool\"}],\"name\":\"TimedOutChallenge\",\"type\":\"event\"}]"

// ChallengeManagerFuncSigs maps the 4-byte function signature to its string representation.
var ChallengeManagerFuncSigs = map[string]string{
	"b814f8c7": "asserterTimedOut(bytes32)",
	"70027cbd": "bisectAssertion(bytes32,bytes32[2],bytes32[],uint256[],uint32,uint64[2],bytes21[],uint256[])",
	"4db0228b": "challengerTimedOut(bytes32)",
	"770a2e56": "continueChallenge(bytes32,uint256,bytes,bytes32,bytes32)",
	"2b50d42b": "initiateChallenge(bytes32,address[2],uint128[2],uint32,bytes32)",
	"832dcf57": "oneStepProof(bytes32,bytes32[2],uint64[2],bytes21[],uint256[],bytes32[5],uint256[],bytes)",
}

// ChallengeManagerBin is the compiled bytecode used for deploying new contracts.
var ChallengeManagerBin = "0x608060405234801561001057600080fd5b506040516119d13803806119d18339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b031990921691909117905561196c806100656000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80632b50d42b146100675780634db0228b146100a257806370027cbd146100bf578063770a2e5614610353578063832dcf5714610408578063b814f8c7146106bc575b600080fd5b6100a0600480360360e081101561007d57600080fd5b508035906020810190606081019063ffffffff60a0820135169060c001356106d9565b005b6100a0600480360360208110156100b857600080fd5b50356108d2565b6100a060048036036101408110156100d657600080fd5b60408051808201825283359392830192916060830191906020840190600290839083908082843760009201919091525091949392602081019250359050600160201b81111561012457600080fd5b82018360208201111561013657600080fd5b803590602001918460208302840111600160201b8311171561015757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156101a657600080fd5b8201836020820111156101b857600080fd5b803590602001918460208302840111600160201b831117156101d957600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080518082018252939663ffffffff8635169690959094606082019450925060200190600290839083908082843760009201919091525091949392602081019250359050600160201b81111561026057600080fd5b82018360208201111561027257600080fd5b803590602001918460208302840111600160201b8311171561029357600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156102e257600080fd5b8201836020820111156102f457600080fd5b803590602001918460208302840111600160201b8311171561031557600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506109d9945050505050565b6100a0600480360360a081101561036957600080fd5b813591602081013591810190606081016040820135600160201b81111561038f57600080fd5b8201836020820111156103a157600080fd5b803590602001918460018302840111600160201b831117156103c257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060200135610be8565b6100a060048036036101c081101561041f57600080fd5b604080518082018252833593928301929160608301919060208401906002908390839080828437600092019190915250506040805180820182529295949381810193925090600290839083908082843760009201919091525091949392602081019250359050600160201b81111561049657600080fd5b8201836020820111156104a857600080fd5b803590602001918460208302840111600160201b831117156104c957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561051857600080fd5b82018360208201111561052a57600080fd5b803590602001918460208302840111600160201b8311171561054b57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250506040805160a0818101909252939695948181019493509150600590839083908082843760009201919091525091949392602081019250359050600160201b8111156105c657600080fd5b8201836020820111156105d857600080fd5b803590602001918460208302840111600160201b831117156105f957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561064857600080fd5b82018360208201111561065a57600080fd5b803590602001918460018302840111600160201b8311171561067b57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610cf0945050505050565b6100a0600480360360208110156106d257600080fd5b503561144a565b6000546001600160a01b031633146107225760405162461bcd60e51b815260040180806020018281038252602e815260200180611850602e913960400191505060405180910390fd5b60008581526001602081905260409091200154156107715760405162461bcd60e51b81526004018080602001828103825260238152602001806119156023913960400191505060405180910390fd5b6040518060e001604052808681526020018281526020018460028060200260405190810160405280929190826002602002808284376000920191909152505050815260408051808201825260209092019190879060029083908390808284376000920191909152505050815263ffffffff841643810167ffffffffffffffff1660208301526040820152606001600190526000868152600160208181526040928390208451815590840151918101919091559082015161083790600280840191906116dc565b50606082015161084d9060038301906002611781565b50608082015160058201805460a085015163ffffffff1668010000000000000000026bffffffff00000000000000001967ffffffffffffffff90941667ffffffffffffffff1990921691909117929092169190911780825560c0840151919060ff60601b1916600160601b8360028111156108c457fe5b021790555050505050505050565b600081815260016020526040902060026005820154600160601b900460ff1660028111156108fc57fe5b146109385760405162461bcd60e51b81526004018080602001828103825260308152602001806118ac6030913960400191505060405180910390fd5b600581015467ffffffffffffffff164311610994576040805162461bcd60e51b8152602060048201526017602482015276111958591b1a5b99481a185cdb89dd08195e1c1a5c9959604a1b604482015290519081900360640190fd5b61099e8282611553565b6040805160008152905183917f4bb5367552b51194ffa95563d907630acc956760789c62c6e92b82dfddf7e680919081900360200190a25050565b6000600160008a8152602001908152602001600020905073__$0a962fcb953632d8e24d6f165261b09420$__631c81699c828a8a8a8a8a8a8a6040518963ffffffff1660e01b81526004018089815260200188600260200280838360005b83811015610a4f578181015183820152602001610a37565b5050505090500180602001806020018763ffffffff1663ffffffff16815260200186600260200280838360005b83811015610a94578181015183820152602001610a7c565b50505050905001806020018060200185810385528b818151815260200191508051906020019060200280838360005b83811015610adb578181015183820152602001610ac3565b5050505090500185810384528a818151815260200191508051906020019060200280838360005b83811015610b1a578181015183820152602001610b02565b50505050905001858103835287818151815260200191508051906020019060200280838360005b83811015610b59578181015183820152602001610b41565b50505050905001858103825286818151815260200191508051906020019060200280838360005b83811015610b98578181015183820152602001610b80565b505050509050019c5050505050505050505050505060006040518083038186803b158015610bc557600080fd5b505af4158015610bd9573d6000803e3d6000fd5b50505050505050505050505050565b600060016000878152602001908152602001600020905073__$0a962fcb953632d8e24d6f165261b09420$__630486211782878787876040518663ffffffff1660e01b81526004018086815260200185815260200180602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b83811015610c82578181015183820152602001610c6a565b50505050905090810190601f168015610caf5780820380516001836020036101000a031916815260200191505b50965050505050505060006040518083038186803b158015610cd057600080fd5b505af4158015610ce4573d6000803e3d6000fd5b50505050505050505050565b6000888152600160208190526040909120906005820154600160601b900460ff166002811115610d1c57fe5b14610d585760405162461bcd60e51b81526004018080602001828103825260398152602001806118dc6039913960400191505060405180910390fd5b600581015467ffffffffffffffff16431115610dbb576040805162461bcd60e51b815260206004820152601e60248201527f4f6e6520737465702070726f6f66206d697373656420646561646c696e650000604482015290519081900360640190fd5b600181015473__$6b4cc75dad3e0abd6ad83b3d907747c608$__633e2855988a600060200201518a8c600160200201518b8b6040518663ffffffff1660e01b81526004018086815260200185600260200280838360005b83811015610e2a578181015183820152602001610e12565b505050509050018481526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b83811015610e77578181015183820152602001610e5f565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015610eb6578181015183820152602001610e9e565b5050505090500197505050505050505060206040518083038186803b158015610ede57600080fd5b505af4158015610ef2573d6000803e3d6000fd5b505050506040513d6020811015610f0857600080fd5b505173__$6b4cc75dad3e0abd6ad83b3d907747c608$__632090372187600060200201516001898160200201518a600260200201518b600360200201518c600460200201518c6040518863ffffffff1660e01b8152600401808881526020018763ffffffff16815260200186815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019060200280838360005b83811015610fc7578181015183820152602001610faf565b505050509050019850505050505050505060206040518083038186803b158015610ff057600080fd5b505af4158015611004573d6000803e3d6000fd5b505050506040513d602081101561101a57600080fd5b50516040805160208181019490945280820192909252805180830382018152606090920190528051910120146110815760405162461bcd60e51b815260040180806020018281038252602681526020018061182a6026913960400191505060405180910390fd5b600073__$98107b176d4310ec680b0534b46d40334a$__630eca9f136040518060e001604052808c6000600281106110b557fe5b602002015181526020018c6001600281106110cc57fe5b60200201518152602001886000600581106110e357fe5b60200201518152602001886001600581106110fa57fe5b602002015181526020018860026005811061111157fe5b602002015181526020018860036005811061112857fe5b602002015181526020018860046005811061113f57fe5b60200201518152508a8a8a89896040518763ffffffff1660e01b81526004018087600760200280838360005b8381101561118357818101518382015260200161116b565b5050505090500186600260200280838360005b838110156111ae578181015183820152602001611196565b5050505090500180602001806020018060200180602001858103855289818151815260200191508051906020019060200280838360005b838110156111fd5781810151838201526020016111e5565b50505050905001858103845288818151815260200191508051906020019060200280838360005b8381101561123c578181015183820152602001611224565b50505050905001858103835287818151815260200191508051906020019060200280838360005b8381101561127b578181015183820152602001611263565b50505050905001858103825286818151815260200191508051906020019080838360005b838110156112b757818101518382015260200161129f565b50505050905090810190601f1680156112e45780820380516001836020036101000a031916815260200191505b509a505050505050505050505060206040518083038186803b15801561130957600080fd5b505af415801561131d573d6000803e3d6000fd5b505050506040513d602081101561133357600080fd5b505190508015611380576040805162461bcd60e51b8152602060048201526013602482015272141c9bdbd9881dd85cc81a5b98dbdc9c9958dd606a1b604482015290519081900360640190fd5b61138a8a83611553565b897fb3b8389d4e8e450396d43f12a728c7a73784c6089340d0e904a24fd366ddaaff338560405180836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b838110156114035781810151838201526020016113eb565b50505050905090810190601f1680156114305780820380516001836020036101000a031916815260200191505b50935050505060405180910390a250505050505050505050565b6000818152600160208190526040909120906005820154600160601b900460ff16600281111561147657fe5b146114b25760405162461bcd60e51b815260040180806020018281038252602e81526020018061187e602e913960400191505060405180910390fd5b600581015467ffffffffffffffff16431161150e576040805162461bcd60e51b8152602060048201526017602482015276111958591b1a5b99481a185cdb89dd08195e1c1a5c9959604a1b604482015290519081900360640190fd5b6115188282611646565b6040805160018152905183917f4bb5367552b51194ffa95563d907630acc956760789c62c6e92b82dfddf7e680919081900360200190a25050565b600080546040805180820182526002858101546001600160801b03808216600160801b909204811692909204011681526020810193909352516363d8463760e01b8152600481018581526001600160a01b03909216926363d8463792869260038701929190606481019084906024015b81546001600160a01b031681526001909101906020018083116115c35750839050604080838360005b838110156116045781810151838201526020016115ec565b505050509050019350505050600060405180830381600087803b15801561162a57600080fd5b505af115801561163e573d6000803e3d6000fd5b505050505050565b60008054604080518082018252928352600284810154600160801b81046001600160801b039081169181169290920401166020840152516363d8463760e01b815260048082018681526003860180546001600160a01b039081166024860152909416946363d8463794889490939192916064820191908901906044018083116115c35750508251815282604080838360206115ec565b6001830191839082156117715791602002820160005b8382111561173c57835183826101000a8154816001600160801b0302191690836001600160801b031602179055509260200192601001602081600f010492830192600103026116f2565b801561176f5782816101000a8154906001600160801b030219169055601001602081600f0104928301926001030261173c565b505b5061177d9291506117d5565b5090565b82600281019282156117c9579160200282015b828111156117c957825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190611794565b5061177d929150611805565b61180291905b8082111561177d5780546fffffffffffffffffffffffffffffffff191681556001016117db565b90565b61180291905b8082111561177d5780546001600160a01b031916815560010161180b56fe4f6e6520737465702070726f6f66207769746820696e76616c696420707265762073746174654368616c6c656e6765206d75737420626520666f727761726465642066726f6d206d61696e20636f6e747261637443616e206f6e6c792074696d65206f7574206173736572746572206966206974206973207468656972207475726e43616e206f6e6c792074696d65206f7574206368616c6c656e676572206966206974206973207468656972207475726e43616e206f6e6c79206f6e6520737465702070726f6f6620666f6c6c6f77696e6720612073696e676c652073746570206368616c6c656e67655468657265206d757374206265206e6f206578697374696e67206368616c6c656e6765a265627a7a7230582042c9e18c5592935f14037cfd0d72e071e73690288ff3926a8fb6edf0a077726a64736f6c634300050a0032"

// DeployChallengeManager deploys a new Ethereum contract, binding an instance of ChallengeManager to it.
func DeployChallengeManager(auth *bind.TransactOpts, backend bind.ContractBackend, _vmTracker common.Address) (common.Address, *types.Transaction, *ChallengeManager, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	bisectionAddr, _, _, _ := DeployBisection(auth, backend)
	ChallengeManagerBin = strings.Replace(ChallengeManagerBin, "__$0a962fcb953632d8e24d6f165261b09420$__", bisectionAddr.String()[2:], -1)

	arbProtocolAddr, _, _, _ := DeployArbProtocol(auth, backend)
	ChallengeManagerBin = strings.Replace(ChallengeManagerBin, "__$6b4cc75dad3e0abd6ad83b3d907747c608$__", arbProtocolAddr.String()[2:], -1)

	oneStepProofAddr, _, _, _ := DeployOneStepProof(auth, backend)
	ChallengeManagerBin = strings.Replace(ChallengeManagerBin, "__$98107b176d4310ec680b0534b46d40334a$__", oneStepProofAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChallengeManagerBin), backend, _vmTracker)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChallengeManager{ChallengeManagerCaller: ChallengeManagerCaller{contract: contract}, ChallengeManagerTransactor: ChallengeManagerTransactor{contract: contract}, ChallengeManagerFilterer: ChallengeManagerFilterer{contract: contract}}, nil
}

// ChallengeManager is an auto generated Go binding around an Ethereum contract.
type ChallengeManager struct {
	ChallengeManagerCaller     // Read-only binding to the contract
	ChallengeManagerTransactor // Write-only binding to the contract
	ChallengeManagerFilterer   // Log filterer for contract events
}

// ChallengeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeManagerSession struct {
	Contract     *ChallengeManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeManagerCallerSession struct {
	Contract *ChallengeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ChallengeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeManagerTransactorSession struct {
	Contract     *ChallengeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ChallengeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeManagerRaw struct {
	Contract *ChallengeManager // Generic contract binding to access the raw methods on
}

// ChallengeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeManagerCallerRaw struct {
	Contract *ChallengeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeManagerTransactorRaw struct {
	Contract *ChallengeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallengeManager creates a new instance of ChallengeManager, bound to a specific deployed contract.
func NewChallengeManager(address common.Address, backend bind.ContractBackend) (*ChallengeManager, error) {
	contract, err := bindChallengeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChallengeManager{ChallengeManagerCaller: ChallengeManagerCaller{contract: contract}, ChallengeManagerTransactor: ChallengeManagerTransactor{contract: contract}, ChallengeManagerFilterer: ChallengeManagerFilterer{contract: contract}}, nil
}

// NewChallengeManagerCaller creates a new read-only instance of ChallengeManager, bound to a specific deployed contract.
func NewChallengeManagerCaller(address common.Address, caller bind.ContractCaller) (*ChallengeManagerCaller, error) {
	contract, err := bindChallengeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerCaller{contract: contract}, nil
}

// NewChallengeManagerTransactor creates a new write-only instance of ChallengeManager, bound to a specific deployed contract.
func NewChallengeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeManagerTransactor, error) {
	contract, err := bindChallengeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerTransactor{contract: contract}, nil
}

// NewChallengeManagerFilterer creates a new log filterer instance of ChallengeManager, bound to a specific deployed contract.
func NewChallengeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeManagerFilterer, error) {
	contract, err := bindChallengeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerFilterer{contract: contract}, nil
}

// bindChallengeManager binds a generic wrapper to an already deployed contract.
func bindChallengeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeManager *ChallengeManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChallengeManager.Contract.ChallengeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeManager *ChallengeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ChallengeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeManager *ChallengeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ChallengeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeManager *ChallengeManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChallengeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeManager *ChallengeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeManager *ChallengeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeManager.Contract.contract.Transact(opts, method, params...)
}

// AsserterTimedOut is a paid mutator transaction binding the contract method 0xb814f8c7.
//
// Solidity: function asserterTimedOut(bytes32 _vmId) returns()
func (_ChallengeManager *ChallengeManagerTransactor) AsserterTimedOut(opts *bind.TransactOpts, _vmId [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "asserterTimedOut", _vmId)
}

// AsserterTimedOut is a paid mutator transaction binding the contract method 0xb814f8c7.
//
// Solidity: function asserterTimedOut(bytes32 _vmId) returns()
func (_ChallengeManager *ChallengeManagerSession) AsserterTimedOut(_vmId [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.AsserterTimedOut(&_ChallengeManager.TransactOpts, _vmId)
}

// AsserterTimedOut is a paid mutator transaction binding the contract method 0xb814f8c7.
//
// Solidity: function asserterTimedOut(bytes32 _vmId) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) AsserterTimedOut(_vmId [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.AsserterTimedOut(&_ChallengeManager.TransactOpts, _vmId)
}

// BisectAssertion is a paid mutator transaction binding the contract method 0x70027cbd.
//
// Solidity: function bisectAssertion(bytes32 _challengeId, bytes32[2] _fields, bytes32[] _afterHashAndMessageAndLogsBisections, uint256[] _totalMessageAmounts, uint32 _totalSteps, uint64[2] _timeBounds, bytes21[] _tokenTypes, uint256[] _beforeBalances) returns()
func (_ChallengeManager *ChallengeManagerTransactor) BisectAssertion(opts *bind.TransactOpts, _challengeId [32]byte, _fields [2][32]byte, _afterHashAndMessageAndLogsBisections [][32]byte, _totalMessageAmounts []*big.Int, _totalSteps uint32, _timeBounds [2]uint64, _tokenTypes [][21]byte, _beforeBalances []*big.Int) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "bisectAssertion", _challengeId, _fields, _afterHashAndMessageAndLogsBisections, _totalMessageAmounts, _totalSteps, _timeBounds, _tokenTypes, _beforeBalances)
}

// BisectAssertion is a paid mutator transaction binding the contract method 0x70027cbd.
//
// Solidity: function bisectAssertion(bytes32 _challengeId, bytes32[2] _fields, bytes32[] _afterHashAndMessageAndLogsBisections, uint256[] _totalMessageAmounts, uint32 _totalSteps, uint64[2] _timeBounds, bytes21[] _tokenTypes, uint256[] _beforeBalances) returns()
func (_ChallengeManager *ChallengeManagerSession) BisectAssertion(_challengeId [32]byte, _fields [2][32]byte, _afterHashAndMessageAndLogsBisections [][32]byte, _totalMessageAmounts []*big.Int, _totalSteps uint32, _timeBounds [2]uint64, _tokenTypes [][21]byte, _beforeBalances []*big.Int) (*types.Transaction, error) {
	return _ChallengeManager.Contract.BisectAssertion(&_ChallengeManager.TransactOpts, _challengeId, _fields, _afterHashAndMessageAndLogsBisections, _totalMessageAmounts, _totalSteps, _timeBounds, _tokenTypes, _beforeBalances)
}

// BisectAssertion is a paid mutator transaction binding the contract method 0x70027cbd.
//
// Solidity: function bisectAssertion(bytes32 _challengeId, bytes32[2] _fields, bytes32[] _afterHashAndMessageAndLogsBisections, uint256[] _totalMessageAmounts, uint32 _totalSteps, uint64[2] _timeBounds, bytes21[] _tokenTypes, uint256[] _beforeBalances) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) BisectAssertion(_challengeId [32]byte, _fields [2][32]byte, _afterHashAndMessageAndLogsBisections [][32]byte, _totalMessageAmounts []*big.Int, _totalSteps uint32, _timeBounds [2]uint64, _tokenTypes [][21]byte, _beforeBalances []*big.Int) (*types.Transaction, error) {
	return _ChallengeManager.Contract.BisectAssertion(&_ChallengeManager.TransactOpts, _challengeId, _fields, _afterHashAndMessageAndLogsBisections, _totalMessageAmounts, _totalSteps, _timeBounds, _tokenTypes, _beforeBalances)
}

// ChallengerTimedOut is a paid mutator transaction binding the contract method 0x4db0228b.
//
// Solidity: function challengerTimedOut(bytes32 _vmId) returns()
func (_ChallengeManager *ChallengeManagerTransactor) ChallengerTimedOut(opts *bind.TransactOpts, _vmId [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "challengerTimedOut", _vmId)
}

// ChallengerTimedOut is a paid mutator transaction binding the contract method 0x4db0228b.
//
// Solidity: function challengerTimedOut(bytes32 _vmId) returns()
func (_ChallengeManager *ChallengeManagerSession) ChallengerTimedOut(_vmId [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ChallengerTimedOut(&_ChallengeManager.TransactOpts, _vmId)
}

// ChallengerTimedOut is a paid mutator transaction binding the contract method 0x4db0228b.
//
// Solidity: function challengerTimedOut(bytes32 _vmId) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) ChallengerTimedOut(_vmId [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ChallengerTimedOut(&_ChallengeManager.TransactOpts, _vmId)
}

// ContinueChallenge is a paid mutator transaction binding the contract method 0x770a2e56.
//
// Solidity: function continueChallenge(bytes32 _vmId, uint256 _assertionToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_ChallengeManager *ChallengeManagerTransactor) ContinueChallenge(opts *bind.TransactOpts, _vmId [32]byte, _assertionToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "continueChallenge", _vmId, _assertionToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// ContinueChallenge is a paid mutator transaction binding the contract method 0x770a2e56.
//
// Solidity: function continueChallenge(bytes32 _vmId, uint256 _assertionToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_ChallengeManager *ChallengeManagerSession) ContinueChallenge(_vmId [32]byte, _assertionToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ContinueChallenge(&_ChallengeManager.TransactOpts, _vmId, _assertionToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// ContinueChallenge is a paid mutator transaction binding the contract method 0x770a2e56.
//
// Solidity: function continueChallenge(bytes32 _vmId, uint256 _assertionToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) ContinueChallenge(_vmId [32]byte, _assertionToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ContinueChallenge(&_ChallengeManager.TransactOpts, _vmId, _assertionToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x2b50d42b.
//
// Solidity: function initiateChallenge(bytes32 _vmId, address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 _challengeRoot) returns()
func (_ChallengeManager *ChallengeManagerTransactor) InitiateChallenge(opts *bind.TransactOpts, _vmId [32]byte, _players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, _challengeRoot [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "initiateChallenge", _vmId, _players, _escrows, _challengePeriod, _challengeRoot)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x2b50d42b.
//
// Solidity: function initiateChallenge(bytes32 _vmId, address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 _challengeRoot) returns()
func (_ChallengeManager *ChallengeManagerSession) InitiateChallenge(_vmId [32]byte, _players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, _challengeRoot [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.InitiateChallenge(&_ChallengeManager.TransactOpts, _vmId, _players, _escrows, _challengePeriod, _challengeRoot)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x2b50d42b.
//
// Solidity: function initiateChallenge(bytes32 _vmId, address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 _challengeRoot) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) InitiateChallenge(_vmId [32]byte, _players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, _challengeRoot [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.InitiateChallenge(&_ChallengeManager.TransactOpts, _vmId, _players, _escrows, _challengePeriod, _challengeRoot)
}

// OneStepProof is a paid mutator transaction binding the contract method 0x832dcf57.
//
// Solidity: function oneStepProof(bytes32 _vmId, bytes32[2] _beforeHashAndInbox, uint64[2] _timeBounds, bytes21[] _tokenTypes, uint256[] _beforeBalances, bytes32[5] _afterHashAndMessages, uint256[] _amounts, bytes _proof) returns()
func (_ChallengeManager *ChallengeManagerTransactor) OneStepProof(opts *bind.TransactOpts, _vmId [32]byte, _beforeHashAndInbox [2][32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _beforeBalances []*big.Int, _afterHashAndMessages [5][32]byte, _amounts []*big.Int, _proof []byte) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "oneStepProof", _vmId, _beforeHashAndInbox, _timeBounds, _tokenTypes, _beforeBalances, _afterHashAndMessages, _amounts, _proof)
}

// OneStepProof is a paid mutator transaction binding the contract method 0x832dcf57.
//
// Solidity: function oneStepProof(bytes32 _vmId, bytes32[2] _beforeHashAndInbox, uint64[2] _timeBounds, bytes21[] _tokenTypes, uint256[] _beforeBalances, bytes32[5] _afterHashAndMessages, uint256[] _amounts, bytes _proof) returns()
func (_ChallengeManager *ChallengeManagerSession) OneStepProof(_vmId [32]byte, _beforeHashAndInbox [2][32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _beforeBalances []*big.Int, _afterHashAndMessages [5][32]byte, _amounts []*big.Int, _proof []byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.OneStepProof(&_ChallengeManager.TransactOpts, _vmId, _beforeHashAndInbox, _timeBounds, _tokenTypes, _beforeBalances, _afterHashAndMessages, _amounts, _proof)
}

// OneStepProof is a paid mutator transaction binding the contract method 0x832dcf57.
//
// Solidity: function oneStepProof(bytes32 _vmId, bytes32[2] _beforeHashAndInbox, uint64[2] _timeBounds, bytes21[] _tokenTypes, uint256[] _beforeBalances, bytes32[5] _afterHashAndMessages, uint256[] _amounts, bytes _proof) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) OneStepProof(_vmId [32]byte, _beforeHashAndInbox [2][32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _beforeBalances []*big.Int, _afterHashAndMessages [5][32]byte, _amounts []*big.Int, _proof []byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.OneStepProof(&_ChallengeManager.TransactOpts, _vmId, _beforeHashAndInbox, _timeBounds, _tokenTypes, _beforeBalances, _afterHashAndMessages, _amounts, _proof)
}

// ChallengeManagerBisectedAssertionIterator is returned from FilterBisectedAssertion and is used to iterate over the raw logs and unpacked data for BisectedAssertion events raised by the ChallengeManager contract.
type ChallengeManagerBisectedAssertionIterator struct {
	Event *ChallengeManagerBisectedAssertion // Event containing the contract specifics and raw log

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
func (it *ChallengeManagerBisectedAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeManagerBisectedAssertion)
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
		it.Event = new(ChallengeManagerBisectedAssertion)
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
func (it *ChallengeManagerBisectedAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeManagerBisectedAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeManagerBisectedAssertion represents a BisectedAssertion event raised by the ChallengeManager contract.
type ChallengeManagerBisectedAssertion struct {
	VmId                                 [32]byte
	Bisecter                             common.Address
	AfterHashAndMessageAndLogsBisections [][32]byte
	TotalSteps                           uint32
	TotalMessageAmounts                  []*big.Int
	Raw                                  types.Log // Blockchain specific contextual infos
}

// FilterBisectedAssertion is a free log retrieval operation binding the contract event 0xa4f8cbbd195d8e69d66b332eb24e05c24a35ee00de81da472b89ecb42a70ef71.
//
// Solidity: event BisectedAssertion(bytes32 indexed vmId, address bisecter, bytes32[] afterHashAndMessageAndLogsBisections, uint32 totalSteps, uint256[] totalMessageAmounts)
func (_ChallengeManager *ChallengeManagerFilterer) FilterBisectedAssertion(opts *bind.FilterOpts, vmId [][32]byte) (*ChallengeManagerBisectedAssertionIterator, error) {

	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "BisectedAssertion", vmIdRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerBisectedAssertionIterator{contract: _ChallengeManager.contract, event: "BisectedAssertion", logs: logs, sub: sub}, nil
}

// WatchBisectedAssertion is a free log subscription operation binding the contract event 0xa4f8cbbd195d8e69d66b332eb24e05c24a35ee00de81da472b89ecb42a70ef71.
//
// Solidity: event BisectedAssertion(bytes32 indexed vmId, address bisecter, bytes32[] afterHashAndMessageAndLogsBisections, uint32 totalSteps, uint256[] totalMessageAmounts)
func (_ChallengeManager *ChallengeManagerFilterer) WatchBisectedAssertion(opts *bind.WatchOpts, sink chan<- *ChallengeManagerBisectedAssertion, vmId [][32]byte) (event.Subscription, error) {

	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "BisectedAssertion", vmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeManagerBisectedAssertion)
				if err := _ChallengeManager.contract.UnpackLog(event, "BisectedAssertion", log); err != nil {
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

// ParseBisectedAssertion is a log parse operation binding the contract event 0xa4f8cbbd195d8e69d66b332eb24e05c24a35ee00de81da472b89ecb42a70ef71.
//
// Solidity: event BisectedAssertion(bytes32 indexed vmId, address bisecter, bytes32[] afterHashAndMessageAndLogsBisections, uint32 totalSteps, uint256[] totalMessageAmounts)
func (_ChallengeManager *ChallengeManagerFilterer) ParseBisectedAssertion(log types.Log) (*ChallengeManagerBisectedAssertion, error) {
	event := new(ChallengeManagerBisectedAssertion)
	if err := _ChallengeManager.contract.UnpackLog(event, "BisectedAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChallengeManagerContinuedChallengeIterator is returned from FilterContinuedChallenge and is used to iterate over the raw logs and unpacked data for ContinuedChallenge events raised by the ChallengeManager contract.
type ChallengeManagerContinuedChallengeIterator struct {
	Event *ChallengeManagerContinuedChallenge // Event containing the contract specifics and raw log

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
func (it *ChallengeManagerContinuedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeManagerContinuedChallenge)
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
		it.Event = new(ChallengeManagerContinuedChallenge)
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
func (it *ChallengeManagerContinuedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeManagerContinuedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeManagerContinuedChallenge represents a ContinuedChallenge event raised by the ChallengeManager contract.
type ChallengeManagerContinuedChallenge struct {
	VmId           [32]byte
	Challenger     common.Address
	AssertionIndex *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterContinuedChallenge is a free log retrieval operation binding the contract event 0xf5a6d1468c6ce7a03663fa2fd47dad9f6693a9289ebdbcf3309941caa75c7cf0.
//
// Solidity: event ContinuedChallenge(bytes32 indexed vmId, address challenger, uint256 assertionIndex)
func (_ChallengeManager *ChallengeManagerFilterer) FilterContinuedChallenge(opts *bind.FilterOpts, vmId [][32]byte) (*ChallengeManagerContinuedChallengeIterator, error) {

	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "ContinuedChallenge", vmIdRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerContinuedChallengeIterator{contract: _ChallengeManager.contract, event: "ContinuedChallenge", logs: logs, sub: sub}, nil
}

// WatchContinuedChallenge is a free log subscription operation binding the contract event 0xf5a6d1468c6ce7a03663fa2fd47dad9f6693a9289ebdbcf3309941caa75c7cf0.
//
// Solidity: event ContinuedChallenge(bytes32 indexed vmId, address challenger, uint256 assertionIndex)
func (_ChallengeManager *ChallengeManagerFilterer) WatchContinuedChallenge(opts *bind.WatchOpts, sink chan<- *ChallengeManagerContinuedChallenge, vmId [][32]byte) (event.Subscription, error) {

	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "ContinuedChallenge", vmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeManagerContinuedChallenge)
				if err := _ChallengeManager.contract.UnpackLog(event, "ContinuedChallenge", log); err != nil {
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

// ParseContinuedChallenge is a log parse operation binding the contract event 0xf5a6d1468c6ce7a03663fa2fd47dad9f6693a9289ebdbcf3309941caa75c7cf0.
//
// Solidity: event ContinuedChallenge(bytes32 indexed vmId, address challenger, uint256 assertionIndex)
func (_ChallengeManager *ChallengeManagerFilterer) ParseContinuedChallenge(log types.Log) (*ChallengeManagerContinuedChallenge, error) {
	event := new(ChallengeManagerContinuedChallenge)
	if err := _ChallengeManager.contract.UnpackLog(event, "ContinuedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChallengeManagerOneStepProofCompletedIterator is returned from FilterOneStepProofCompleted and is used to iterate over the raw logs and unpacked data for OneStepProofCompleted events raised by the ChallengeManager contract.
type ChallengeManagerOneStepProofCompletedIterator struct {
	Event *ChallengeManagerOneStepProofCompleted // Event containing the contract specifics and raw log

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
func (it *ChallengeManagerOneStepProofCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeManagerOneStepProofCompleted)
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
		it.Event = new(ChallengeManagerOneStepProofCompleted)
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
func (it *ChallengeManagerOneStepProofCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeManagerOneStepProofCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeManagerOneStepProofCompleted represents a OneStepProofCompleted event raised by the ChallengeManager contract.
type ChallengeManagerOneStepProofCompleted struct {
	VmId     [32]byte
	Asserter common.Address
	Proof    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOneStepProofCompleted is a free log retrieval operation binding the contract event 0xb3b8389d4e8e450396d43f12a728c7a73784c6089340d0e904a24fd366ddaaff.
//
// Solidity: event OneStepProofCompleted(bytes32 indexed vmId, address asserter, bytes proof)
func (_ChallengeManager *ChallengeManagerFilterer) FilterOneStepProofCompleted(opts *bind.FilterOpts, vmId [][32]byte) (*ChallengeManagerOneStepProofCompletedIterator, error) {

	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "OneStepProofCompleted", vmIdRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerOneStepProofCompletedIterator{contract: _ChallengeManager.contract, event: "OneStepProofCompleted", logs: logs, sub: sub}, nil
}

// WatchOneStepProofCompleted is a free log subscription operation binding the contract event 0xb3b8389d4e8e450396d43f12a728c7a73784c6089340d0e904a24fd366ddaaff.
//
// Solidity: event OneStepProofCompleted(bytes32 indexed vmId, address asserter, bytes proof)
func (_ChallengeManager *ChallengeManagerFilterer) WatchOneStepProofCompleted(opts *bind.WatchOpts, sink chan<- *ChallengeManagerOneStepProofCompleted, vmId [][32]byte) (event.Subscription, error) {

	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "OneStepProofCompleted", vmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeManagerOneStepProofCompleted)
				if err := _ChallengeManager.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
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

// ParseOneStepProofCompleted is a log parse operation binding the contract event 0xb3b8389d4e8e450396d43f12a728c7a73784c6089340d0e904a24fd366ddaaff.
//
// Solidity: event OneStepProofCompleted(bytes32 indexed vmId, address asserter, bytes proof)
func (_ChallengeManager *ChallengeManagerFilterer) ParseOneStepProofCompleted(log types.Log) (*ChallengeManagerOneStepProofCompleted, error) {
	event := new(ChallengeManagerOneStepProofCompleted)
	if err := _ChallengeManager.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChallengeManagerTimedOutChallengeIterator is returned from FilterTimedOutChallenge and is used to iterate over the raw logs and unpacked data for TimedOutChallenge events raised by the ChallengeManager contract.
type ChallengeManagerTimedOutChallengeIterator struct {
	Event *ChallengeManagerTimedOutChallenge // Event containing the contract specifics and raw log

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
func (it *ChallengeManagerTimedOutChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeManagerTimedOutChallenge)
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
		it.Event = new(ChallengeManagerTimedOutChallenge)
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
func (it *ChallengeManagerTimedOutChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeManagerTimedOutChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeManagerTimedOutChallenge represents a TimedOutChallenge event raised by the ChallengeManager contract.
type ChallengeManagerTimedOutChallenge struct {
	VmId            [32]byte
	ChallengerWrong bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimedOutChallenge is a free log retrieval operation binding the contract event 0x4bb5367552b51194ffa95563d907630acc956760789c62c6e92b82dfddf7e680.
//
// Solidity: event TimedOutChallenge(bytes32 indexed vmId, bool challengerWrong)
func (_ChallengeManager *ChallengeManagerFilterer) FilterTimedOutChallenge(opts *bind.FilterOpts, vmId [][32]byte) (*ChallengeManagerTimedOutChallengeIterator, error) {

	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "TimedOutChallenge", vmIdRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerTimedOutChallengeIterator{contract: _ChallengeManager.contract, event: "TimedOutChallenge", logs: logs, sub: sub}, nil
}

// WatchTimedOutChallenge is a free log subscription operation binding the contract event 0x4bb5367552b51194ffa95563d907630acc956760789c62c6e92b82dfddf7e680.
//
// Solidity: event TimedOutChallenge(bytes32 indexed vmId, bool challengerWrong)
func (_ChallengeManager *ChallengeManagerFilterer) WatchTimedOutChallenge(opts *bind.WatchOpts, sink chan<- *ChallengeManagerTimedOutChallenge, vmId [][32]byte) (event.Subscription, error) {

	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "TimedOutChallenge", vmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeManagerTimedOutChallenge)
				if err := _ChallengeManager.contract.UnpackLog(event, "TimedOutChallenge", log); err != nil {
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

// ParseTimedOutChallenge is a log parse operation binding the contract event 0x4bb5367552b51194ffa95563d907630acc956760789c62c6e92b82dfddf7e680.
//
// Solidity: event TimedOutChallenge(bytes32 indexed vmId, bool challengerWrong)
func (_ChallengeManager *ChallengeManagerFilterer) ParseTimedOutChallenge(log types.Log) (*ChallengeManagerTimedOutChallenge, error) {
	event := new(ChallengeManagerTimedOutChallenge)
	if err := _ChallengeManager.contract.UnpackLog(event, "TimedOutChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IChallengeManagerABI is the input ABI used to generate the binding from.
const IChallengeManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"vmId\",\"type\":\"bytes32\"},{\"name\":\"players\",\"type\":\"address[2]\"},{\"name\":\"escrows\",\"type\":\"uint128[2]\"},{\"name\":\"challengePeriod\",\"type\":\"uint32\"},{\"name\":\"challengeRoot\",\"type\":\"bytes32\"}],\"name\":\"initiateChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IChallengeManagerFuncSigs maps the 4-byte function signature to its string representation.
var IChallengeManagerFuncSigs = map[string]string{
	"2b50d42b": "initiateChallenge(bytes32,address[2],uint128[2],uint32,bytes32)",
}

// IChallengeManager is an auto generated Go binding around an Ethereum contract.
type IChallengeManager struct {
	IChallengeManagerCaller     // Read-only binding to the contract
	IChallengeManagerTransactor // Write-only binding to the contract
	IChallengeManagerFilterer   // Log filterer for contract events
}

// IChallengeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IChallengeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IChallengeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IChallengeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IChallengeManagerSession struct {
	Contract     *IChallengeManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IChallengeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IChallengeManagerCallerSession struct {
	Contract *IChallengeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IChallengeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IChallengeManagerTransactorSession struct {
	Contract     *IChallengeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IChallengeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IChallengeManagerRaw struct {
	Contract *IChallengeManager // Generic contract binding to access the raw methods on
}

// IChallengeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IChallengeManagerCallerRaw struct {
	Contract *IChallengeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IChallengeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IChallengeManagerTransactorRaw struct {
	Contract *IChallengeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIChallengeManager creates a new instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManager(address common.Address, backend bind.ContractBackend) (*IChallengeManager, error) {
	contract, err := bindIChallengeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IChallengeManager{IChallengeManagerCaller: IChallengeManagerCaller{contract: contract}, IChallengeManagerTransactor: IChallengeManagerTransactor{contract: contract}, IChallengeManagerFilterer: IChallengeManagerFilterer{contract: contract}}, nil
}

// NewIChallengeManagerCaller creates a new read-only instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManagerCaller(address common.Address, caller bind.ContractCaller) (*IChallengeManagerCaller, error) {
	contract, err := bindIChallengeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerCaller{contract: contract}, nil
}

// NewIChallengeManagerTransactor creates a new write-only instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IChallengeManagerTransactor, error) {
	contract, err := bindIChallengeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerTransactor{contract: contract}, nil
}

// NewIChallengeManagerFilterer creates a new log filterer instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IChallengeManagerFilterer, error) {
	contract, err := bindIChallengeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerFilterer{contract: contract}, nil
}

// bindIChallengeManager binds a generic wrapper to an already deployed contract.
func bindIChallengeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IChallengeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeManager *IChallengeManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IChallengeManager.Contract.IChallengeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeManager *IChallengeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeManager.Contract.IChallengeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeManager *IChallengeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeManager.Contract.IChallengeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeManager *IChallengeManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IChallengeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeManager *IChallengeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeManager *IChallengeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeManager.Contract.contract.Transact(opts, method, params...)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x2b50d42b.
//
// Solidity: function initiateChallenge(bytes32 vmId, address[2] players, uint128[2] escrows, uint32 challengePeriod, bytes32 challengeRoot) returns()
func (_IChallengeManager *IChallengeManagerTransactor) InitiateChallenge(opts *bind.TransactOpts, vmId [32]byte, players [2]common.Address, escrows [2]*big.Int, challengePeriod uint32, challengeRoot [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "initiateChallenge", vmId, players, escrows, challengePeriod, challengeRoot)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x2b50d42b.
//
// Solidity: function initiateChallenge(bytes32 vmId, address[2] players, uint128[2] escrows, uint32 challengePeriod, bytes32 challengeRoot) returns()
func (_IChallengeManager *IChallengeManagerSession) InitiateChallenge(vmId [32]byte, players [2]common.Address, escrows [2]*big.Int, challengePeriod uint32, challengeRoot [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.InitiateChallenge(&_IChallengeManager.TransactOpts, vmId, players, escrows, challengePeriod, challengeRoot)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x2b50d42b.
//
// Solidity: function initiateChallenge(bytes32 vmId, address[2] players, uint128[2] escrows, uint32 challengePeriod, bytes32 challengeRoot) returns()
func (_IChallengeManager *IChallengeManagerTransactorSession) InitiateChallenge(vmId [32]byte, players [2]common.Address, escrows [2]*big.Int, challengePeriod uint32, challengeRoot [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.InitiateChallenge(&_IChallengeManager.TransactOpts, vmId, players, escrows, challengePeriod, challengeRoot)
}

// IVMTrackerABI is the input ABI used to generate the binding from.
const IVMTrackerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"bytes32\"},{\"name\":\"_players\",\"type\":\"address[2]\"},{\"name\":\"_rewards\",\"type\":\"uint128[2]\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IVMTrackerFuncSigs maps the 4-byte function signature to its string representation.
var IVMTrackerFuncSigs = map[string]string{
	"63d84637": "completeChallenge(bytes32,address[2],uint128[2])",
}

// IVMTracker is an auto generated Go binding around an Ethereum contract.
type IVMTracker struct {
	IVMTrackerCaller     // Read-only binding to the contract
	IVMTrackerTransactor // Write-only binding to the contract
	IVMTrackerFilterer   // Log filterer for contract events
}

// IVMTrackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IVMTrackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVMTrackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IVMTrackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVMTrackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IVMTrackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVMTrackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IVMTrackerSession struct {
	Contract     *IVMTracker       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVMTrackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IVMTrackerCallerSession struct {
	Contract *IVMTrackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IVMTrackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IVMTrackerTransactorSession struct {
	Contract     *IVMTrackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IVMTrackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IVMTrackerRaw struct {
	Contract *IVMTracker // Generic contract binding to access the raw methods on
}

// IVMTrackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IVMTrackerCallerRaw struct {
	Contract *IVMTrackerCaller // Generic read-only contract binding to access the raw methods on
}

// IVMTrackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IVMTrackerTransactorRaw struct {
	Contract *IVMTrackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIVMTracker creates a new instance of IVMTracker, bound to a specific deployed contract.
func NewIVMTracker(address common.Address, backend bind.ContractBackend) (*IVMTracker, error) {
	contract, err := bindIVMTracker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVMTracker{IVMTrackerCaller: IVMTrackerCaller{contract: contract}, IVMTrackerTransactor: IVMTrackerTransactor{contract: contract}, IVMTrackerFilterer: IVMTrackerFilterer{contract: contract}}, nil
}

// NewIVMTrackerCaller creates a new read-only instance of IVMTracker, bound to a specific deployed contract.
func NewIVMTrackerCaller(address common.Address, caller bind.ContractCaller) (*IVMTrackerCaller, error) {
	contract, err := bindIVMTracker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVMTrackerCaller{contract: contract}, nil
}

// NewIVMTrackerTransactor creates a new write-only instance of IVMTracker, bound to a specific deployed contract.
func NewIVMTrackerTransactor(address common.Address, transactor bind.ContractTransactor) (*IVMTrackerTransactor, error) {
	contract, err := bindIVMTracker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVMTrackerTransactor{contract: contract}, nil
}

// NewIVMTrackerFilterer creates a new log filterer instance of IVMTracker, bound to a specific deployed contract.
func NewIVMTrackerFilterer(address common.Address, filterer bind.ContractFilterer) (*IVMTrackerFilterer, error) {
	contract, err := bindIVMTracker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVMTrackerFilterer{contract: contract}, nil
}

// bindIVMTracker binds a generic wrapper to an already deployed contract.
func bindIVMTracker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IVMTrackerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVMTracker *IVMTrackerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IVMTracker.Contract.IVMTrackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVMTracker *IVMTrackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVMTracker.Contract.IVMTrackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVMTracker *IVMTrackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVMTracker.Contract.IVMTrackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVMTracker *IVMTrackerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IVMTracker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVMTracker *IVMTrackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVMTracker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVMTracker *IVMTrackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVMTracker.Contract.contract.Transact(opts, method, params...)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x63d84637.
//
// Solidity: function completeChallenge(bytes32 _vmId, address[2] _players, uint128[2] _rewards) returns()
func (_IVMTracker *IVMTrackerTransactor) CompleteChallenge(opts *bind.TransactOpts, _vmId [32]byte, _players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _IVMTracker.contract.Transact(opts, "completeChallenge", _vmId, _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x63d84637.
//
// Solidity: function completeChallenge(bytes32 _vmId, address[2] _players, uint128[2] _rewards) returns()
func (_IVMTracker *IVMTrackerSession) CompleteChallenge(_vmId [32]byte, _players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _IVMTracker.Contract.CompleteChallenge(&_IVMTracker.TransactOpts, _vmId, _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x63d84637.
//
// Solidity: function completeChallenge(bytes32 _vmId, address[2] _players, uint128[2] _rewards) returns()
func (_IVMTracker *IVMTrackerTransactorSession) CompleteChallenge(_vmId [32]byte, _players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _IVMTracker.Contract.CompleteChallenge(&_IVMTracker.TransactOpts, _vmId, _players, _rewards)
}

// MerkleLibABI is the input ABI used to generate the binding from.
const MerkleLibABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"generateAddressRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_hashes\",\"type\":\"bytes32[]\"}],\"name\":\"generateRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"root\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"verifyProof\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// MerkleLibFuncSigs maps the 4-byte function signature to its string representation.
var MerkleLibFuncSigs = map[string]string{
	"6a2dda67": "generateAddressRoot(address[])",
	"9898dc10": "generateRoot(bytes32[])",
	"b792d767": "verifyProof(bytes,bytes32,bytes32,uint256)",
}

// MerkleLibBin is the compiled bytecode used for deploying new contracts.
var MerkleLibBin = "0x610575610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004b5760003560e01c80636a2dda67146100505780639898dc1014610105578063b792d767146101a8575b600080fd5b6100f36004803603602081101561006657600080fd5b81019060208101813564010000000081111561008157600080fd5b82018360208201111561009357600080fd5b803590602001918460208302840111640100000000831117156100b557600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061026d945050505050565b60408051918252519081900360200190f35b6100f36004803603602081101561011b57600080fd5b81019060208101813564010000000081111561013657600080fd5b82018360208201111561014857600080fd5b8035906020019184602083028401116401000000008311171561016a57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610301945050505050565b610259600480360360808110156101be57600080fd5b8101906020810181356401000000008111156101d957600080fd5b8201836020820111156101eb57600080fd5b8035906020019184600183028401116401000000008311171561020d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550508235935050506020810135906040013561043f565b604080519115158252519081900360200190f35b60006060825160405190808252806020026020018201604052801561029c578160200160208202803883390190505b50905060005b83518110156102f0578381815181106102b757fe5b602002602001015160601b6bffffffffffffffffffffffff19168282815181106102dd57fe5b60209081029190910101526001016102a2565b506102fa81610301565b9392505050565b6000815b600181511115610422576060600282516001018161031f57fe5b04604051908082528060200260200182016040528015610349578160200160208202803883390190505b50905060005b815181101561041a5782518160020260010110156103e25782816002028151811061037657fe5b602002602001015183826002026001018151811061039057fe5b60200260200101516040516020018083815260200182815260200192505050604051602081830303815290604052805190602001208282815181106103d157fe5b602002602001018181525050610412565b8281600202815181106103f157fe5b602002602001015182828151811061040557fe5b6020026020010181815250505b60010161034f565b509050610305565b8060008151811061042f57fe5b6020026020010151915050919050565b600080838160205b88518111610532578089015193506020818a51036020018161046557fe5b0491505b60008211801561047c5750600286066001145b801561048a57508160020a86115b1561049d57600286046001019550610469565b600286066104e85783836040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209250600286816104e057fe5b04955061052a565b828460405160200180838152602001828152602001925050506040516020818303038152906040528051906020012092506002868161052357fe5b0460010195505b602001610447565b50509094149594505050505056fea265627a7a723058209d7af0b10f14fb3c4510a87a685097ce94041ed1d974962866d79395034b110064736f6c634300050a0032"

// DeployMerkleLib deploys a new Ethereum contract, binding an instance of MerkleLib to it.
func DeployMerkleLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MerkleLib, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MerkleLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleLib{MerkleLibCaller: MerkleLibCaller{contract: contract}, MerkleLibTransactor: MerkleLibTransactor{contract: contract}, MerkleLibFilterer: MerkleLibFilterer{contract: contract}}, nil
}

// MerkleLib is an auto generated Go binding around an Ethereum contract.
type MerkleLib struct {
	MerkleLibCaller     // Read-only binding to the contract
	MerkleLibTransactor // Write-only binding to the contract
	MerkleLibFilterer   // Log filterer for contract events
}

// MerkleLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleLibSession struct {
	Contract     *MerkleLib        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleLibCallerSession struct {
	Contract *MerkleLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MerkleLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleLibTransactorSession struct {
	Contract     *MerkleLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MerkleLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleLibRaw struct {
	Contract *MerkleLib // Generic contract binding to access the raw methods on
}

// MerkleLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleLibCallerRaw struct {
	Contract *MerkleLibCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleLibTransactorRaw struct {
	Contract *MerkleLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleLib creates a new instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLib(address common.Address, backend bind.ContractBackend) (*MerkleLib, error) {
	contract, err := bindMerkleLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleLib{MerkleLibCaller: MerkleLibCaller{contract: contract}, MerkleLibTransactor: MerkleLibTransactor{contract: contract}, MerkleLibFilterer: MerkleLibFilterer{contract: contract}}, nil
}

// NewMerkleLibCaller creates a new read-only instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLibCaller(address common.Address, caller bind.ContractCaller) (*MerkleLibCaller, error) {
	contract, err := bindMerkleLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleLibCaller{contract: contract}, nil
}

// NewMerkleLibTransactor creates a new write-only instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLibTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleLibTransactor, error) {
	contract, err := bindMerkleLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleLibTransactor{contract: contract}, nil
}

// NewMerkleLibFilterer creates a new log filterer instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLibFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleLibFilterer, error) {
	contract, err := bindMerkleLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleLibFilterer{contract: contract}, nil
}

// bindMerkleLib binds a generic wrapper to an already deployed contract.
func bindMerkleLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleLib *MerkleLibRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MerkleLib.Contract.MerkleLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleLib *MerkleLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleLib.Contract.MerkleLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleLib *MerkleLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleLib.Contract.MerkleLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleLib *MerkleLibCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MerkleLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleLib *MerkleLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleLib *MerkleLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleLib.Contract.contract.Transact(opts, method, params...)
}

// GenerateAddressRoot is a free data retrieval call binding the contract method 0x6a2dda67.
//
// Solidity: function generateAddressRoot(address[] _addresses) constant returns(bytes32)
func (_MerkleLib *MerkleLibCaller) GenerateAddressRoot(opts *bind.CallOpts, _addresses []common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MerkleLib.contract.Call(opts, out, "generateAddressRoot", _addresses)
	return *ret0, err
}

// GenerateAddressRoot is a free data retrieval call binding the contract method 0x6a2dda67.
//
// Solidity: function generateAddressRoot(address[] _addresses) constant returns(bytes32)
func (_MerkleLib *MerkleLibSession) GenerateAddressRoot(_addresses []common.Address) ([32]byte, error) {
	return _MerkleLib.Contract.GenerateAddressRoot(&_MerkleLib.CallOpts, _addresses)
}

// GenerateAddressRoot is a free data retrieval call binding the contract method 0x6a2dda67.
//
// Solidity: function generateAddressRoot(address[] _addresses) constant returns(bytes32)
func (_MerkleLib *MerkleLibCallerSession) GenerateAddressRoot(_addresses []common.Address) ([32]byte, error) {
	return _MerkleLib.Contract.GenerateAddressRoot(&_MerkleLib.CallOpts, _addresses)
}

// GenerateRoot is a free data retrieval call binding the contract method 0x9898dc10.
//
// Solidity: function generateRoot(bytes32[] _hashes) constant returns(bytes32)
func (_MerkleLib *MerkleLibCaller) GenerateRoot(opts *bind.CallOpts, _hashes [][32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MerkleLib.contract.Call(opts, out, "generateRoot", _hashes)
	return *ret0, err
}

// GenerateRoot is a free data retrieval call binding the contract method 0x9898dc10.
//
// Solidity: function generateRoot(bytes32[] _hashes) constant returns(bytes32)
func (_MerkleLib *MerkleLibSession) GenerateRoot(_hashes [][32]byte) ([32]byte, error) {
	return _MerkleLib.Contract.GenerateRoot(&_MerkleLib.CallOpts, _hashes)
}

// GenerateRoot is a free data retrieval call binding the contract method 0x9898dc10.
//
// Solidity: function generateRoot(bytes32[] _hashes) constant returns(bytes32)
func (_MerkleLib *MerkleLibCallerSession) GenerateRoot(_hashes [][32]byte) ([32]byte, error) {
	return _MerkleLib.Contract.GenerateRoot(&_MerkleLib.CallOpts, _hashes)
}

// VerifyProof is a free data retrieval call binding the contract method 0xb792d767.
//
// Solidity: function verifyProof(bytes proof, bytes32 root, bytes32 hash, uint256 index) constant returns(bool)
func (_MerkleLib *MerkleLibCaller) VerifyProof(opts *bind.CallOpts, proof []byte, root [32]byte, hash [32]byte, index *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MerkleLib.contract.Call(opts, out, "verifyProof", proof, root, hash, index)
	return *ret0, err
}

// VerifyProof is a free data retrieval call binding the contract method 0xb792d767.
//
// Solidity: function verifyProof(bytes proof, bytes32 root, bytes32 hash, uint256 index) constant returns(bool)
func (_MerkleLib *MerkleLibSession) VerifyProof(proof []byte, root [32]byte, hash [32]byte, index *big.Int) (bool, error) {
	return _MerkleLib.Contract.VerifyProof(&_MerkleLib.CallOpts, proof, root, hash, index)
}

// VerifyProof is a free data retrieval call binding the contract method 0xb792d767.
//
// Solidity: function verifyProof(bytes proof, bytes32 root, bytes32 hash, uint256 index) constant returns(bool)
func (_MerkleLib *MerkleLibCallerSession) VerifyProof(proof []byte, root [32]byte, hash [32]byte, index *big.Int) (bool, error) {
	return _MerkleLib.Contract.VerifyProof(&_MerkleLib.CallOpts, proof, root, hash, index)
}

// OneStepProofABI is the input ABI used to generate the binding from.
const OneStepProofABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"fields\",\"type\":\"bytes32[7]\"},{\"name\":\"timeBounds\",\"type\":\"uint64[2]\"},{\"name\":\"tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"beforeValues\",\"type\":\"uint256[]\"},{\"name\":\"messageValue\",\"type\":\"uint256[]\"},{\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"validateProof\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"instructionStack\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"dataStack\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"auxStack\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"register\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"staticHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"errHandler\",\"type\":\"bytes32\"}],\"name\":\"SawMachine\",\"type\":\"event\"}]"

// OneStepProofFuncSigs maps the 4-byte function signature to its string representation.
var OneStepProofFuncSigs = map[string]string{
	"0eca9f13": "validateProof(bytes32[7],uint64[2],bytes21[],uint256[],uint256[],bytes)",
}

// OneStepProofBin is the compiled bytecode used for deploying new contracts.
var OneStepProofBin = "0x613633610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c80630eca9f131461003a575b600080fd5b6102c160048036036101a081101561005157600080fd5b810190808060e001906007806020026040519081016040528092919082600760200280828437600092019190915250506040805180820182529295949381810193925090600290839083908082843760009201919091525091949392602081019250359050600160201b8111156100c757600080fd5b8201836020820111156100d957600080fd5b803590602001918460208302840111600160201b831117156100fa57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561014957600080fd5b82018360208201111561015b57600080fd5b803590602001918460208302840111600160201b8311171561017c57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156101cb57600080fd5b8201836020820111156101dd57600080fd5b803590602001918460208302840111600160201b831117156101fe57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561024d57600080fd5b82018360208201111561025f57600080fd5b803590602001918460018302840111600160201b8311171561028057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506102d3945050505050565b60408051918252519081900360200190f35b6080860151606087015160009182918291829114801590600019906104e85760005b88518167ffffffffffffffff16101561038a57888167ffffffffffffffff168151811061031e57fe5b6020026020010151600014610382578160070b6000191461037e576040805162461bcd60e51b81526020600482015260156024820152746d756c7469706c65206f7574206d6573736167657360581b604482015290519081900360640190fd5b8091505b6001016102f5565b508060070b600019146104e357878160070b815181106103a657fe5b60200260200101519350898160070b815181106103bf57fe5b6020026020010151945060019250898160070b815181106103dc57fe5b60200260200101516014601581106103f057fe5b1a60f81b6001600160f81b031916600160f81b14156104785783898260070b8151811061041957fe5b602002602001015114610473576040805162461bcd60e51b815260206004820152601a60248201527f707265636f6e646974696f6e206d7573742068617665206e6674000000000000604482015290519081900360640190fd5b6104e3565b888160070b8151811061048757fe5b60200260200101518411156104e3576040805162461bcd60e51b815260206004820152601c60248201527f707265636f6e646974696f6e206d75737420686176652076616c756500000000604482015290519081900360640190fd5b61057a565b60005b88518167ffffffffffffffff16101561057857888167ffffffffffffffff168151811061051457fe5b6020026020010151600014610570576040805162461bcd60e51b815260206004820152601b60248201527f4d7573742068617665206e6f206d6573736167652076616c7565730000000000604482015290519081900360640190fd5b6001016104eb565b505b61065b6040518061018001604052808e60006007811061059657fe5b602002015181526020018d81526020018e6001600781106105b357fe5b602002015181526020018e6002600781106105ca57fe5b602002015181526020018e6003600781106105e157fe5b602002015181526020018e6004600781106105f857fe5b602002015181526020018e60056007811061060f57fe5b602002015181526020018e60066007811061062657fe5b60200201518152602001876affffffffffffffffffffff1916815260200186815260200185151581526020018981525061066b565b9c9b505050505050505050505050565b6000808080606061067a613469565b610682613469565b61068b88611515565b939950929650909450925090506001600060ff88168214156106e1576106da83866000815181106106b857fe5b6020026020010151876001815181106106cd57fe5b6020026020010151611955565b9150611369565b60ff881660021415610720576106da83866000815181106106fe57fe5b60200260200101518760018151811061071357fe5b60200260200101516119a5565b60ff88166003141561075f576106da838660008151811061073d57fe5b60200260200101518760018151811061075257fe5b60200260200101516119e6565b60ff88166004141561079e576106da838660008151811061077c57fe5b60200260200101518760018151811061079157fe5b6020026020010151611a27565b60ff8816600514156107dd576106da83866000815181106107bb57fe5b6020026020010151876001815181106107d057fe5b6020026020010151611a68565b60ff88166006141561081c576106da83866000815181106107fa57fe5b60200260200101518760018151811061080f57fe5b6020026020010151611aa9565b60ff88166007141561085b576106da838660008151811061083957fe5b60200260200101518760018151811061084e57fe5b6020026020010151611aea565b60ff8816600814156108af576106da838660008151811061087857fe5b60200260200101518760018151811061088d57fe5b6020026020010151886002815181106108a257fe5b6020026020010151611b2b565b60ff881660091415610903576106da83866000815181106108cc57fe5b6020026020010151876001815181106108e157fe5b6020026020010151886002815181106108f657fe5b6020026020010151611b84565b60ff8816600a1415610942576106da838660008151811061092057fe5b60200260200101518760018151811061093557fe5b6020026020010151611bcc565b60ff881660101415610981576106da838660008151811061095f57fe5b60200260200101518760018151811061097457fe5b6020026020010151611c0d565b60ff8816601114156109c0576106da838660008151811061099e57fe5b6020026020010151876001815181106109b357fe5b6020026020010151611c4e565b60ff8816601214156109ff576106da83866000815181106109dd57fe5b6020026020010151876001815181106109f257fe5b6020026020010151611c8f565b60ff881660131415610a1c576106da83866000815181106109dd57fe5b60ff881660141415610a5b576106da8386600081518110610a3957fe5b602002602001015187600181518110610a4e57fe5b6020026020010151611cd0565b60ff881660151415610a85576106da8386600081518110610a7857fe5b6020026020010151611cfc565b60ff881660161415610ac4576106da8386600081518110610aa257fe5b602002602001015187600181518110610ab757fe5b6020026020010151611d42565b60ff881660171415610b03576106da8386600081518110610ae157fe5b602002602001015187600181518110610af657fe5b6020026020010151611d83565b60ff881660181415610b42576106da8386600081518110610b2057fe5b602002602001015187600181518110610b3557fe5b6020026020010151611dc4565b60ff881660191415610b6c576106da8386600081518110610b5f57fe5b6020026020010151611e05565b60ff8816601a1415610bab576106da8386600081518110610b8957fe5b602002602001015187600181518110610b9e57fe5b6020026020010151611e3b565b60ff8816601b1415610bea576106da8386600081518110610bc857fe5b602002602001015187600181518110610bdd57fe5b6020026020010151611e7c565b60ff881660201415610c14576106da8386600081518110610c0757fe5b6020026020010151611ebd565b60ff881660211415610c3e576106da8386600081518110610c3157fe5b6020026020010151611ed9565b60ff881660301415610c68576106da8386600081518110610c5b57fe5b6020026020010151611ef4565b60ff881660311415610c7d576106da83611efc565b60ff881660321415610c92576106da83611f1d565b60ff881660331415610cbc576106da8386600081518110610caf57fe5b6020026020010151611f36565b60ff881660341415610ce6576106da8386600081518110610cd957fe5b6020026020010151611f4f565b60ff881660351415610d25576106da8386600081518110610d0357fe5b602002602001015187600181518110610d1857fe5b6020026020010151611f65565b60ff881660361415610d3a576106da83611f98565b60ff881660371415610d54576106da838560000151611fca565b60ff881660381415610d7e576106da8386600081518110610d7157fe5b6020026020010151611fdc565b60ff881660391415610e0b57610d926134ca565b610da18b610160015188611fee565b919950975090508715610de55760405162461bcd60e51b81526004018080602001828103825260218152602001806135de6021913960400191505060405180910390fd5b610df5858263ffffffff61214316565b610e05848263ffffffff61216516565b50611369565b60ff8816603b1415610e1c57611369565b60ff881660401415610e46576106da8386600081518110610e3957fe5b6020026020010151612182565b60ff881660411415610e85576106da8386600081518110610e6357fe5b602002602001015187600181518110610e7857fe5b60200260200101516121a4565b60ff881660421415610ed9576106da8386600081518110610ea257fe5b602002602001015187600181518110610eb757fe5b602002602001015188600281518110610ecc57fe5b60200260200101516121d6565b60ff881660431415610f18576106da8386600081518110610ef657fe5b602002602001015187600181518110610f0b57fe5b6020026020010151612218565b60ff881660441415610f6c576106da8386600081518110610f3557fe5b602002602001015187600181518110610f4a57fe5b602002602001015188600281518110610f5f57fe5b602002602001015161222a565b60ff881660501415610fab576106da8386600081518110610f8957fe5b602002602001015187600181518110610f9e57fe5b602002602001015161224c565b60ff881660511415610fff576106da8386600081518110610fc857fe5b602002602001015187600181518110610fdd57fe5b602002602001015188600281518110610ff257fe5b60200260200101516122c3565b60ff881660521415611029576106da838660008151811061101c57fe5b602002602001015161233c565b60ff88166060141561103e576106da8361236f565b60ff88166061141561112857611068838660008151811061105b57fe5b6020026020010151612375565b60e08c015160c08d01516040805160208082019390935280820185905281518082038301815260609091019091528051910120929450909250146110dd5760405162461bcd60e51b815260040180806020018281038252602581526020018061356a6025913960400191505060405180910390fd5b8960a001518a60800151146111235760405162461bcd60e51b815260040180806020018281038252602781526020018061358f6027913960400191505060405180910390fd5b611369565b60ff88166070141561122657600080611155858860008151811061114857fe5b6020026020010151612397565b809450819550829650839750505050508b60a001518c6080015184604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120146111d95760405162461bcd60e51b81526004018080602001828103825260288152602001806135b66028913960400191505060405180910390fd5b8b60e001518c60c001511461121f5760405162461bcd60e51b81526004018080602001828103825260268152602001806135236026913960400191505060405180910390fd5b5050611369565b60ff88166071141561124657600080611155858860008151811061114857fe5b60ff881660721415611302576040805160028082526060828101909352816020015b6112706134ca565b81526020019060019003908161126857505060208c01519091506112a59060005b602002015167ffffffffffffffff16612575565b816000815181106112b257fe5b60200260200101819052506112d18b6020015160016002811061129157fe5b816001815181106112de57fe5b6020026020010181905250610e056112f5826125cf565b859063ffffffff61216516565b60ff88166073141561133f576106da838660008151811061131f57fe5b602002602001015160405180602001604052808e60400151815250612657565b60ff8816607414156113545760009150611369565b60ff88166075141561136957611369836126c9565b806113fa578960a001518a60800151146113b45760405162461bcd60e51b815260040180806020018281038252602781526020018061358f6027913960400191505060405180910390fd5b8960e001518a60c00151146113fa5760405162461bcd60e51b81526004018080602001828103825260268152602001806135236026913960400191505060405180910390fd5b8161145c5760408051600160f81b602080830191909152600060218301819052602280840191909152835180840390910181526042909201909252805191012060a08401515114156114545761144f836126d3565b61145c565b60a083015183525b611465846126dd565b8a51146114a35760405162461bcd60e51b81526004018080602001828103825260228152602001806135016022913960400191505060405180910390fd5b6114ac836126dd565b8a6060015114611503576040805162461bcd60e51b815260206004820181905260248201527f50726f6f6620686164206e6f6e206d61746368696e6720656e64207374617465604482015290519081900360640190fd5b6000985050505050505050505b919050565b60006060611521613469565b611529613469565b60008080611535613469565b61153e81612772565b61154d8961016001518461277c565b909450909250905061155d613469565b61156682612881565b905060008a6101600151858151811061157b57fe5b602001015160f81c60f81b60f81c905060008b610160015186600101815181106115a157fe5b016020015160f81c905060006115b6826128df565b90506060816040519080825280602002602001820160405280156115f457816020015b6115e16134ca565b8152602001906001900390816115d95790505b5090506002880197508360ff166000148061161257508360ff166001145b611663576040805162461bcd60e51b815260206004820152601c60248201527f50726f6f662068616420626164206f7065726174696f6e207479706500000000604482015290519081900360640190fd5b60ff8416611706576040805160208082018084528951516353409fab60e01b90915260ff87166024840152604483015291519091829173__$0d86abb4a722a612872fb80f4c7e7e95bd$__916353409fab916064808601929190818703018186803b1580156116d157600080fd5b505af41580156116e5573d6000803e3d6000fd5b505050506040513d60208110156116fb57600080fd5b50519052865261185d565b61170e6134ca565b61171d8f61016001518a611fee565b909a5090985090508715611778576040805162461bcd60e51b815260206004820152601d60248201527f50726f6f66206861642062616420696d6d6564696174652076616c7565000000604482015290519081900360640190fd5b821561179c57808260008151811061178c57fe5b60200260200101819052506117ac565b6117ac868263ffffffff61216516565b604051806020016040528073__$0d86abb4a722a612872fb80f4c7e7e95bd$__63264f384b876117db866128f9565b518c5151604080516001600160e01b031960e087901b16815260ff909416600485015260248401929092526044830152516064808301926020929190829003018186803b15801561182b57600080fd5b505af415801561183f573d6000803e3d6000fd5b505050506040513d602081101561185557600080fd5b505190528752505b60ff84165b828110156118f1576118798f61016001518a611fee565b845185908590811061188757fe5b60209081029190910101529950975087156118e9576040805162461bcd60e51b815260206004820152601960248201527f50726f6f66206861642062616420737461636b2076616c756500000000000000604482015290519081900360640190fd5b600101611862565b81511561193e575060005b8460ff1682510381101561193e5761193682826001855103038151811061191f57fe5b60200260200101518861216590919063ffffffff16565b6001016118fc565b50919d919c50939a50919850939650945050505050565b6000611960836129e8565b15806119725750611970826129e8565b155b1561197f5750600061199e565b82518251808201611996878263ffffffff6129f316565b600193505050505b9392505050565b60006119b0836129e8565b15806119c257506119c0826129e8565b155b156119cf5750600061199e565b82518251808202611996878263ffffffff6129f316565b60006119f1836129e8565b1580611a035750611a01826129e8565b155b15611a105750600061199e565b82518251808203611996878263ffffffff6129f316565b6000611a32836129e8565b1580611a445750611a42826129e8565b155b15611a515750600061199e565b82518251808204611996878263ffffffff6129f316565b6000611a73836129e8565b1580611a855750611a83826129e8565b155b15611a925750600061199e565b82518251808205611996878263ffffffff6129f316565b6000611ab4836129e8565b1580611ac65750611ac4826129e8565b155b15611ad35750600061199e565b82518251808206611996878263ffffffff6129f316565b6000611af5836129e8565b1580611b075750611b05826129e8565b155b15611b145750600061199e565b82518251808207611996878263ffffffff6129f316565b6000611b36846129e8565b1580611b485750611b46836129e8565b155b15611b5557506000611b7c565b8351835183516000818385089050611b73898263ffffffff6129f316565b60019450505050505b949350505050565b6000611b8f846129e8565b1580611ba15750611b9f836129e8565b155b15611bae57506000611b7c565b8351835183516000818385099050611b73898263ffffffff6129f316565b6000611bd7836129e8565b1580611be95750611be7826129e8565b155b15611bf65750600061199e565b8251825180820a611996878263ffffffff6129f316565b6000611c18836129e8565b1580611c2a5750611c28826129e8565b155b15611c375750600061199e565b82518251808210611996878263ffffffff6129f316565b6000611c59836129e8565b1580611c6b5750611c69826129e8565b155b15611c785750600061199e565b82518251808211611996878263ffffffff6129f316565b6000611c9a836129e8565b1580611cac5750611caa826129e8565b155b15611cb95750600061199e565b82518251808212611996878263ffffffff6129f316565b6000611cf26112f5611ce1846128f9565b51611ceb866128f9565b5114612a07565b5060019392505050565b6000611d07826129e8565b611d2157611d1c83600063ffffffff6129f316565b611d38565b81518015611d35858263ffffffff6129f316565b50505b5060015b92915050565b6000611d4d836129e8565b1580611d5f5750611d5d826129e8565b155b15611d6c5750600061199e565b82518251808216611996878263ffffffff6129f316565b6000611d8e836129e8565b1580611da05750611d9e826129e8565b155b15611dad5750600061199e565b82518251808217611996878263ffffffff6129f316565b6000611dcf836129e8565b1580611de15750611ddf826129e8565b155b15611dee5750600061199e565b82518251808218611996878263ffffffff6129f316565b6000611e10826129e8565b611e1c57506000611d3c565b81518019611e30858263ffffffff6129f316565b506001949350505050565b6000611e46836129e8565b1580611e585750611e56826129e8565b155b15611e655750600061199e565b8251825180821a611996878263ffffffff6129f316565b6000611e87836129e8565b1580611e995750611e97826129e8565b155b15611ea65750600061199e565b8251825180820b611996878263ffffffff6129f316565b6000611d38611ecb836128f9565b51849063ffffffff6129f316565b6000611d38611ee783612a30565b849063ffffffff61216516565b600192915050565b6000611f15826080015183612ab990919063ffffffff16565b506001919050565b6000611f15826060015183612ab990919063ffffffff16565b6000611f41826128f9565b606084015250600192915050565b6000611f5a826128f9565b835250600192915050565b6000611f70826129e8565b611f7c5750600061199e565b825115611cf257611f8c836128f9565b84525060019392505050565b6000611f15611fbd611fb0611fab612ac7565b6128f9565b5160208501515114612a07565b839063ffffffff61216516565b6000611d38838363ffffffff612ab916565b6000611d38838363ffffffff61214316565b600080611ff96134ca565b8451841061204e576040805162461bcd60e51b815260206004820152601960248201527f44617461206f6666736574206f7574206f6620626f756e647300000000000000604482015290519081900360640190fd5b6000849050600086828151811061206157fe5b016020015160019092019160f81c90506000816120a3576120828884612b21565b909350905060008361209383612575565b9197509550935061213c92505050565b60ff8216600214156120ca576120b98884612b21565b909350905060008361209383612b48565b600360ff8316108015906120e15750600c60ff8316105b1561211d576002198201606060006120fa838c88612ba2565b90975092509050808661210c846125cf565b98509850985050505050505061213c565b8160ff166127100160006121316000612575565b919750955093505050505b9250925092565b6121598260400151612154836128f9565b612c5d565b82604001819052505050565b6121768260200151612154836128f9565b82602001819052505050565b6000612194838363ffffffff61216516565b611d38838363ffffffff61216516565b60006121b6848363ffffffff61216516565b6121c6848463ffffffff61216516565b611cf2848363ffffffff61216516565b60006121e8858363ffffffff61216516565b6121f8858463ffffffff61216516565b612208858563ffffffff61216516565b611e30858363ffffffff61216516565b60006121c6848463ffffffff61216516565b600061223c858563ffffffff61216516565b612208858463ffffffff61216516565b6000612257836129e8565b1580612269575061226782612d13565b155b156122765750600061199e565b61227f82612d22565b60ff16836000015111156122955750600061199e565b611cf282602001518460000151815181106122ac57fe5b60200260200101518561216590919063ffffffff16565b60006122ce83612d13565b15806122e057506122de846129e8565b155b156122ed57506000611b7c565b6122f683612d22565b60ff168460000151111561230c57506000611b7c565b81836020015185600001518151811061232157fe5b6020908102919091010152611e30858463ffffffff61216516565b600061234782612d13565b61235357506000611d3c565b611d3861235f83612d22565b849060ff1663ffffffff6129f316565b50600190565b6000806123806134ee565b612389846128f9565b516001969095509350505050565b60008060008060008060006123ab88612d13565b6123bf57600096509450909250905061256c565b6123e088602001516001815181106123d357fe5b60200260200101516129e8565b6123f457600096509450909250905061256c565b61240888602001516002815181106123d357fe5b61241c57600096509450909250905061256c565b61243088602001516003815181106123d357fe5b61244457600096509450909250905061256c565b876020015160018151811061245557fe5b60200260200101516000015160001b9250876020015160028151811061247757fe5b602002602001015160000151915073__$6b4cc75dad3e0abd6ad83b3d907747c608$__63ccf69dd76124a88a6128f9565b6000015185858c602001516003815181106124bf57fe5b60200260200101516000015160001b6040518563ffffffff1660e01b815260040180858152602001846affffffffffffffffffffff19166affffffffffffffffffffff1916815260200183815260200182815260200194505050505060206040518083038186803b15801561253357600080fd5b505af4158015612547573d6000803e3d6000fd5b505050506040513d602081101561255d57600080fd5b50516001975095509193509150505b92959194509250565b61257d6134ca565b6040805160608101825283815281516000808252602082810190945291928301916125be565b6125ab6134ca565b8152602001906001900390816125a35790505b508152600060209091015292915050565b6125d76134ca565b6125e18251612d31565b612632576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b50604080516060810182526000815260208101839052915160030160ff169082015290565b8051600090612665846128f9565b5114156126b9576040805162461bcd60e51b815260206004820152601d60248201527f496e626f7820696e737472756374696f6e2077617320626c6f636b6564000000604482015290519081900360640190fd5b611cf2848363ffffffff612ab916565b600260c090910152565b600160c090910152565b600060028260c0015114156126f457506000611510565b60018260c00151141561270957506001611510565b508051516020808301515160408085015151606080870151516080808901515160a0808b0151518751808b019b909b528a8801989098529389019490945287015285015260c0808501929092528051808503909201825260e09093019092528151910120611510565b600060c090910152565b600080612787613469565b61278f613469565b600060c082018190526127a28787612d38565b84529650905080156127ba579350849250905061213c565b6127c48787612d38565b60208501529650905080156127df579350849250905061213c565b6127e98787612d38565b6040850152965090508015612804579350849250905061213c565b61280e8787612d38565b6060850152965090508015612829579350849250905061213c565b6128338787612d38565b608085015296509050801561284e579350849250905061213c565b6128588787612d38565b60a0850152965090508015612873579350849250905061213c565b506000969495509392505050565b612889613469565b6040518060e0016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c001518152509050919050565b60008060006128f08460ff16612d76565b50949350505050565b6129016134ee565b6040820151600c60ff90911610612953576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b604082015160ff16612980576040518060200160405280612977846000015161321d565b90529050611510565b604082015160ff16600214156129a55750604080516020810190915281518152611510565b600360ff16826040015160ff16101580156129c957506040820151600c60ff909116105b156129e65760405180602001604052806129778460200151613241565bfe5b6040015160ff161590565b6121768260200151612154611fab84612575565b612a0f6134ca565b8115612a2657612a1f6001612575565b9050611510565b612a1f6000612575565b612a386134ca565b816040015160ff1660021415612a7f5760405162461bcd60e51b81526004018080602001828103825260218152602001806135496021913960400191505060405180910390fd5b604082015160ff16612a9557612a1f6000612575565b816040015160ff1660011415612aaf57612a1f6001612575565b612a1f6003612575565b612176826020015182612c5d565b612acf6134ca565b60408051606081018252600080825282518181526020818101909452919283019190612b11565b612afe6134ca565b815260200190600190039081612af65790505b5081526003602090910152905090565b6000808281612b36868363ffffffff61338d16565b60209290920196919550909350505050565b612b506134ca565b604080516060810182528381528151600080825260208281019094529192830191612b91565b612b7e6134ca565b815260200190600190039081612b765790505b508152600260209091015292915050565b60008060606000849050600060608860ff16604051908082528060200260200182016040528015612bed57816020015b612bda6134ca565b815260200190600190039081612bd25790505b50905060005b8960ff168160ff161015612c4757612c0b8985611fee565b8451859060ff8616908110612c1c57fe5b6020908102919091010152945092508215612c3f57509094509092509050612c54565b600101612bf3565b5060009550919350909150505b93509350939050565b612c656134ee565b6040805160028082526060828101909352816020015b612c836134ee565b815260200190600190039081612c7b5790505090508281600081518110612ca657fe5b60200260200101819052508381600181518110612cbf57fe5b60200260200101819052506040518060200160405280612d096040518060400160405280612cf08860000151612b48565b8152602001612d028960000151612b48565b90526133a9565b9052949350505050565b6000611d3c8260400151613428565b6000611d3c8260400151613446565b6008101590565b600080612d436134ee565b836000612d56878363ffffffff61338d16565b604080516020808201909252918252600099930197509550909350505050565b6000806001831415612d8e5750600290506001613218565b6002831415612da35750600290506001613218565b6003831415612db85750600290506001613218565b6004831415612dcd5750600290506001613218565b6005831415612de25750600290506001613218565b6006831415612df75750600290506001613218565b6007831415612e0c5750600290506001613218565b6008831415612e215750600390506001613218565b6009831415612e365750600390506001613218565b600a831415612e4b5750600290506001613218565b6010831415612e605750600290506001613218565b6011831415612e755750600290506001613218565b6012831415612e8a5750600290506001613218565b6013831415612e9f5750600290506001613218565b6014831415612eb45750600290506001613218565b6015831415612ec857506001905080613218565b6016831415612edd5750600290506001613218565b6017831415612ef25750600290506001613218565b6018831415612f075750600290506001613218565b6019831415612f1b57506001905080613218565b601a831415612f305750600290506001613218565b601b831415612f455750600290506001613218565b6020831415612f5957506001905080613218565b6021831415612f6d57506001905080613218565b6030831415612f825750600190506000613218565b6031831415612f975750600090506001613218565b6032831415612fac5750600090506001613218565b6033831415612fc15750600190506000613218565b6034831415612fd65750600190506000613218565b6035831415612feb5750600290506000613218565b60368314156130005750600090506001613218565b60378314156130155750600090506001613218565b603883141561302a5750600190506000613218565b603983141561303f5750600090506001613218565b603a8314156130545750600090506001613218565b603b83141561306857506000905080613218565b603c83141561307d5750600090506001613218565b603d8314156130925750600190506000613218565b60408314156130a75750600190506002613218565b60418314156130bc5750600290506003613218565b60428314156130d15750600390506004613218565b60438314156130e557506002905080613218565b60448314156130f957506003905080613218565b605083141561310e5750600290506001613218565b60518314156131235750600390506001613218565b605283141561313757506001905080613218565b606083141561314b57506000905080613218565b60618314156131605750600190506000613218565b60708314156131755750600190506000613218565b607183141561318957506001905080613218565b607283141561319e5750600090506001613218565b60738314156131b257506001905080613218565b60748314156131c657506000905080613218565b60758314156131da57506000905080613218565b6040805162461bcd60e51b815260206004820152600e60248201526d496e76616c6964206f70636f646560901b604482015290519081900360640190fd5b915091565b60408051602080820193909352815180820384018152908201909152805191012090565b6000600882511115613291576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b606082516040519080825280602002602001820160405280156132be578160200160208202803883390190505b50805190915060005b8181101561331a576132d76134ee565b6132f38683815181106132e657fe5b60200260200101516128f9565b9050806000015184838151811061330657fe5b6020908102919091010152506001016132c7565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b8381101561336357818101518382015260200161334b565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b600081602001835110156133a057600080fd5b50016020015190565b60408051600280825260608281019093526000929190816020015b6133cc6134ca565b8152602001906001900390816133c4575050805190915060005b8181101561341e578481600281106133fa57fe5b602002015183828151811061340b57fe5b60209081029190910101526001016133e6565b50611b7c82613241565b6000600c60ff8316108015611d3c575050600360ff91909116101590565b600061345182613428565b1561346157506002198101611510565b506001611510565b6040518060e0016040528061347c6134ee565b81526020016134896134ee565b81526020016134966134ee565b81526020016134a36134ee565b81526020016134b06134ee565b81526020016134bd6134ee565b8152602001600081525090565b60405180606001604052806000815260200160608152602001600060ff1681525090565b6040805160208101909152600081529056fe50726f6f6620686164206e6f6e206d61746368696e672073746172742073746174654c6f67206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f56616c7565206d757374206861766520612076616c6964207479706520636f64654c6f676765642076616c756520646f65736e2774206d61746368206f7574707574206c6f6753656e64206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f73656e74206d65737361676520646f65736e2774206d61746368206f7574707574206d657361676550726f6f66206f6620617578706f702068616420626164206175782076616c7565a265627a7a72305820ba3ae81076abec3e2b54e8da456aab25c7e5d6eaa7e24923ece0df0237e5561c64736f6c634300050a0032"

// DeployOneStepProof deploys a new Ethereum contract, binding an instance of OneStepProof to it.
func DeployOneStepProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProof, error) {
	parsed, err := abi.JSON(strings.NewReader(OneStepProofABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	arbValueAddr, _, _, _ := DeployArbValue(auth, backend)
	OneStepProofBin = strings.Replace(OneStepProofBin, "__$0d86abb4a722a612872fb80f4c7e7e95bd$__", arbValueAddr.String()[2:], -1)

	arbProtocolAddr, _, _, _ := DeployArbProtocol(auth, backend)
	OneStepProofBin = strings.Replace(OneStepProofBin, "__$6b4cc75dad3e0abd6ad83b3d907747c608$__", arbProtocolAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OneStepProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProof{OneStepProofCaller: OneStepProofCaller{contract: contract}, OneStepProofTransactor: OneStepProofTransactor{contract: contract}, OneStepProofFilterer: OneStepProofFilterer{contract: contract}}, nil
}

// OneStepProof is an auto generated Go binding around an Ethereum contract.
type OneStepProof struct {
	OneStepProofCaller     // Read-only binding to the contract
	OneStepProofTransactor // Write-only binding to the contract
	OneStepProofFilterer   // Log filterer for contract events
}

// OneStepProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProofSession struct {
	Contract     *OneStepProof     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OneStepProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProofCallerSession struct {
	Contract *OneStepProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// OneStepProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProofTransactorSession struct {
	Contract     *OneStepProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OneStepProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProofRaw struct {
	Contract *OneStepProof // Generic contract binding to access the raw methods on
}

// OneStepProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProofCallerRaw struct {
	Contract *OneStepProofCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProofTransactorRaw struct {
	Contract *OneStepProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProof creates a new instance of OneStepProof, bound to a specific deployed contract.
func NewOneStepProof(address common.Address, backend bind.ContractBackend) (*OneStepProof, error) {
	contract, err := bindOneStepProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProof{OneStepProofCaller: OneStepProofCaller{contract: contract}, OneStepProofTransactor: OneStepProofTransactor{contract: contract}, OneStepProofFilterer: OneStepProofFilterer{contract: contract}}, nil
}

// NewOneStepProofCaller creates a new read-only instance of OneStepProof, bound to a specific deployed contract.
func NewOneStepProofCaller(address common.Address, caller bind.ContractCaller) (*OneStepProofCaller, error) {
	contract, err := bindOneStepProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofCaller{contract: contract}, nil
}

// NewOneStepProofTransactor creates a new write-only instance of OneStepProof, bound to a specific deployed contract.
func NewOneStepProofTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProofTransactor, error) {
	contract, err := bindOneStepProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofTransactor{contract: contract}, nil
}

// NewOneStepProofFilterer creates a new log filterer instance of OneStepProof, bound to a specific deployed contract.
func NewOneStepProofFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProofFilterer, error) {
	contract, err := bindOneStepProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProofFilterer{contract: contract}, nil
}

// bindOneStepProof binds a generic wrapper to an already deployed contract.
func bindOneStepProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OneStepProofABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProof *OneStepProofRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OneStepProof.Contract.OneStepProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProof *OneStepProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProof.Contract.OneStepProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProof *OneStepProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProof.Contract.OneStepProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProof *OneStepProofCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OneStepProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProof *OneStepProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProof *OneStepProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProof.Contract.contract.Transact(opts, method, params...)
}

// ValidateProof is a free data retrieval call binding the contract method 0x0eca9f13.
//
// Solidity: function validateProof(bytes32[7] fields, uint64[2] timeBounds, bytes21[] tokenTypes, uint256[] beforeValues, uint256[] messageValue, bytes proof) constant returns(uint256)
func (_OneStepProof *OneStepProofCaller) ValidateProof(opts *bind.CallOpts, fields [7][32]byte, timeBounds [2]uint64, tokenTypes [][21]byte, beforeValues []*big.Int, messageValue []*big.Int, proof []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OneStepProof.contract.Call(opts, out, "validateProof", fields, timeBounds, tokenTypes, beforeValues, messageValue, proof)
	return *ret0, err
}

// ValidateProof is a free data retrieval call binding the contract method 0x0eca9f13.
//
// Solidity: function validateProof(bytes32[7] fields, uint64[2] timeBounds, bytes21[] tokenTypes, uint256[] beforeValues, uint256[] messageValue, bytes proof) constant returns(uint256)
func (_OneStepProof *OneStepProofSession) ValidateProof(fields [7][32]byte, timeBounds [2]uint64, tokenTypes [][21]byte, beforeValues []*big.Int, messageValue []*big.Int, proof []byte) (*big.Int, error) {
	return _OneStepProof.Contract.ValidateProof(&_OneStepProof.CallOpts, fields, timeBounds, tokenTypes, beforeValues, messageValue, proof)
}

// ValidateProof is a free data retrieval call binding the contract method 0x0eca9f13.
//
// Solidity: function validateProof(bytes32[7] fields, uint64[2] timeBounds, bytes21[] tokenTypes, uint256[] beforeValues, uint256[] messageValue, bytes proof) constant returns(uint256)
func (_OneStepProof *OneStepProofCallerSession) ValidateProof(fields [7][32]byte, timeBounds [2]uint64, tokenTypes [][21]byte, beforeValues []*big.Int, messageValue []*big.Int, proof []byte) (*big.Int, error) {
	return _OneStepProof.Contract.ValidateProof(&_OneStepProof.CallOpts, fields, timeBounds, tokenTypes, beforeValues, messageValue, proof)
}

// OneStepProofSawMachineIterator is returned from FilterSawMachine and is used to iterate over the raw logs and unpacked data for SawMachine events raised by the OneStepProof contract.
type OneStepProofSawMachineIterator struct {
	Event *OneStepProofSawMachine // Event containing the contract specifics and raw log

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
func (it *OneStepProofSawMachineIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneStepProofSawMachine)
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
		it.Event = new(OneStepProofSawMachine)
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
func (it *OneStepProofSawMachineIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneStepProofSawMachineIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneStepProofSawMachine represents a SawMachine event raised by the OneStepProof contract.
type OneStepProofSawMachine struct {
	InstructionStack [32]byte
	DataStack        [32]byte
	AuxStack         [32]byte
	Register         [32]byte
	StaticHash       [32]byte
	ErrHandler       [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSawMachine is a free log retrieval operation binding the contract event 0x10d11f456a57c1ced446abc92b6cfa3854cd21f069b29f283252b23223d03080.
//
// Solidity: event SawMachine(bytes32 instructionStack, bytes32 dataStack, bytes32 auxStack, bytes32 register, bytes32 staticHash, bytes32 errHandler)
func (_OneStepProof *OneStepProofFilterer) FilterSawMachine(opts *bind.FilterOpts) (*OneStepProofSawMachineIterator, error) {

	logs, sub, err := _OneStepProof.contract.FilterLogs(opts, "SawMachine")
	if err != nil {
		return nil, err
	}
	return &OneStepProofSawMachineIterator{contract: _OneStepProof.contract, event: "SawMachine", logs: logs, sub: sub}, nil
}

// WatchSawMachine is a free log subscription operation binding the contract event 0x10d11f456a57c1ced446abc92b6cfa3854cd21f069b29f283252b23223d03080.
//
// Solidity: event SawMachine(bytes32 instructionStack, bytes32 dataStack, bytes32 auxStack, bytes32 register, bytes32 staticHash, bytes32 errHandler)
func (_OneStepProof *OneStepProofFilterer) WatchSawMachine(opts *bind.WatchOpts, sink chan<- *OneStepProofSawMachine) (event.Subscription, error) {

	logs, sub, err := _OneStepProof.contract.WatchLogs(opts, "SawMachine")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneStepProofSawMachine)
				if err := _OneStepProof.contract.UnpackLog(event, "SawMachine", log); err != nil {
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

// ParseSawMachine is a log parse operation binding the contract event 0x10d11f456a57c1ced446abc92b6cfa3854cd21f069b29f283252b23223d03080.
//
// Solidity: event SawMachine(bytes32 instructionStack, bytes32 dataStack, bytes32 auxStack, bytes32 register, bytes32 staticHash, bytes32 errHandler)
func (_OneStepProof *OneStepProofFilterer) ParseSawMachine(log types.Log) (*OneStepProofSawMachine, error) {
	event := new(OneStepProofSawMachine)
	if err := _OneStepProof.contract.UnpackLog(event, "SawMachine", log); err != nil {
		return nil, err
	}
	return event, nil
}
