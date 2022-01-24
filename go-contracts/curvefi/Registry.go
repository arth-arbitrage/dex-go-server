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

// RegistryABI is the input ABI used to generate the binding from.
const RegistryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"rate_method_id\",\"type\":\"bytes\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"inputs\":[{\"name\":\"_address_provider\",\"type\":\"address\"},{\"name\":\"_gauge_controller\",\"type\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"find_pool_for_coins\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"find_pool_for_coins\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":1704,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"get_n_coins\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":12285,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"get_coins\",\"outputs\":[{\"name\":\"\",\"type\":\"address[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":12347,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"get_underlying_coins\",\"outputs\":[{\"name\":\"\",\"type\":\"address[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":8199,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"get_decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":8261,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"get_underlying_decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":34780,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"get_rates\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":20310,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"get_gauges\",\"outputs\":[{\"name\":\"\",\"type\":\"address[10]\"},{\"name\":\"\",\"type\":\"int128[10]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":16818,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"get_balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":158953,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"get_underlying_balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":2080,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"get_virtual_price_from_lp_token\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":1198,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"get_A\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":6458,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"get_parameters\",\"outputs\":[{\"name\":\"A\",\"type\":\"uint256\"},{\"name\":\"future_A\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"},{\"name\":\"admin_fee\",\"type\":\"uint256\"},{\"name\":\"future_fee\",\"type\":\"uint256\"},{\"name\":\"future_admin_fee\",\"type\":\"uint256\"},{\"name\":\"future_owner\",\"type\":\"address\"},{\"name\":\"initial_A\",\"type\":\"uint256\"},{\"name\":\"initial_A_time\",\"type\":\"uint256\"},{\"name\":\"future_A_time\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":1603,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"get_fees\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":36719,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"get_admin_balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":27456,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"get_coin_indices\",\"outputs\":[{\"name\":\"\",\"type\":\"int128\"},{\"name\":\"\",\"type\":\"int128\"},{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":32329,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"estimate_gas_used\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":10196577,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_n_coins\",\"type\":\"uint256\"},{\"name\":\"_lp_token\",\"type\":\"address\"},{\"name\":\"_rate_method_id\",\"type\":\"bytes32\"},{\"name\":\"_decimals\",\"type\":\"uint256\"},{\"name\":\"_underlying_decimals\",\"type\":\"uint256\"},{\"name\":\"_has_initial_A\",\"type\":\"bool\"},{\"name\":\"_is_v1\",\"type\":\"bool\"}],\"name\":\"add_pool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":5590664,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_n_coins\",\"type\":\"uint256\"},{\"name\":\"_lp_token\",\"type\":\"address\"},{\"name\":\"_rate_method_id\",\"type\":\"bytes32\"},{\"name\":\"_decimals\",\"type\":\"uint256\"},{\"name\":\"_use_rates\",\"type\":\"uint256\"},{\"name\":\"_has_initial_A\",\"type\":\"bool\"},{\"name\":\"_is_v1\",\"type\":\"bool\"}],\"name\":\"add_pool_without_underlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":10226976,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_n_coins\",\"type\":\"uint256\"},{\"name\":\"_lp_token\",\"type\":\"address\"},{\"name\":\"_decimals\",\"type\":\"uint256\"}],\"name\":\"add_metapool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":779646579509,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"remove_pool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":355578,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address[5]\"},{\"name\":\"_amount\",\"type\":\"uint256[2][5]\"}],\"name\":\"set_pool_gas_estimates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":357165,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address[10]\"},{\"name\":\"_amount\",\"type\":\"uint256[10]\"}],\"name\":\"set_coin_gas_estimates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":37747,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_estimator\",\"type\":\"address\"}],\"name\":\"set_gas_estimate_contract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":365793,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_liquidity_gauges\",\"type\":\"address[10]\"}],\"name\":\"set_liquidity_gauges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":2111,\"inputs\":[],\"name\":\"address_provider\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":2141,\"inputs\":[],\"name\":\"gauge_controller\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":2280,\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"name\":\"pool_list\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":2201,\"inputs\":[],\"name\":\"pool_count\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":2446,\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"name\":\"get_pool_from_lp_token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":2476,\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"name\":\"get_lp_token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// AddressProvider is a free data retrieval call binding the contract method 0xce50c2e7.
//
// Solidity: function address_provider() view returns(address)
func (_Registry *RegistryCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "address_provider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0xce50c2e7.
//
// Solidity: function address_provider() view returns(address)
func (_Registry *RegistrySession) AddressProvider() (common.Address, error) {
	return _Registry.Contract.AddressProvider(&_Registry.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0xce50c2e7.
//
// Solidity: function address_provider() view returns(address)
func (_Registry *RegistryCallerSession) AddressProvider() (common.Address, error) {
	return _Registry.Contract.AddressProvider(&_Registry.CallOpts)
}

// EstimateGasUsed is a free data retrieval call binding the contract method 0xb0bb365b.
//
// Solidity: function estimate_gas_used(address _pool, address _from, address _to) view returns(uint256)
func (_Registry *RegistryCaller) EstimateGasUsed(opts *bind.CallOpts, _pool common.Address, _from common.Address, _to common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "estimate_gas_used", _pool, _from, _to)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateGasUsed is a free data retrieval call binding the contract method 0xb0bb365b.
//
// Solidity: function estimate_gas_used(address _pool, address _from, address _to) view returns(uint256)
func (_Registry *RegistrySession) EstimateGasUsed(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, error) {
	return _Registry.Contract.EstimateGasUsed(&_Registry.CallOpts, _pool, _from, _to)
}

// EstimateGasUsed is a free data retrieval call binding the contract method 0xb0bb365b.
//
// Solidity: function estimate_gas_used(address _pool, address _from, address _to) view returns(uint256)
func (_Registry *RegistryCallerSession) EstimateGasUsed(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, error) {
	return _Registry.Contract.EstimateGasUsed(&_Registry.CallOpts, _pool, _from, _to)
}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_Registry *RegistryCaller) FindPoolForCoins(opts *bind.CallOpts, _from common.Address, _to common.Address) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "find_pool_for_coins", _from, _to)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_Registry *RegistrySession) FindPoolForCoins(_from common.Address, _to common.Address) (common.Address, error) {
	return _Registry.Contract.FindPoolForCoins(&_Registry.CallOpts, _from, _to)
}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_Registry *RegistryCallerSession) FindPoolForCoins(_from common.Address, _to common.Address) (common.Address, error) {
	return _Registry.Contract.FindPoolForCoins(&_Registry.CallOpts, _from, _to)
}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_Registry *RegistryCaller) FindPoolForCoins0(opts *bind.CallOpts, _from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "find_pool_for_coins0", _from, _to, i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_Registry *RegistrySession) FindPoolForCoins0(_from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	return _Registry.Contract.FindPoolForCoins0(&_Registry.CallOpts, _from, _to, i)
}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_Registry *RegistryCallerSession) FindPoolForCoins0(_from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	return _Registry.Contract.FindPoolForCoins0(&_Registry.CallOpts, _from, _to, i)
}

// GaugeController is a free data retrieval call binding the contract method 0xd8b9a018.
//
// Solidity: function gauge_controller() view returns(address)
func (_Registry *RegistryCaller) GaugeController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "gauge_controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GaugeController is a free data retrieval call binding the contract method 0xd8b9a018.
//
// Solidity: function gauge_controller() view returns(address)
func (_Registry *RegistrySession) GaugeController() (common.Address, error) {
	return _Registry.Contract.GaugeController(&_Registry.CallOpts)
}

// GaugeController is a free data retrieval call binding the contract method 0xd8b9a018.
//
// Solidity: function gauge_controller() view returns(address)
func (_Registry *RegistryCallerSession) GaugeController() (common.Address, error) {
	return _Registry.Contract.GaugeController(&_Registry.CallOpts)
}

// GetA is a free data retrieval call binding the contract method 0x55b30b19.
//
// Solidity: function get_A(address _pool) view returns(uint256)
func (_Registry *RegistryCaller) GetA(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "get_A", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetA is a free data retrieval call binding the contract method 0x55b30b19.
//
// Solidity: function get_A(address _pool) view returns(uint256)
func (_Registry *RegistrySession) GetA(_pool common.Address) (*big.Int, error) {
	return _Registry.Contract.GetA(&_Registry.CallOpts, _pool)
}

// GetA is a free data retrieval call binding the contract method 0x55b30b19.
//
// Solidity: function get_A(address _pool) view returns(uint256)
func (_Registry *RegistryCallerSession) GetA(_pool common.Address) (*big.Int, error) {
	return _Registry.Contract.GetA(&_Registry.CallOpts, _pool)
}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[8])
func (_Registry *RegistryCaller) GetAdminBalances(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "get_admin_balances", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[8])
func (_Registry *RegistrySession) GetAdminBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Registry.Contract.GetAdminBalances(&_Registry.CallOpts, _pool)
}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[8])
func (_Registry *RegistryCallerSession) GetAdminBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Registry.Contract.GetAdminBalances(&_Registry.CallOpts, _pool)
}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[8])
func (_Registry *RegistryCaller) GetBalances(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "get_balances", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[8])
func (_Registry *RegistrySession) GetBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Registry.Contract.GetBalances(&_Registry.CallOpts, _pool)
}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[8])
func (_Registry *RegistryCallerSession) GetBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Registry.Contract.GetBalances(&_Registry.CallOpts, _pool)
}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(int128, int128, bool)
func (_Registry *RegistryCaller) GetCoinIndices(opts *bind.CallOpts, _pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "get_coin_indices", _pool, _from, _to)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)

	return out0, out1, out2, err

}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(int128, int128, bool)
func (_Registry *RegistrySession) GetCoinIndices(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, bool, error) {
	return _Registry.Contract.GetCoinIndices(&_Registry.CallOpts, _pool, _from, _to)
}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(int128, int128, bool)
func (_Registry *RegistryCallerSession) GetCoinIndices(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, bool, error) {
	return _Registry.Contract.GetCoinIndices(&_Registry.CallOpts, _pool, _from, _to)
}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[8])
func (_Registry *RegistryCaller) GetCoins(opts *bind.CallOpts, _pool common.Address) ([8]common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "get_coins", _pool)

	if err != nil {
		return *new([8]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([8]common.Address)).(*[8]common.Address)

	return out0, err

}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[8])
func (_Registry *RegistrySession) GetCoins(_pool common.Address) ([8]common.Address, error) {
	return _Registry.Contract.GetCoins(&_Registry.CallOpts, _pool)
}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[8])
func (_Registry *RegistryCallerSession) GetCoins(_pool common.Address) ([8]common.Address, error) {
	return _Registry.Contract.GetCoins(&_Registry.CallOpts, _pool)
}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[8])
func (_Registry *RegistryCaller) GetDecimals(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "get_decimals", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[8])
func (_Registry *RegistrySession) GetDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _Registry.Contract.GetDecimals(&_Registry.CallOpts, _pool)
}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[8])
func (_Registry *RegistryCallerSession) GetDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _Registry.Contract.GetDecimals(&_Registry.CallOpts, _pool)
}

// GetFees is a free data retrieval call binding the contract method 0x7cdb72b0.
//
// Solidity: function get_fees(address _pool) view returns(uint256[2])
func (_Registry *RegistryCaller) GetFees(opts *bind.CallOpts, _pool common.Address) ([2]*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "get_fees", _pool)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetFees is a free data retrieval call binding the contract method 0x7cdb72b0.
//
// Solidity: function get_fees(address _pool) view returns(uint256[2])
func (_Registry *RegistrySession) GetFees(_pool common.Address) ([2]*big.Int, error) {
	return _Registry.Contract.GetFees(&_Registry.CallOpts, _pool)
}

// GetFees is a free data retrieval call binding the contract method 0x7cdb72b0.
//
// Solidity: function get_fees(address _pool) view returns(uint256[2])
func (_Registry *RegistryCallerSession) GetFees(_pool common.Address) ([2]*big.Int, error) {
	return _Registry.Contract.GetFees(&_Registry.CallOpts, _pool)
}

// GetGauges is a free data retrieval call binding the contract method 0x56059ffb.
//
// Solidity: function get_gauges(address _pool) view returns(address[10], int128[10])
func (_Registry *RegistryCaller) GetGauges(opts *bind.CallOpts, _pool common.Address) ([10]common.Address, [10]*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "get_gauges", _pool)

	if err != nil {
		return *new([10]common.Address), *new([10]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([10]common.Address)).(*[10]common.Address)
	out1 := *abi.ConvertType(out[1], new([10]*big.Int)).(*[10]*big.Int)

	return out0, out1, err

}

// GetGauges is a free data retrieval call binding the contract method 0x56059ffb.
//
// Solidity: function get_gauges(address _pool) view returns(address[10], int128[10])
func (_Registry *RegistrySession) GetGauges(_pool common.Address) ([10]common.Address, [10]*big.Int, error) {
	return _Registry.Contract.GetGauges(&_Registry.CallOpts, _pool)
}

// GetGauges is a free data retrieval call binding the contract method 0x56059ffb.
//
// Solidity: function get_gauges(address _pool) view returns(address[10], int128[10])
func (_Registry *RegistryCallerSession) GetGauges(_pool common.Address) ([10]common.Address, [10]*big.Int, error) {
	return _Registry.Contract.GetGauges(&_Registry.CallOpts, _pool)
}

// GetLpToken is a free data retrieval call binding the contract method 0x37951049.
//
// Solidity: function get_lp_token(address arg0) view returns(address)
func (_Registry *RegistryCaller) GetLpToken(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "get_lp_token", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLpToken is a free data retrieval call binding the contract method 0x37951049.
//
// Solidity: function get_lp_token(address arg0) view returns(address)
func (_Registry *RegistrySession) GetLpToken(arg0 common.Address) (common.Address, error) {
	return _Registry.Contract.GetLpToken(&_Registry.CallOpts, arg0)
}

// GetLpToken is a free data retrieval call binding the contract method 0x37951049.
//
// Solidity: function get_lp_token(address arg0) view returns(address)
func (_Registry *RegistryCallerSession) GetLpToken(arg0 common.Address) (common.Address, error) {
	return _Registry.Contract.GetLpToken(&_Registry.CallOpts, arg0)
}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256[2])
func (_Registry *RegistryCaller) GetNCoins(opts *bind.CallOpts, _pool common.Address) ([2]*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "get_n_coins", _pool)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256[2])
func (_Registry *RegistrySession) GetNCoins(_pool common.Address) ([2]*big.Int, error) {
	return _Registry.Contract.GetNCoins(&_Registry.CallOpts, _pool)
}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256[2])
func (_Registry *RegistryCallerSession) GetNCoins(_pool common.Address) ([2]*big.Int, error) {
	return _Registry.Contract.GetNCoins(&_Registry.CallOpts, _pool)
}

// GetParameters is a free data retrieval call binding the contract method 0x1f80a957.
//
// Solidity: function get_parameters(address _pool) view returns(uint256 A, uint256 future_A, uint256 fee, uint256 admin_fee, uint256 future_fee, uint256 future_admin_fee, address future_owner, uint256 initial_A, uint256 initial_A_time, uint256 future_A_time)
func (_Registry *RegistryCaller) GetParameters(opts *bind.CallOpts, _pool common.Address) (struct {
	A              *big.Int
	FutureA        *big.Int
	Fee            *big.Int
	AdminFee       *big.Int
	FutureFee      *big.Int
	FutureAdminFee *big.Int
	FutureOwner    common.Address
	InitialA       *big.Int
	InitialATime   *big.Int
	FutureATime    *big.Int
}, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "get_parameters", _pool)

	outstruct := new(struct {
		A              *big.Int
		FutureA        *big.Int
		Fee            *big.Int
		AdminFee       *big.Int
		FutureFee      *big.Int
		FutureAdminFee *big.Int
		FutureOwner    common.Address
		InitialA       *big.Int
		InitialATime   *big.Int
		FutureATime    *big.Int
	})

	outstruct.A = out[0].(*big.Int)
	outstruct.FutureA = out[1].(*big.Int)
	outstruct.Fee = out[2].(*big.Int)
	outstruct.AdminFee = out[3].(*big.Int)
	outstruct.FutureFee = out[4].(*big.Int)
	outstruct.FutureAdminFee = out[5].(*big.Int)
	outstruct.FutureOwner = out[6].(common.Address)
	outstruct.InitialA = out[7].(*big.Int)
	outstruct.InitialATime = out[8].(*big.Int)
	outstruct.FutureATime = out[9].(*big.Int)

	return *outstruct, err

}

// GetParameters is a free data retrieval call binding the contract method 0x1f80a957.
//
// Solidity: function get_parameters(address _pool) view returns(uint256 A, uint256 future_A, uint256 fee, uint256 admin_fee, uint256 future_fee, uint256 future_admin_fee, address future_owner, uint256 initial_A, uint256 initial_A_time, uint256 future_A_time)
func (_Registry *RegistrySession) GetParameters(_pool common.Address) (struct {
	A              *big.Int
	FutureA        *big.Int
	Fee            *big.Int
	AdminFee       *big.Int
	FutureFee      *big.Int
	FutureAdminFee *big.Int
	FutureOwner    common.Address
	InitialA       *big.Int
	InitialATime   *big.Int
	FutureATime    *big.Int
}, error) {
	return _Registry.Contract.GetParameters(&_Registry.CallOpts, _pool)
}

// GetParameters is a free data retrieval call binding the contract method 0x1f80a957.
//
// Solidity: function get_parameters(address _pool) view returns(uint256 A, uint256 future_A, uint256 fee, uint256 admin_fee, uint256 future_fee, uint256 future_admin_fee, address future_owner, uint256 initial_A, uint256 initial_A_time, uint256 future_A_time)
func (_Registry *RegistryCallerSession) GetParameters(_pool common.Address) (struct {
	A              *big.Int
	FutureA        *big.Int
	Fee            *big.Int
	AdminFee       *big.Int
	FutureFee      *big.Int
	FutureAdminFee *big.Int
	FutureOwner    common.Address
	InitialA       *big.Int
	InitialATime   *big.Int
	FutureATime    *big.Int
}, error) {
	return _Registry.Contract.GetParameters(&_Registry.CallOpts, _pool)
}

// GetPoolFromLpToken is a free data retrieval call binding the contract method 0xbdf475c3.
//
// Solidity: function get_pool_from_lp_token(address arg0) view returns(address)
func (_Registry *RegistryCaller) GetPoolFromLpToken(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "get_pool_from_lp_token", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolFromLpToken is a free data retrieval call binding the contract method 0xbdf475c3.
//
// Solidity: function get_pool_from_lp_token(address arg0) view returns(address)
func (_Registry *RegistrySession) GetPoolFromLpToken(arg0 common.Address) (common.Address, error) {
	return _Registry.Contract.GetPoolFromLpToken(&_Registry.CallOpts, arg0)
}

// GetPoolFromLpToken is a free data retrieval call binding the contract method 0xbdf475c3.
//
// Solidity: function get_pool_from_lp_token(address arg0) view returns(address)
func (_Registry *RegistryCallerSession) GetPoolFromLpToken(arg0 common.Address) (common.Address, error) {
	return _Registry.Contract.GetPoolFromLpToken(&_Registry.CallOpts, arg0)
}

// GetRates is a free data retrieval call binding the contract method 0xce99e45a.
//
// Solidity: function get_rates(address _pool) view returns(uint256[8])
func (_Registry *RegistryCaller) GetRates(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "get_rates", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetRates is a free data retrieval call binding the contract method 0xce99e45a.
//
// Solidity: function get_rates(address _pool) view returns(uint256[8])
func (_Registry *RegistrySession) GetRates(_pool common.Address) ([8]*big.Int, error) {
	return _Registry.Contract.GetRates(&_Registry.CallOpts, _pool)
}

// GetRates is a free data retrieval call binding the contract method 0xce99e45a.
//
// Solidity: function get_rates(address _pool) view returns(uint256[8])
func (_Registry *RegistryCallerSession) GetRates(_pool common.Address) ([8]*big.Int, error) {
	return _Registry.Contract.GetRates(&_Registry.CallOpts, _pool)
}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x59f4f351.
//
// Solidity: function get_underlying_balances(address _pool) view returns(uint256[8])
func (_Registry *RegistryCaller) GetUnderlyingBalances(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "get_underlying_balances", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x59f4f351.
//
// Solidity: function get_underlying_balances(address _pool) view returns(uint256[8])
func (_Registry *RegistrySession) GetUnderlyingBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Registry.Contract.GetUnderlyingBalances(&_Registry.CallOpts, _pool)
}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x59f4f351.
//
// Solidity: function get_underlying_balances(address _pool) view returns(uint256[8])
func (_Registry *RegistryCallerSession) GetUnderlyingBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Registry.Contract.GetUnderlyingBalances(&_Registry.CallOpts, _pool)
}

// GetUnderlyingCoins is a free data retrieval call binding the contract method 0xa77576ef.
//
// Solidity: function get_underlying_coins(address _pool) view returns(address[8])
func (_Registry *RegistryCaller) GetUnderlyingCoins(opts *bind.CallOpts, _pool common.Address) ([8]common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "get_underlying_coins", _pool)

	if err != nil {
		return *new([8]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([8]common.Address)).(*[8]common.Address)

	return out0, err

}

// GetUnderlyingCoins is a free data retrieval call binding the contract method 0xa77576ef.
//
// Solidity: function get_underlying_coins(address _pool) view returns(address[8])
func (_Registry *RegistrySession) GetUnderlyingCoins(_pool common.Address) ([8]common.Address, error) {
	return _Registry.Contract.GetUnderlyingCoins(&_Registry.CallOpts, _pool)
}

// GetUnderlyingCoins is a free data retrieval call binding the contract method 0xa77576ef.
//
// Solidity: function get_underlying_coins(address _pool) view returns(address[8])
func (_Registry *RegistryCallerSession) GetUnderlyingCoins(_pool common.Address) ([8]common.Address, error) {
	return _Registry.Contract.GetUnderlyingCoins(&_Registry.CallOpts, _pool)
}

// GetUnderlyingDecimals is a free data retrieval call binding the contract method 0x4cb088f1.
//
// Solidity: function get_underlying_decimals(address _pool) view returns(uint256[8])
func (_Registry *RegistryCaller) GetUnderlyingDecimals(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "get_underlying_decimals", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetUnderlyingDecimals is a free data retrieval call binding the contract method 0x4cb088f1.
//
// Solidity: function get_underlying_decimals(address _pool) view returns(uint256[8])
func (_Registry *RegistrySession) GetUnderlyingDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _Registry.Contract.GetUnderlyingDecimals(&_Registry.CallOpts, _pool)
}

// GetUnderlyingDecimals is a free data retrieval call binding the contract method 0x4cb088f1.
//
// Solidity: function get_underlying_decimals(address _pool) view returns(uint256[8])
func (_Registry *RegistryCallerSession) GetUnderlyingDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _Registry.Contract.GetUnderlyingDecimals(&_Registry.CallOpts, _pool)
}

// GetVirtualPriceFromLpToken is a free data retrieval call binding the contract method 0xc5b7074a.
//
// Solidity: function get_virtual_price_from_lp_token(address _token) view returns(uint256)
func (_Registry *RegistryCaller) GetVirtualPriceFromLpToken(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "get_virtual_price_from_lp_token", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPriceFromLpToken is a free data retrieval call binding the contract method 0xc5b7074a.
//
// Solidity: function get_virtual_price_from_lp_token(address _token) view returns(uint256)
func (_Registry *RegistrySession) GetVirtualPriceFromLpToken(_token common.Address) (*big.Int, error) {
	return _Registry.Contract.GetVirtualPriceFromLpToken(&_Registry.CallOpts, _token)
}

// GetVirtualPriceFromLpToken is a free data retrieval call binding the contract method 0xc5b7074a.
//
// Solidity: function get_virtual_price_from_lp_token(address _token) view returns(uint256)
func (_Registry *RegistryCallerSession) GetVirtualPriceFromLpToken(_token common.Address) (*big.Int, error) {
	return _Registry.Contract.GetVirtualPriceFromLpToken(&_Registry.CallOpts, _token)
}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_Registry *RegistryCaller) PoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "pool_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_Registry *RegistrySession) PoolCount() (*big.Int, error) {
	return _Registry.Contract.PoolCount(&_Registry.CallOpts)
}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_Registry *RegistryCallerSession) PoolCount() (*big.Int, error) {
	return _Registry.Contract.PoolCount(&_Registry.CallOpts)
}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_Registry *RegistryCaller) PoolList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "pool_list", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_Registry *RegistrySession) PoolList(arg0 *big.Int) (common.Address, error) {
	return _Registry.Contract.PoolList(&_Registry.CallOpts, arg0)
}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_Registry *RegistryCallerSession) PoolList(arg0 *big.Int) (common.Address, error) {
	return _Registry.Contract.PoolList(&_Registry.CallOpts, arg0)
}

// AddMetapool is a paid mutator transaction binding the contract method 0xcfedad3c.
//
// Solidity: function add_metapool(address _pool, uint256 _n_coins, address _lp_token, uint256 _decimals) returns()
func (_Registry *RegistryTransactor) AddMetapool(opts *bind.TransactOpts, _pool common.Address, _n_coins *big.Int, _lp_token common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "add_metapool", _pool, _n_coins, _lp_token, _decimals)
}

// AddMetapool is a paid mutator transaction binding the contract method 0xcfedad3c.
//
// Solidity: function add_metapool(address _pool, uint256 _n_coins, address _lp_token, uint256 _decimals) returns()
func (_Registry *RegistrySession) AddMetapool(_pool common.Address, _n_coins *big.Int, _lp_token common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.AddMetapool(&_Registry.TransactOpts, _pool, _n_coins, _lp_token, _decimals)
}

// AddMetapool is a paid mutator transaction binding the contract method 0xcfedad3c.
//
// Solidity: function add_metapool(address _pool, uint256 _n_coins, address _lp_token, uint256 _decimals) returns()
func (_Registry *RegistryTransactorSession) AddMetapool(_pool common.Address, _n_coins *big.Int, _lp_token common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.AddMetapool(&_Registry.TransactOpts, _pool, _n_coins, _lp_token, _decimals)
}

// AddPool is a paid mutator transaction binding the contract method 0x290b1e48.
//
// Solidity: function add_pool(address _pool, uint256 _n_coins, address _lp_token, bytes32 _rate_method_id, uint256 _decimals, uint256 _underlying_decimals, bool _has_initial_A, bool _is_v1) returns()
func (_Registry *RegistryTransactor) AddPool(opts *bind.TransactOpts, _pool common.Address, _n_coins *big.Int, _lp_token common.Address, _rate_method_id [32]byte, _decimals *big.Int, _underlying_decimals *big.Int, _has_initial_A bool, _is_v1 bool) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "add_pool", _pool, _n_coins, _lp_token, _rate_method_id, _decimals, _underlying_decimals, _has_initial_A, _is_v1)
}

// AddPool is a paid mutator transaction binding the contract method 0x290b1e48.
//
// Solidity: function add_pool(address _pool, uint256 _n_coins, address _lp_token, bytes32 _rate_method_id, uint256 _decimals, uint256 _underlying_decimals, bool _has_initial_A, bool _is_v1) returns()
func (_Registry *RegistrySession) AddPool(_pool common.Address, _n_coins *big.Int, _lp_token common.Address, _rate_method_id [32]byte, _decimals *big.Int, _underlying_decimals *big.Int, _has_initial_A bool, _is_v1 bool) (*types.Transaction, error) {
	return _Registry.Contract.AddPool(&_Registry.TransactOpts, _pool, _n_coins, _lp_token, _rate_method_id, _decimals, _underlying_decimals, _has_initial_A, _is_v1)
}

// AddPool is a paid mutator transaction binding the contract method 0x290b1e48.
//
// Solidity: function add_pool(address _pool, uint256 _n_coins, address _lp_token, bytes32 _rate_method_id, uint256 _decimals, uint256 _underlying_decimals, bool _has_initial_A, bool _is_v1) returns()
func (_Registry *RegistryTransactorSession) AddPool(_pool common.Address, _n_coins *big.Int, _lp_token common.Address, _rate_method_id [32]byte, _decimals *big.Int, _underlying_decimals *big.Int, _has_initial_A bool, _is_v1 bool) (*types.Transaction, error) {
	return _Registry.Contract.AddPool(&_Registry.TransactOpts, _pool, _n_coins, _lp_token, _rate_method_id, _decimals, _underlying_decimals, _has_initial_A, _is_v1)
}

// AddPoolWithoutUnderlying is a paid mutator transaction binding the contract method 0x60c06f89.
//
// Solidity: function add_pool_without_underlying(address _pool, uint256 _n_coins, address _lp_token, bytes32 _rate_method_id, uint256 _decimals, uint256 _use_rates, bool _has_initial_A, bool _is_v1) returns()
func (_Registry *RegistryTransactor) AddPoolWithoutUnderlying(opts *bind.TransactOpts, _pool common.Address, _n_coins *big.Int, _lp_token common.Address, _rate_method_id [32]byte, _decimals *big.Int, _use_rates *big.Int, _has_initial_A bool, _is_v1 bool) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "add_pool_without_underlying", _pool, _n_coins, _lp_token, _rate_method_id, _decimals, _use_rates, _has_initial_A, _is_v1)
}

// AddPoolWithoutUnderlying is a paid mutator transaction binding the contract method 0x60c06f89.
//
// Solidity: function add_pool_without_underlying(address _pool, uint256 _n_coins, address _lp_token, bytes32 _rate_method_id, uint256 _decimals, uint256 _use_rates, bool _has_initial_A, bool _is_v1) returns()
func (_Registry *RegistrySession) AddPoolWithoutUnderlying(_pool common.Address, _n_coins *big.Int, _lp_token common.Address, _rate_method_id [32]byte, _decimals *big.Int, _use_rates *big.Int, _has_initial_A bool, _is_v1 bool) (*types.Transaction, error) {
	return _Registry.Contract.AddPoolWithoutUnderlying(&_Registry.TransactOpts, _pool, _n_coins, _lp_token, _rate_method_id, _decimals, _use_rates, _has_initial_A, _is_v1)
}

// AddPoolWithoutUnderlying is a paid mutator transaction binding the contract method 0x60c06f89.
//
// Solidity: function add_pool_without_underlying(address _pool, uint256 _n_coins, address _lp_token, bytes32 _rate_method_id, uint256 _decimals, uint256 _use_rates, bool _has_initial_A, bool _is_v1) returns()
func (_Registry *RegistryTransactorSession) AddPoolWithoutUnderlying(_pool common.Address, _n_coins *big.Int, _lp_token common.Address, _rate_method_id [32]byte, _decimals *big.Int, _use_rates *big.Int, _has_initial_A bool, _is_v1 bool) (*types.Transaction, error) {
	return _Registry.Contract.AddPoolWithoutUnderlying(&_Registry.TransactOpts, _pool, _n_coins, _lp_token, _rate_method_id, _decimals, _use_rates, _has_initial_A, _is_v1)
}

// RemovePool is a paid mutator transaction binding the contract method 0x474932b0.
//
// Solidity: function remove_pool(address _pool) returns()
func (_Registry *RegistryTransactor) RemovePool(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "remove_pool", _pool)
}

// RemovePool is a paid mutator transaction binding the contract method 0x474932b0.
//
// Solidity: function remove_pool(address _pool) returns()
func (_Registry *RegistrySession) RemovePool(_pool common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RemovePool(&_Registry.TransactOpts, _pool)
}

// RemovePool is a paid mutator transaction binding the contract method 0x474932b0.
//
// Solidity: function remove_pool(address _pool) returns()
func (_Registry *RegistryTransactorSession) RemovePool(_pool common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RemovePool(&_Registry.TransactOpts, _pool)
}

// SetCoinGasEstimates is a paid mutator transaction binding the contract method 0x237f89f2.
//
// Solidity: function set_coin_gas_estimates(address[10] _addr, uint256[10] _amount) returns()
func (_Registry *RegistryTransactor) SetCoinGasEstimates(opts *bind.TransactOpts, _addr [10]common.Address, _amount [10]*big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "set_coin_gas_estimates", _addr, _amount)
}

// SetCoinGasEstimates is a paid mutator transaction binding the contract method 0x237f89f2.
//
// Solidity: function set_coin_gas_estimates(address[10] _addr, uint256[10] _amount) returns()
func (_Registry *RegistrySession) SetCoinGasEstimates(_addr [10]common.Address, _amount [10]*big.Int) (*types.Transaction, error) {
	return _Registry.Contract.SetCoinGasEstimates(&_Registry.TransactOpts, _addr, _amount)
}

// SetCoinGasEstimates is a paid mutator transaction binding the contract method 0x237f89f2.
//
// Solidity: function set_coin_gas_estimates(address[10] _addr, uint256[10] _amount) returns()
func (_Registry *RegistryTransactorSession) SetCoinGasEstimates(_addr [10]common.Address, _amount [10]*big.Int) (*types.Transaction, error) {
	return _Registry.Contract.SetCoinGasEstimates(&_Registry.TransactOpts, _addr, _amount)
}

// SetGasEstimateContract is a paid mutator transaction binding the contract method 0xca991b14.
//
// Solidity: function set_gas_estimate_contract(address _pool, address _estimator) returns()
func (_Registry *RegistryTransactor) SetGasEstimateContract(opts *bind.TransactOpts, _pool common.Address, _estimator common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "set_gas_estimate_contract", _pool, _estimator)
}

// SetGasEstimateContract is a paid mutator transaction binding the contract method 0xca991b14.
//
// Solidity: function set_gas_estimate_contract(address _pool, address _estimator) returns()
func (_Registry *RegistrySession) SetGasEstimateContract(_pool common.Address, _estimator common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetGasEstimateContract(&_Registry.TransactOpts, _pool, _estimator)
}

// SetGasEstimateContract is a paid mutator transaction binding the contract method 0xca991b14.
//
// Solidity: function set_gas_estimate_contract(address _pool, address _estimator) returns()
func (_Registry *RegistryTransactorSession) SetGasEstimateContract(_pool common.Address, _estimator common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetGasEstimateContract(&_Registry.TransactOpts, _pool, _estimator)
}

// SetLiquidityGauges is a paid mutator transaction binding the contract method 0xef6b9788.
//
// Solidity: function set_liquidity_gauges(address _pool, address[10] _liquidity_gauges) returns()
func (_Registry *RegistryTransactor) SetLiquidityGauges(opts *bind.TransactOpts, _pool common.Address, _liquidity_gauges [10]common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "set_liquidity_gauges", _pool, _liquidity_gauges)
}

// SetLiquidityGauges is a paid mutator transaction binding the contract method 0xef6b9788.
//
// Solidity: function set_liquidity_gauges(address _pool, address[10] _liquidity_gauges) returns()
func (_Registry *RegistrySession) SetLiquidityGauges(_pool common.Address, _liquidity_gauges [10]common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetLiquidityGauges(&_Registry.TransactOpts, _pool, _liquidity_gauges)
}

// SetLiquidityGauges is a paid mutator transaction binding the contract method 0xef6b9788.
//
// Solidity: function set_liquidity_gauges(address _pool, address[10] _liquidity_gauges) returns()
func (_Registry *RegistryTransactorSession) SetLiquidityGauges(_pool common.Address, _liquidity_gauges [10]common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetLiquidityGauges(&_Registry.TransactOpts, _pool, _liquidity_gauges)
}

// SetPoolGasEstimates is a paid mutator transaction binding the contract method 0x0733b67a.
//
// Solidity: function set_pool_gas_estimates(address[5] _addr, uint256[2][5] _amount) returns()
func (_Registry *RegistryTransactor) SetPoolGasEstimates(opts *bind.TransactOpts, _addr [5]common.Address, _amount [5][2]*big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "set_pool_gas_estimates", _addr, _amount)
}

// SetPoolGasEstimates is a paid mutator transaction binding the contract method 0x0733b67a.
//
// Solidity: function set_pool_gas_estimates(address[5] _addr, uint256[2][5] _amount) returns()
func (_Registry *RegistrySession) SetPoolGasEstimates(_addr [5]common.Address, _amount [5][2]*big.Int) (*types.Transaction, error) {
	return _Registry.Contract.SetPoolGasEstimates(&_Registry.TransactOpts, _addr, _amount)
}

// SetPoolGasEstimates is a paid mutator transaction binding the contract method 0x0733b67a.
//
// Solidity: function set_pool_gas_estimates(address[5] _addr, uint256[2][5] _amount) returns()
func (_Registry *RegistryTransactorSession) SetPoolGasEstimates(_addr [5]common.Address, _amount [5][2]*big.Int) (*types.Transaction, error) {
	return _Registry.Contract.SetPoolGasEstimates(&_Registry.TransactOpts, _addr, _amount)
}

// RegistryPoolAddedIterator is returned from FilterPoolAdded and is used to iterate over the raw logs and unpacked data for PoolAdded events raised by the Registry contract.
type RegistryPoolAddedIterator struct {
	Event *RegistryPoolAdded // Event containing the contract specifics and raw log

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
func (it *RegistryPoolAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryPoolAdded)
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
		it.Event = new(RegistryPoolAdded)
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
func (it *RegistryPoolAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryPoolAdded represents a PoolAdded event raised by the Registry contract.
type RegistryPoolAdded struct {
	Pool         common.Address
	RateMethodId []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPoolAdded is a free log retrieval operation binding the contract event 0xe485c16479ab7092c0b3fc4649843c06be7f072194675261590c84473ab0aea9.
//
// Solidity: event PoolAdded(address indexed pool, bytes rate_method_id)
func (_Registry *RegistryFilterer) FilterPoolAdded(opts *bind.FilterOpts, pool []common.Address) (*RegistryPoolAddedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "PoolAdded", poolRule)
	if err != nil {
		return nil, err
	}
	return &RegistryPoolAddedIterator{contract: _Registry.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

// WatchPoolAdded is a free log subscription operation binding the contract event 0xe485c16479ab7092c0b3fc4649843c06be7f072194675261590c84473ab0aea9.
//
// Solidity: event PoolAdded(address indexed pool, bytes rate_method_id)
func (_Registry *RegistryFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *RegistryPoolAdded, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "PoolAdded", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryPoolAdded)
				if err := _Registry.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

// ParsePoolAdded is a log parse operation binding the contract event 0xe485c16479ab7092c0b3fc4649843c06be7f072194675261590c84473ab0aea9.
//
// Solidity: event PoolAdded(address indexed pool, bytes rate_method_id)
func (_Registry *RegistryFilterer) ParsePoolAdded(log types.Log) (*RegistryPoolAdded, error) {
	event := new(RegistryPoolAdded)
	if err := _Registry.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryPoolRemovedIterator is returned from FilterPoolRemoved and is used to iterate over the raw logs and unpacked data for PoolRemoved events raised by the Registry contract.
type RegistryPoolRemovedIterator struct {
	Event *RegistryPoolRemoved // Event containing the contract specifics and raw log

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
func (it *RegistryPoolRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryPoolRemoved)
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
		it.Event = new(RegistryPoolRemoved)
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
func (it *RegistryPoolRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryPoolRemoved represents a PoolRemoved event raised by the Registry contract.
type RegistryPoolRemoved struct {
	Pool common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPoolRemoved is a free log retrieval operation binding the contract event 0x4106dfdaa577573db51c0ca93f766dbedfa0758faa2e7f5bcdb7c142be803c3f.
//
// Solidity: event PoolRemoved(address indexed pool)
func (_Registry *RegistryFilterer) FilterPoolRemoved(opts *bind.FilterOpts, pool []common.Address) (*RegistryPoolRemovedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "PoolRemoved", poolRule)
	if err != nil {
		return nil, err
	}
	return &RegistryPoolRemovedIterator{contract: _Registry.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

// WatchPoolRemoved is a free log subscription operation binding the contract event 0x4106dfdaa577573db51c0ca93f766dbedfa0758faa2e7f5bcdb7c142be803c3f.
//
// Solidity: event PoolRemoved(address indexed pool)
func (_Registry *RegistryFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *RegistryPoolRemoved, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "PoolRemoved", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryPoolRemoved)
				if err := _Registry.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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

// ParsePoolRemoved is a log parse operation binding the contract event 0x4106dfdaa577573db51c0ca93f766dbedfa0758faa2e7f5bcdb7c142be803c3f.
//
// Solidity: event PoolRemoved(address indexed pool)
func (_Registry *RegistryFilterer) ParsePoolRemoved(log types.Log) (*RegistryPoolRemoved, error) {
	event := new(RegistryPoolRemoved)
	if err := _Registry.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
