// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curvefi

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

// SwapsABI is the input ABI used to generate the binding from.
const SwapsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"token_sold\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"token_bought\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount_sold\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount_bought\",\"type\":\"uint256\"}],\"name\":\"TokenExchange\",\"type\":\"event\"},{\"inputs\":[{\"name\":\"_address_provider\",\"type\":\"address\"},{\"name\":\"_calculator\",\"type\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_expected\",\"type\":\"uint256\"}],\"name\":\"exchange_with_best_rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_expected\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"exchange_with_best_rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_expected\",\"type\":\"uint256\"}],\"name\":\"exchange\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_expected\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"exchange\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"gas\":395840312,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"get_best_rate\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":5104,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"get_exchange_amount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":15358,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"get_input_amount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":17925,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amounts\",\"type\":\"uint256[100]\"}],\"name\":\"get_exchange_amounts\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[100]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":2429,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"get_calculator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":72066,\"inputs\":[],\"name\":\"update_registry_address\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":37159,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_calculator\",\"type\":\"address\"}],\"name\":\"set_calculator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":36974,\"inputs\":[{\"name\":\"_calculator\",\"type\":\"address\"}],\"name\":\"set_default_calculator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":37798,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"claim_balance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":37034,\"inputs\":[{\"name\":\"_is_killed\",\"type\":\"bool\"}],\"name\":\"set_killed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":1538,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":1568,\"inputs\":[],\"name\":\"factory_registry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":1598,\"inputs\":[],\"name\":\"default_calculator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":1628,\"inputs\":[],\"name\":\"is_killed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Swaps is an auto generated Go binding around an Ethereum contract.
type Swaps struct {
	SwapsCaller     // Read-only binding to the contract
	SwapsTransactor // Write-only binding to the contract
	SwapsFilterer   // Log filterer for contract events
}

// SwapsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapsSession struct {
	Contract     *Swaps            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapsCallerSession struct {
	Contract *SwapsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SwapsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapsTransactorSession struct {
	Contract     *SwapsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapsRaw struct {
	Contract *Swaps // Generic contract binding to access the raw methods on
}

// SwapsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapsCallerRaw struct {
	Contract *SwapsCaller // Generic read-only contract binding to access the raw methods on
}

// SwapsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapsTransactorRaw struct {
	Contract *SwapsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwaps creates a new instance of Swaps, bound to a specific deployed contract.
func NewSwaps(address common.Address, backend bind.ContractBackend) (*Swaps, error) {
	contract, err := bindSwaps(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Swaps{SwapsCaller: SwapsCaller{contract: contract}, SwapsTransactor: SwapsTransactor{contract: contract}, SwapsFilterer: SwapsFilterer{contract: contract}}, nil
}

// NewSwapsCaller creates a new read-only instance of Swaps, bound to a specific deployed contract.
func NewSwapsCaller(address common.Address, caller bind.ContractCaller) (*SwapsCaller, error) {
	contract, err := bindSwaps(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapsCaller{contract: contract}, nil
}

// NewSwapsTransactor creates a new write-only instance of Swaps, bound to a specific deployed contract.
func NewSwapsTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapsTransactor, error) {
	contract, err := bindSwaps(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapsTransactor{contract: contract}, nil
}

// NewSwapsFilterer creates a new log filterer instance of Swaps, bound to a specific deployed contract.
func NewSwapsFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapsFilterer, error) {
	contract, err := bindSwaps(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapsFilterer{contract: contract}, nil
}

// bindSwaps binds a generic wrapper to an already deployed contract.
func bindSwaps(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swaps *SwapsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swaps.Contract.SwapsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swaps *SwapsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swaps.Contract.SwapsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swaps *SwapsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swaps.Contract.SwapsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swaps *SwapsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swaps.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swaps *SwapsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swaps.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swaps *SwapsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swaps.Contract.contract.Transact(opts, method, params...)
}

// DefaultCalculator is a free data retrieval call binding the contract method 0x3b359fc8.
//
// Solidity: function default_calculator() view returns(address)
func (_Swaps *SwapsCaller) DefaultCalculator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swaps.contract.Call(opts, &out, "default_calculator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultCalculator is a free data retrieval call binding the contract method 0x3b359fc8.
//
// Solidity: function default_calculator() view returns(address)
func (_Swaps *SwapsSession) DefaultCalculator() (common.Address, error) {
	return _Swaps.Contract.DefaultCalculator(&_Swaps.CallOpts)
}

// DefaultCalculator is a free data retrieval call binding the contract method 0x3b359fc8.
//
// Solidity: function default_calculator() view returns(address)
func (_Swaps *SwapsCallerSession) DefaultCalculator() (common.Address, error) {
	return _Swaps.Contract.DefaultCalculator(&_Swaps.CallOpts)
}

// FactoryRegistry is a free data retrieval call binding the contract method 0xf7cbf4c6.
//
// Solidity: function factory_registry() view returns(address)
func (_Swaps *SwapsCaller) FactoryRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swaps.contract.Call(opts, &out, "factory_registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryRegistry is a free data retrieval call binding the contract method 0xf7cbf4c6.
//
// Solidity: function factory_registry() view returns(address)
func (_Swaps *SwapsSession) FactoryRegistry() (common.Address, error) {
	return _Swaps.Contract.FactoryRegistry(&_Swaps.CallOpts)
}

// FactoryRegistry is a free data retrieval call binding the contract method 0xf7cbf4c6.
//
// Solidity: function factory_registry() view returns(address)
func (_Swaps *SwapsCallerSession) FactoryRegistry() (common.Address, error) {
	return _Swaps.Contract.FactoryRegistry(&_Swaps.CallOpts)
}

// GetBestRate is a free data retrieval call binding the contract method 0x4e21df75.
//
// Solidity: function get_best_rate(address _from, address _to, uint256 _amount) view returns(address, uint256)
func (_Swaps *SwapsCaller) GetBestRate(opts *bind.CallOpts, _from common.Address, _to common.Address, _amount *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _Swaps.contract.Call(opts, &out, "get_best_rate", _from, _to, _amount)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetBestRate is a free data retrieval call binding the contract method 0x4e21df75.
//
// Solidity: function get_best_rate(address _from, address _to, uint256 _amount) view returns(address, uint256)
func (_Swaps *SwapsSession) GetBestRate(_from common.Address, _to common.Address, _amount *big.Int) (common.Address, *big.Int, error) {
	return _Swaps.Contract.GetBestRate(&_Swaps.CallOpts, _from, _to, _amount)
}

// GetBestRate is a free data retrieval call binding the contract method 0x4e21df75.
//
// Solidity: function get_best_rate(address _from, address _to, uint256 _amount) view returns(address, uint256)
func (_Swaps *SwapsCallerSession) GetBestRate(_from common.Address, _to common.Address, _amount *big.Int) (common.Address, *big.Int, error) {
	return _Swaps.Contract.GetBestRate(&_Swaps.CallOpts, _from, _to, _amount)
}

// GetCalculator is a free data retrieval call binding the contract method 0x5d7dc825.
//
// Solidity: function get_calculator(address _pool) view returns(address)
func (_Swaps *SwapsCaller) GetCalculator(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _Swaps.contract.Call(opts, &out, "get_calculator", _pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCalculator is a free data retrieval call binding the contract method 0x5d7dc825.
//
// Solidity: function get_calculator(address _pool) view returns(address)
func (_Swaps *SwapsSession) GetCalculator(_pool common.Address) (common.Address, error) {
	return _Swaps.Contract.GetCalculator(&_Swaps.CallOpts, _pool)
}

// GetCalculator is a free data retrieval call binding the contract method 0x5d7dc825.
//
// Solidity: function get_calculator(address _pool) view returns(address)
func (_Swaps *SwapsCallerSession) GetCalculator(_pool common.Address) (common.Address, error) {
	return _Swaps.Contract.GetCalculator(&_Swaps.CallOpts, _pool)
}

// GetExchangeAmount is a free data retrieval call binding the contract method 0x3973e834.
//
// Solidity: function get_exchange_amount(address _pool, address _from, address _to, uint256 _amount) view returns(uint256)
func (_Swaps *SwapsCaller) GetExchangeAmount(opts *bind.CallOpts, _pool common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Swaps.contract.Call(opts, &out, "get_exchange_amount", _pool, _from, _to, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExchangeAmount is a free data retrieval call binding the contract method 0x3973e834.
//
// Solidity: function get_exchange_amount(address _pool, address _from, address _to, uint256 _amount) view returns(uint256)
func (_Swaps *SwapsSession) GetExchangeAmount(_pool common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*big.Int, error) {
	return _Swaps.Contract.GetExchangeAmount(&_Swaps.CallOpts, _pool, _from, _to, _amount)
}

// GetExchangeAmount is a free data retrieval call binding the contract method 0x3973e834.
//
// Solidity: function get_exchange_amount(address _pool, address _from, address _to, uint256 _amount) view returns(uint256)
func (_Swaps *SwapsCallerSession) GetExchangeAmount(_pool common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*big.Int, error) {
	return _Swaps.Contract.GetExchangeAmount(&_Swaps.CallOpts, _pool, _from, _to, _amount)
}

// GetExchangeAmounts is a free data retrieval call binding the contract method 0x4be9ae42.
//
// Solidity: function get_exchange_amounts(address _pool, address _from, address _to, uint256[100] _amounts) view returns(uint256[100])
func (_Swaps *SwapsCaller) GetExchangeAmounts(opts *bind.CallOpts, _pool common.Address, _from common.Address, _to common.Address, _amounts [100]*big.Int) ([100]*big.Int, error) {
	var out []interface{}
	err := _Swaps.contract.Call(opts, &out, "get_exchange_amounts", _pool, _from, _to, _amounts)

	if err != nil {
		return *new([100]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([100]*big.Int)).(*[100]*big.Int)

	return out0, err

}

// GetExchangeAmounts is a free data retrieval call binding the contract method 0x4be9ae42.
//
// Solidity: function get_exchange_amounts(address _pool, address _from, address _to, uint256[100] _amounts) view returns(uint256[100])
func (_Swaps *SwapsSession) GetExchangeAmounts(_pool common.Address, _from common.Address, _to common.Address, _amounts [100]*big.Int) ([100]*big.Int, error) {
	return _Swaps.Contract.GetExchangeAmounts(&_Swaps.CallOpts, _pool, _from, _to, _amounts)
}

// GetExchangeAmounts is a free data retrieval call binding the contract method 0x4be9ae42.
//
// Solidity: function get_exchange_amounts(address _pool, address _from, address _to, uint256[100] _amounts) view returns(uint256[100])
func (_Swaps *SwapsCallerSession) GetExchangeAmounts(_pool common.Address, _from common.Address, _to common.Address, _amounts [100]*big.Int) ([100]*big.Int, error) {
	return _Swaps.Contract.GetExchangeAmounts(&_Swaps.CallOpts, _pool, _from, _to, _amounts)
}

// GetInputAmount is a free data retrieval call binding the contract method 0x7fa5a654.
//
// Solidity: function get_input_amount(address _pool, address _from, address _to, uint256 _amount) view returns(uint256)
func (_Swaps *SwapsCaller) GetInputAmount(opts *bind.CallOpts, _pool common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Swaps.contract.Call(opts, &out, "get_input_amount", _pool, _from, _to, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInputAmount is a free data retrieval call binding the contract method 0x7fa5a654.
//
// Solidity: function get_input_amount(address _pool, address _from, address _to, uint256 _amount) view returns(uint256)
func (_Swaps *SwapsSession) GetInputAmount(_pool common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*big.Int, error) {
	return _Swaps.Contract.GetInputAmount(&_Swaps.CallOpts, _pool, _from, _to, _amount)
}

// GetInputAmount is a free data retrieval call binding the contract method 0x7fa5a654.
//
// Solidity: function get_input_amount(address _pool, address _from, address _to, uint256 _amount) view returns(uint256)
func (_Swaps *SwapsCallerSession) GetInputAmount(_pool common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*big.Int, error) {
	return _Swaps.Contract.GetInputAmount(&_Swaps.CallOpts, _pool, _from, _to, _amount)
}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_Swaps *SwapsCaller) IsKilled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Swaps.contract.Call(opts, &out, "is_killed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_Swaps *SwapsSession) IsKilled() (bool, error) {
	return _Swaps.Contract.IsKilled(&_Swaps.CallOpts)
}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_Swaps *SwapsCallerSession) IsKilled() (bool, error) {
	return _Swaps.Contract.IsKilled(&_Swaps.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Swaps *SwapsCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swaps.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Swaps *SwapsSession) Registry() (common.Address, error) {
	return _Swaps.Contract.Registry(&_Swaps.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Swaps *SwapsCallerSession) Registry() (common.Address, error) {
	return _Swaps.Contract.Registry(&_Swaps.CallOpts)
}

// ClaimBalance is a paid mutator transaction binding the contract method 0x752d53c6.
//
// Solidity: function claim_balance(address _token) returns(bool)
func (_Swaps *SwapsTransactor) ClaimBalance(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Swaps.contract.Transact(opts, "claim_balance", _token)
}

// ClaimBalance is a paid mutator transaction binding the contract method 0x752d53c6.
//
// Solidity: function claim_balance(address _token) returns(bool)
func (_Swaps *SwapsSession) ClaimBalance(_token common.Address) (*types.Transaction, error) {
	return _Swaps.Contract.ClaimBalance(&_Swaps.TransactOpts, _token)
}

// ClaimBalance is a paid mutator transaction binding the contract method 0x752d53c6.
//
// Solidity: function claim_balance(address _token) returns(bool)
func (_Swaps *SwapsTransactorSession) ClaimBalance(_token common.Address) (*types.Transaction, error) {
	return _Swaps.Contract.ClaimBalance(&_Swaps.TransactOpts, _token)
}

// Exchange is a paid mutator transaction binding the contract method 0x4798ce5b.
//
// Solidity: function exchange(address _pool, address _from, address _to, uint256 _amount, uint256 _expected) payable returns(uint256)
func (_Swaps *SwapsTransactor) Exchange(opts *bind.TransactOpts, _pool common.Address, _from common.Address, _to common.Address, _amount *big.Int, _expected *big.Int) (*types.Transaction, error) {
	return _Swaps.contract.Transact(opts, "exchange", _pool, _from, _to, _amount, _expected)
}

// Exchange is a paid mutator transaction binding the contract method 0x4798ce5b.
//
// Solidity: function exchange(address _pool, address _from, address _to, uint256 _amount, uint256 _expected) payable returns(uint256)
func (_Swaps *SwapsSession) Exchange(_pool common.Address, _from common.Address, _to common.Address, _amount *big.Int, _expected *big.Int) (*types.Transaction, error) {
	return _Swaps.Contract.Exchange(&_Swaps.TransactOpts, _pool, _from, _to, _amount, _expected)
}

// Exchange is a paid mutator transaction binding the contract method 0x4798ce5b.
//
// Solidity: function exchange(address _pool, address _from, address _to, uint256 _amount, uint256 _expected) payable returns(uint256)
func (_Swaps *SwapsTransactorSession) Exchange(_pool common.Address, _from common.Address, _to common.Address, _amount *big.Int, _expected *big.Int) (*types.Transaction, error) {
	return _Swaps.Contract.Exchange(&_Swaps.TransactOpts, _pool, _from, _to, _amount, _expected)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x1a4c1ca3.
//
// Solidity: function exchange(address _pool, address _from, address _to, uint256 _amount, uint256 _expected, address _receiver) payable returns(uint256)
func (_Swaps *SwapsTransactor) Exchange0(opts *bind.TransactOpts, _pool common.Address, _from common.Address, _to common.Address, _amount *big.Int, _expected *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Swaps.contract.Transact(opts, "exchange0", _pool, _from, _to, _amount, _expected, _receiver)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x1a4c1ca3.
//
// Solidity: function exchange(address _pool, address _from, address _to, uint256 _amount, uint256 _expected, address _receiver) payable returns(uint256)
func (_Swaps *SwapsSession) Exchange0(_pool common.Address, _from common.Address, _to common.Address, _amount *big.Int, _expected *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Swaps.Contract.Exchange0(&_Swaps.TransactOpts, _pool, _from, _to, _amount, _expected, _receiver)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x1a4c1ca3.
//
// Solidity: function exchange(address _pool, address _from, address _to, uint256 _amount, uint256 _expected, address _receiver) payable returns(uint256)
func (_Swaps *SwapsTransactorSession) Exchange0(_pool common.Address, _from common.Address, _to common.Address, _amount *big.Int, _expected *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Swaps.Contract.Exchange0(&_Swaps.TransactOpts, _pool, _from, _to, _amount, _expected, _receiver)
}

// ExchangeWithBestRate is a paid mutator transaction binding the contract method 0x10e5e303.
//
// Solidity: function exchange_with_best_rate(address _from, address _to, uint256 _amount, uint256 _expected) payable returns(uint256)
func (_Swaps *SwapsTransactor) ExchangeWithBestRate(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int, _expected *big.Int) (*types.Transaction, error) {
	return _Swaps.contract.Transact(opts, "exchange_with_best_rate", _from, _to, _amount, _expected)
}

// ExchangeWithBestRate is a paid mutator transaction binding the contract method 0x10e5e303.
//
// Solidity: function exchange_with_best_rate(address _from, address _to, uint256 _amount, uint256 _expected) payable returns(uint256)
func (_Swaps *SwapsSession) ExchangeWithBestRate(_from common.Address, _to common.Address, _amount *big.Int, _expected *big.Int) (*types.Transaction, error) {
	return _Swaps.Contract.ExchangeWithBestRate(&_Swaps.TransactOpts, _from, _to, _amount, _expected)
}

// ExchangeWithBestRate is a paid mutator transaction binding the contract method 0x10e5e303.
//
// Solidity: function exchange_with_best_rate(address _from, address _to, uint256 _amount, uint256 _expected) payable returns(uint256)
func (_Swaps *SwapsTransactorSession) ExchangeWithBestRate(_from common.Address, _to common.Address, _amount *big.Int, _expected *big.Int) (*types.Transaction, error) {
	return _Swaps.Contract.ExchangeWithBestRate(&_Swaps.TransactOpts, _from, _to, _amount, _expected)
}

// ExchangeWithBestRate0 is a paid mutator transaction binding the contract method 0x9f69a6a6.
//
// Solidity: function exchange_with_best_rate(address _from, address _to, uint256 _amount, uint256 _expected, address _receiver) payable returns(uint256)
func (_Swaps *SwapsTransactor) ExchangeWithBestRate0(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int, _expected *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Swaps.contract.Transact(opts, "exchange_with_best_rate0", _from, _to, _amount, _expected, _receiver)
}

// ExchangeWithBestRate0 is a paid mutator transaction binding the contract method 0x9f69a6a6.
//
// Solidity: function exchange_with_best_rate(address _from, address _to, uint256 _amount, uint256 _expected, address _receiver) payable returns(uint256)
func (_Swaps *SwapsSession) ExchangeWithBestRate0(_from common.Address, _to common.Address, _amount *big.Int, _expected *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Swaps.Contract.ExchangeWithBestRate0(&_Swaps.TransactOpts, _from, _to, _amount, _expected, _receiver)
}

// ExchangeWithBestRate0 is a paid mutator transaction binding the contract method 0x9f69a6a6.
//
// Solidity: function exchange_with_best_rate(address _from, address _to, uint256 _amount, uint256 _expected, address _receiver) payable returns(uint256)
func (_Swaps *SwapsTransactorSession) ExchangeWithBestRate0(_from common.Address, _to common.Address, _amount *big.Int, _expected *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Swaps.Contract.ExchangeWithBestRate0(&_Swaps.TransactOpts, _from, _to, _amount, _expected, _receiver)
}

// SetCalculator is a paid mutator transaction binding the contract method 0x188c7ee5.
//
// Solidity: function set_calculator(address _pool, address _calculator) returns(bool)
func (_Swaps *SwapsTransactor) SetCalculator(opts *bind.TransactOpts, _pool common.Address, _calculator common.Address) (*types.Transaction, error) {
	return _Swaps.contract.Transact(opts, "set_calculator", _pool, _calculator)
}

// SetCalculator is a paid mutator transaction binding the contract method 0x188c7ee5.
//
// Solidity: function set_calculator(address _pool, address _calculator) returns(bool)
func (_Swaps *SwapsSession) SetCalculator(_pool common.Address, _calculator common.Address) (*types.Transaction, error) {
	return _Swaps.Contract.SetCalculator(&_Swaps.TransactOpts, _pool, _calculator)
}

// SetCalculator is a paid mutator transaction binding the contract method 0x188c7ee5.
//
// Solidity: function set_calculator(address _pool, address _calculator) returns(bool)
func (_Swaps *SwapsTransactorSession) SetCalculator(_pool common.Address, _calculator common.Address) (*types.Transaction, error) {
	return _Swaps.Contract.SetCalculator(&_Swaps.TransactOpts, _pool, _calculator)
}

// SetDefaultCalculator is a paid mutator transaction binding the contract method 0xda3fb2ab.
//
// Solidity: function set_default_calculator(address _calculator) returns(bool)
func (_Swaps *SwapsTransactor) SetDefaultCalculator(opts *bind.TransactOpts, _calculator common.Address) (*types.Transaction, error) {
	return _Swaps.contract.Transact(opts, "set_default_calculator", _calculator)
}

// SetDefaultCalculator is a paid mutator transaction binding the contract method 0xda3fb2ab.
//
// Solidity: function set_default_calculator(address _calculator) returns(bool)
func (_Swaps *SwapsSession) SetDefaultCalculator(_calculator common.Address) (*types.Transaction, error) {
	return _Swaps.Contract.SetDefaultCalculator(&_Swaps.TransactOpts, _calculator)
}

// SetDefaultCalculator is a paid mutator transaction binding the contract method 0xda3fb2ab.
//
// Solidity: function set_default_calculator(address _calculator) returns(bool)
func (_Swaps *SwapsTransactorSession) SetDefaultCalculator(_calculator common.Address) (*types.Transaction, error) {
	return _Swaps.Contract.SetDefaultCalculator(&_Swaps.TransactOpts, _calculator)
}

// SetKilled is a paid mutator transaction binding the contract method 0x90b22997.
//
// Solidity: function set_killed(bool _is_killed) returns(bool)
func (_Swaps *SwapsTransactor) SetKilled(opts *bind.TransactOpts, _is_killed bool) (*types.Transaction, error) {
	return _Swaps.contract.Transact(opts, "set_killed", _is_killed)
}

// SetKilled is a paid mutator transaction binding the contract method 0x90b22997.
//
// Solidity: function set_killed(bool _is_killed) returns(bool)
func (_Swaps *SwapsSession) SetKilled(_is_killed bool) (*types.Transaction, error) {
	return _Swaps.Contract.SetKilled(&_Swaps.TransactOpts, _is_killed)
}

// SetKilled is a paid mutator transaction binding the contract method 0x90b22997.
//
// Solidity: function set_killed(bool _is_killed) returns(bool)
func (_Swaps *SwapsTransactorSession) SetKilled(_is_killed bool) (*types.Transaction, error) {
	return _Swaps.Contract.SetKilled(&_Swaps.TransactOpts, _is_killed)
}

// UpdateRegistryAddress is a paid mutator transaction binding the contract method 0x4bbc5b1f.
//
// Solidity: function update_registry_address() returns(bool)
func (_Swaps *SwapsTransactor) UpdateRegistryAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swaps.contract.Transact(opts, "update_registry_address")
}

// UpdateRegistryAddress is a paid mutator transaction binding the contract method 0x4bbc5b1f.
//
// Solidity: function update_registry_address() returns(bool)
func (_Swaps *SwapsSession) UpdateRegistryAddress() (*types.Transaction, error) {
	return _Swaps.Contract.UpdateRegistryAddress(&_Swaps.TransactOpts)
}

// UpdateRegistryAddress is a paid mutator transaction binding the contract method 0x4bbc5b1f.
//
// Solidity: function update_registry_address() returns(bool)
func (_Swaps *SwapsTransactorSession) UpdateRegistryAddress() (*types.Transaction, error) {
	return _Swaps.Contract.UpdateRegistryAddress(&_Swaps.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Swaps *SwapsTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Swaps.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Swaps *SwapsSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Swaps.Contract.Fallback(&_Swaps.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Swaps *SwapsTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Swaps.Contract.Fallback(&_Swaps.TransactOpts, calldata)
}

// SwapsTokenExchangeIterator is returned from FilterTokenExchange and is used to iterate over the raw logs and unpacked data for TokenExchange events raised by the Swaps contract.
type SwapsTokenExchangeIterator struct {
	Event *SwapsTokenExchange // Event containing the contract specifics and raw log

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
func (it *SwapsTokenExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapsTokenExchange)
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
		it.Event = new(SwapsTokenExchange)
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
func (it *SwapsTokenExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapsTokenExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapsTokenExchange represents a TokenExchange event raised by the Swaps contract.
type SwapsTokenExchange struct {
	Buyer        common.Address
	Receiver     common.Address
	Pool         common.Address
	TokenSold    common.Address
	TokenBought  common.Address
	AmountSold   *big.Int
	AmountBought *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenExchange is a free log retrieval operation binding the contract event 0xbd3eb7bcfdd1721a4eb4f00d0df3ed91bd6f17222f82b2d7bce519d8cab3fe46.
//
// Solidity: event TokenExchange(address indexed buyer, address indexed receiver, address indexed pool, address token_sold, address token_bought, uint256 amount_sold, uint256 amount_bought)
func (_Swaps *SwapsFilterer) FilterTokenExchange(opts *bind.FilterOpts, buyer []common.Address, receiver []common.Address, pool []common.Address) (*SwapsTokenExchangeIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Swaps.contract.FilterLogs(opts, "TokenExchange", buyerRule, receiverRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &SwapsTokenExchangeIterator{contract: _Swaps.contract, event: "TokenExchange", logs: logs, sub: sub}, nil
}

// WatchTokenExchange is a free log subscription operation binding the contract event 0xbd3eb7bcfdd1721a4eb4f00d0df3ed91bd6f17222f82b2d7bce519d8cab3fe46.
//
// Solidity: event TokenExchange(address indexed buyer, address indexed receiver, address indexed pool, address token_sold, address token_bought, uint256 amount_sold, uint256 amount_bought)
func (_Swaps *SwapsFilterer) WatchTokenExchange(opts *bind.WatchOpts, sink chan<- *SwapsTokenExchange, buyer []common.Address, receiver []common.Address, pool []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Swaps.contract.WatchLogs(opts, "TokenExchange", buyerRule, receiverRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapsTokenExchange)
				if err := _Swaps.contract.UnpackLog(event, "TokenExchange", log); err != nil {
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

// ParseTokenExchange is a log parse operation binding the contract event 0xbd3eb7bcfdd1721a4eb4f00d0df3ed91bd6f17222f82b2d7bce519d8cab3fe46.
//
// Solidity: event TokenExchange(address indexed buyer, address indexed receiver, address indexed pool, address token_sold, address token_bought, uint256 amount_sold, uint256 amount_bought)
func (_Swaps *SwapsFilterer) ParseTokenExchange(log types.Log) (*SwapsTokenExchange, error) {
	event := new(SwapsTokenExchange)
	if err := _Swaps.contract.UnpackLog(event, "TokenExchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
