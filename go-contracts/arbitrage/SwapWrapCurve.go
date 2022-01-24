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

// SwapWrapCurveABI is the input ABI used to generate the binding from.
const SwapWrapCurveABI = "[{\"inputs\":[{\"name\":\"_address_provider\",\"type\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"gas\":4136,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_intoken_addr\",\"type\":\"address\"},{\"name\":\"_outtoken_addr\",\"type\":\"address\"},{\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"name\":\"_amountExpOut\",\"type\":\"uint256\"},{\"name\":\"_expected\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":1091,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":1121,\"inputs\":[],\"name\":\"address_provider\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SwapWrapCurve is an auto generated Go binding around an Ethereum contract.
type SwapWrapCurve struct {
	SwapWrapCurveCaller     // Read-only binding to the contract
	SwapWrapCurveTransactor // Write-only binding to the contract
	SwapWrapCurveFilterer   // Log filterer for contract events
}

// SwapWrapCurveCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapWrapCurveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapWrapCurveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapWrapCurveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapWrapCurveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapWrapCurveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapWrapCurveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapWrapCurveSession struct {
	Contract     *SwapWrapCurve    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapWrapCurveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapWrapCurveCallerSession struct {
	Contract *SwapWrapCurveCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SwapWrapCurveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapWrapCurveTransactorSession struct {
	Contract     *SwapWrapCurveTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SwapWrapCurveRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapWrapCurveRaw struct {
	Contract *SwapWrapCurve // Generic contract binding to access the raw methods on
}

// SwapWrapCurveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapWrapCurveCallerRaw struct {
	Contract *SwapWrapCurveCaller // Generic read-only contract binding to access the raw methods on
}

// SwapWrapCurveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapWrapCurveTransactorRaw struct {
	Contract *SwapWrapCurveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapWrapCurve creates a new instance of SwapWrapCurve, bound to a specific deployed contract.
func NewSwapWrapCurve(address common.Address, backend bind.ContractBackend) (*SwapWrapCurve, error) {
	contract, err := bindSwapWrapCurve(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapWrapCurve{SwapWrapCurveCaller: SwapWrapCurveCaller{contract: contract}, SwapWrapCurveTransactor: SwapWrapCurveTransactor{contract: contract}, SwapWrapCurveFilterer: SwapWrapCurveFilterer{contract: contract}}, nil
}

// NewSwapWrapCurveCaller creates a new read-only instance of SwapWrapCurve, bound to a specific deployed contract.
func NewSwapWrapCurveCaller(address common.Address, caller bind.ContractCaller) (*SwapWrapCurveCaller, error) {
	contract, err := bindSwapWrapCurve(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapWrapCurveCaller{contract: contract}, nil
}

// NewSwapWrapCurveTransactor creates a new write-only instance of SwapWrapCurve, bound to a specific deployed contract.
func NewSwapWrapCurveTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapWrapCurveTransactor, error) {
	contract, err := bindSwapWrapCurve(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapWrapCurveTransactor{contract: contract}, nil
}

// NewSwapWrapCurveFilterer creates a new log filterer instance of SwapWrapCurve, bound to a specific deployed contract.
func NewSwapWrapCurveFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapWrapCurveFilterer, error) {
	contract, err := bindSwapWrapCurve(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapWrapCurveFilterer{contract: contract}, nil
}

// bindSwapWrapCurve binds a generic wrapper to an already deployed contract.
func bindSwapWrapCurve(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapWrapCurveABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapWrapCurve *SwapWrapCurveRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapWrapCurve.Contract.SwapWrapCurveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapWrapCurve *SwapWrapCurveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapWrapCurve.Contract.SwapWrapCurveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapWrapCurve *SwapWrapCurveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapWrapCurve.Contract.SwapWrapCurveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapWrapCurve *SwapWrapCurveCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapWrapCurve.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapWrapCurve *SwapWrapCurveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapWrapCurve.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapWrapCurve *SwapWrapCurveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapWrapCurve.Contract.contract.Transact(opts, method, params...)
}

// AddressProvider is a free data retrieval call binding the contract method 0xce50c2e7.
//
// Solidity: function address_provider() view returns(address)
func (_SwapWrapCurve *SwapWrapCurveCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapWrapCurve.contract.Call(opts, &out, "address_provider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0xce50c2e7.
//
// Solidity: function address_provider() view returns(address)
func (_SwapWrapCurve *SwapWrapCurveSession) AddressProvider() (common.Address, error) {
	return _SwapWrapCurve.Contract.AddressProvider(&_SwapWrapCurve.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0xce50c2e7.
//
// Solidity: function address_provider() view returns(address)
func (_SwapWrapCurve *SwapWrapCurveCallerSession) AddressProvider() (common.Address, error) {
	return _SwapWrapCurve.Contract.AddressProvider(&_SwapWrapCurve.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_SwapWrapCurve *SwapWrapCurveCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapWrapCurve.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_SwapWrapCurve *SwapWrapCurveSession) Admin() (common.Address, error) {
	return _SwapWrapCurve.Contract.Admin(&_SwapWrapCurve.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_SwapWrapCurve *SwapWrapCurveCallerSession) Admin() (common.Address, error) {
	return _SwapWrapCurve.Contract.Admin(&_SwapWrapCurve.CallOpts)
}

// Swap is a paid mutator transaction binding the contract method 0x75ce8b83.
//
// Solidity: function swap(address _pool, address _intoken_addr, address _outtoken_addr, uint256 _amountIn, uint256 _amountExpOut, uint256 _expected) returns(uint256)
func (_SwapWrapCurve *SwapWrapCurveTransactor) Swap(opts *bind.TransactOpts, _pool common.Address, _intoken_addr common.Address, _outtoken_addr common.Address, _amountIn *big.Int, _amountExpOut *big.Int, _expected *big.Int) (*types.Transaction, error) {
	return _SwapWrapCurve.contract.Transact(opts, "swap", _pool, _intoken_addr, _outtoken_addr, _amountIn, _amountExpOut, _expected)
}

// Swap is a paid mutator transaction binding the contract method 0x75ce8b83.
//
// Solidity: function swap(address _pool, address _intoken_addr, address _outtoken_addr, uint256 _amountIn, uint256 _amountExpOut, uint256 _expected) returns(uint256)
func (_SwapWrapCurve *SwapWrapCurveSession) Swap(_pool common.Address, _intoken_addr common.Address, _outtoken_addr common.Address, _amountIn *big.Int, _amountExpOut *big.Int, _expected *big.Int) (*types.Transaction, error) {
	return _SwapWrapCurve.Contract.Swap(&_SwapWrapCurve.TransactOpts, _pool, _intoken_addr, _outtoken_addr, _amountIn, _amountExpOut, _expected)
}

// Swap is a paid mutator transaction binding the contract method 0x75ce8b83.
//
// Solidity: function swap(address _pool, address _intoken_addr, address _outtoken_addr, uint256 _amountIn, uint256 _amountExpOut, uint256 _expected) returns(uint256)
func (_SwapWrapCurve *SwapWrapCurveTransactorSession) Swap(_pool common.Address, _intoken_addr common.Address, _outtoken_addr common.Address, _amountIn *big.Int, _amountExpOut *big.Int, _expected *big.Int) (*types.Transaction, error) {
	return _SwapWrapCurve.Contract.Swap(&_SwapWrapCurve.TransactOpts, _pool, _intoken_addr, _outtoken_addr, _amountIn, _amountExpOut, _expected)
}
