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

// SwapWrapUniv2ABI is the input ABI used to generate the binding from.
const SwapWrapUniv2ABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0In\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1Out\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_expectedOut\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SwapWrapUniv2 is an auto generated Go binding around an Ethereum contract.
type SwapWrapUniv2 struct {
	SwapWrapUniv2Caller     // Read-only binding to the contract
	SwapWrapUniv2Transactor // Write-only binding to the contract
	SwapWrapUniv2Filterer   // Log filterer for contract events
}

// SwapWrapUniv2Caller is an auto generated read-only Go binding around an Ethereum contract.
type SwapWrapUniv2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapWrapUniv2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapWrapUniv2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapWrapUniv2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapWrapUniv2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapWrapUniv2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapWrapUniv2Session struct {
	Contract     *SwapWrapUniv2    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapWrapUniv2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapWrapUniv2CallerSession struct {
	Contract *SwapWrapUniv2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SwapWrapUniv2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapWrapUniv2TransactorSession struct {
	Contract     *SwapWrapUniv2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SwapWrapUniv2Raw is an auto generated low-level Go binding around an Ethereum contract.
type SwapWrapUniv2Raw struct {
	Contract *SwapWrapUniv2 // Generic contract binding to access the raw methods on
}

// SwapWrapUniv2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapWrapUniv2CallerRaw struct {
	Contract *SwapWrapUniv2Caller // Generic read-only contract binding to access the raw methods on
}

// SwapWrapUniv2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapWrapUniv2TransactorRaw struct {
	Contract *SwapWrapUniv2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapWrapUniv2 creates a new instance of SwapWrapUniv2, bound to a specific deployed contract.
func NewSwapWrapUniv2(address common.Address, backend bind.ContractBackend) (*SwapWrapUniv2, error) {
	contract, err := bindSwapWrapUniv2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapWrapUniv2{SwapWrapUniv2Caller: SwapWrapUniv2Caller{contract: contract}, SwapWrapUniv2Transactor: SwapWrapUniv2Transactor{contract: contract}, SwapWrapUniv2Filterer: SwapWrapUniv2Filterer{contract: contract}}, nil
}

// NewSwapWrapUniv2Caller creates a new read-only instance of SwapWrapUniv2, bound to a specific deployed contract.
func NewSwapWrapUniv2Caller(address common.Address, caller bind.ContractCaller) (*SwapWrapUniv2Caller, error) {
	contract, err := bindSwapWrapUniv2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapWrapUniv2Caller{contract: contract}, nil
}

// NewSwapWrapUniv2Transactor creates a new write-only instance of SwapWrapUniv2, bound to a specific deployed contract.
func NewSwapWrapUniv2Transactor(address common.Address, transactor bind.ContractTransactor) (*SwapWrapUniv2Transactor, error) {
	contract, err := bindSwapWrapUniv2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapWrapUniv2Transactor{contract: contract}, nil
}

// NewSwapWrapUniv2Filterer creates a new log filterer instance of SwapWrapUniv2, bound to a specific deployed contract.
func NewSwapWrapUniv2Filterer(address common.Address, filterer bind.ContractFilterer) (*SwapWrapUniv2Filterer, error) {
	contract, err := bindSwapWrapUniv2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapWrapUniv2Filterer{contract: contract}, nil
}

// bindSwapWrapUniv2 binds a generic wrapper to an already deployed contract.
func bindSwapWrapUniv2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapWrapUniv2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapWrapUniv2 *SwapWrapUniv2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapWrapUniv2.Contract.SwapWrapUniv2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapWrapUniv2 *SwapWrapUniv2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapWrapUniv2.Contract.SwapWrapUniv2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapWrapUniv2 *SwapWrapUniv2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapWrapUniv2.Contract.SwapWrapUniv2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapWrapUniv2 *SwapWrapUniv2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapWrapUniv2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapWrapUniv2 *SwapWrapUniv2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapWrapUniv2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapWrapUniv2 *SwapWrapUniv2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapWrapUniv2.Contract.contract.Transact(opts, method, params...)
}

// Swap is a paid mutator transaction binding the contract method 0xe343fe12.
//
// Solidity: function swap(address pair, address token0In, address token1Out, uint256 _amountIn, uint256 _expectedOut) returns(uint256)
func (_SwapWrapUniv2 *SwapWrapUniv2Transactor) Swap(opts *bind.TransactOpts, pair common.Address, token0In common.Address, token1Out common.Address, _amountIn *big.Int, _expectedOut *big.Int) (*types.Transaction, error) {
	return _SwapWrapUniv2.contract.Transact(opts, "swap", pair, token0In, token1Out, _amountIn, _expectedOut)
}

// Swap is a paid mutator transaction binding the contract method 0xe343fe12.
//
// Solidity: function swap(address pair, address token0In, address token1Out, uint256 _amountIn, uint256 _expectedOut) returns(uint256)
func (_SwapWrapUniv2 *SwapWrapUniv2Session) Swap(pair common.Address, token0In common.Address, token1Out common.Address, _amountIn *big.Int, _expectedOut *big.Int) (*types.Transaction, error) {
	return _SwapWrapUniv2.Contract.Swap(&_SwapWrapUniv2.TransactOpts, pair, token0In, token1Out, _amountIn, _expectedOut)
}

// Swap is a paid mutator transaction binding the contract method 0xe343fe12.
//
// Solidity: function swap(address pair, address token0In, address token1Out, uint256 _amountIn, uint256 _expectedOut) returns(uint256)
func (_SwapWrapUniv2 *SwapWrapUniv2TransactorSession) Swap(pair common.Address, token0In common.Address, token1Out common.Address, _amountIn *big.Int, _expectedOut *big.Int) (*types.Transaction, error) {
	return _SwapWrapUniv2.Contract.Swap(&_SwapWrapUniv2.TransactOpts, pair, token0In, token1Out, _amountIn, _expectedOut)
}

// SwapWrapUniv2SwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the SwapWrapUniv2 contract.
type SwapWrapUniv2SwapIterator struct {
	Event *SwapWrapUniv2Swap // Event containing the contract specifics and raw log

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
func (it *SwapWrapUniv2SwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapWrapUniv2Swap)
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
		it.Event = new(SwapWrapUniv2Swap)
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
func (it *SwapWrapUniv2SwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapWrapUniv2SwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapWrapUniv2Swap represents a Swap event raised by the SwapWrapUniv2 contract.
type SwapWrapUniv2Swap struct {
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
func (_SwapWrapUniv2 *SwapWrapUniv2Filterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address) (*SwapWrapUniv2SwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SwapWrapUniv2.contract.FilterLogs(opts, "Swap", senderRule)
	if err != nil {
		return nil, err
	}
	return &SwapWrapUniv2SwapIterator{contract: _SwapWrapUniv2.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xd6d34547c69c5ee3d2667625c188acf1006abb93e0ee7cf03925c67cf7760413.
//
// Solidity: event Swap(address indexed sender, address token0, address token1, uint256 amount0Out, uint256 amount1Out, uint256 amountOut)
func (_SwapWrapUniv2 *SwapWrapUniv2Filterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *SwapWrapUniv2Swap, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SwapWrapUniv2.contract.WatchLogs(opts, "Swap", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapWrapUniv2Swap)
				if err := _SwapWrapUniv2.contract.UnpackLog(event, "Swap", log); err != nil {
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
func (_SwapWrapUniv2 *SwapWrapUniv2Filterer) ParseSwap(log types.Log) (*SwapWrapUniv2Swap, error) {
	event := new(SwapWrapUniv2Swap)
	if err := _SwapWrapUniv2.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
