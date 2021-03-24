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

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
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
}

// PoolInfoABI is the input ABI used to generate the binding from.
const PoolInfoABI = "[{\"inputs\":[{\"name\":\"_provider\",\"type\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"gas\":15876,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"get_pool_coins\",\"outputs\":[{\"name\":\"coins\",\"type\":\"address[8]\"},{\"name\":\"underlying_coins\",\"type\":\"address[8]\"},{\"name\":\"decimals\",\"type\":\"uint256[8]\"},{\"name\":\"underlying_decimals\",\"type\":\"uint256[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":35142,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"get_pool_info\",\"outputs\":[{\"name\":\"balances\",\"type\":\"uint256[8]\"},{\"name\":\"underlying_balances\",\"type\":\"uint256[8]\"},{\"name\":\"decimals\",\"type\":\"uint256[8]\"},{\"name\":\"underlying_decimals\",\"type\":\"uint256[8]\"},{\"name\":\"rates\",\"type\":\"uint256[8]\"},{\"name\":\"lp_token\",\"type\":\"address\"},{\"components\":[{\"name\":\"A\",\"type\":\"uint256\"},{\"name\":\"future_A\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"},{\"name\":\"admin_fee\",\"type\":\"uint256\"},{\"name\":\"future_fee\",\"type\":\"uint256\"},{\"name\":\"future_admin_fee\",\"type\":\"uint256\"},{\"name\":\"future_owner\",\"type\":\"address\"},{\"name\":\"initial_A\",\"type\":\"uint256\"},{\"name\":\"initial_A_time\",\"type\":\"uint256\"},{\"name\":\"future_A_time\",\"type\":\"uint256\"}],\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":1121,\"inputs\":[],\"name\":\"address_provider\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PoolInfo is an auto generated Go binding around an Ethereum contract.
type PoolInfo struct {
	PoolInfoCaller     // Read-only binding to the contract
	PoolInfoTransactor // Write-only binding to the contract
	PoolInfoFilterer   // Log filterer for contract events
}

// PoolInfoCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolInfoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolInfoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolInfoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolInfoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolInfoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolInfoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolInfoSession struct {
	Contract     *PoolInfo         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolInfoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolInfoCallerSession struct {
	Contract *PoolInfoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PoolInfoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolInfoTransactorSession struct {
	Contract     *PoolInfoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PoolInfoRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolInfoRaw struct {
	Contract *PoolInfo // Generic contract binding to access the raw methods on
}

// PoolInfoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolInfoCallerRaw struct {
	Contract *PoolInfoCaller // Generic read-only contract binding to access the raw methods on
}

// PoolInfoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolInfoTransactorRaw struct {
	Contract *PoolInfoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolInfo creates a new instance of PoolInfo, bound to a specific deployed contract.
func NewPoolInfo(address common.Address, backend bind.ContractBackend) (*PoolInfo, error) {
	contract, err := bindPoolInfo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolInfo{PoolInfoCaller: PoolInfoCaller{contract: contract}, PoolInfoTransactor: PoolInfoTransactor{contract: contract}, PoolInfoFilterer: PoolInfoFilterer{contract: contract}}, nil
}

// NewPoolInfoCaller creates a new read-only instance of PoolInfo, bound to a specific deployed contract.
func NewPoolInfoCaller(address common.Address, caller bind.ContractCaller) (*PoolInfoCaller, error) {
	contract, err := bindPoolInfo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolInfoCaller{contract: contract}, nil
}

// NewPoolInfoTransactor creates a new write-only instance of PoolInfo, bound to a specific deployed contract.
func NewPoolInfoTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolInfoTransactor, error) {
	contract, err := bindPoolInfo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolInfoTransactor{contract: contract}, nil
}

// NewPoolInfoFilterer creates a new log filterer instance of PoolInfo, bound to a specific deployed contract.
func NewPoolInfoFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolInfoFilterer, error) {
	contract, err := bindPoolInfo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolInfoFilterer{contract: contract}, nil
}

// bindPoolInfo binds a generic wrapper to an already deployed contract.
func bindPoolInfo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolInfoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolInfo *PoolInfoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolInfo.Contract.PoolInfoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolInfo *PoolInfoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolInfo.Contract.PoolInfoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolInfo *PoolInfoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolInfo.Contract.PoolInfoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolInfo *PoolInfoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolInfo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolInfo *PoolInfoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolInfo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolInfo *PoolInfoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolInfo.Contract.contract.Transact(opts, method, params...)
}

// AddressProvider is a free data retrieval call binding the contract method 0xce50c2e7.
//
// Solidity: function address_provider() view returns(address)
func (_PoolInfo *PoolInfoCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolInfo.contract.Call(opts, &out, "address_provider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0xce50c2e7.
//
// Solidity: function address_provider() view returns(address)
func (_PoolInfo *PoolInfoSession) AddressProvider() (common.Address, error) {
	return _PoolInfo.Contract.AddressProvider(&_PoolInfo.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0xce50c2e7.
//
// Solidity: function address_provider() view returns(address)
func (_PoolInfo *PoolInfoCallerSession) AddressProvider() (common.Address, error) {
	return _PoolInfo.Contract.AddressProvider(&_PoolInfo.CallOpts)
}

// GetPoolCoins is a free data retrieval call binding the contract method 0xe030afb8.
//
// Solidity: function get_pool_coins(address _pool) view returns(address[8] coins, address[8] underlying_coins, uint256[8] decimals, uint256[8] underlying_decimals)
func (_PoolInfo *PoolInfoCaller) GetPoolCoins(opts *bind.CallOpts, _pool common.Address) (struct {
	Coins              [8]common.Address
	UnderlyingCoins    [8]common.Address
	Decimals           [8]*big.Int
	UnderlyingDecimals [8]*big.Int
}, error) {
	var out []interface{}
	err := _PoolInfo.contract.Call(opts, &out, "get_pool_coins", _pool)

	outstruct := new(struct {
		Coins              [8]common.Address
		UnderlyingCoins    [8]common.Address
		Decimals           [8]*big.Int
		UnderlyingDecimals [8]*big.Int
	})

	outstruct.Coins = out[0].([8]common.Address)
	outstruct.UnderlyingCoins = out[1].([8]common.Address)
	outstruct.Decimals = out[2].([8]*big.Int)
	outstruct.UnderlyingDecimals = out[3].([8]*big.Int)

	return *outstruct, err

}

// GetPoolCoins is a free data retrieval call binding the contract method 0xe030afb8.
//
// Solidity: function get_pool_coins(address _pool) view returns(address[8] coins, address[8] underlying_coins, uint256[8] decimals, uint256[8] underlying_decimals)
func (_PoolInfo *PoolInfoSession) GetPoolCoins(_pool common.Address) (struct {
	Coins              [8]common.Address
	UnderlyingCoins    [8]common.Address
	Decimals           [8]*big.Int
	UnderlyingDecimals [8]*big.Int
}, error) {
	return _PoolInfo.Contract.GetPoolCoins(&_PoolInfo.CallOpts, _pool)
}

// GetPoolCoins is a free data retrieval call binding the contract method 0xe030afb8.
//
// Solidity: function get_pool_coins(address _pool) view returns(address[8] coins, address[8] underlying_coins, uint256[8] decimals, uint256[8] underlying_decimals)
func (_PoolInfo *PoolInfoCallerSession) GetPoolCoins(_pool common.Address) (struct {
	Coins              [8]common.Address
	UnderlyingCoins    [8]common.Address
	Decimals           [8]*big.Int
	UnderlyingDecimals [8]*big.Int
}, error) {
	return _PoolInfo.Contract.GetPoolCoins(&_PoolInfo.CallOpts, _pool)
}

// GetPoolInfo is a free data retrieval call binding the contract method 0x100f2c00.
//
// Solidity: function get_pool_info(address _pool) view returns(uint256[8] balances, uint256[8] underlying_balances, uint256[8] decimals, uint256[8] underlying_decimals, uint256[8] rates, address lp_token, (uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256) params)
func (_PoolInfo *PoolInfoCaller) GetPoolInfo(opts *bind.CallOpts, _pool common.Address) (struct {
	Balances           [8]*big.Int
	UnderlyingBalances [8]*big.Int
	Decimals           [8]*big.Int
	UnderlyingDecimals [8]*big.Int
	Rates              [8]*big.Int
	LpToken            common.Address
	Params             Struct0
}, error) {
	var out []interface{}
	err := _PoolInfo.contract.Call(opts, &out, "get_pool_info", _pool)

	outstruct := new(struct {
		Balances           [8]*big.Int
		UnderlyingBalances [8]*big.Int
		Decimals           [8]*big.Int
		UnderlyingDecimals [8]*big.Int
		Rates              [8]*big.Int
		LpToken            common.Address
		Params             Struct0
	})

	outstruct.Balances = out[0].([8]*big.Int)
	outstruct.UnderlyingBalances = out[1].([8]*big.Int)
	outstruct.Decimals = out[2].([8]*big.Int)
	outstruct.UnderlyingDecimals = out[3].([8]*big.Int)
	outstruct.Rates = out[4].([8]*big.Int)
	outstruct.LpToken = out[5].(common.Address)
	outstruct.Params = out[6].(Struct0)

	return *outstruct, err

}

// GetPoolInfo is a free data retrieval call binding the contract method 0x100f2c00.
//
// Solidity: function get_pool_info(address _pool) view returns(uint256[8] balances, uint256[8] underlying_balances, uint256[8] decimals, uint256[8] underlying_decimals, uint256[8] rates, address lp_token, (uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256) params)
func (_PoolInfo *PoolInfoSession) GetPoolInfo(_pool common.Address) (struct {
	Balances           [8]*big.Int
	UnderlyingBalances [8]*big.Int
	Decimals           [8]*big.Int
	UnderlyingDecimals [8]*big.Int
	Rates              [8]*big.Int
	LpToken            common.Address
	Params             Struct0
}, error) {
	return _PoolInfo.Contract.GetPoolInfo(&_PoolInfo.CallOpts, _pool)
}

// GetPoolInfo is a free data retrieval call binding the contract method 0x100f2c00.
//
// Solidity: function get_pool_info(address _pool) view returns(uint256[8] balances, uint256[8] underlying_balances, uint256[8] decimals, uint256[8] underlying_decimals, uint256[8] rates, address lp_token, (uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256) params)
func (_PoolInfo *PoolInfoCallerSession) GetPoolInfo(_pool common.Address) (struct {
	Balances           [8]*big.Int
	UnderlyingBalances [8]*big.Int
	Decimals           [8]*big.Int
	UnderlyingDecimals [8]*big.Int
	Rates              [8]*big.Int
	LpToken            common.Address
	Params             Struct0
}, error) {
	return _PoolInfo.Contract.GetPoolInfo(&_PoolInfo.CallOpts, _pool)
}
