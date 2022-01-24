// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aave

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

// ILendingPoolAddressesProviderABI is the input ABI used to generate the binding from.
const ILendingPoolAddressesProviderABI = "[{\"inputs\":[],\"name\":\"getLendingPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLendingPoolCore\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ILendingPoolAddressesProvider is an auto generated Go binding around an Ethereum contract.
type ILendingPoolAddressesProvider struct {
	ILendingPoolAddressesProviderCaller     // Read-only binding to the contract
	ILendingPoolAddressesProviderTransactor // Write-only binding to the contract
	ILendingPoolAddressesProviderFilterer   // Log filterer for contract events
}

// ILendingPoolAddressesProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILendingPoolAddressesProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILendingPoolAddressesProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILendingPoolAddressesProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILendingPoolAddressesProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILendingPoolAddressesProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILendingPoolAddressesProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILendingPoolAddressesProviderSession struct {
	Contract     *ILendingPoolAddressesProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ILendingPoolAddressesProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILendingPoolAddressesProviderCallerSession struct {
	Contract *ILendingPoolAddressesProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// ILendingPoolAddressesProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILendingPoolAddressesProviderTransactorSession struct {
	Contract     *ILendingPoolAddressesProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// ILendingPoolAddressesProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILendingPoolAddressesProviderRaw struct {
	Contract *ILendingPoolAddressesProvider // Generic contract binding to access the raw methods on
}

// ILendingPoolAddressesProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILendingPoolAddressesProviderCallerRaw struct {
	Contract *ILendingPoolAddressesProviderCaller // Generic read-only contract binding to access the raw methods on
}

// ILendingPoolAddressesProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILendingPoolAddressesProviderTransactorRaw struct {
	Contract *ILendingPoolAddressesProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewILendingPoolAddressesProvider creates a new instance of ILendingPoolAddressesProvider, bound to a specific deployed contract.
func NewILendingPoolAddressesProvider(address common.Address, backend bind.ContractBackend) (*ILendingPoolAddressesProvider, error) {
	contract, err := bindILendingPoolAddressesProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolAddressesProvider{ILendingPoolAddressesProviderCaller: ILendingPoolAddressesProviderCaller{contract: contract}, ILendingPoolAddressesProviderTransactor: ILendingPoolAddressesProviderTransactor{contract: contract}, ILendingPoolAddressesProviderFilterer: ILendingPoolAddressesProviderFilterer{contract: contract}}, nil
}

// NewILendingPoolAddressesProviderCaller creates a new read-only instance of ILendingPoolAddressesProvider, bound to a specific deployed contract.
func NewILendingPoolAddressesProviderCaller(address common.Address, caller bind.ContractCaller) (*ILendingPoolAddressesProviderCaller, error) {
	contract, err := bindILendingPoolAddressesProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolAddressesProviderCaller{contract: contract}, nil
}

// NewILendingPoolAddressesProviderTransactor creates a new write-only instance of ILendingPoolAddressesProvider, bound to a specific deployed contract.
func NewILendingPoolAddressesProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*ILendingPoolAddressesProviderTransactor, error) {
	contract, err := bindILendingPoolAddressesProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolAddressesProviderTransactor{contract: contract}, nil
}

// NewILendingPoolAddressesProviderFilterer creates a new log filterer instance of ILendingPoolAddressesProvider, bound to a specific deployed contract.
func NewILendingPoolAddressesProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*ILendingPoolAddressesProviderFilterer, error) {
	contract, err := bindILendingPoolAddressesProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolAddressesProviderFilterer{contract: contract}, nil
}

// bindILendingPoolAddressesProvider binds a generic wrapper to an already deployed contract.
func bindILendingPoolAddressesProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ILendingPoolAddressesProviderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILendingPoolAddressesProvider.Contract.ILendingPoolAddressesProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.Contract.ILendingPoolAddressesProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.Contract.ILendingPoolAddressesProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILendingPoolAddressesProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.Contract.contract.Transact(opts, method, params...)
}

// GetLendingPool is a free data retrieval call binding the contract method 0x0261bf8b.
//
// Solidity: function getLendingPool() view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderCaller) GetLendingPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILendingPoolAddressesProvider.contract.Call(opts, &out, "getLendingPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPool is a free data retrieval call binding the contract method 0x0261bf8b.
//
// Solidity: function getLendingPool() view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderSession) GetLendingPool() (common.Address, error) {
	return _ILendingPoolAddressesProvider.Contract.GetLendingPool(&_ILendingPoolAddressesProvider.CallOpts)
}

// GetLendingPool is a free data retrieval call binding the contract method 0x0261bf8b.
//
// Solidity: function getLendingPool() view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderCallerSession) GetLendingPool() (common.Address, error) {
	return _ILendingPoolAddressesProvider.Contract.GetLendingPool(&_ILendingPoolAddressesProvider.CallOpts)
}

// GetLendingPoolCore is a free data retrieval call binding the contract method 0xed6ff760.
//
// Solidity: function getLendingPoolCore() view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderCaller) GetLendingPoolCore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILendingPoolAddressesProvider.contract.Call(opts, &out, "getLendingPoolCore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPoolCore is a free data retrieval call binding the contract method 0xed6ff760.
//
// Solidity: function getLendingPoolCore() view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderSession) GetLendingPoolCore() (common.Address, error) {
	return _ILendingPoolAddressesProvider.Contract.GetLendingPoolCore(&_ILendingPoolAddressesProvider.CallOpts)
}

// GetLendingPoolCore is a free data retrieval call binding the contract method 0xed6ff760.
//
// Solidity: function getLendingPoolCore() view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderCallerSession) GetLendingPoolCore() (common.Address, error) {
	return _ILendingPoolAddressesProvider.Contract.GetLendingPoolCore(&_ILendingPoolAddressesProvider.CallOpts)
}
