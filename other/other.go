// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// AbiABI is the input ABI used to generate the binding from.
const AbiABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"Login\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"Register\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"SetPassword\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"}],\"name\":\"compare\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"}],\"name\":\"login\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_passwdAgain\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_passwdNew\",\"type\":\"string\"}],\"name\":\"setPassword\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Abi is an auto generated Go binding around an Ethereum contract.
type Abi struct {
	AbiCaller     // Read-only binding to the contract
	AbiTransactor // Write-only binding to the contract
	AbiFilterer   // Log filterer for contract events
}

// AbiCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbiSession struct {
	Contract     *Abi              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbiCallerSession struct {
	Contract *AbiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AbiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbiTransactorSession struct {
	Contract     *AbiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbiRaw struct {
	Contract *Abi // Generic contract binding to access the raw methods on
}

// AbiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbiCallerRaw struct {
	Contract *AbiCaller // Generic read-only contract binding to access the raw methods on
}

// AbiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbiTransactorRaw struct {
	Contract *AbiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbi creates a new instance of Abi, bound to a specific deployed contract.
func NewAbi(address common.Address, backend bind.ContractBackend) (*Abi, error) {
	contract, err := bindAbi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abi{AbiCaller: AbiCaller{contract: contract}, AbiTransactor: AbiTransactor{contract: contract}, AbiFilterer: AbiFilterer{contract: contract}}, nil
}

// NewAbiCaller creates a new read-only instance of Abi, bound to a specific deployed contract.
func NewAbiCaller(address common.Address, caller bind.ContractCaller) (*AbiCaller, error) {
	contract, err := bindAbi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbiCaller{contract: contract}, nil
}

// NewAbiTransactor creates a new write-only instance of Abi, bound to a specific deployed contract.
func NewAbiTransactor(address common.Address, transactor bind.ContractTransactor) (*AbiTransactor, error) {
	contract, err := bindAbi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbiTransactor{contract: contract}, nil
}

// NewAbiFilterer creates a new log filterer instance of Abi, bound to a specific deployed contract.
func NewAbiFilterer(address common.Address, filterer bind.ContractFilterer) (*AbiFilterer, error) {
	contract, err := bindAbi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbiFilterer{contract: contract}, nil
}

// bindAbi binds a generic wrapper to an already deployed contract.
func bindAbi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abi *AbiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abi.Contract.AbiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abi *AbiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.Contract.AbiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abi *AbiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abi.Contract.AbiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abi *AbiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abi *AbiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abi *AbiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abi.Contract.contract.Transact(opts, method, params...)
}

// Compare is a free data retrieval call binding the contract method 0x46a36a1c.
//
// Solidity: function compare(uint8 id, string passwd) view returns()
func (_Abi *AbiCaller) Compare(opts *bind.CallOpts, id uint8, passwd string) error {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "compare", id, passwd)

	if err != nil {
		return err
	}

	return err

}

// Compare is a free data retrieval call binding the contract method 0x46a36a1c.
//
// Solidity: function compare(uint8 id, string passwd) view returns()
func (_Abi *AbiSession) Compare(id uint8, passwd string) error {
	return _Abi.Contract.Compare(&_Abi.CallOpts, id, passwd)
}

// Compare is a free data retrieval call binding the contract method 0x46a36a1c.
//
// Solidity: function compare(uint8 id, string passwd) view returns()
func (_Abi *AbiCallerSession) Compare(id uint8, passwd string) error {
	return _Abi.Contract.Compare(&_Abi.CallOpts, id, passwd)
}

// Login is a paid mutator transaction binding the contract method 0x9d1cba1a.
//
// Solidity: function login(uint8 id, string passwd) returns(bool)
func (_Abi *AbiTransactor) Login(opts *bind.TransactOpts, id uint8, passwd string) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "login", id, passwd)
}

// Login is a paid mutator transaction binding the contract method 0x9d1cba1a.
//
// Solidity: function login(uint8 id, string passwd) returns(bool)
func (_Abi *AbiSession) Login(id uint8, passwd string) (*types.Transaction, error) {
	return _Abi.Contract.Login(&_Abi.TransactOpts, id, passwd)
}

// Login is a paid mutator transaction binding the contract method 0x9d1cba1a.
//
// Solidity: function login(uint8 id, string passwd) returns(bool)
func (_Abi *AbiTransactorSession) Login(id uint8, passwd string) (*types.Transaction, error) {
	return _Abi.Contract.Login(&_Abi.TransactOpts, id, passwd)
}

// Register is a paid mutator transaction binding the contract method 0x89dc7021.
//
// Solidity: function register(uint8 id, string passwd, string _passwdAgain) returns(bool)
func (_Abi *AbiTransactor) Register(opts *bind.TransactOpts, id uint8, passwd string, _passwdAgain string) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "register", id, passwd, _passwdAgain)
}

// Register is a paid mutator transaction binding the contract method 0x89dc7021.
//
// Solidity: function register(uint8 id, string passwd, string _passwdAgain) returns(bool)
func (_Abi *AbiSession) Register(id uint8, passwd string, _passwdAgain string) (*types.Transaction, error) {
	return _Abi.Contract.Register(&_Abi.TransactOpts, id, passwd, _passwdAgain)
}

// Register is a paid mutator transaction binding the contract method 0x89dc7021.
//
// Solidity: function register(uint8 id, string passwd, string _passwdAgain) returns(bool)
func (_Abi *AbiTransactorSession) Register(id uint8, passwd string, _passwdAgain string) (*types.Transaction, error) {
	return _Abi.Contract.Register(&_Abi.TransactOpts, id, passwd, _passwdAgain)
}

// SetPassword is a paid mutator transaction binding the contract method 0xd9efca05.
//
// Solidity: function setPassword(uint8 id, string passwd, string _passwdNew) returns(bool)
func (_Abi *AbiTransactor) SetPassword(opts *bind.TransactOpts, id uint8, passwd string, _passwdNew string) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "setPassword", id, passwd, _passwdNew)
}

// SetPassword is a paid mutator transaction binding the contract method 0xd9efca05.
//
// Solidity: function setPassword(uint8 id, string passwd, string _passwdNew) returns(bool)
func (_Abi *AbiSession) SetPassword(id uint8, passwd string, _passwdNew string) (*types.Transaction, error) {
	return _Abi.Contract.SetPassword(&_Abi.TransactOpts, id, passwd, _passwdNew)
}

// SetPassword is a paid mutator transaction binding the contract method 0xd9efca05.
//
// Solidity: function setPassword(uint8 id, string passwd, string _passwdNew) returns(bool)
func (_Abi *AbiTransactorSession) SetPassword(id uint8, passwd string, _passwdNew string) (*types.Transaction, error) {
	return _Abi.Contract.SetPassword(&_Abi.TransactOpts, id, passwd, _passwdNew)
}

// AbiLoginIterator is returned from FilterLogin and is used to iterate over the raw logs and unpacked data for Login events raised by the Abi contract.
type AbiLoginIterator struct {
	Event *AbiLogin // Event containing the contract specifics and raw log

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
func (it *AbiLoginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiLogin)
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
		it.Event = new(AbiLogin)
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
func (it *AbiLoginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiLoginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiLogin represents a Login event raised by the Abi contract.
type AbiLogin struct {
	Id   uint8
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogin is a free log retrieval operation binding the contract event 0x554145b27dd70d8d66a4693ccd8b2b4793b89e6fd4e0c1596c71f5ff4844b042.
//
// Solidity: event Login(uint8 id, uint256 time)
func (_Abi *AbiFilterer) FilterLogin(opts *bind.FilterOpts) (*AbiLoginIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Login")
	if err != nil {
		return nil, err
	}
	return &AbiLoginIterator{contract: _Abi.contract, event: "Login", logs: logs, sub: sub}, nil
}

// WatchLogin is a free log subscription operation binding the contract event 0x554145b27dd70d8d66a4693ccd8b2b4793b89e6fd4e0c1596c71f5ff4844b042.
//
// Solidity: event Login(uint8 id, uint256 time)
func (_Abi *AbiFilterer) WatchLogin(opts *bind.WatchOpts, sink chan<- *AbiLogin) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Login")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiLogin)
				if err := _Abi.contract.UnpackLog(event, "Login", log); err != nil {
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

// ParseLogin is a log parse operation binding the contract event 0x554145b27dd70d8d66a4693ccd8b2b4793b89e6fd4e0c1596c71f5ff4844b042.
//
// Solidity: event Login(uint8 id, uint256 time)
func (_Abi *AbiFilterer) ParseLogin(log types.Log) (*AbiLogin, error) {
	event := new(AbiLogin)
	if err := _Abi.contract.UnpackLog(event, "Login", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AbiRegisterIterator is returned from FilterRegister and is used to iterate over the raw logs and unpacked data for Register events raised by the Abi contract.
type AbiRegisterIterator struct {
	Event *AbiRegister // Event containing the contract specifics and raw log

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
func (it *AbiRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiRegister)
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
		it.Event = new(AbiRegister)
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
func (it *AbiRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiRegister represents a Register event raised by the Abi contract.
type AbiRegister struct {
	Id   uint8
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRegister is a free log retrieval operation binding the contract event 0xee4a376ed37ae060b62f7fa13533b1bc1c364bcf9e23fc63075d3e691fa16cd0.
//
// Solidity: event Register(uint8 id, uint256 time)
func (_Abi *AbiFilterer) FilterRegister(opts *bind.FilterOpts) (*AbiRegisterIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Register")
	if err != nil {
		return nil, err
	}
	return &AbiRegisterIterator{contract: _Abi.contract, event: "Register", logs: logs, sub: sub}, nil
}

// WatchRegister is a free log subscription operation binding the contract event 0xee4a376ed37ae060b62f7fa13533b1bc1c364bcf9e23fc63075d3e691fa16cd0.
//
// Solidity: event Register(uint8 id, uint256 time)
func (_Abi *AbiFilterer) WatchRegister(opts *bind.WatchOpts, sink chan<- *AbiRegister) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Register")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiRegister)
				if err := _Abi.contract.UnpackLog(event, "Register", log); err != nil {
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

// ParseRegister is a log parse operation binding the contract event 0xee4a376ed37ae060b62f7fa13533b1bc1c364bcf9e23fc63075d3e691fa16cd0.
//
// Solidity: event Register(uint8 id, uint256 time)
func (_Abi *AbiFilterer) ParseRegister(log types.Log) (*AbiRegister, error) {
	event := new(AbiRegister)
	if err := _Abi.contract.UnpackLog(event, "Register", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AbiSetPasswordIterator is returned from FilterSetPassword and is used to iterate over the raw logs and unpacked data for SetPassword events raised by the Abi contract.
type AbiSetPasswordIterator struct {
	Event *AbiSetPassword // Event containing the contract specifics and raw log

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
func (it *AbiSetPasswordIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiSetPassword)
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
		it.Event = new(AbiSetPassword)
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
func (it *AbiSetPasswordIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiSetPasswordIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiSetPassword represents a SetPassword event raised by the Abi contract.
type AbiSetPassword struct {
	Id   uint8
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetPassword is a free log retrieval operation binding the contract event 0x70de75321ca0214c75e5784f8c57b5474012c521cfee2e1c6c88b8c3d439a3cc.
//
// Solidity: event SetPassword(uint8 id, uint256 time)
func (_Abi *AbiFilterer) FilterSetPassword(opts *bind.FilterOpts) (*AbiSetPasswordIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "SetPassword")
	if err != nil {
		return nil, err
	}
	return &AbiSetPasswordIterator{contract: _Abi.contract, event: "SetPassword", logs: logs, sub: sub}, nil
}

// WatchSetPassword is a free log subscription operation binding the contract event 0x70de75321ca0214c75e5784f8c57b5474012c521cfee2e1c6c88b8c3d439a3cc.
//
// Solidity: event SetPassword(uint8 id, uint256 time)
func (_Abi *AbiFilterer) WatchSetPassword(opts *bind.WatchOpts, sink chan<- *AbiSetPassword) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "SetPassword")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiSetPassword)
				if err := _Abi.contract.UnpackLog(event, "SetPassword", log); err != nil {
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

// ParseSetPassword is a log parse operation binding the contract event 0x70de75321ca0214c75e5784f8c57b5474012c521cfee2e1c6c88b8c3d439a3cc.
//
// Solidity: event SetPassword(uint8 id, uint256 time)
func (_Abi *AbiFilterer) ParseSetPassword(log types.Log) (*AbiSetPassword, error) {
	event := new(AbiSetPassword)
	if err := _Abi.contract.UnpackLog(event, "SetPassword", log); err != nil {
		return nil, err
	}
	return event, nil
}
