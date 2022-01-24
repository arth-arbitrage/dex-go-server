// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapv2

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IUniswapv2SwapABI is the input ABI used to generate the binding from.
const IUniswapv2SwapABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedOut\",\"type\":\"uint256\"}],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IUniswapv2Swap is an auto generated Go binding around an Ethereum contract.
type IUniswapv2Swap struct {
	IUniswapv2SwapCaller     // Read-only binding to the contract
	IUniswapv2SwapTransactor // Write-only binding to the contract
	IUniswapv2SwapFilterer   // Log filterer for contract events
}

// IUniswapv2SwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUniswapv2SwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapv2SwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUniswapv2SwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapv2SwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUniswapv2SwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapv2SwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUniswapv2SwapSession struct {
	Contract     *IUniswapv2Swap   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IUniswapv2SwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUniswapv2SwapCallerSession struct {
	Contract *IUniswapv2SwapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IUniswapv2SwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUniswapv2SwapTransactorSession struct {
	Contract     *IUniswapv2SwapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IUniswapv2SwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUniswapv2SwapRaw struct {
	Contract *IUniswapv2Swap // Generic contract binding to access the raw methods on
}

// IUniswapv2SwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUniswapv2SwapCallerRaw struct {
	Contract *IUniswapv2SwapCaller // Generic read-only contract binding to access the raw methods on
}

// IUniswapv2SwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUniswapv2SwapTransactorRaw struct {
	Contract *IUniswapv2SwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUniswapv2Swap creates a new instance of IUniswapv2Swap, bound to a specific deployed contract.
func NewIUniswapv2Swap(address common.Address, backend bind.ContractBackend) (*IUniswapv2Swap, error) {
	contract, err := bindIUniswapv2Swap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUniswapv2Swap{IUniswapv2SwapCaller: IUniswapv2SwapCaller{contract: contract}, IUniswapv2SwapTransactor: IUniswapv2SwapTransactor{contract: contract}, IUniswapv2SwapFilterer: IUniswapv2SwapFilterer{contract: contract}}, nil
}

// NewIUniswapv2SwapCaller creates a new read-only instance of IUniswapv2Swap, bound to a specific deployed contract.
func NewIUniswapv2SwapCaller(address common.Address, caller bind.ContractCaller) (*IUniswapv2SwapCaller, error) {
	contract, err := bindIUniswapv2Swap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapv2SwapCaller{contract: contract}, nil
}

// NewIUniswapv2SwapTransactor creates a new write-only instance of IUniswapv2Swap, bound to a specific deployed contract.
func NewIUniswapv2SwapTransactor(address common.Address, transactor bind.ContractTransactor) (*IUniswapv2SwapTransactor, error) {
	contract, err := bindIUniswapv2Swap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapv2SwapTransactor{contract: contract}, nil
}

// NewIUniswapv2SwapFilterer creates a new log filterer instance of IUniswapv2Swap, bound to a specific deployed contract.
func NewIUniswapv2SwapFilterer(address common.Address, filterer bind.ContractFilterer) (*IUniswapv2SwapFilterer, error) {
	contract, err := bindIUniswapv2Swap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUniswapv2SwapFilterer{contract: contract}, nil
}

// bindIUniswapv2Swap binds a generic wrapper to an already deployed contract.
func bindIUniswapv2Swap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IUniswapv2SwapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapv2Swap *IUniswapv2SwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapv2Swap.Contract.IUniswapv2SwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapv2Swap *IUniswapv2SwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapv2Swap.Contract.IUniswapv2SwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapv2Swap *IUniswapv2SwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapv2Swap.Contract.IUniswapv2SwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapv2Swap *IUniswapv2SwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapv2Swap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapv2Swap *IUniswapv2SwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapv2Swap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapv2Swap *IUniswapv2SwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapv2Swap.Contract.contract.Transact(opts, method, params...)
}

// Exchange is a paid mutator transaction binding the contract method 0x236e06f6.
//
// Solidity: function exchange(address pair, uint256 swapAmount, uint256 expectedOut) returns(uint256)
func (_IUniswapv2Swap *IUniswapv2SwapTransactor) Exchange(opts *bind.TransactOpts, pair common.Address, swapAmount *big.Int, expectedOut *big.Int) (*types.Transaction, error) {
	return _IUniswapv2Swap.contract.Transact(opts, "exchange", pair, swapAmount, expectedOut)
}

// Exchange is a paid mutator transaction binding the contract method 0x236e06f6.
//
// Solidity: function exchange(address pair, uint256 swapAmount, uint256 expectedOut) returns(uint256)
func (_IUniswapv2Swap *IUniswapv2SwapSession) Exchange(pair common.Address, swapAmount *big.Int, expectedOut *big.Int) (*types.Transaction, error) {
	return _IUniswapv2Swap.Contract.Exchange(&_IUniswapv2Swap.TransactOpts, pair, swapAmount, expectedOut)
}

// Exchange is a paid mutator transaction binding the contract method 0x236e06f6.
//
// Solidity: function exchange(address pair, uint256 swapAmount, uint256 expectedOut) returns(uint256)
func (_IUniswapv2Swap *IUniswapv2SwapTransactorSession) Exchange(pair common.Address, swapAmount *big.Int, expectedOut *big.Int) (*types.Transaction, error) {
	return _IUniswapv2Swap.Contract.Exchange(&_IUniswapv2Swap.TransactOpts, pair, swapAmount, expectedOut)
}
