// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbitrage

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

// IArthArbV1MultiSwapABI is the input ABI used to generate the binding from.
const IArthArbV1MultiSwapABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"loanAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"arbitrage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IArthArbV1MultiSwap is an auto generated Go binding around an Ethereum contract.
type IArthArbV1MultiSwap struct {
	IArthArbV1MultiSwapCaller     // Read-only binding to the contract
	IArthArbV1MultiSwapTransactor // Write-only binding to the contract
	IArthArbV1MultiSwapFilterer   // Log filterer for contract events
}

// IArthArbV1MultiSwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type IArthArbV1MultiSwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArthArbV1MultiSwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IArthArbV1MultiSwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArthArbV1MultiSwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IArthArbV1MultiSwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArthArbV1MultiSwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IArthArbV1MultiSwapSession struct {
	Contract     *IArthArbV1MultiSwap // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IArthArbV1MultiSwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IArthArbV1MultiSwapCallerSession struct {
	Contract *IArthArbV1MultiSwapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IArthArbV1MultiSwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IArthArbV1MultiSwapTransactorSession struct {
	Contract     *IArthArbV1MultiSwapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IArthArbV1MultiSwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type IArthArbV1MultiSwapRaw struct {
	Contract *IArthArbV1MultiSwap // Generic contract binding to access the raw methods on
}

// IArthArbV1MultiSwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IArthArbV1MultiSwapCallerRaw struct {
	Contract *IArthArbV1MultiSwapCaller // Generic read-only contract binding to access the raw methods on
}

// IArthArbV1MultiSwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IArthArbV1MultiSwapTransactorRaw struct {
	Contract *IArthArbV1MultiSwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIArthArbV1MultiSwap creates a new instance of IArthArbV1MultiSwap, bound to a specific deployed contract.
func NewIArthArbV1MultiSwap(address common.Address, backend bind.ContractBackend) (*IArthArbV1MultiSwap, error) {
	contract, err := bindIArthArbV1MultiSwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IArthArbV1MultiSwap{IArthArbV1MultiSwapCaller: IArthArbV1MultiSwapCaller{contract: contract}, IArthArbV1MultiSwapTransactor: IArthArbV1MultiSwapTransactor{contract: contract}, IArthArbV1MultiSwapFilterer: IArthArbV1MultiSwapFilterer{contract: contract}}, nil
}

// NewIArthArbV1MultiSwapCaller creates a new read-only instance of IArthArbV1MultiSwap, bound to a specific deployed contract.
func NewIArthArbV1MultiSwapCaller(address common.Address, caller bind.ContractCaller) (*IArthArbV1MultiSwapCaller, error) {
	contract, err := bindIArthArbV1MultiSwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IArthArbV1MultiSwapCaller{contract: contract}, nil
}

// NewIArthArbV1MultiSwapTransactor creates a new write-only instance of IArthArbV1MultiSwap, bound to a specific deployed contract.
func NewIArthArbV1MultiSwapTransactor(address common.Address, transactor bind.ContractTransactor) (*IArthArbV1MultiSwapTransactor, error) {
	contract, err := bindIArthArbV1MultiSwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IArthArbV1MultiSwapTransactor{contract: contract}, nil
}

// NewIArthArbV1MultiSwapFilterer creates a new log filterer instance of IArthArbV1MultiSwap, bound to a specific deployed contract.
func NewIArthArbV1MultiSwapFilterer(address common.Address, filterer bind.ContractFilterer) (*IArthArbV1MultiSwapFilterer, error) {
	contract, err := bindIArthArbV1MultiSwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IArthArbV1MultiSwapFilterer{contract: contract}, nil
}

// bindIArthArbV1MultiSwap binds a generic wrapper to an already deployed contract.
func bindIArthArbV1MultiSwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IArthArbV1MultiSwapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IArthArbV1MultiSwap *IArthArbV1MultiSwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IArthArbV1MultiSwap.Contract.IArthArbV1MultiSwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IArthArbV1MultiSwap *IArthArbV1MultiSwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IArthArbV1MultiSwap.Contract.IArthArbV1MultiSwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IArthArbV1MultiSwap *IArthArbV1MultiSwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IArthArbV1MultiSwap.Contract.IArthArbV1MultiSwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IArthArbV1MultiSwap *IArthArbV1MultiSwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IArthArbV1MultiSwap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IArthArbV1MultiSwap *IArthArbV1MultiSwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IArthArbV1MultiSwap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IArthArbV1MultiSwap *IArthArbV1MultiSwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IArthArbV1MultiSwap.Contract.contract.Transact(opts, method, params...)
}

// Arbitrage is a paid mutator transaction binding the contract method 0x14dd1de6.
//
// Solidity: function arbitrage(address lender, address loanAsset, uint256 amount, address[] path, uint256[] amounts) returns()
func (_IArthArbV1MultiSwap *IArthArbV1MultiSwapTransactor) Arbitrage(opts *bind.TransactOpts, lender common.Address, loanAsset common.Address, amount *big.Int, path []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _IArthArbV1MultiSwap.contract.Transact(opts, "arbitrage", lender, loanAsset, amount, path, amounts)
}

// Arbitrage is a paid mutator transaction binding the contract method 0x14dd1de6.
//
// Solidity: function arbitrage(address lender, address loanAsset, uint256 amount, address[] path, uint256[] amounts) returns()
func (_IArthArbV1MultiSwap *IArthArbV1MultiSwapSession) Arbitrage(lender common.Address, loanAsset common.Address, amount *big.Int, path []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _IArthArbV1MultiSwap.Contract.Arbitrage(&_IArthArbV1MultiSwap.TransactOpts, lender, loanAsset, amount, path, amounts)
}

// Arbitrage is a paid mutator transaction binding the contract method 0x14dd1de6.
//
// Solidity: function arbitrage(address lender, address loanAsset, uint256 amount, address[] path, uint256[] amounts) returns()
func (_IArthArbV1MultiSwap *IArthArbV1MultiSwapTransactorSession) Arbitrage(lender common.Address, loanAsset common.Address, amount *big.Int, path []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _IArthArbV1MultiSwap.Contract.Arbitrage(&_IArthArbV1MultiSwap.TransactOpts, lender, loanAsset, amount, path, amounts)
}
