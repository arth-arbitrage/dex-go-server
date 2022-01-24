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

// SwapWrapArthV1ABI is the input ABI used to generate the binding from.
const SwapWrapArthV1ABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_expectedOut\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SwapWrapArthV1 is an auto generated Go binding around an Ethereum contract.
type SwapWrapArthV1 struct {
	SwapWrapArthV1Caller     // Read-only binding to the contract
	SwapWrapArthV1Transactor // Write-only binding to the contract
	SwapWrapArthV1Filterer   // Log filterer for contract events
}

// SwapWrapArthV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type SwapWrapArthV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapWrapArthV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapWrapArthV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapWrapArthV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapWrapArthV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapWrapArthV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapWrapArthV1Session struct {
	Contract     *SwapWrapArthV1   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapWrapArthV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapWrapArthV1CallerSession struct {
	Contract *SwapWrapArthV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SwapWrapArthV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapWrapArthV1TransactorSession struct {
	Contract     *SwapWrapArthV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SwapWrapArthV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type SwapWrapArthV1Raw struct {
	Contract *SwapWrapArthV1 // Generic contract binding to access the raw methods on
}

// SwapWrapArthV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapWrapArthV1CallerRaw struct {
	Contract *SwapWrapArthV1Caller // Generic read-only contract binding to access the raw methods on
}

// SwapWrapArthV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapWrapArthV1TransactorRaw struct {
	Contract *SwapWrapArthV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapWrapArthV1 creates a new instance of SwapWrapArthV1, bound to a specific deployed contract.
func NewSwapWrapArthV1(address common.Address, backend bind.ContractBackend) (*SwapWrapArthV1, error) {
	contract, err := bindSwapWrapArthV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapWrapArthV1{SwapWrapArthV1Caller: SwapWrapArthV1Caller{contract: contract}, SwapWrapArthV1Transactor: SwapWrapArthV1Transactor{contract: contract}, SwapWrapArthV1Filterer: SwapWrapArthV1Filterer{contract: contract}}, nil
}

// NewSwapWrapArthV1Caller creates a new read-only instance of SwapWrapArthV1, bound to a specific deployed contract.
func NewSwapWrapArthV1Caller(address common.Address, caller bind.ContractCaller) (*SwapWrapArthV1Caller, error) {
	contract, err := bindSwapWrapArthV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapWrapArthV1Caller{contract: contract}, nil
}

// NewSwapWrapArthV1Transactor creates a new write-only instance of SwapWrapArthV1, bound to a specific deployed contract.
func NewSwapWrapArthV1Transactor(address common.Address, transactor bind.ContractTransactor) (*SwapWrapArthV1Transactor, error) {
	contract, err := bindSwapWrapArthV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapWrapArthV1Transactor{contract: contract}, nil
}

// NewSwapWrapArthV1Filterer creates a new log filterer instance of SwapWrapArthV1, bound to a specific deployed contract.
func NewSwapWrapArthV1Filterer(address common.Address, filterer bind.ContractFilterer) (*SwapWrapArthV1Filterer, error) {
	contract, err := bindSwapWrapArthV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapWrapArthV1Filterer{contract: contract}, nil
}

// bindSwapWrapArthV1 binds a generic wrapper to an already deployed contract.
func bindSwapWrapArthV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapWrapArthV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapWrapArthV1 *SwapWrapArthV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapWrapArthV1.Contract.SwapWrapArthV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapWrapArthV1 *SwapWrapArthV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapWrapArthV1.Contract.SwapWrapArthV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapWrapArthV1 *SwapWrapArthV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapWrapArthV1.Contract.SwapWrapArthV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapWrapArthV1 *SwapWrapArthV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapWrapArthV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapWrapArthV1 *SwapWrapArthV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapWrapArthV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapWrapArthV1 *SwapWrapArthV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapWrapArthV1.Contract.contract.Transact(opts, method, params...)
}

// Swap is a paid mutator transaction binding the contract method 0xe343fe12.
//
// Solidity: function swap(address pair, address _token0, address _token1, uint256 _amountIn, uint256 _expectedOut) returns(uint256 amountOut)
func (_SwapWrapArthV1 *SwapWrapArthV1Transactor) Swap(opts *bind.TransactOpts, pair common.Address, _token0 common.Address, _token1 common.Address, _amountIn *big.Int, _expectedOut *big.Int) (*types.Transaction, error) {
	return _SwapWrapArthV1.contract.Transact(opts, "swap", pair, _token0, _token1, _amountIn, _expectedOut)
}

// Swap is a paid mutator transaction binding the contract method 0xe343fe12.
//
// Solidity: function swap(address pair, address _token0, address _token1, uint256 _amountIn, uint256 _expectedOut) returns(uint256 amountOut)
func (_SwapWrapArthV1 *SwapWrapArthV1Session) Swap(pair common.Address, _token0 common.Address, _token1 common.Address, _amountIn *big.Int, _expectedOut *big.Int) (*types.Transaction, error) {
	return _SwapWrapArthV1.Contract.Swap(&_SwapWrapArthV1.TransactOpts, pair, _token0, _token1, _amountIn, _expectedOut)
}

// Swap is a paid mutator transaction binding the contract method 0xe343fe12.
//
// Solidity: function swap(address pair, address _token0, address _token1, uint256 _amountIn, uint256 _expectedOut) returns(uint256 amountOut)
func (_SwapWrapArthV1 *SwapWrapArthV1TransactorSession) Swap(pair common.Address, _token0 common.Address, _token1 common.Address, _amountIn *big.Int, _expectedOut *big.Int) (*types.Transaction, error) {
	return _SwapWrapArthV1.Contract.Swap(&_SwapWrapArthV1.TransactOpts, pair, _token0, _token1, _amountIn, _expectedOut)
}

// SwapWrapArthV1SwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the SwapWrapArthV1 contract.
type SwapWrapArthV1SwapIterator struct {
	Event *SwapWrapArthV1Swap // Event containing the contract specifics and raw log

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
func (it *SwapWrapArthV1SwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapWrapArthV1Swap)
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
		it.Event = new(SwapWrapArthV1Swap)
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
func (it *SwapWrapArthV1SwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapWrapArthV1SwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapWrapArthV1Swap represents a Swap event raised by the SwapWrapArthV1 contract.
type SwapWrapArthV1Swap struct {
	Sender     common.Address
	Token0     common.Address
	Token1     common.Address
	Amount0Out *big.Int
	Amount1Out *big.Int
	AmountOut  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xd6d34547c69c5ee3d2667625c188acf1006abb93e0ee7cf03925c67cf7760413.
//
// Solidity: event Swap(address indexed sender, address token0, address token1, uint256 amount0Out, uint256 amount1Out, uint256 amountOut)
func (_SwapWrapArthV1 *SwapWrapArthV1Filterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address) (*SwapWrapArthV1SwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SwapWrapArthV1.contract.FilterLogs(opts, "Swap", senderRule)
	if err != nil {
		return nil, err
	}
	return &SwapWrapArthV1SwapIterator{contract: _SwapWrapArthV1.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xd6d34547c69c5ee3d2667625c188acf1006abb93e0ee7cf03925c67cf7760413.
//
// Solidity: event Swap(address indexed sender, address token0, address token1, uint256 amount0Out, uint256 amount1Out, uint256 amountOut)
func (_SwapWrapArthV1 *SwapWrapArthV1Filterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *SwapWrapArthV1Swap, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SwapWrapArthV1.contract.WatchLogs(opts, "Swap", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapWrapArthV1Swap)
				if err := _SwapWrapArthV1.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xd6d34547c69c5ee3d2667625c188acf1006abb93e0ee7cf03925c67cf7760413.
//
// Solidity: event Swap(address indexed sender, address token0, address token1, uint256 amount0Out, uint256 amount1Out, uint256 amountOut)
func (_SwapWrapArthV1 *SwapWrapArthV1Filterer) ParseSwap(log types.Log) (*SwapWrapArthV1Swap, error) {
	event := new(SwapWrapArthV1Swap)
	if err := _SwapWrapArthV1.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
