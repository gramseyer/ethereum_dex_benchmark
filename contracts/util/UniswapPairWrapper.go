// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package UniswapPairWrapper

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

// UniswapPairWrapperMetaData contains all meta data concerning the UniswapPairWrapper contract.
var UniswapPairWrapperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair_addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tok0_addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tok1_addr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0in\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1in\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1out\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516109763803806109768339818101604052810190610032919061015f565b826000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050506101b2565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061012c82610101565b9050919050565b61013c81610121565b811461014757600080fd5b50565b60008151905061015981610133565b92915050565b600080600060608486031215610178576101776100fc565b5b60006101868682870161014a565b93505060206101978682870161014a565b92505060406101a88682870161014a565b9150509250925092565b6107b5806101c16000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80631b2ef1ca1461003b5780635673b02d14610057575b600080fd5b610055600480360381019061005091906103fe565b610073565b005b610071600480360381019061006c919061043e565b61011e565b005b61007d82826101bd565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636a627842336040518263ffffffff1660e01b81526004016100d691906104e6565b6020604051808303816000875af11580156100f5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101199190610516565b505050565b61012884846101bd565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663022c0d9f8383336040518463ffffffff1660e01b815260040161018593929190610589565b600060405180830381600087803b15801561019f57600080fd5b505af11580156101b3573d6000803e3d6000fd5b5050505050505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3360008054906101000a900473ffffffffffffffffffffffffffffffffffffffff16856040518463ffffffff1660e01b815260040161023c939291906105d3565b6020604051808303816000875af115801561025b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061027f9190610642565b6102be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102b5906106cc565b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3360008054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518463ffffffff1660e01b815260040161033d939291906105d3565b6020604051808303816000875af115801561035c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103809190610642565b6103bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103b690610738565b60405180910390fd5b5050565b600080fd5b6000819050919050565b6103db816103c8565b81146103e657600080fd5b50565b6000813590506103f8816103d2565b92915050565b60008060408385031215610415576104146103c3565b5b6000610423858286016103e9565b9250506020610434858286016103e9565b9150509250929050565b60008060008060808587031215610458576104576103c3565b5b6000610466878288016103e9565b9450506020610477878288016103e9565b9350506040610488878288016103e9565b9250506060610499878288016103e9565b91505092959194509250565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006104d0826104a5565b9050919050565b6104e0816104c5565b82525050565b60006020820190506104fb60008301846104d7565b92915050565b600081519050610510816103d2565b92915050565b60006020828403121561052c5761052b6103c3565b5b600061053a84828501610501565b91505092915050565b61054c816103c8565b82525050565b600082825260208201905092915050565b50565b6000610573600083610552565b915061057e82610563565b600082019050919050565b600060808201905061059e6000830186610543565b6105ab6020830185610543565b6105b860408301846104d7565b81810360608301526105c981610566565b9050949350505050565b60006060820190506105e860008301866104d7565b6105f560208301856104d7565b6106026040830184610543565b949350505050565b60008115159050919050565b61061f8161060a565b811461062a57600080fd5b50565b60008151905061063c81610616565b92915050565b600060208284031215610658576106576103c3565b5b60006106668482850161062d565b91505092915050565b600082825260208201905092915050565b7f5452414e53464552204641494c45440000000000000000000000000000000000600082015250565b60006106b6600f8361066f565b91506106c182610680565b602082019050919050565b600060208201905081810360008301526106e5816106a9565b9050919050565b7f5452414e534645525f4641494c45440000000000000000000000000000000000600082015250565b6000610722600f8361066f565b915061072d826106ec565b602082019050919050565b6000602082019050818103600083015261075181610715565b905091905056fea26469706673582212208748890ce6b1d43076698f0f2c16db46f33a6e21cc8a09780ff9132a5c0cc52264736f6c637829302e382e31382d646576656c6f702e323032322e31302e31312b636f6d6d69742e3233386163346664005a",
}

// UniswapPairWrapperABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapPairWrapperMetaData.ABI instead.
var UniswapPairWrapperABI = UniswapPairWrapperMetaData.ABI

// UniswapPairWrapperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UniswapPairWrapperMetaData.Bin instead.
var UniswapPairWrapperBin = UniswapPairWrapperMetaData.Bin

// DeployUniswapPairWrapper deploys a new Ethereum contract, binding an instance of UniswapPairWrapper to it.
func DeployUniswapPairWrapper(auth *bind.TransactOpts, backend bind.ContractBackend, pair_addr common.Address, tok0_addr common.Address, tok1_addr common.Address) (common.Address, *types.Transaction, *UniswapPairWrapper, error) {
	parsed, err := UniswapPairWrapperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UniswapPairWrapperBin), backend, pair_addr, tok0_addr, tok1_addr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UniswapPairWrapper{UniswapPairWrapperCaller: UniswapPairWrapperCaller{contract: contract}, UniswapPairWrapperTransactor: UniswapPairWrapperTransactor{contract: contract}, UniswapPairWrapperFilterer: UniswapPairWrapperFilterer{contract: contract}}, nil
}

// UniswapPairWrapper is an auto generated Go binding around an Ethereum contract.
type UniswapPairWrapper struct {
	UniswapPairWrapperCaller     // Read-only binding to the contract
	UniswapPairWrapperTransactor // Write-only binding to the contract
	UniswapPairWrapperFilterer   // Log filterer for contract events
}

// UniswapPairWrapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapPairWrapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapPairWrapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapPairWrapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapPairWrapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapPairWrapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapPairWrapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapPairWrapperSession struct {
	Contract     *UniswapPairWrapper // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// UniswapPairWrapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapPairWrapperCallerSession struct {
	Contract *UniswapPairWrapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// UniswapPairWrapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapPairWrapperTransactorSession struct {
	Contract     *UniswapPairWrapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// UniswapPairWrapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapPairWrapperRaw struct {
	Contract *UniswapPairWrapper // Generic contract binding to access the raw methods on
}

// UniswapPairWrapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapPairWrapperCallerRaw struct {
	Contract *UniswapPairWrapperCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapPairWrapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapPairWrapperTransactorRaw struct {
	Contract *UniswapPairWrapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapPairWrapper creates a new instance of UniswapPairWrapper, bound to a specific deployed contract.
func NewUniswapPairWrapper(address common.Address, backend bind.ContractBackend) (*UniswapPairWrapper, error) {
	contract, err := bindUniswapPairWrapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapPairWrapper{UniswapPairWrapperCaller: UniswapPairWrapperCaller{contract: contract}, UniswapPairWrapperTransactor: UniswapPairWrapperTransactor{contract: contract}, UniswapPairWrapperFilterer: UniswapPairWrapperFilterer{contract: contract}}, nil
}

// NewUniswapPairWrapperCaller creates a new read-only instance of UniswapPairWrapper, bound to a specific deployed contract.
func NewUniswapPairWrapperCaller(address common.Address, caller bind.ContractCaller) (*UniswapPairWrapperCaller, error) {
	contract, err := bindUniswapPairWrapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapPairWrapperCaller{contract: contract}, nil
}

// NewUniswapPairWrapperTransactor creates a new write-only instance of UniswapPairWrapper, bound to a specific deployed contract.
func NewUniswapPairWrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapPairWrapperTransactor, error) {
	contract, err := bindUniswapPairWrapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapPairWrapperTransactor{contract: contract}, nil
}

// NewUniswapPairWrapperFilterer creates a new log filterer instance of UniswapPairWrapper, bound to a specific deployed contract.
func NewUniswapPairWrapperFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapPairWrapperFilterer, error) {
	contract, err := bindUniswapPairWrapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapPairWrapperFilterer{contract: contract}, nil
}

// bindUniswapPairWrapper binds a generic wrapper to an already deployed contract.
func bindUniswapPairWrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapPairWrapperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapPairWrapper *UniswapPairWrapperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapPairWrapper.Contract.UniswapPairWrapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapPairWrapper *UniswapPairWrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapPairWrapper.Contract.UniswapPairWrapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapPairWrapper *UniswapPairWrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapPairWrapper.Contract.UniswapPairWrapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapPairWrapper *UniswapPairWrapperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapPairWrapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapPairWrapper *UniswapPairWrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapPairWrapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapPairWrapper *UniswapPairWrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapPairWrapper.Contract.contract.Transact(opts, method, params...)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 amount0, uint256 amount1) returns()
func (_UniswapPairWrapper *UniswapPairWrapperTransactor) Mint(opts *bind.TransactOpts, amount0 *big.Int, amount1 *big.Int) (*types.Transaction, error) {
	return _UniswapPairWrapper.contract.Transact(opts, "mint", amount0, amount1)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 amount0, uint256 amount1) returns()
func (_UniswapPairWrapper *UniswapPairWrapperSession) Mint(amount0 *big.Int, amount1 *big.Int) (*types.Transaction, error) {
	return _UniswapPairWrapper.Contract.Mint(&_UniswapPairWrapper.TransactOpts, amount0, amount1)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 amount0, uint256 amount1) returns()
func (_UniswapPairWrapper *UniswapPairWrapperTransactorSession) Mint(amount0 *big.Int, amount1 *big.Int) (*types.Transaction, error) {
	return _UniswapPairWrapper.Contract.Mint(&_UniswapPairWrapper.TransactOpts, amount0, amount1)
}

// Swap is a paid mutator transaction binding the contract method 0x5673b02d.
//
// Solidity: function swap(uint256 amount0in, uint256 amount1in, uint256 amount0out, uint256 amount1out) returns()
func (_UniswapPairWrapper *UniswapPairWrapperTransactor) Swap(opts *bind.TransactOpts, amount0in *big.Int, amount1in *big.Int, amount0out *big.Int, amount1out *big.Int) (*types.Transaction, error) {
	return _UniswapPairWrapper.contract.Transact(opts, "swap", amount0in, amount1in, amount0out, amount1out)
}

// Swap is a paid mutator transaction binding the contract method 0x5673b02d.
//
// Solidity: function swap(uint256 amount0in, uint256 amount1in, uint256 amount0out, uint256 amount1out) returns()
func (_UniswapPairWrapper *UniswapPairWrapperSession) Swap(amount0in *big.Int, amount1in *big.Int, amount0out *big.Int, amount1out *big.Int) (*types.Transaction, error) {
	return _UniswapPairWrapper.Contract.Swap(&_UniswapPairWrapper.TransactOpts, amount0in, amount1in, amount0out, amount1out)
}

// Swap is a paid mutator transaction binding the contract method 0x5673b02d.
//
// Solidity: function swap(uint256 amount0in, uint256 amount1in, uint256 amount0out, uint256 amount1out) returns()
func (_UniswapPairWrapper *UniswapPairWrapperTransactorSession) Swap(amount0in *big.Int, amount1in *big.Int, amount0out *big.Int, amount1out *big.Int) (*types.Transaction, error) {
	return _UniswapPairWrapper.Contract.Swap(&_UniswapPairWrapper.TransactOpts, amount0in, amount1in, amount0out, amount1out)
}
