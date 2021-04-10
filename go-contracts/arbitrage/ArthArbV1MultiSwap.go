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

// ArthArbV1MultiSwapABI is the input ABI used to generate the binding from.
const ArthArbV1MultiSwapABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registryAddressr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"borrowMade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"tradeMade\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"loanAsset\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"arbitrage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"swapWrapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"exchange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_params\",\"type\":\"bytes\"}],\"name\":\"executeOperation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// ArthArbV1MultiSwap is an auto generated Go binding around an Ethereum contract.
type ArthArbV1MultiSwap struct {
	ArthArbV1MultiSwapCaller     // Read-only binding to the contract
	ArthArbV1MultiSwapTransactor // Write-only binding to the contract
	ArthArbV1MultiSwapFilterer   // Log filterer for contract events
}

// ArthArbV1MultiSwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArthArbV1MultiSwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArthArbV1MultiSwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArthArbV1MultiSwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArthArbV1MultiSwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArthArbV1MultiSwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArthArbV1MultiSwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArthArbV1MultiSwapSession struct {
	Contract     *ArthArbV1MultiSwap // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ArthArbV1MultiSwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArthArbV1MultiSwapCallerSession struct {
	Contract *ArthArbV1MultiSwapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ArthArbV1MultiSwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArthArbV1MultiSwapTransactorSession struct {
	Contract     *ArthArbV1MultiSwapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ArthArbV1MultiSwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArthArbV1MultiSwapRaw struct {
	Contract *ArthArbV1MultiSwap // Generic contract binding to access the raw methods on
}

// ArthArbV1MultiSwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArthArbV1MultiSwapCallerRaw struct {
	Contract *ArthArbV1MultiSwapCaller // Generic read-only contract binding to access the raw methods on
}

// ArthArbV1MultiSwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArthArbV1MultiSwapTransactorRaw struct {
	Contract *ArthArbV1MultiSwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArthArbV1MultiSwap creates a new instance of ArthArbV1MultiSwap, bound to a specific deployed contract.
func NewArthArbV1MultiSwap(address common.Address, backend bind.ContractBackend) (*ArthArbV1MultiSwap, error) {
	contract, err := bindArthArbV1MultiSwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArthArbV1MultiSwap{ArthArbV1MultiSwapCaller: ArthArbV1MultiSwapCaller{contract: contract}, ArthArbV1MultiSwapTransactor: ArthArbV1MultiSwapTransactor{contract: contract}, ArthArbV1MultiSwapFilterer: ArthArbV1MultiSwapFilterer{contract: contract}}, nil
}

// NewArthArbV1MultiSwapCaller creates a new read-only instance of ArthArbV1MultiSwap, bound to a specific deployed contract.
func NewArthArbV1MultiSwapCaller(address common.Address, caller bind.ContractCaller) (*ArthArbV1MultiSwapCaller, error) {
	contract, err := bindArthArbV1MultiSwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArthArbV1MultiSwapCaller{contract: contract}, nil
}

// NewArthArbV1MultiSwapTransactor creates a new write-only instance of ArthArbV1MultiSwap, bound to a specific deployed contract.
func NewArthArbV1MultiSwapTransactor(address common.Address, transactor bind.ContractTransactor) (*ArthArbV1MultiSwapTransactor, error) {
	contract, err := bindArthArbV1MultiSwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArthArbV1MultiSwapTransactor{contract: contract}, nil
}

// NewArthArbV1MultiSwapFilterer creates a new log filterer instance of ArthArbV1MultiSwap, bound to a specific deployed contract.
func NewArthArbV1MultiSwapFilterer(address common.Address, filterer bind.ContractFilterer) (*ArthArbV1MultiSwapFilterer, error) {
	contract, err := bindArthArbV1MultiSwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArthArbV1MultiSwapFilterer{contract: contract}, nil
}

// bindArthArbV1MultiSwap binds a generic wrapper to an already deployed contract.
func bindArthArbV1MultiSwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArthArbV1MultiSwapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArthArbV1MultiSwap *ArthArbV1MultiSwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArthArbV1MultiSwap.Contract.ArthArbV1MultiSwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArthArbV1MultiSwap *ArthArbV1MultiSwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArthArbV1MultiSwap.Contract.ArthArbV1MultiSwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArthArbV1MultiSwap *ArthArbV1MultiSwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArthArbV1MultiSwap.Contract.ArthArbV1MultiSwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArthArbV1MultiSwap *ArthArbV1MultiSwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArthArbV1MultiSwap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArthArbV1MultiSwap *ArthArbV1MultiSwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArthArbV1MultiSwap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArthArbV1MultiSwap *ArthArbV1MultiSwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArthArbV1MultiSwap.Contract.contract.Transact(opts, method, params...)
}

// Arbitrage is a paid mutator transaction binding the contract method 0xe983cac1.
//
// Solidity: function arbitrage(address lender, address loanAsset, address[] path, uint256[] amounts) returns()
func (_ArthArbV1MultiSwap *ArthArbV1MultiSwapTransactor) Arbitrage(opts *bind.TransactOpts, lender common.Address, loanAsset common.Address, path []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _ArthArbV1MultiSwap.contract.Transact(opts, "arbitrage", lender, loanAsset, path, amounts)
}

// Arbitrage is a paid mutator transaction binding the contract method 0xe983cac1.
//
// Solidity: function arbitrage(address lender, address loanAsset, address[] path, uint256[] amounts) returns()
func (_ArthArbV1MultiSwap *ArthArbV1MultiSwapSession) Arbitrage(lender common.Address, loanAsset common.Address, path []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _ArthArbV1MultiSwap.Contract.Arbitrage(&_ArthArbV1MultiSwap.TransactOpts, lender, loanAsset, path, amounts)
}

// Arbitrage is a paid mutator transaction binding the contract method 0xe983cac1.
//
// Solidity: function arbitrage(address lender, address loanAsset, address[] path, uint256[] amounts) returns()
func (_ArthArbV1MultiSwap *ArthArbV1MultiSwapTransactorSession) Arbitrage(lender common.Address, loanAsset common.Address, path []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _ArthArbV1MultiSwap.Contract.Arbitrage(&_ArthArbV1MultiSwap.TransactOpts, lender, loanAsset, path, amounts)
}

// Exchange is a paid mutator transaction binding the contract method 0x74c5bbf7.
//
// Solidity: function exchange(address swapWrapp, address pool, address token0, address token1, uint256 amountIn, uint256 amountOut) returns()
func (_ArthArbV1MultiSwap *ArthArbV1MultiSwapTransactor) Exchange(opts *bind.TransactOpts, swapWrapp common.Address, pool common.Address, token0 common.Address, token1 common.Address, amountIn *big.Int, amountOut *big.Int) (*types.Transaction, error) {
	return _ArthArbV1MultiSwap.contract.Transact(opts, "exchange", swapWrapp, pool, token0, token1, amountIn, amountOut)
}

// Exchange is a paid mutator transaction binding the contract method 0x74c5bbf7.
//
// Solidity: function exchange(address swapWrapp, address pool, address token0, address token1, uint256 amountIn, uint256 amountOut) returns()
func (_ArthArbV1MultiSwap *ArthArbV1MultiSwapSession) Exchange(swapWrapp common.Address, pool common.Address, token0 common.Address, token1 common.Address, amountIn *big.Int, amountOut *big.Int) (*types.Transaction, error) {
	return _ArthArbV1MultiSwap.Contract.Exchange(&_ArthArbV1MultiSwap.TransactOpts, swapWrapp, pool, token0, token1, amountIn, amountOut)
}

// Exchange is a paid mutator transaction binding the contract method 0x74c5bbf7.
//
// Solidity: function exchange(address swapWrapp, address pool, address token0, address token1, uint256 amountIn, uint256 amountOut) returns()
func (_ArthArbV1MultiSwap *ArthArbV1MultiSwapTransactorSession) Exchange(swapWrapp common.Address, pool common.Address, token0 common.Address, token1 common.Address, amountIn *big.Int, amountOut *big.Int) (*types.Transaction, error) {
	return _ArthArbV1MultiSwap.Contract.Exchange(&_ArthArbV1MultiSwap.TransactOpts, swapWrapp, pool, token0, token1, amountIn, amountOut)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0xee872558.
//
// Solidity: function executeOperation(address _reserve, uint256 _amount, uint256 _fee, bytes _params) returns()
func (_ArthArbV1MultiSwap *ArthArbV1MultiSwapTransactor) ExecuteOperation(opts *bind.TransactOpts, _reserve common.Address, _amount *big.Int, _fee *big.Int, _params []byte) (*types.Transaction, error) {
	return _ArthArbV1MultiSwap.contract.Transact(opts, "executeOperation", _reserve, _amount, _fee, _params)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0xee872558.
//
// Solidity: function executeOperation(address _reserve, uint256 _amount, uint256 _fee, bytes _params) returns()
func (_ArthArbV1MultiSwap *ArthArbV1MultiSwapSession) ExecuteOperation(_reserve common.Address, _amount *big.Int, _fee *big.Int, _params []byte) (*types.Transaction, error) {
	return _ArthArbV1MultiSwap.Contract.ExecuteOperation(&_ArthArbV1MultiSwap.TransactOpts, _reserve, _amount, _fee, _params)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0xee872558.
//
// Solidity: function executeOperation(address _reserve, uint256 _amount, uint256 _fee, bytes _params) returns()
func (_ArthArbV1MultiSwap *ArthArbV1MultiSwapTransactorSession) ExecuteOperation(_reserve common.Address, _amount *big.Int, _fee *big.Int, _params []byte) (*types.Transaction, error) {
	return _ArthArbV1MultiSwap.Contract.ExecuteOperation(&_ArthArbV1MultiSwap.TransactOpts, _reserve, _amount, _fee, _params)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ArthArbV1MultiSwap *ArthArbV1MultiSwapTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArthArbV1MultiSwap.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ArthArbV1MultiSwap *ArthArbV1MultiSwapSession) Receive() (*types.Transaction, error) {
	return _ArthArbV1MultiSwap.Contract.Receive(&_ArthArbV1MultiSwap.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ArthArbV1MultiSwap *ArthArbV1MultiSwapTransactorSession) Receive() (*types.Transaction, error) {
	return _ArthArbV1MultiSwap.Contract.Receive(&_ArthArbV1MultiSwap.TransactOpts)
}

// ArthArbV1MultiSwapReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the ArthArbV1MultiSwap contract.
type ArthArbV1MultiSwapReceivedIterator struct {
	Event *ArthArbV1MultiSwapReceived // Event containing the contract specifics and raw log

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
func (it *ArthArbV1MultiSwapReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArthArbV1MultiSwapReceived)
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
		it.Event = new(ArthArbV1MultiSwapReceived)
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
func (it *ArthArbV1MultiSwapReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArthArbV1MultiSwapReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArthArbV1MultiSwapReceived represents a Received event raised by the ArthArbV1MultiSwap contract.
type ArthArbV1MultiSwapReceived struct {
	Caller  common.Address
	Amount  *big.Int
	Message string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x59e04c3f0d44b7caf6e8ef854b61d9a51cf1960d7a88ff6356cc5e946b4b5832.
//
// Solidity: event Received(address caller, uint256 amount, string message)
func (_ArthArbV1MultiSwap *ArthArbV1MultiSwapFilterer) FilterReceived(opts *bind.FilterOpts) (*ArthArbV1MultiSwapReceivedIterator, error) {

	logs, sub, err := _ArthArbV1MultiSwap.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &ArthArbV1MultiSwapReceivedIterator{contract: _ArthArbV1MultiSwap.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x59e04c3f0d44b7caf6e8ef854b61d9a51cf1960d7a88ff6356cc5e946b4b5832.
//
// Solidity: event Received(address caller, uint256 amount, string message)
func (_ArthArbV1MultiSwap *ArthArbV1MultiSwapFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *ArthArbV1MultiSwapReceived) (event.Subscription, error) {

	logs, sub, err := _ArthArbV1MultiSwap.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArthArbV1MultiSwapReceived)
				if err := _ArthArbV1MultiSwap.contract.UnpackLog(event, "Received", log); err != nil {
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

// ParseReceived is a log parse operation binding the contract event 0x59e04c3f0d44b7caf6e8ef854b61d9a51cf1960d7a88ff6356cc5e946b4b5832.
//
// Solidity: event Received(address caller, uint256 amount, string message)
func (_ArthArbV1MultiSwap *ArthArbV1MultiSwapFilterer) ParseReceived(log types.Log) (*ArthArbV1MultiSwapReceived, error) {
	event := new(ArthArbV1MultiSwapReceived)
	if err := _ArthArbV1MultiSwap.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArthArbV1MultiSwapBorrowMadeIterator is returned from FilterBorrowMade and is used to iterate over the raw logs and unpacked data for BorrowMade events raised by the ArthArbV1MultiSwap contract.
type ArthArbV1MultiSwapBorrowMadeIterator struct {
	Event *ArthArbV1MultiSwapBorrowMade // Event containing the contract specifics and raw log

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
func (it *ArthArbV1MultiSwapBorrowMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArthArbV1MultiSwapBorrowMade)
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
		it.Event = new(ArthArbV1MultiSwapBorrowMade)
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
func (it *ArthArbV1MultiSwapBorrowMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArthArbV1MultiSwapBorrowMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArthArbV1MultiSwapBorrowMade represents a BorrowMade event raised by the ArthArbV1MultiSwap contract.
type ArthArbV1MultiSwapBorrowMade struct {
	Reserve common.Address
	Amount  *big.Int
	Fee     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBorrowMade is a free log retrieval operation binding the contract event 0xcd8f01cf11df8ac6e385851ab2495f5370eff19deb3c99a29b542b49d32ba0ba.
//
// Solidity: event borrowMade(address _reserve, uint256 _amount, uint256 _fee)
func (_ArthArbV1MultiSwap *ArthArbV1MultiSwapFilterer) FilterBorrowMade(opts *bind.FilterOpts) (*ArthArbV1MultiSwapBorrowMadeIterator, error) {

	logs, sub, err := _ArthArbV1MultiSwap.contract.FilterLogs(opts, "borrowMade")
	if err != nil {
		return nil, err
	}
	return &ArthArbV1MultiSwapBorrowMadeIterator{contract: _ArthArbV1MultiSwap.contract, event: "borrowMade", logs: logs, sub: sub}, nil
}

// WatchBorrowMade is a free log subscription operation binding the contract event 0xcd8f01cf11df8ac6e385851ab2495f5370eff19deb3c99a29b542b49d32ba0ba.
//
// Solidity: event borrowMade(address _reserve, uint256 _amount, uint256 _fee)
func (_ArthArbV1MultiSwap *ArthArbV1MultiSwapFilterer) WatchBorrowMade(opts *bind.WatchOpts, sink chan<- *ArthArbV1MultiSwapBorrowMade) (event.Subscription, error) {

	logs, sub, err := _ArthArbV1MultiSwap.contract.WatchLogs(opts, "borrowMade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArthArbV1MultiSwapBorrowMade)
				if err := _ArthArbV1MultiSwap.contract.UnpackLog(event, "borrowMade", log); err != nil {
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
func (_ArthArbV1MultiSwap *ArthArbV1MultiSwapFilterer) ParseBorrowMade(log types.Log) (*ArthArbV1MultiSwapBorrowMade, error) {
	event := new(ArthArbV1MultiSwapBorrowMade)
	if err := _ArthArbV1MultiSwap.contract.UnpackLog(event, "borrowMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArthArbV1MultiSwapTradeMadeIterator is returned from FilterTradeMade and is used to iterate over the raw logs and unpacked data for TradeMade events raised by the ArthArbV1MultiSwap contract.
type ArthArbV1MultiSwapTradeMadeIterator struct {
	Event *ArthArbV1MultiSwapTradeMade // Event containing the contract specifics and raw log

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
func (it *ArthArbV1MultiSwapTradeMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArthArbV1MultiSwapTradeMade)
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
		it.Event = new(ArthArbV1MultiSwapTradeMade)
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
func (it *ArthArbV1MultiSwapTradeMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArthArbV1MultiSwapTradeMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArthArbV1MultiSwapTradeMade represents a TradeMade event raised by the ArthArbV1MultiSwap contract.
type ArthArbV1MultiSwapTradeMade struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTradeMade is a free log retrieval operation binding the contract event 0xfe94c5aba7ab62637d0f86536cc1e8f7d3650c78fc65a37e09506fdd075bf007.
//
// Solidity: event tradeMade(uint256 _amount)
func (_ArthArbV1MultiSwap *ArthArbV1MultiSwapFilterer) FilterTradeMade(opts *bind.FilterOpts) (*ArthArbV1MultiSwapTradeMadeIterator, error) {

	logs, sub, err := _ArthArbV1MultiSwap.contract.FilterLogs(opts, "tradeMade")
	if err != nil {
		return nil, err
	}
	return &ArthArbV1MultiSwapTradeMadeIterator{contract: _ArthArbV1MultiSwap.contract, event: "tradeMade", logs: logs, sub: sub}, nil
}

// WatchTradeMade is a free log subscription operation binding the contract event 0xfe94c5aba7ab62637d0f86536cc1e8f7d3650c78fc65a37e09506fdd075bf007.
//
// Solidity: event tradeMade(uint256 _amount)
func (_ArthArbV1MultiSwap *ArthArbV1MultiSwapFilterer) WatchTradeMade(opts *bind.WatchOpts, sink chan<- *ArthArbV1MultiSwapTradeMade) (event.Subscription, error) {

	logs, sub, err := _ArthArbV1MultiSwap.contract.WatchLogs(opts, "tradeMade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArthArbV1MultiSwapTradeMade)
				if err := _ArthArbV1MultiSwap.contract.UnpackLog(event, "tradeMade", log); err != nil {
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

// ParseTradeMade is a log parse operation binding the contract event 0xfe94c5aba7ab62637d0f86536cc1e8f7d3650c78fc65a37e09506fdd075bf007.
//
// Solidity: event tradeMade(uint256 _amount)
func (_ArthArbV1MultiSwap *ArthArbV1MultiSwapFilterer) ParseTradeMade(log types.Log) (*ArthArbV1MultiSwapTradeMade, error) {
	event := new(ArthArbV1MultiSwapTradeMade)
	if err := _ArthArbV1MultiSwap.contract.UnpackLog(event, "tradeMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
