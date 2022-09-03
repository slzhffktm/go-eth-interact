// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package helloworld

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// HelloworldMetaData contains all meta data concerning the Helloworld contract.
var HelloworldMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"greetOwner\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hello\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newName\",\"type\":\"string\"}],\"name\":\"updateOwnerName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// HelloworldABI is the input ABI used to generate the binding from.
// Deprecated: Use HelloworldMetaData.ABI instead.
var HelloworldABI = HelloworldMetaData.ABI

// Helloworld is an auto generated Go binding around an Ethereum contract.
type Helloworld struct {
	HelloworldCaller     // Read-only binding to the contract
	HelloworldTransactor // Write-only binding to the contract
	HelloworldFilterer   // Log filterer for contract events
}

// HelloworldCaller is an auto generated read-only Go binding around an Ethereum contract.
type HelloworldCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloworldTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HelloworldTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloworldFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HelloworldFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloworldSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HelloworldSession struct {
	Contract     *Helloworld       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelloworldCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HelloworldCallerSession struct {
	Contract *HelloworldCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// HelloworldTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HelloworldTransactorSession struct {
	Contract     *HelloworldTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// HelloworldRaw is an auto generated low-level Go binding around an Ethereum contract.
type HelloworldRaw struct {
	Contract *Helloworld // Generic contract binding to access the raw methods on
}

// HelloworldCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HelloworldCallerRaw struct {
	Contract *HelloworldCaller // Generic read-only contract binding to access the raw methods on
}

// HelloworldTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HelloworldTransactorRaw struct {
	Contract *HelloworldTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHelloworld creates a new instance of Helloworld, bound to a specific deployed contract.
func NewHelloworld(address common.Address, backend bind.ContractBackend) (*Helloworld, error) {
	contract, err := bindHelloworld(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Helloworld{HelloworldCaller: HelloworldCaller{contract: contract}, HelloworldTransactor: HelloworldTransactor{contract: contract}, HelloworldFilterer: HelloworldFilterer{contract: contract}}, nil
}

// NewHelloworldCaller creates a new read-only instance of Helloworld, bound to a specific deployed contract.
func NewHelloworldCaller(address common.Address, caller bind.ContractCaller) (*HelloworldCaller, error) {
	contract, err := bindHelloworld(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HelloworldCaller{contract: contract}, nil
}

// NewHelloworldTransactor creates a new write-only instance of Helloworld, bound to a specific deployed contract.
func NewHelloworldTransactor(address common.Address, transactor bind.ContractTransactor) (*HelloworldTransactor, error) {
	contract, err := bindHelloworld(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HelloworldTransactor{contract: contract}, nil
}

// NewHelloworldFilterer creates a new log filterer instance of Helloworld, bound to a specific deployed contract.
func NewHelloworldFilterer(address common.Address, filterer bind.ContractFilterer) (*HelloworldFilterer, error) {
	contract, err := bindHelloworld(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HelloworldFilterer{contract: contract}, nil
}

// bindHelloworld binds a generic wrapper to an already deployed contract.
func bindHelloworld(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HelloworldABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Helloworld *HelloworldRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Helloworld.Contract.HelloworldCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Helloworld *HelloworldRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Helloworld.Contract.HelloworldTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Helloworld *HelloworldRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Helloworld.Contract.HelloworldTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Helloworld *HelloworldCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Helloworld.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Helloworld *HelloworldTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Helloworld.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Helloworld *HelloworldTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Helloworld.Contract.contract.Transact(opts, method, params...)
}

// GreetOwner is a free data retrieval call binding the contract method 0x897640ec.
//
// Solidity: function greetOwner() view returns(string)
func (_Helloworld *HelloworldCaller) GreetOwner(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Helloworld.contract.Call(opts, &out, "greetOwner")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GreetOwner is a free data retrieval call binding the contract method 0x897640ec.
//
// Solidity: function greetOwner() view returns(string)
func (_Helloworld *HelloworldSession) GreetOwner() (string, error) {
	return _Helloworld.Contract.GreetOwner(&_Helloworld.CallOpts)
}

// GreetOwner is a free data retrieval call binding the contract method 0x897640ec.
//
// Solidity: function greetOwner() view returns(string)
func (_Helloworld *HelloworldCallerSession) GreetOwner() (string, error) {
	return _Helloworld.Contract.GreetOwner(&_Helloworld.CallOpts)
}

// Hello is a free data retrieval call binding the contract method 0x19ff1d21.
//
// Solidity: function hello() pure returns(string)
func (_Helloworld *HelloworldCaller) Hello(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Helloworld.contract.Call(opts, &out, "hello")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Hello is a free data retrieval call binding the contract method 0x19ff1d21.
//
// Solidity: function hello() pure returns(string)
func (_Helloworld *HelloworldSession) Hello() (string, error) {
	return _Helloworld.Contract.Hello(&_Helloworld.CallOpts)
}

// Hello is a free data retrieval call binding the contract method 0x19ff1d21.
//
// Solidity: function hello() pure returns(string)
func (_Helloworld *HelloworldCallerSession) Hello() (string, error) {
	return _Helloworld.Contract.Hello(&_Helloworld.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Helloworld *HelloworldCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Helloworld.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Helloworld *HelloworldSession) Owner() (common.Address, error) {
	return _Helloworld.Contract.Owner(&_Helloworld.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Helloworld *HelloworldCallerSession) Owner() (common.Address, error) {
	return _Helloworld.Contract.Owner(&_Helloworld.CallOpts)
}

// OwnerName is a free data retrieval call binding the contract method 0x65ac2320.
//
// Solidity: function ownerName() view returns(string)
func (_Helloworld *HelloworldCaller) OwnerName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Helloworld.contract.Call(opts, &out, "ownerName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// OwnerName is a free data retrieval call binding the contract method 0x65ac2320.
//
// Solidity: function ownerName() view returns(string)
func (_Helloworld *HelloworldSession) OwnerName() (string, error) {
	return _Helloworld.Contract.OwnerName(&_Helloworld.CallOpts)
}

// OwnerName is a free data retrieval call binding the contract method 0x65ac2320.
//
// Solidity: function ownerName() view returns(string)
func (_Helloworld *HelloworldCallerSession) OwnerName() (string, error) {
	return _Helloworld.Contract.OwnerName(&_Helloworld.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Helloworld *HelloworldTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Helloworld.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Helloworld *HelloworldSession) RenounceOwnership() (*types.Transaction, error) {
	return _Helloworld.Contract.RenounceOwnership(&_Helloworld.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Helloworld *HelloworldTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Helloworld.Contract.RenounceOwnership(&_Helloworld.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Helloworld *HelloworldTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Helloworld.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Helloworld *HelloworldSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Helloworld.Contract.TransferOwnership(&_Helloworld.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Helloworld *HelloworldTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Helloworld.Contract.TransferOwnership(&_Helloworld.TransactOpts, newOwner)
}

// UpdateOwnerName is a paid mutator transaction binding the contract method 0xbbc8dbf0.
//
// Solidity: function updateOwnerName(string newName) returns()
func (_Helloworld *HelloworldTransactor) UpdateOwnerName(opts *bind.TransactOpts, newName string) (*types.Transaction, error) {
	return _Helloworld.contract.Transact(opts, "updateOwnerName", newName)
}

// UpdateOwnerName is a paid mutator transaction binding the contract method 0xbbc8dbf0.
//
// Solidity: function updateOwnerName(string newName) returns()
func (_Helloworld *HelloworldSession) UpdateOwnerName(newName string) (*types.Transaction, error) {
	return _Helloworld.Contract.UpdateOwnerName(&_Helloworld.TransactOpts, newName)
}

// UpdateOwnerName is a paid mutator transaction binding the contract method 0xbbc8dbf0.
//
// Solidity: function updateOwnerName(string newName) returns()
func (_Helloworld *HelloworldTransactorSession) UpdateOwnerName(newName string) (*types.Transaction, error) {
	return _Helloworld.Contract.UpdateOwnerName(&_Helloworld.TransactOpts, newName)
}

// HelloworldOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Helloworld contract.
type HelloworldOwnershipTransferredIterator struct {
	Event *HelloworldOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HelloworldOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HelloworldOwnershipTransferred)
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
		it.Event = new(HelloworldOwnershipTransferred)
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
func (it *HelloworldOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HelloworldOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HelloworldOwnershipTransferred represents a OwnershipTransferred event raised by the Helloworld contract.
type HelloworldOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Helloworld *HelloworldFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HelloworldOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Helloworld.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HelloworldOwnershipTransferredIterator{contract: _Helloworld.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Helloworld *HelloworldFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HelloworldOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Helloworld.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HelloworldOwnershipTransferred)
				if err := _Helloworld.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Helloworld *HelloworldFilterer) ParseOwnershipTransferred(log types.Log) (*HelloworldOwnershipTransferred, error) {
	event := new(HelloworldOwnershipTransferred)
	if err := _Helloworld.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
