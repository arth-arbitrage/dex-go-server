// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arth

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

// ArthLendingABI is the input ABI used to generate the binding from.
const ArthLendingABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"_referral\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_protocolFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"FlashLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_referralCode\",\"type\":\"uint16\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_params\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveAvailableLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lendingAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalFeeBips\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lendingAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_aTokenBalanceAfterRedeem\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ArthLending is an auto generated Go binding around an Ethereum contract.
type ArthLending struct {
	ArthLendingCaller     // Read-only binding to the contract
	ArthLendingTransactor // Write-only binding to the contract
	ArthLendingFilterer   // Log filterer for contract events
}

// ArthLendingCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArthLendingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArthLendingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArthLendingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArthLendingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArthLendingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArthLendingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArthLendingSession struct {
	Contract     *ArthLending      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArthLendingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArthLendingCallerSession struct {
	Contract *ArthLendingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ArthLendingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArthLendingTransactorSession struct {
	Contract     *ArthLendingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ArthLendingRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArthLendingRaw struct {
	Contract *ArthLending // Generic contract binding to access the raw methods on
}

// ArthLendingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArthLendingCallerRaw struct {
	Contract *ArthLendingCaller // Generic read-only contract binding to access the raw methods on
}

// ArthLendingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArthLendingTransactorRaw struct {
	Contract *ArthLendingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArthLending creates a new instance of ArthLending, bound to a specific deployed contract.
func NewArthLending(address common.Address, backend bind.ContractBackend) (*ArthLending, error) {
	contract, err := bindArthLending(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArthLending{ArthLendingCaller: ArthLendingCaller{contract: contract}, ArthLendingTransactor: ArthLendingTransactor{contract: contract}, ArthLendingFilterer: ArthLendingFilterer{contract: contract}}, nil
}

// NewArthLendingCaller creates a new read-only instance of ArthLending, bound to a specific deployed contract.
func NewArthLendingCaller(address common.Address, caller bind.ContractCaller) (*ArthLendingCaller, error) {
	contract, err := bindArthLending(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArthLendingCaller{contract: contract}, nil
}

// NewArthLendingTransactor creates a new write-only instance of ArthLending, bound to a specific deployed contract.
func NewArthLendingTransactor(address common.Address, transactor bind.ContractTransactor) (*ArthLendingTransactor, error) {
	contract, err := bindArthLending(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArthLendingTransactor{contract: contract}, nil
}

// NewArthLendingFilterer creates a new log filterer instance of ArthLending, bound to a specific deployed contract.
func NewArthLendingFilterer(address common.Address, filterer bind.ContractFilterer) (*ArthLendingFilterer, error) {
	contract, err := bindArthLending(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArthLendingFilterer{contract: contract}, nil
}

// bindArthLending binds a generic wrapper to an already deployed contract.
func bindArthLending(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArthLendingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArthLending *ArthLendingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArthLending.Contract.ArthLendingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArthLending *ArthLendingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArthLending.Contract.ArthLendingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArthLending *ArthLendingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArthLending.Contract.ArthLendingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArthLending *ArthLendingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArthLending.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArthLending *ArthLendingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArthLending.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArthLending *ArthLendingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArthLending.Contract.contract.Transact(opts, method, params...)
}

// GetReserveAvailableLiquidity is a free data retrieval call binding the contract method 0xe2403019.
//
// Solidity: function getReserveAvailableLiquidity(address _reserve) view returns(uint256)
func (_ArthLending *ArthLendingCaller) GetReserveAvailableLiquidity(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ArthLending.contract.Call(opts, &out, "getReserveAvailableLiquidity", _reserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveAvailableLiquidity is a free data retrieval call binding the contract method 0xe2403019.
//
// Solidity: function getReserveAvailableLiquidity(address _reserve) view returns(uint256)
func (_ArthLending *ArthLendingSession) GetReserveAvailableLiquidity(_reserve common.Address) (*big.Int, error) {
	return _ArthLending.Contract.GetReserveAvailableLiquidity(&_ArthLending.CallOpts, _reserve)
}

// GetReserveAvailableLiquidity is a free data retrieval call binding the contract method 0xe2403019.
//
// Solidity: function getReserveAvailableLiquidity(address _reserve) view returns(uint256)
func (_ArthLending *ArthLendingCallerSession) GetReserveAvailableLiquidity(_reserve common.Address) (*big.Int, error) {
	return _ArthLending.Contract.GetReserveAvailableLiquidity(&_ArthLending.CallOpts, _reserve)
}

// LendingAddress is a free data retrieval call binding the contract method 0x78fe9a5b.
//
// Solidity: function lendingAddress() view returns(address)
func (_ArthLending *ArthLendingCaller) LendingAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArthLending.contract.Call(opts, &out, "lendingAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LendingAddress is a free data retrieval call binding the contract method 0x78fe9a5b.
//
// Solidity: function lendingAddress() view returns(address)
func (_ArthLending *ArthLendingSession) LendingAddress() (common.Address, error) {
	return _ArthLending.Contract.LendingAddress(&_ArthLending.CallOpts)
}

// LendingAddress is a free data retrieval call binding the contract method 0x78fe9a5b.
//
// Solidity: function lendingAddress() view returns(address)
func (_ArthLending *ArthLendingCallerSession) LendingAddress() (common.Address, error) {
	return _ArthLending.Contract.LendingAddress(&_ArthLending.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd2d0e066.
//
// Solidity: function deposit(address _reserve, uint256 _amount, uint16 _referralCode) payable returns()
func (_ArthLending *ArthLendingTransactor) Deposit(opts *bind.TransactOpts, _reserve common.Address, _amount *big.Int, _referralCode uint16) (*types.Transaction, error) {
	return _ArthLending.contract.Transact(opts, "deposit", _reserve, _amount, _referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xd2d0e066.
//
// Solidity: function deposit(address _reserve, uint256 _amount, uint16 _referralCode) payable returns()
func (_ArthLending *ArthLendingSession) Deposit(_reserve common.Address, _amount *big.Int, _referralCode uint16) (*types.Transaction, error) {
	return _ArthLending.Contract.Deposit(&_ArthLending.TransactOpts, _reserve, _amount, _referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xd2d0e066.
//
// Solidity: function deposit(address _reserve, uint256 _amount, uint16 _referralCode) payable returns()
func (_ArthLending *ArthLendingTransactorSession) Deposit(_reserve common.Address, _amount *big.Int, _referralCode uint16) (*types.Transaction, error) {
	return _ArthLending.Contract.Deposit(&_ArthLending.TransactOpts, _reserve, _amount, _referralCode)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address _receiver, address _reserve, uint256 _amount, bytes _params) payable returns()
func (_ArthLending *ArthLendingTransactor) FlashLoan(opts *bind.TransactOpts, _receiver common.Address, _reserve common.Address, _amount *big.Int, _params []byte) (*types.Transaction, error) {
	return _ArthLending.contract.Transact(opts, "flashLoan", _receiver, _reserve, _amount, _params)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address _receiver, address _reserve, uint256 _amount, bytes _params) payable returns()
func (_ArthLending *ArthLendingSession) FlashLoan(_receiver common.Address, _reserve common.Address, _amount *big.Int, _params []byte) (*types.Transaction, error) {
	return _ArthLending.Contract.FlashLoan(&_ArthLending.TransactOpts, _receiver, _reserve, _amount, _params)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address _receiver, address _reserve, uint256 _amount, bytes _params) payable returns()
func (_ArthLending *ArthLendingTransactorSession) FlashLoan(_receiver common.Address, _reserve common.Address, _amount *big.Int, _params []byte) (*types.Transaction, error) {
	return _ArthLending.Contract.FlashLoan(&_ArthLending.TransactOpts, _receiver, _reserve, _amount, _params)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _lendingAddress, uint256 _totalFeeBips) returns()
func (_ArthLending *ArthLendingTransactor) Initialize(opts *bind.TransactOpts, _lendingAddress common.Address, _totalFeeBips *big.Int) (*types.Transaction, error) {
	return _ArthLending.contract.Transact(opts, "initialize", _lendingAddress, _totalFeeBips)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _lendingAddress, uint256 _totalFeeBips) returns()
func (_ArthLending *ArthLendingSession) Initialize(_lendingAddress common.Address, _totalFeeBips *big.Int) (*types.Transaction, error) {
	return _ArthLending.Contract.Initialize(&_ArthLending.TransactOpts, _lendingAddress, _totalFeeBips)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _lendingAddress, uint256 _totalFeeBips) returns()
func (_ArthLending *ArthLendingTransactorSession) Initialize(_lendingAddress common.Address, _totalFeeBips *big.Int) (*types.Transaction, error) {
	return _ArthLending.Contract.Initialize(&_ArthLending.TransactOpts, _lendingAddress, _totalFeeBips)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb5c5f672.
//
// Solidity: function withdraw(address _reserve, uint256 _amount, uint256 _aTokenBalanceAfterRedeem) returns()
func (_ArthLending *ArthLendingTransactor) Withdraw(opts *bind.TransactOpts, _reserve common.Address, _amount *big.Int, _aTokenBalanceAfterRedeem *big.Int) (*types.Transaction, error) {
	return _ArthLending.contract.Transact(opts, "withdraw", _reserve, _amount, _aTokenBalanceAfterRedeem)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb5c5f672.
//
// Solidity: function withdraw(address _reserve, uint256 _amount, uint256 _aTokenBalanceAfterRedeem) returns()
func (_ArthLending *ArthLendingSession) Withdraw(_reserve common.Address, _amount *big.Int, _aTokenBalanceAfterRedeem *big.Int) (*types.Transaction, error) {
	return _ArthLending.Contract.Withdraw(&_ArthLending.TransactOpts, _reserve, _amount, _aTokenBalanceAfterRedeem)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb5c5f672.
//
// Solidity: function withdraw(address _reserve, uint256 _amount, uint256 _aTokenBalanceAfterRedeem) returns()
func (_ArthLending *ArthLendingTransactorSession) Withdraw(_reserve common.Address, _amount *big.Int, _aTokenBalanceAfterRedeem *big.Int) (*types.Transaction, error) {
	return _ArthLending.Contract.Withdraw(&_ArthLending.TransactOpts, _reserve, _amount, _aTokenBalanceAfterRedeem)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ArthLending *ArthLendingTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ArthLending.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ArthLending *ArthLendingSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ArthLending.Contract.Fallback(&_ArthLending.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ArthLending *ArthLendingTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ArthLending.Contract.Fallback(&_ArthLending.TransactOpts, calldata)
}

// ArthLendingDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ArthLending contract.
type ArthLendingDepositIterator struct {
	Event *ArthLendingDeposit // Event containing the contract specifics and raw log

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
func (it *ArthLendingDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArthLendingDeposit)
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
		it.Event = new(ArthLendingDeposit)
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
func (it *ArthLendingDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArthLendingDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArthLendingDeposit represents a Deposit event raised by the ArthLending contract.
type ArthLendingDeposit struct {
	Reserve   common.Address
	User      common.Address
	Amount    *big.Int
	Referral  uint16
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xc12c57b1c73a2c3a2ea4613e9476abb3d8d146857aab7329e24243fb59710c82.
//
// Solidity: event Deposit(address indexed _reserve, address indexed _user, uint256 _amount, uint16 indexed _referral, uint256 _timestamp)
func (_ArthLending *ArthLendingFilterer) FilterDeposit(opts *bind.FilterOpts, _reserve []common.Address, _user []common.Address, _referral []uint16) (*ArthLendingDepositIterator, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	var _referralRule []interface{}
	for _, _referralItem := range _referral {
		_referralRule = append(_referralRule, _referralItem)
	}

	logs, sub, err := _ArthLending.contract.FilterLogs(opts, "Deposit", _reserveRule, _userRule, _referralRule)
	if err != nil {
		return nil, err
	}
	return &ArthLendingDepositIterator{contract: _ArthLending.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xc12c57b1c73a2c3a2ea4613e9476abb3d8d146857aab7329e24243fb59710c82.
//
// Solidity: event Deposit(address indexed _reserve, address indexed _user, uint256 _amount, uint16 indexed _referral, uint256 _timestamp)
func (_ArthLending *ArthLendingFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ArthLendingDeposit, _reserve []common.Address, _user []common.Address, _referral []uint16) (event.Subscription, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	var _referralRule []interface{}
	for _, _referralItem := range _referral {
		_referralRule = append(_referralRule, _referralItem)
	}

	logs, sub, err := _ArthLending.contract.WatchLogs(opts, "Deposit", _reserveRule, _userRule, _referralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArthLendingDeposit)
				if err := _ArthLending.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xc12c57b1c73a2c3a2ea4613e9476abb3d8d146857aab7329e24243fb59710c82.
//
// Solidity: event Deposit(address indexed _reserve, address indexed _user, uint256 _amount, uint16 indexed _referral, uint256 _timestamp)
func (_ArthLending *ArthLendingFilterer) ParseDeposit(log types.Log) (*ArthLendingDeposit, error) {
	event := new(ArthLendingDeposit)
	if err := _ArthLending.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArthLendingFlashLoanIterator is returned from FilterFlashLoan and is used to iterate over the raw logs and unpacked data for FlashLoan events raised by the ArthLending contract.
type ArthLendingFlashLoanIterator struct {
	Event *ArthLendingFlashLoan // Event containing the contract specifics and raw log

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
func (it *ArthLendingFlashLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArthLendingFlashLoan)
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
		it.Event = new(ArthLendingFlashLoan)
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
func (it *ArthLendingFlashLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArthLendingFlashLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArthLendingFlashLoan represents a FlashLoan event raised by the ArthLending contract.
type ArthLendingFlashLoan struct {
	Target      common.Address
	Reserve     common.Address
	Amount      *big.Int
	TotalFee    *big.Int
	ProtocolFee *big.Int
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFlashLoan is a free log retrieval operation binding the contract event 0x5b8f46461c1dd69fb968f1a003acee221ea3e19540e350233b612ddb43433b55.
//
// Solidity: event FlashLoan(address indexed _target, address indexed _reserve, uint256 _amount, uint256 _totalFee, uint256 _protocolFee, uint256 _timestamp)
func (_ArthLending *ArthLendingFilterer) FilterFlashLoan(opts *bind.FilterOpts, _target []common.Address, _reserve []common.Address) (*ArthLendingFlashLoanIterator, error) {

	var _targetRule []interface{}
	for _, _targetItem := range _target {
		_targetRule = append(_targetRule, _targetItem)
	}
	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}

	logs, sub, err := _ArthLending.contract.FilterLogs(opts, "FlashLoan", _targetRule, _reserveRule)
	if err != nil {
		return nil, err
	}
	return &ArthLendingFlashLoanIterator{contract: _ArthLending.contract, event: "FlashLoan", logs: logs, sub: sub}, nil
}

// WatchFlashLoan is a free log subscription operation binding the contract event 0x5b8f46461c1dd69fb968f1a003acee221ea3e19540e350233b612ddb43433b55.
//
// Solidity: event FlashLoan(address indexed _target, address indexed _reserve, uint256 _amount, uint256 _totalFee, uint256 _protocolFee, uint256 _timestamp)
func (_ArthLending *ArthLendingFilterer) WatchFlashLoan(opts *bind.WatchOpts, sink chan<- *ArthLendingFlashLoan, _target []common.Address, _reserve []common.Address) (event.Subscription, error) {

	var _targetRule []interface{}
	for _, _targetItem := range _target {
		_targetRule = append(_targetRule, _targetItem)
	}
	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}

	logs, sub, err := _ArthLending.contract.WatchLogs(opts, "FlashLoan", _targetRule, _reserveRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArthLendingFlashLoan)
				if err := _ArthLending.contract.UnpackLog(event, "FlashLoan", log); err != nil {
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

// ParseFlashLoan is a log parse operation binding the contract event 0x5b8f46461c1dd69fb968f1a003acee221ea3e19540e350233b612ddb43433b55.
//
// Solidity: event FlashLoan(address indexed _target, address indexed _reserve, uint256 _amount, uint256 _totalFee, uint256 _protocolFee, uint256 _timestamp)
func (_ArthLending *ArthLendingFilterer) ParseFlashLoan(log types.Log) (*ArthLendingFlashLoan, error) {
	event := new(ArthLendingFlashLoan)
	if err := _ArthLending.contract.UnpackLog(event, "FlashLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArthLendingWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the ArthLending contract.
type ArthLendingWithdrawIterator struct {
	Event *ArthLendingWithdraw // Event containing the contract specifics and raw log

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
func (it *ArthLendingWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArthLendingWithdraw)
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
		it.Event = new(ArthLendingWithdraw)
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
func (it *ArthLendingWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArthLendingWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArthLendingWithdraw represents a Withdraw event raised by the ArthLending contract.
type ArthLendingWithdraw struct {
	Reserve   common.Address
	User      common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf341246adaac6f497bc2a656f546ab9e182111d630394f0c57c710a59a2cb567.
//
// Solidity: event Withdraw(address indexed _reserve, address indexed _user, uint256 _amount, uint256 _timestamp)
func (_ArthLending *ArthLendingFilterer) FilterWithdraw(opts *bind.FilterOpts, _reserve []common.Address, _user []common.Address) (*ArthLendingWithdrawIterator, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _ArthLending.contract.FilterLogs(opts, "Withdraw", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return &ArthLendingWithdrawIterator{contract: _ArthLending.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf341246adaac6f497bc2a656f546ab9e182111d630394f0c57c710a59a2cb567.
//
// Solidity: event Withdraw(address indexed _reserve, address indexed _user, uint256 _amount, uint256 _timestamp)
func (_ArthLending *ArthLendingFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ArthLendingWithdraw, _reserve []common.Address, _user []common.Address) (event.Subscription, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _ArthLending.contract.WatchLogs(opts, "Withdraw", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArthLendingWithdraw)
				if err := _ArthLending.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf341246adaac6f497bc2a656f546ab9e182111d630394f0c57c710a59a2cb567.
//
// Solidity: event Withdraw(address indexed _reserve, address indexed _user, uint256 _amount, uint256 _timestamp)
func (_ArthLending *ArthLendingFilterer) ParseWithdraw(log types.Log) (*ArthLendingWithdraw, error) {
	event := new(ArthLendingWithdraw)
	if err := _ArthLending.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
