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

// AaveArbMultiSwapV1ABI is the input ABI used to generate the binding from.
const AaveArbMultiSwapV1ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressProvider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_assetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LogWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"borrowMade\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addressesProvider\",\"outputs\":[{\"internalType\":\"contractILendingPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"loanAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"arbitrage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_params\",\"type\":\"bytes\"}],\"name\":\"executeOperation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_assetAddress\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// AaveArbMultiSwapV1 is an auto generated Go binding around an Ethereum contract.
type AaveArbMultiSwapV1 struct {
	AaveArbMultiSwapV1Caller     // Read-only binding to the contract
	AaveArbMultiSwapV1Transactor // Write-only binding to the contract
	AaveArbMultiSwapV1Filterer   // Log filterer for contract events
}

// AaveArbMultiSwapV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type AaveArbMultiSwapV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveArbMultiSwapV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveArbMultiSwapV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveArbMultiSwapV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveArbMultiSwapV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveArbMultiSwapV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveArbMultiSwapV1Session struct {
	Contract     *AaveArbMultiSwapV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AaveArbMultiSwapV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveArbMultiSwapV1CallerSession struct {
	Contract *AaveArbMultiSwapV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AaveArbMultiSwapV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveArbMultiSwapV1TransactorSession struct {
	Contract     *AaveArbMultiSwapV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AaveArbMultiSwapV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type AaveArbMultiSwapV1Raw struct {
	Contract *AaveArbMultiSwapV1 // Generic contract binding to access the raw methods on
}

// AaveArbMultiSwapV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveArbMultiSwapV1CallerRaw struct {
	Contract *AaveArbMultiSwapV1Caller // Generic read-only contract binding to access the raw methods on
}

// AaveArbMultiSwapV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveArbMultiSwapV1TransactorRaw struct {
	Contract *AaveArbMultiSwapV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveArbMultiSwapV1 creates a new instance of AaveArbMultiSwapV1, bound to a specific deployed contract.
func NewAaveArbMultiSwapV1(address common.Address, backend bind.ContractBackend) (*AaveArbMultiSwapV1, error) {
	contract, err := bindAaveArbMultiSwapV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveArbMultiSwapV1{AaveArbMultiSwapV1Caller: AaveArbMultiSwapV1Caller{contract: contract}, AaveArbMultiSwapV1Transactor: AaveArbMultiSwapV1Transactor{contract: contract}, AaveArbMultiSwapV1Filterer: AaveArbMultiSwapV1Filterer{contract: contract}}, nil
}

// NewAaveArbMultiSwapV1Caller creates a new read-only instance of AaveArbMultiSwapV1, bound to a specific deployed contract.
func NewAaveArbMultiSwapV1Caller(address common.Address, caller bind.ContractCaller) (*AaveArbMultiSwapV1Caller, error) {
	contract, err := bindAaveArbMultiSwapV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveArbMultiSwapV1Caller{contract: contract}, nil
}

// NewAaveArbMultiSwapV1Transactor creates a new write-only instance of AaveArbMultiSwapV1, bound to a specific deployed contract.
func NewAaveArbMultiSwapV1Transactor(address common.Address, transactor bind.ContractTransactor) (*AaveArbMultiSwapV1Transactor, error) {
	contract, err := bindAaveArbMultiSwapV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveArbMultiSwapV1Transactor{contract: contract}, nil
}

// NewAaveArbMultiSwapV1Filterer creates a new log filterer instance of AaveArbMultiSwapV1, bound to a specific deployed contract.
func NewAaveArbMultiSwapV1Filterer(address common.Address, filterer bind.ContractFilterer) (*AaveArbMultiSwapV1Filterer, error) {
	contract, err := bindAaveArbMultiSwapV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveArbMultiSwapV1Filterer{contract: contract}, nil
}

// bindAaveArbMultiSwapV1 binds a generic wrapper to an already deployed contract.
func bindAaveArbMultiSwapV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AaveArbMultiSwapV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveArbMultiSwapV1.Contract.AaveArbMultiSwapV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveArbMultiSwapV1.Contract.AaveArbMultiSwapV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveArbMultiSwapV1.Contract.AaveArbMultiSwapV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveArbMultiSwapV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveArbMultiSwapV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveArbMultiSwapV1.Contract.contract.Transact(opts, method, params...)
}

// AddressesProvider is a free data retrieval call binding the contract method 0xc72c4d10.
//
// Solidity: function addressesProvider() view returns(address)
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1Caller) AddressesProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveArbMultiSwapV1.contract.Call(opts, &out, "addressesProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressesProvider is a free data retrieval call binding the contract method 0xc72c4d10.
//
// Solidity: function addressesProvider() view returns(address)
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1Session) AddressesProvider() (common.Address, error) {
	return _AaveArbMultiSwapV1.Contract.AddressesProvider(&_AaveArbMultiSwapV1.CallOpts)
}

// AddressesProvider is a free data retrieval call binding the contract method 0xc72c4d10.
//
// Solidity: function addressesProvider() view returns(address)
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1CallerSession) AddressesProvider() (common.Address, error) {
	return _AaveArbMultiSwapV1.Contract.AddressesProvider(&_AaveArbMultiSwapV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveArbMultiSwapV1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1Session) Owner() (common.Address, error) {
	return _AaveArbMultiSwapV1.Contract.Owner(&_AaveArbMultiSwapV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1CallerSession) Owner() (common.Address, error) {
	return _AaveArbMultiSwapV1.Contract.Owner(&_AaveArbMultiSwapV1.CallOpts)
}

// Arbitrage is a paid mutator transaction binding the contract method 0x14dd1de6.
//
// Solidity: function arbitrage(address lender, address loanAsset, uint256 amount, address[] path, uint256[] amounts) returns()
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1Transactor) Arbitrage(opts *bind.TransactOpts, lender common.Address, loanAsset common.Address, amount *big.Int, path []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _AaveArbMultiSwapV1.contract.Transact(opts, "arbitrage", lender, loanAsset, amount, path, amounts)
}

// Arbitrage is a paid mutator transaction binding the contract method 0x14dd1de6.
//
// Solidity: function arbitrage(address lender, address loanAsset, uint256 amount, address[] path, uint256[] amounts) returns()
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1Session) Arbitrage(lender common.Address, loanAsset common.Address, amount *big.Int, path []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _AaveArbMultiSwapV1.Contract.Arbitrage(&_AaveArbMultiSwapV1.TransactOpts, lender, loanAsset, amount, path, amounts)
}

// Arbitrage is a paid mutator transaction binding the contract method 0x14dd1de6.
//
// Solidity: function arbitrage(address lender, address loanAsset, uint256 amount, address[] path, uint256[] amounts) returns()
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1TransactorSession) Arbitrage(lender common.Address, loanAsset common.Address, amount *big.Int, path []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _AaveArbMultiSwapV1.Contract.Arbitrage(&_AaveArbMultiSwapV1.TransactOpts, lender, loanAsset, amount, path, amounts)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0xee872558.
//
// Solidity: function executeOperation(address _reserve, uint256 _amount, uint256 _fee, bytes _params) returns()
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1Transactor) ExecuteOperation(opts *bind.TransactOpts, _reserve common.Address, _amount *big.Int, _fee *big.Int, _params []byte) (*types.Transaction, error) {
	return _AaveArbMultiSwapV1.contract.Transact(opts, "executeOperation", _reserve, _amount, _fee, _params)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0xee872558.
//
// Solidity: function executeOperation(address _reserve, uint256 _amount, uint256 _fee, bytes _params) returns()
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1Session) ExecuteOperation(_reserve common.Address, _amount *big.Int, _fee *big.Int, _params []byte) (*types.Transaction, error) {
	return _AaveArbMultiSwapV1.Contract.ExecuteOperation(&_AaveArbMultiSwapV1.TransactOpts, _reserve, _amount, _fee, _params)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0xee872558.
//
// Solidity: function executeOperation(address _reserve, uint256 _amount, uint256 _fee, bytes _params) returns()
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1TransactorSession) ExecuteOperation(_reserve common.Address, _amount *big.Int, _fee *big.Int, _params []byte) (*types.Transaction, error) {
	return _AaveArbMultiSwapV1.Contract.ExecuteOperation(&_AaveArbMultiSwapV1.TransactOpts, _reserve, _amount, _fee, _params)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveArbMultiSwapV1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1Session) RenounceOwnership() (*types.Transaction, error) {
	return _AaveArbMultiSwapV1.Contract.RenounceOwnership(&_AaveArbMultiSwapV1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AaveArbMultiSwapV1.Contract.RenounceOwnership(&_AaveArbMultiSwapV1.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AaveArbMultiSwapV1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AaveArbMultiSwapV1.Contract.TransferOwnership(&_AaveArbMultiSwapV1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AaveArbMultiSwapV1.Contract.TransferOwnership(&_AaveArbMultiSwapV1.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _assetAddress) returns()
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1Transactor) Withdraw(opts *bind.TransactOpts, _assetAddress common.Address) (*types.Transaction, error) {
	return _AaveArbMultiSwapV1.contract.Transact(opts, "withdraw", _assetAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _assetAddress) returns()
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1Session) Withdraw(_assetAddress common.Address) (*types.Transaction, error) {
	return _AaveArbMultiSwapV1.Contract.Withdraw(&_AaveArbMultiSwapV1.TransactOpts, _assetAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _assetAddress) returns()
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1TransactorSession) Withdraw(_assetAddress common.Address) (*types.Transaction, error) {
	return _AaveArbMultiSwapV1.Contract.Withdraw(&_AaveArbMultiSwapV1.TransactOpts, _assetAddress)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveArbMultiSwapV1.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1Session) Receive() (*types.Transaction, error) {
	return _AaveArbMultiSwapV1.Contract.Receive(&_AaveArbMultiSwapV1.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1TransactorSession) Receive() (*types.Transaction, error) {
	return _AaveArbMultiSwapV1.Contract.Receive(&_AaveArbMultiSwapV1.TransactOpts)
}

// AaveArbMultiSwapV1LogWithdrawIterator is returned from FilterLogWithdraw and is used to iterate over the raw logs and unpacked data for LogWithdraw events raised by the AaveArbMultiSwapV1 contract.
type AaveArbMultiSwapV1LogWithdrawIterator struct {
	Event *AaveArbMultiSwapV1LogWithdraw // Event containing the contract specifics and raw log

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
func (it *AaveArbMultiSwapV1LogWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveArbMultiSwapV1LogWithdraw)
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
		it.Event = new(AaveArbMultiSwapV1LogWithdraw)
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
func (it *AaveArbMultiSwapV1LogWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveArbMultiSwapV1LogWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveArbMultiSwapV1LogWithdraw represents a LogWithdraw event raised by the AaveArbMultiSwapV1 contract.
type AaveArbMultiSwapV1LogWithdraw struct {
	From         common.Address
	AssetAddress common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLogWithdraw is a free log retrieval operation binding the contract event 0x9207361cc2a04b9c7a06691df1eb87c6a63957ae88bf01d0d18c81e3d1272099.
//
// Solidity: event LogWithdraw(address indexed _from, address indexed _assetAddress, uint256 amount)
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1Filterer) FilterLogWithdraw(opts *bind.FilterOpts, _from []common.Address, _assetAddress []common.Address) (*AaveArbMultiSwapV1LogWithdrawIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _assetAddressRule []interface{}
	for _, _assetAddressItem := range _assetAddress {
		_assetAddressRule = append(_assetAddressRule, _assetAddressItem)
	}

	logs, sub, err := _AaveArbMultiSwapV1.contract.FilterLogs(opts, "LogWithdraw", _fromRule, _assetAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveArbMultiSwapV1LogWithdrawIterator{contract: _AaveArbMultiSwapV1.contract, event: "LogWithdraw", logs: logs, sub: sub}, nil
}

// WatchLogWithdraw is a free log subscription operation binding the contract event 0x9207361cc2a04b9c7a06691df1eb87c6a63957ae88bf01d0d18c81e3d1272099.
//
// Solidity: event LogWithdraw(address indexed _from, address indexed _assetAddress, uint256 amount)
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1Filterer) WatchLogWithdraw(opts *bind.WatchOpts, sink chan<- *AaveArbMultiSwapV1LogWithdraw, _from []common.Address, _assetAddress []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _assetAddressRule []interface{}
	for _, _assetAddressItem := range _assetAddress {
		_assetAddressRule = append(_assetAddressRule, _assetAddressItem)
	}

	logs, sub, err := _AaveArbMultiSwapV1.contract.WatchLogs(opts, "LogWithdraw", _fromRule, _assetAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveArbMultiSwapV1LogWithdraw)
				if err := _AaveArbMultiSwapV1.contract.UnpackLog(event, "LogWithdraw", log); err != nil {
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

// ParseLogWithdraw is a log parse operation binding the contract event 0x9207361cc2a04b9c7a06691df1eb87c6a63957ae88bf01d0d18c81e3d1272099.
//
// Solidity: event LogWithdraw(address indexed _from, address indexed _assetAddress, uint256 amount)
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1Filterer) ParseLogWithdraw(log types.Log) (*AaveArbMultiSwapV1LogWithdraw, error) {
	event := new(AaveArbMultiSwapV1LogWithdraw)
	if err := _AaveArbMultiSwapV1.contract.UnpackLog(event, "LogWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveArbMultiSwapV1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AaveArbMultiSwapV1 contract.
type AaveArbMultiSwapV1OwnershipTransferredIterator struct {
	Event *AaveArbMultiSwapV1OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AaveArbMultiSwapV1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveArbMultiSwapV1OwnershipTransferred)
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
		it.Event = new(AaveArbMultiSwapV1OwnershipTransferred)
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
func (it *AaveArbMultiSwapV1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveArbMultiSwapV1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveArbMultiSwapV1OwnershipTransferred represents a OwnershipTransferred event raised by the AaveArbMultiSwapV1 contract.
type AaveArbMultiSwapV1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AaveArbMultiSwapV1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AaveArbMultiSwapV1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AaveArbMultiSwapV1OwnershipTransferredIterator{contract: _AaveArbMultiSwapV1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AaveArbMultiSwapV1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AaveArbMultiSwapV1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveArbMultiSwapV1OwnershipTransferred)
				if err := _AaveArbMultiSwapV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1Filterer) ParseOwnershipTransferred(log types.Log) (*AaveArbMultiSwapV1OwnershipTransferred, error) {
	event := new(AaveArbMultiSwapV1OwnershipTransferred)
	if err := _AaveArbMultiSwapV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveArbMultiSwapV1BorrowMadeIterator is returned from FilterBorrowMade and is used to iterate over the raw logs and unpacked data for BorrowMade events raised by the AaveArbMultiSwapV1 contract.
type AaveArbMultiSwapV1BorrowMadeIterator struct {
	Event *AaveArbMultiSwapV1BorrowMade // Event containing the contract specifics and raw log

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
func (it *AaveArbMultiSwapV1BorrowMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveArbMultiSwapV1BorrowMade)
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
		it.Event = new(AaveArbMultiSwapV1BorrowMade)
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
func (it *AaveArbMultiSwapV1BorrowMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveArbMultiSwapV1BorrowMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveArbMultiSwapV1BorrowMade represents a BorrowMade event raised by the AaveArbMultiSwapV1 contract.
type AaveArbMultiSwapV1BorrowMade struct {
	Reserve common.Address
	Amount  *big.Int
	Fee     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBorrowMade is a free log retrieval operation binding the contract event 0xcd8f01cf11df8ac6e385851ab2495f5370eff19deb3c99a29b542b49d32ba0ba.
//
// Solidity: event borrowMade(address _reserve, uint256 _amount, uint256 _fee)
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1Filterer) FilterBorrowMade(opts *bind.FilterOpts) (*AaveArbMultiSwapV1BorrowMadeIterator, error) {

	logs, sub, err := _AaveArbMultiSwapV1.contract.FilterLogs(opts, "borrowMade")
	if err != nil {
		return nil, err
	}
	return &AaveArbMultiSwapV1BorrowMadeIterator{contract: _AaveArbMultiSwapV1.contract, event: "borrowMade", logs: logs, sub: sub}, nil
}

// WatchBorrowMade is a free log subscription operation binding the contract event 0xcd8f01cf11df8ac6e385851ab2495f5370eff19deb3c99a29b542b49d32ba0ba.
//
// Solidity: event borrowMade(address _reserve, uint256 _amount, uint256 _fee)
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1Filterer) WatchBorrowMade(opts *bind.WatchOpts, sink chan<- *AaveArbMultiSwapV1BorrowMade) (event.Subscription, error) {

	logs, sub, err := _AaveArbMultiSwapV1.contract.WatchLogs(opts, "borrowMade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveArbMultiSwapV1BorrowMade)
				if err := _AaveArbMultiSwapV1.contract.UnpackLog(event, "borrowMade", log); err != nil {
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

// ParseBorrowMade is a log parse operation binding the contract event 0xcd8f01cf11df8ac6e385851ab2495f5370eff19deb3c99a29b542b49d32ba0ba.
//
// Solidity: event borrowMade(address _reserve, uint256 _amount, uint256 _fee)
func (_AaveArbMultiSwapV1 *AaveArbMultiSwapV1Filterer) ParseBorrowMade(log types.Log) (*AaveArbMultiSwapV1BorrowMade, error) {
	event := new(AaveArbMultiSwapV1BorrowMade)
	if err := _AaveArbMultiSwapV1.contract.UnpackLog(event, "borrowMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
