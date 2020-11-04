// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package validatorset

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

// BaseOwnedSetABI is the input ABI used to generate the binding from.
const BaseOwnedSetABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"recentBlocks\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPending\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_new\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"addValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recentBlocks\",\"type\":\"uint256\"}],\"name\":\"setRecentBlocks\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"finalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_initial\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"currentSet\",\"type\":\"address[]\"}],\"name\":\"ChangeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"reporter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"reported\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"malicious\",\"type\":\"bool\"}],\"name\":\"Report\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"old\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"current\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"}]"

// BaseOwnedSetFuncSigs maps the 4-byte function signature to its string representation.
var BaseOwnedSetFuncSigs = map[string]string{
	"4d238c8e": "addValidator(address)",
	"b3f05b97": "finalized()",
	"11ae9ed2": "getPending()",
	"b7ab4db5": "getValidators()",
	"8da5cb5b": "owner()",
	"0c857805": "recentBlocks()",
	"40a141ff": "removeValidator(address)",
	"13af4035": "setOwner(address)",
	"8c90ac07": "setRecentBlocks(uint256)",
}

// BaseOwnedSet is an auto generated Go binding around an Ethereum contract.
type BaseOwnedSet struct {
	BaseOwnedSetCaller     // Read-only binding to the contract
	BaseOwnedSetTransactor // Write-only binding to the contract
	BaseOwnedSetFilterer   // Log filterer for contract events
}

// BaseOwnedSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseOwnedSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseOwnedSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseOwnedSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseOwnedSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseOwnedSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseOwnedSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseOwnedSetSession struct {
	Contract     *BaseOwnedSet     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseOwnedSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseOwnedSetCallerSession struct {
	Contract *BaseOwnedSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BaseOwnedSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseOwnedSetTransactorSession struct {
	Contract     *BaseOwnedSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BaseOwnedSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseOwnedSetRaw struct {
	Contract *BaseOwnedSet // Generic contract binding to access the raw methods on
}

// BaseOwnedSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseOwnedSetCallerRaw struct {
	Contract *BaseOwnedSetCaller // Generic read-only contract binding to access the raw methods on
}

// BaseOwnedSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseOwnedSetTransactorRaw struct {
	Contract *BaseOwnedSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseOwnedSet creates a new instance of BaseOwnedSet, bound to a specific deployed contract.
func NewBaseOwnedSet(address common.Address, backend bind.ContractBackend) (*BaseOwnedSet, error) {
	contract, err := bindBaseOwnedSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseOwnedSet{BaseOwnedSetCaller: BaseOwnedSetCaller{contract: contract}, BaseOwnedSetTransactor: BaseOwnedSetTransactor{contract: contract}, BaseOwnedSetFilterer: BaseOwnedSetFilterer{contract: contract}}, nil
}

// NewBaseOwnedSetCaller creates a new read-only instance of BaseOwnedSet, bound to a specific deployed contract.
func NewBaseOwnedSetCaller(address common.Address, caller bind.ContractCaller) (*BaseOwnedSetCaller, error) {
	contract, err := bindBaseOwnedSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseOwnedSetCaller{contract: contract}, nil
}

// NewBaseOwnedSetTransactor creates a new write-only instance of BaseOwnedSet, bound to a specific deployed contract.
func NewBaseOwnedSetTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseOwnedSetTransactor, error) {
	contract, err := bindBaseOwnedSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseOwnedSetTransactor{contract: contract}, nil
}

// NewBaseOwnedSetFilterer creates a new log filterer instance of BaseOwnedSet, bound to a specific deployed contract.
func NewBaseOwnedSetFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseOwnedSetFilterer, error) {
	contract, err := bindBaseOwnedSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseOwnedSetFilterer{contract: contract}, nil
}

// bindBaseOwnedSet binds a generic wrapper to an already deployed contract.
func bindBaseOwnedSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseOwnedSetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseOwnedSet *BaseOwnedSetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaseOwnedSet.Contract.BaseOwnedSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseOwnedSet *BaseOwnedSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseOwnedSet.Contract.BaseOwnedSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseOwnedSet *BaseOwnedSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseOwnedSet.Contract.BaseOwnedSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseOwnedSet *BaseOwnedSetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaseOwnedSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseOwnedSet *BaseOwnedSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseOwnedSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseOwnedSet *BaseOwnedSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseOwnedSet.Contract.contract.Transact(opts, method, params...)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_BaseOwnedSet *BaseOwnedSetCaller) Finalized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseOwnedSet.contract.Call(opts, out, "finalized")
	return *ret0, err
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_BaseOwnedSet *BaseOwnedSetSession) Finalized() (bool, error) {
	return _BaseOwnedSet.Contract.Finalized(&_BaseOwnedSet.CallOpts)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_BaseOwnedSet *BaseOwnedSetCallerSession) Finalized() (bool, error) {
	return _BaseOwnedSet.Contract.Finalized(&_BaseOwnedSet.CallOpts)
}

// GetPending is a free data retrieval call binding the contract method 0x11ae9ed2.
//
// Solidity: function getPending() view returns(address[])
func (_BaseOwnedSet *BaseOwnedSetCaller) GetPending(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _BaseOwnedSet.contract.Call(opts, out, "getPending")
	return *ret0, err
}

// GetPending is a free data retrieval call binding the contract method 0x11ae9ed2.
//
// Solidity: function getPending() view returns(address[])
func (_BaseOwnedSet *BaseOwnedSetSession) GetPending() ([]common.Address, error) {
	return _BaseOwnedSet.Contract.GetPending(&_BaseOwnedSet.CallOpts)
}

// GetPending is a free data retrieval call binding the contract method 0x11ae9ed2.
//
// Solidity: function getPending() view returns(address[])
func (_BaseOwnedSet *BaseOwnedSetCallerSession) GetPending() ([]common.Address, error) {
	return _BaseOwnedSet.Contract.GetPending(&_BaseOwnedSet.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_BaseOwnedSet *BaseOwnedSetCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _BaseOwnedSet.contract.Call(opts, out, "getValidators")
	return *ret0, err
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_BaseOwnedSet *BaseOwnedSetSession) GetValidators() ([]common.Address, error) {
	return _BaseOwnedSet.Contract.GetValidators(&_BaseOwnedSet.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_BaseOwnedSet *BaseOwnedSetCallerSession) GetValidators() ([]common.Address, error) {
	return _BaseOwnedSet.Contract.GetValidators(&_BaseOwnedSet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BaseOwnedSet *BaseOwnedSetCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseOwnedSet.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BaseOwnedSet *BaseOwnedSetSession) Owner() (common.Address, error) {
	return _BaseOwnedSet.Contract.Owner(&_BaseOwnedSet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BaseOwnedSet *BaseOwnedSetCallerSession) Owner() (common.Address, error) {
	return _BaseOwnedSet.Contract.Owner(&_BaseOwnedSet.CallOpts)
}

// RecentBlocks is a free data retrieval call binding the contract method 0x0c857805.
//
// Solidity: function recentBlocks() view returns(uint256)
func (_BaseOwnedSet *BaseOwnedSetCaller) RecentBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaseOwnedSet.contract.Call(opts, out, "recentBlocks")
	return *ret0, err
}

// RecentBlocks is a free data retrieval call binding the contract method 0x0c857805.
//
// Solidity: function recentBlocks() view returns(uint256)
func (_BaseOwnedSet *BaseOwnedSetSession) RecentBlocks() (*big.Int, error) {
	return _BaseOwnedSet.Contract.RecentBlocks(&_BaseOwnedSet.CallOpts)
}

// RecentBlocks is a free data retrieval call binding the contract method 0x0c857805.
//
// Solidity: function recentBlocks() view returns(uint256)
func (_BaseOwnedSet *BaseOwnedSetCallerSession) RecentBlocks() (*big.Int, error) {
	return _BaseOwnedSet.Contract.RecentBlocks(&_BaseOwnedSet.CallOpts)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address _validator) returns()
func (_BaseOwnedSet *BaseOwnedSetTransactor) AddValidator(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _BaseOwnedSet.contract.Transact(opts, "addValidator", _validator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address _validator) returns()
func (_BaseOwnedSet *BaseOwnedSetSession) AddValidator(_validator common.Address) (*types.Transaction, error) {
	return _BaseOwnedSet.Contract.AddValidator(&_BaseOwnedSet.TransactOpts, _validator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address _validator) returns()
func (_BaseOwnedSet *BaseOwnedSetTransactorSession) AddValidator(_validator common.Address) (*types.Transaction, error) {
	return _BaseOwnedSet.Contract.AddValidator(&_BaseOwnedSet.TransactOpts, _validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validator) returns()
func (_BaseOwnedSet *BaseOwnedSetTransactor) RemoveValidator(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _BaseOwnedSet.contract.Transact(opts, "removeValidator", _validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validator) returns()
func (_BaseOwnedSet *BaseOwnedSetSession) RemoveValidator(_validator common.Address) (*types.Transaction, error) {
	return _BaseOwnedSet.Contract.RemoveValidator(&_BaseOwnedSet.TransactOpts, _validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validator) returns()
func (_BaseOwnedSet *BaseOwnedSetTransactorSession) RemoveValidator(_validator common.Address) (*types.Transaction, error) {
	return _BaseOwnedSet.Contract.RemoveValidator(&_BaseOwnedSet.TransactOpts, _validator)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _new) returns()
func (_BaseOwnedSet *BaseOwnedSetTransactor) SetOwner(opts *bind.TransactOpts, _new common.Address) (*types.Transaction, error) {
	return _BaseOwnedSet.contract.Transact(opts, "setOwner", _new)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _new) returns()
func (_BaseOwnedSet *BaseOwnedSetSession) SetOwner(_new common.Address) (*types.Transaction, error) {
	return _BaseOwnedSet.Contract.SetOwner(&_BaseOwnedSet.TransactOpts, _new)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _new) returns()
func (_BaseOwnedSet *BaseOwnedSetTransactorSession) SetOwner(_new common.Address) (*types.Transaction, error) {
	return _BaseOwnedSet.Contract.SetOwner(&_BaseOwnedSet.TransactOpts, _new)
}

// SetRecentBlocks is a paid mutator transaction binding the contract method 0x8c90ac07.
//
// Solidity: function setRecentBlocks(uint256 _recentBlocks) returns()
func (_BaseOwnedSet *BaseOwnedSetTransactor) SetRecentBlocks(opts *bind.TransactOpts, _recentBlocks *big.Int) (*types.Transaction, error) {
	return _BaseOwnedSet.contract.Transact(opts, "setRecentBlocks", _recentBlocks)
}

// SetRecentBlocks is a paid mutator transaction binding the contract method 0x8c90ac07.
//
// Solidity: function setRecentBlocks(uint256 _recentBlocks) returns()
func (_BaseOwnedSet *BaseOwnedSetSession) SetRecentBlocks(_recentBlocks *big.Int) (*types.Transaction, error) {
	return _BaseOwnedSet.Contract.SetRecentBlocks(&_BaseOwnedSet.TransactOpts, _recentBlocks)
}

// SetRecentBlocks is a paid mutator transaction binding the contract method 0x8c90ac07.
//
// Solidity: function setRecentBlocks(uint256 _recentBlocks) returns()
func (_BaseOwnedSet *BaseOwnedSetTransactorSession) SetRecentBlocks(_recentBlocks *big.Int) (*types.Transaction, error) {
	return _BaseOwnedSet.Contract.SetRecentBlocks(&_BaseOwnedSet.TransactOpts, _recentBlocks)
}

// BaseOwnedSetChangeFinalizedIterator is returned from FilterChangeFinalized and is used to iterate over the raw logs and unpacked data for ChangeFinalized events raised by the BaseOwnedSet contract.
type BaseOwnedSetChangeFinalizedIterator struct {
	Event *BaseOwnedSetChangeFinalized // Event containing the contract specifics and raw log

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
func (it *BaseOwnedSetChangeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseOwnedSetChangeFinalized)
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
		it.Event = new(BaseOwnedSetChangeFinalized)
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
func (it *BaseOwnedSetChangeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseOwnedSetChangeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseOwnedSetChangeFinalized represents a ChangeFinalized event raised by the BaseOwnedSet contract.
type BaseOwnedSetChangeFinalized struct {
	CurrentSet []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChangeFinalized is a free log retrieval operation binding the contract event 0x8564cd629b15f47dc310d45bcbfc9bcf5420b0d51bf0659a16c67f91d2763253.
//
// Solidity: event ChangeFinalized(address[] currentSet)
func (_BaseOwnedSet *BaseOwnedSetFilterer) FilterChangeFinalized(opts *bind.FilterOpts) (*BaseOwnedSetChangeFinalizedIterator, error) {

	logs, sub, err := _BaseOwnedSet.contract.FilterLogs(opts, "ChangeFinalized")
	if err != nil {
		return nil, err
	}
	return &BaseOwnedSetChangeFinalizedIterator{contract: _BaseOwnedSet.contract, event: "ChangeFinalized", logs: logs, sub: sub}, nil
}

// WatchChangeFinalized is a free log subscription operation binding the contract event 0x8564cd629b15f47dc310d45bcbfc9bcf5420b0d51bf0659a16c67f91d2763253.
//
// Solidity: event ChangeFinalized(address[] currentSet)
func (_BaseOwnedSet *BaseOwnedSetFilterer) WatchChangeFinalized(opts *bind.WatchOpts, sink chan<- *BaseOwnedSetChangeFinalized) (event.Subscription, error) {

	logs, sub, err := _BaseOwnedSet.contract.WatchLogs(opts, "ChangeFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseOwnedSetChangeFinalized)
				if err := _BaseOwnedSet.contract.UnpackLog(event, "ChangeFinalized", log); err != nil {
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

// ParseChangeFinalized is a log parse operation binding the contract event 0x8564cd629b15f47dc310d45bcbfc9bcf5420b0d51bf0659a16c67f91d2763253.
//
// Solidity: event ChangeFinalized(address[] currentSet)
func (_BaseOwnedSet *BaseOwnedSetFilterer) ParseChangeFinalized(log types.Log) (*BaseOwnedSetChangeFinalized, error) {
	event := new(BaseOwnedSetChangeFinalized)
	if err := _BaseOwnedSet.contract.UnpackLog(event, "ChangeFinalized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BaseOwnedSetNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the BaseOwnedSet contract.
type BaseOwnedSetNewOwnerIterator struct {
	Event *BaseOwnedSetNewOwner // Event containing the contract specifics and raw log

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
func (it *BaseOwnedSetNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseOwnedSetNewOwner)
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
		it.Event = new(BaseOwnedSetNewOwner)
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
func (it *BaseOwnedSetNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseOwnedSetNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseOwnedSetNewOwner represents a NewOwner event raised by the BaseOwnedSet contract.
type BaseOwnedSetNewOwner struct {
	Old     common.Address
	Current common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed old, address indexed current)
func (_BaseOwnedSet *BaseOwnedSetFilterer) FilterNewOwner(opts *bind.FilterOpts, old []common.Address, current []common.Address) (*BaseOwnedSetNewOwnerIterator, error) {

	var oldRule []interface{}
	for _, oldItem := range old {
		oldRule = append(oldRule, oldItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _BaseOwnedSet.contract.FilterLogs(opts, "NewOwner", oldRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &BaseOwnedSetNewOwnerIterator{contract: _BaseOwnedSet.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed old, address indexed current)
func (_BaseOwnedSet *BaseOwnedSetFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *BaseOwnedSetNewOwner, old []common.Address, current []common.Address) (event.Subscription, error) {

	var oldRule []interface{}
	for _, oldItem := range old {
		oldRule = append(oldRule, oldItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _BaseOwnedSet.contract.WatchLogs(opts, "NewOwner", oldRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseOwnedSetNewOwner)
				if err := _BaseOwnedSet.contract.UnpackLog(event, "NewOwner", log); err != nil {
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

// ParseNewOwner is a log parse operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed old, address indexed current)
func (_BaseOwnedSet *BaseOwnedSetFilterer) ParseNewOwner(log types.Log) (*BaseOwnedSetNewOwner, error) {
	event := new(BaseOwnedSetNewOwner)
	if err := _BaseOwnedSet.contract.UnpackLog(event, "NewOwner", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BaseOwnedSetReportIterator is returned from FilterReport and is used to iterate over the raw logs and unpacked data for Report events raised by the BaseOwnedSet contract.
type BaseOwnedSetReportIterator struct {
	Event *BaseOwnedSetReport // Event containing the contract specifics and raw log

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
func (it *BaseOwnedSetReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseOwnedSetReport)
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
		it.Event = new(BaseOwnedSetReport)
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
func (it *BaseOwnedSetReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseOwnedSetReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseOwnedSetReport represents a Report event raised by the BaseOwnedSet contract.
type BaseOwnedSetReport struct {
	Reporter  common.Address
	Reported  common.Address
	Malicious bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReport is a free log retrieval operation binding the contract event 0x32c78b6140c46745a46e88cd883707d70dbd2f06d13dd76fe5f499c01290da4f.
//
// Solidity: event Report(address indexed reporter, address indexed reported, bool indexed malicious)
func (_BaseOwnedSet *BaseOwnedSetFilterer) FilterReport(opts *bind.FilterOpts, reporter []common.Address, reported []common.Address, malicious []bool) (*BaseOwnedSetReportIterator, error) {

	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}
	var reportedRule []interface{}
	for _, reportedItem := range reported {
		reportedRule = append(reportedRule, reportedItem)
	}
	var maliciousRule []interface{}
	for _, maliciousItem := range malicious {
		maliciousRule = append(maliciousRule, maliciousItem)
	}

	logs, sub, err := _BaseOwnedSet.contract.FilterLogs(opts, "Report", reporterRule, reportedRule, maliciousRule)
	if err != nil {
		return nil, err
	}
	return &BaseOwnedSetReportIterator{contract: _BaseOwnedSet.contract, event: "Report", logs: logs, sub: sub}, nil
}

// WatchReport is a free log subscription operation binding the contract event 0x32c78b6140c46745a46e88cd883707d70dbd2f06d13dd76fe5f499c01290da4f.
//
// Solidity: event Report(address indexed reporter, address indexed reported, bool indexed malicious)
func (_BaseOwnedSet *BaseOwnedSetFilterer) WatchReport(opts *bind.WatchOpts, sink chan<- *BaseOwnedSetReport, reporter []common.Address, reported []common.Address, malicious []bool) (event.Subscription, error) {

	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}
	var reportedRule []interface{}
	for _, reportedItem := range reported {
		reportedRule = append(reportedRule, reportedItem)
	}
	var maliciousRule []interface{}
	for _, maliciousItem := range malicious {
		maliciousRule = append(maliciousRule, maliciousItem)
	}

	logs, sub, err := _BaseOwnedSet.contract.WatchLogs(opts, "Report", reporterRule, reportedRule, maliciousRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseOwnedSetReport)
				if err := _BaseOwnedSet.contract.UnpackLog(event, "Report", log); err != nil {
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

// ParseReport is a log parse operation binding the contract event 0x32c78b6140c46745a46e88cd883707d70dbd2f06d13dd76fe5f499c01290da4f.
//
// Solidity: event Report(address indexed reporter, address indexed reported, bool indexed malicious)
func (_BaseOwnedSet *BaseOwnedSetFilterer) ParseReport(log types.Log) (*BaseOwnedSetReport, error) {
	event := new(BaseOwnedSetReport)
	if err := _BaseOwnedSet.contract.UnpackLog(event, "Report", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OwnedABI is the input ABI used to generate the binding from.
const OwnedABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_new\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"old\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"current\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"}]"

// OwnedFuncSigs maps the 4-byte function signature to its string representation.
var OwnedFuncSigs = map[string]string{
	"8da5cb5b": "owner()",
	"13af4035": "setOwner(address)",
}

// OwnedBin is the compiled bytecode used for deploying new contracts.
var OwnedBin = "0x608060405260008054600160a060020a0319163317905534801561002257600080fd5b506101ac806100326000396000f30060806040526004361061004b5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166313af403581146100505780638da5cb5b14610080575b600080fd5b34801561005c57600080fd5b5061007e73ffffffffffffffffffffffffffffffffffffffff600435166100be565b005b34801561008c57600080fd5b50610095610164565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff1633146100e257600080fd5b6000805460405173ffffffffffffffffffffffffffffffffffffffff808516939216917f70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b236491a36000805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a7230582016bab962a4772751af56ba843471a6c1517fa6cfb7345794393325db0cef73560029"

// DeployOwned deploys a new Ethereum contract, binding an instance of Owned to it.
func DeployOwned(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Owned, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnedBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}, OwnedFilterer: OwnedFilterer{contract: contract}}, nil
}

// Owned is an auto generated Go binding around an Ethereum contract.
type Owned struct {
	OwnedCaller     // Read-only binding to the contract
	OwnedTransactor // Write-only binding to the contract
	OwnedFilterer   // Log filterer for contract events
}

// OwnedCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnedSession struct {
	Contract     *Owned            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnedCallerSession struct {
	Contract *OwnedCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OwnedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnedTransactorSession struct {
	Contract     *OwnedTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnedRaw struct {
	Contract *Owned // Generic contract binding to access the raw methods on
}

// OwnedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnedCallerRaw struct {
	Contract *OwnedCaller // Generic read-only contract binding to access the raw methods on
}

// OwnedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnedTransactorRaw struct {
	Contract *OwnedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwned creates a new instance of Owned, bound to a specific deployed contract.
func NewOwned(address common.Address, backend bind.ContractBackend) (*Owned, error) {
	contract, err := bindOwned(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}, OwnedFilterer: OwnedFilterer{contract: contract}}, nil
}

// NewOwnedCaller creates a new read-only instance of Owned, bound to a specific deployed contract.
func NewOwnedCaller(address common.Address, caller bind.ContractCaller) (*OwnedCaller, error) {
	contract, err := bindOwned(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedCaller{contract: contract}, nil
}

// NewOwnedTransactor creates a new write-only instance of Owned, bound to a specific deployed contract.
func NewOwnedTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnedTransactor, error) {
	contract, err := bindOwned(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedTransactor{contract: contract}, nil
}

// NewOwnedFilterer creates a new log filterer instance of Owned, bound to a specific deployed contract.
func NewOwnedFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnedFilterer, error) {
	contract, err := bindOwned(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnedFilterer{contract: contract}, nil
}

// bindOwned binds a generic wrapper to an already deployed contract.
func bindOwned(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.OwnedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Owned *OwnedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Owned.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Owned *OwnedSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Owned *OwnedCallerSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _new) returns()
func (_Owned *OwnedTransactor) SetOwner(opts *bind.TransactOpts, _new common.Address) (*types.Transaction, error) {
	return _Owned.contract.Transact(opts, "setOwner", _new)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _new) returns()
func (_Owned *OwnedSession) SetOwner(_new common.Address) (*types.Transaction, error) {
	return _Owned.Contract.SetOwner(&_Owned.TransactOpts, _new)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _new) returns()
func (_Owned *OwnedTransactorSession) SetOwner(_new common.Address) (*types.Transaction, error) {
	return _Owned.Contract.SetOwner(&_Owned.TransactOpts, _new)
}

// OwnedNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the Owned contract.
type OwnedNewOwnerIterator struct {
	Event *OwnedNewOwner // Event containing the contract specifics and raw log

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
func (it *OwnedNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnedNewOwner)
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
		it.Event = new(OwnedNewOwner)
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
func (it *OwnedNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnedNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnedNewOwner represents a NewOwner event raised by the Owned contract.
type OwnedNewOwner struct {
	Old     common.Address
	Current common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed old, address indexed current)
func (_Owned *OwnedFilterer) FilterNewOwner(opts *bind.FilterOpts, old []common.Address, current []common.Address) (*OwnedNewOwnerIterator, error) {

	var oldRule []interface{}
	for _, oldItem := range old {
		oldRule = append(oldRule, oldItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _Owned.contract.FilterLogs(opts, "NewOwner", oldRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &OwnedNewOwnerIterator{contract: _Owned.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed old, address indexed current)
func (_Owned *OwnedFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *OwnedNewOwner, old []common.Address, current []common.Address) (event.Subscription, error) {

	var oldRule []interface{}
	for _, oldItem := range old {
		oldRule = append(oldRule, oldItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _Owned.contract.WatchLogs(opts, "NewOwner", oldRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnedNewOwner)
				if err := _Owned.contract.UnpackLog(event, "NewOwner", log); err != nil {
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

// ParseNewOwner is a log parse operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed old, address indexed current)
func (_Owned *OwnedFilterer) ParseNewOwner(log types.Log) (*OwnedNewOwner, error) {
	event := new(OwnedNewOwner)
	if err := _Owned.contract.UnpackLog(event, "NewOwner", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelaySetABI is the input ABI used to generate the binding from.
const RelaySetABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_new\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalizeChange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_parentHash\",\"type\":\"bytes32\"},{\"name\":\"_newSet\",\"type\":\"address[]\"}],\"name\":\"initiateChange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"relayedSet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_relayedSet\",\"type\":\"address\"}],\"name\":\"setRelayed\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_validator\",\"type\":\"address\"},{\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"reportMalicious\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"systemAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_validator\",\"type\":\"address\"},{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"reportBenign\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"old\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"current\",\"type\":\"address\"}],\"name\":\"NewRelayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_parentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_newSet\",\"type\":\"address[]\"}],\"name\":\"InitiateChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"old\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"current\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"}]"

// RelaySetFuncSigs maps the 4-byte function signature to its string representation.
var RelaySetFuncSigs = map[string]string{
	"75286211": "finalizeChange()",
	"b7ab4db5": "getValidators()",
	"a0285c01": "initiateChange(bytes32,address[])",
	"8da5cb5b": "owner()",
	"ae3783d6": "relayedSet()",
	"d69f13bb": "reportBenign(address,uint256)",
	"c476dd40": "reportMalicious(address,uint256,bytes)",
	"13af4035": "setOwner(address)",
	"bd965677": "setRelayed(address)",
	"d3e848f1": "systemAddress()",
}

// RelaySetBin is the compiled bytecode used for deploying new contracts.
var RelaySetBin = "0x608060405260008054600160a060020a0319163317905534801561002257600080fd5b5060018054600160a060020a031916600260a060020a031790556106db8061004b6000396000f3006080604052600436106100a35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166313af403581146100a857806375286211146100cb5780638da5cb5b146100e0578063a0285c0114610111578063ae3783d614610135578063b7ab4db51461014a578063bd965677146101af578063c476dd40146101d0578063d3e848f114610201578063d69f13bb14610216575b600080fd5b3480156100b457600080fd5b506100c9600160a060020a036004351661023a565b005b3480156100d757600080fd5b506100c96102b9565b3480156100ec57600080fd5b506100f5610356565b60408051600160a060020a039092168252519081900360200190f35b34801561011d57600080fd5b506100c9600480359060248035908101910135610365565b34801561014157600080fd5b506100f56103d9565b34801561015657600080fd5b5061015f6103e8565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561019b578181015183820152602001610183565b505050509050019250505060405180910390f35b3480156101bb57600080fd5b506100c9600160a060020a03600435166104d7565b3480156101dc57600080fd5b506100c960048035600160a060020a0316906024803591604435918201910135610557565b34801561020d57600080fd5b506100f5610610565b34801561022257600080fd5b506100c9600160a060020a036004351660243561061f565b600054600160a060020a0316331461025157600080fd5b60008054604051600160a060020a03808516939216917f70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b236491a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154600160a060020a031633146102d057600080fd5b600260009054906101000a9004600160a060020a0316600160a060020a031663752862116040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b15801561033c57600080fd5b505af1158015610350573d6000803e3d6000fd5b50505050565b600054600160a060020a031681565b600254600160a060020a0316331461037c57600080fd5b82600019167f55252fa6eee4741b4e24a74a70e9c11fd2c2281df8d6ea13126ff845f7825c8983836040518080602001828103825284848281815260200192506020028082843760405192018290039550909350505050a2505050565b600254600160a060020a031681565b600254604080517fb7ab4db50000000000000000000000000000000000000000000000000000000081529051606092600160a060020a03169163b7ab4db591600480830192600092919082900301818387803b15801561044757600080fd5b505af115801561045b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561048457600080fd5b81019080805164010000000081111561049c57600080fd5b820160208101848111156104af57600080fd5b81518560208202830111640100000000821117156104cc57600080fd5b509094505050505090565b600054600160a060020a031633146104ee57600080fd5b600254604051600160a060020a038084169216907f4fea88aaf04c303804bb211ecc32a00ac8e5f0656bb854cad8a4a2e438256b7490600090a36002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6002546040517f5518edf80000000000000000000000000000000000000000000000000000000081523360048201818152600160a060020a0388811660248501526044840188905260806064850190815260848501879052941693635518edf8938992899289928992919060a40184848082843782019150509650505050505050600060405180830381600087803b1580156105f257600080fd5b505af1158015610606573d6000803e3d6000fd5b5050505050505050565b600154600160a060020a031681565b600254604080517ff981dfca000000000000000000000000000000000000000000000000000000008152336004820152600160a060020a038581166024830152604482018590529151919092169163f981dfca91606480830192600092919082900301818387803b15801561069357600080fd5b505af11580156106a7573d6000803e3d6000fd5b5050505050505600a165627a7a72305820f155cfee0b4b51b7631ecb495eb977d3b516745e4725d816285ac67e00a044400029"

// DeployRelaySet deploys a new Ethereum contract, binding an instance of RelaySet to it.
func DeployRelaySet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RelaySet, error) {
	parsed, err := abi.JSON(strings.NewReader(RelaySetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RelaySetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RelaySet{RelaySetCaller: RelaySetCaller{contract: contract}, RelaySetTransactor: RelaySetTransactor{contract: contract}, RelaySetFilterer: RelaySetFilterer{contract: contract}}, nil
}

// RelaySet is an auto generated Go binding around an Ethereum contract.
type RelaySet struct {
	RelaySetCaller     // Read-only binding to the contract
	RelaySetTransactor // Write-only binding to the contract
	RelaySetFilterer   // Log filterer for contract events
}

// RelaySetCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelaySetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelaySetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelaySetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelaySetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RelaySetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelaySetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelaySetSession struct {
	Contract     *RelaySet         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelaySetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelaySetCallerSession struct {
	Contract *RelaySetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RelaySetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelaySetTransactorSession struct {
	Contract     *RelaySetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RelaySetRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelaySetRaw struct {
	Contract *RelaySet // Generic contract binding to access the raw methods on
}

// RelaySetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelaySetCallerRaw struct {
	Contract *RelaySetCaller // Generic read-only contract binding to access the raw methods on
}

// RelaySetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelaySetTransactorRaw struct {
	Contract *RelaySetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelaySet creates a new instance of RelaySet, bound to a specific deployed contract.
func NewRelaySet(address common.Address, backend bind.ContractBackend) (*RelaySet, error) {
	contract, err := bindRelaySet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RelaySet{RelaySetCaller: RelaySetCaller{contract: contract}, RelaySetTransactor: RelaySetTransactor{contract: contract}, RelaySetFilterer: RelaySetFilterer{contract: contract}}, nil
}

// NewRelaySetCaller creates a new read-only instance of RelaySet, bound to a specific deployed contract.
func NewRelaySetCaller(address common.Address, caller bind.ContractCaller) (*RelaySetCaller, error) {
	contract, err := bindRelaySet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RelaySetCaller{contract: contract}, nil
}

// NewRelaySetTransactor creates a new write-only instance of RelaySet, bound to a specific deployed contract.
func NewRelaySetTransactor(address common.Address, transactor bind.ContractTransactor) (*RelaySetTransactor, error) {
	contract, err := bindRelaySet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RelaySetTransactor{contract: contract}, nil
}

// NewRelaySetFilterer creates a new log filterer instance of RelaySet, bound to a specific deployed contract.
func NewRelaySetFilterer(address common.Address, filterer bind.ContractFilterer) (*RelaySetFilterer, error) {
	contract, err := bindRelaySet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RelaySetFilterer{contract: contract}, nil
}

// bindRelaySet binds a generic wrapper to an already deployed contract.
func bindRelaySet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RelaySetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelaySet *RelaySetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RelaySet.Contract.RelaySetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelaySet *RelaySetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelaySet.Contract.RelaySetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelaySet *RelaySetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelaySet.Contract.RelaySetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelaySet *RelaySetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RelaySet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelaySet *RelaySetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelaySet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelaySet *RelaySetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelaySet.Contract.contract.Transact(opts, method, params...)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_RelaySet *RelaySetCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _RelaySet.contract.Call(opts, out, "getValidators")
	return *ret0, err
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_RelaySet *RelaySetSession) GetValidators() ([]common.Address, error) {
	return _RelaySet.Contract.GetValidators(&_RelaySet.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_RelaySet *RelaySetCallerSession) GetValidators() ([]common.Address, error) {
	return _RelaySet.Contract.GetValidators(&_RelaySet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RelaySet *RelaySetCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RelaySet.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RelaySet *RelaySetSession) Owner() (common.Address, error) {
	return _RelaySet.Contract.Owner(&_RelaySet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RelaySet *RelaySetCallerSession) Owner() (common.Address, error) {
	return _RelaySet.Contract.Owner(&_RelaySet.CallOpts)
}

// RelayedSet is a free data retrieval call binding the contract method 0xae3783d6.
//
// Solidity: function relayedSet() view returns(address)
func (_RelaySet *RelaySetCaller) RelayedSet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RelaySet.contract.Call(opts, out, "relayedSet")
	return *ret0, err
}

// RelayedSet is a free data retrieval call binding the contract method 0xae3783d6.
//
// Solidity: function relayedSet() view returns(address)
func (_RelaySet *RelaySetSession) RelayedSet() (common.Address, error) {
	return _RelaySet.Contract.RelayedSet(&_RelaySet.CallOpts)
}

// RelayedSet is a free data retrieval call binding the contract method 0xae3783d6.
//
// Solidity: function relayedSet() view returns(address)
func (_RelaySet *RelaySetCallerSession) RelayedSet() (common.Address, error) {
	return _RelaySet.Contract.RelayedSet(&_RelaySet.CallOpts)
}

// SystemAddress is a free data retrieval call binding the contract method 0xd3e848f1.
//
// Solidity: function systemAddress() view returns(address)
func (_RelaySet *RelaySetCaller) SystemAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RelaySet.contract.Call(opts, out, "systemAddress")
	return *ret0, err
}

// SystemAddress is a free data retrieval call binding the contract method 0xd3e848f1.
//
// Solidity: function systemAddress() view returns(address)
func (_RelaySet *RelaySetSession) SystemAddress() (common.Address, error) {
	return _RelaySet.Contract.SystemAddress(&_RelaySet.CallOpts)
}

// SystemAddress is a free data retrieval call binding the contract method 0xd3e848f1.
//
// Solidity: function systemAddress() view returns(address)
func (_RelaySet *RelaySetCallerSession) SystemAddress() (common.Address, error) {
	return _RelaySet.Contract.SystemAddress(&_RelaySet.CallOpts)
}

// FinalizeChange is a paid mutator transaction binding the contract method 0x75286211.
//
// Solidity: function finalizeChange() returns()
func (_RelaySet *RelaySetTransactor) FinalizeChange(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelaySet.contract.Transact(opts, "finalizeChange")
}

// FinalizeChange is a paid mutator transaction binding the contract method 0x75286211.
//
// Solidity: function finalizeChange() returns()
func (_RelaySet *RelaySetSession) FinalizeChange() (*types.Transaction, error) {
	return _RelaySet.Contract.FinalizeChange(&_RelaySet.TransactOpts)
}

// FinalizeChange is a paid mutator transaction binding the contract method 0x75286211.
//
// Solidity: function finalizeChange() returns()
func (_RelaySet *RelaySetTransactorSession) FinalizeChange() (*types.Transaction, error) {
	return _RelaySet.Contract.FinalizeChange(&_RelaySet.TransactOpts)
}

// InitiateChange is a paid mutator transaction binding the contract method 0xa0285c01.
//
// Solidity: function initiateChange(bytes32 _parentHash, address[] _newSet) returns()
func (_RelaySet *RelaySetTransactor) InitiateChange(opts *bind.TransactOpts, _parentHash [32]byte, _newSet []common.Address) (*types.Transaction, error) {
	return _RelaySet.contract.Transact(opts, "initiateChange", _parentHash, _newSet)
}

// InitiateChange is a paid mutator transaction binding the contract method 0xa0285c01.
//
// Solidity: function initiateChange(bytes32 _parentHash, address[] _newSet) returns()
func (_RelaySet *RelaySetSession) InitiateChange(_parentHash [32]byte, _newSet []common.Address) (*types.Transaction, error) {
	return _RelaySet.Contract.InitiateChange(&_RelaySet.TransactOpts, _parentHash, _newSet)
}

// InitiateChange is a paid mutator transaction binding the contract method 0xa0285c01.
//
// Solidity: function initiateChange(bytes32 _parentHash, address[] _newSet) returns()
func (_RelaySet *RelaySetTransactorSession) InitiateChange(_parentHash [32]byte, _newSet []common.Address) (*types.Transaction, error) {
	return _RelaySet.Contract.InitiateChange(&_RelaySet.TransactOpts, _parentHash, _newSet)
}

// ReportBenign is a paid mutator transaction binding the contract method 0xd69f13bb.
//
// Solidity: function reportBenign(address _validator, uint256 _blockNumber) returns()
func (_RelaySet *RelaySetTransactor) ReportBenign(opts *bind.TransactOpts, _validator common.Address, _blockNumber *big.Int) (*types.Transaction, error) {
	return _RelaySet.contract.Transact(opts, "reportBenign", _validator, _blockNumber)
}

// ReportBenign is a paid mutator transaction binding the contract method 0xd69f13bb.
//
// Solidity: function reportBenign(address _validator, uint256 _blockNumber) returns()
func (_RelaySet *RelaySetSession) ReportBenign(_validator common.Address, _blockNumber *big.Int) (*types.Transaction, error) {
	return _RelaySet.Contract.ReportBenign(&_RelaySet.TransactOpts, _validator, _blockNumber)
}

// ReportBenign is a paid mutator transaction binding the contract method 0xd69f13bb.
//
// Solidity: function reportBenign(address _validator, uint256 _blockNumber) returns()
func (_RelaySet *RelaySetTransactorSession) ReportBenign(_validator common.Address, _blockNumber *big.Int) (*types.Transaction, error) {
	return _RelaySet.Contract.ReportBenign(&_RelaySet.TransactOpts, _validator, _blockNumber)
}

// ReportMalicious is a paid mutator transaction binding the contract method 0xc476dd40.
//
// Solidity: function reportMalicious(address _validator, uint256 _blockNumber, bytes _proof) returns()
func (_RelaySet *RelaySetTransactor) ReportMalicious(opts *bind.TransactOpts, _validator common.Address, _blockNumber *big.Int, _proof []byte) (*types.Transaction, error) {
	return _RelaySet.contract.Transact(opts, "reportMalicious", _validator, _blockNumber, _proof)
}

// ReportMalicious is a paid mutator transaction binding the contract method 0xc476dd40.
//
// Solidity: function reportMalicious(address _validator, uint256 _blockNumber, bytes _proof) returns()
func (_RelaySet *RelaySetSession) ReportMalicious(_validator common.Address, _blockNumber *big.Int, _proof []byte) (*types.Transaction, error) {
	return _RelaySet.Contract.ReportMalicious(&_RelaySet.TransactOpts, _validator, _blockNumber, _proof)
}

// ReportMalicious is a paid mutator transaction binding the contract method 0xc476dd40.
//
// Solidity: function reportMalicious(address _validator, uint256 _blockNumber, bytes _proof) returns()
func (_RelaySet *RelaySetTransactorSession) ReportMalicious(_validator common.Address, _blockNumber *big.Int, _proof []byte) (*types.Transaction, error) {
	return _RelaySet.Contract.ReportMalicious(&_RelaySet.TransactOpts, _validator, _blockNumber, _proof)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _new) returns()
func (_RelaySet *RelaySetTransactor) SetOwner(opts *bind.TransactOpts, _new common.Address) (*types.Transaction, error) {
	return _RelaySet.contract.Transact(opts, "setOwner", _new)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _new) returns()
func (_RelaySet *RelaySetSession) SetOwner(_new common.Address) (*types.Transaction, error) {
	return _RelaySet.Contract.SetOwner(&_RelaySet.TransactOpts, _new)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _new) returns()
func (_RelaySet *RelaySetTransactorSession) SetOwner(_new common.Address) (*types.Transaction, error) {
	return _RelaySet.Contract.SetOwner(&_RelaySet.TransactOpts, _new)
}

// SetRelayed is a paid mutator transaction binding the contract method 0xbd965677.
//
// Solidity: function setRelayed(address _relayedSet) returns()
func (_RelaySet *RelaySetTransactor) SetRelayed(opts *bind.TransactOpts, _relayedSet common.Address) (*types.Transaction, error) {
	return _RelaySet.contract.Transact(opts, "setRelayed", _relayedSet)
}

// SetRelayed is a paid mutator transaction binding the contract method 0xbd965677.
//
// Solidity: function setRelayed(address _relayedSet) returns()
func (_RelaySet *RelaySetSession) SetRelayed(_relayedSet common.Address) (*types.Transaction, error) {
	return _RelaySet.Contract.SetRelayed(&_RelaySet.TransactOpts, _relayedSet)
}

// SetRelayed is a paid mutator transaction binding the contract method 0xbd965677.
//
// Solidity: function setRelayed(address _relayedSet) returns()
func (_RelaySet *RelaySetTransactorSession) SetRelayed(_relayedSet common.Address) (*types.Transaction, error) {
	return _RelaySet.Contract.SetRelayed(&_RelaySet.TransactOpts, _relayedSet)
}

// RelaySetInitiateChangeIterator is returned from FilterInitiateChange and is used to iterate over the raw logs and unpacked data for InitiateChange events raised by the RelaySet contract.
type RelaySetInitiateChangeIterator struct {
	Event *RelaySetInitiateChange // Event containing the contract specifics and raw log

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
func (it *RelaySetInitiateChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelaySetInitiateChange)
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
		it.Event = new(RelaySetInitiateChange)
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
func (it *RelaySetInitiateChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelaySetInitiateChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelaySetInitiateChange represents a InitiateChange event raised by the RelaySet contract.
type RelaySetInitiateChange struct {
	ParentHash [32]byte
	NewSet     []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInitiateChange is a free log retrieval operation binding the contract event 0x55252fa6eee4741b4e24a74a70e9c11fd2c2281df8d6ea13126ff845f7825c89.
//
// Solidity: event InitiateChange(bytes32 indexed _parentHash, address[] _newSet)
func (_RelaySet *RelaySetFilterer) FilterInitiateChange(opts *bind.FilterOpts, _parentHash [][32]byte) (*RelaySetInitiateChangeIterator, error) {

	var _parentHashRule []interface{}
	for _, _parentHashItem := range _parentHash {
		_parentHashRule = append(_parentHashRule, _parentHashItem)
	}

	logs, sub, err := _RelaySet.contract.FilterLogs(opts, "InitiateChange", _parentHashRule)
	if err != nil {
		return nil, err
	}
	return &RelaySetInitiateChangeIterator{contract: _RelaySet.contract, event: "InitiateChange", logs: logs, sub: sub}, nil
}

// WatchInitiateChange is a free log subscription operation binding the contract event 0x55252fa6eee4741b4e24a74a70e9c11fd2c2281df8d6ea13126ff845f7825c89.
//
// Solidity: event InitiateChange(bytes32 indexed _parentHash, address[] _newSet)
func (_RelaySet *RelaySetFilterer) WatchInitiateChange(opts *bind.WatchOpts, sink chan<- *RelaySetInitiateChange, _parentHash [][32]byte) (event.Subscription, error) {

	var _parentHashRule []interface{}
	for _, _parentHashItem := range _parentHash {
		_parentHashRule = append(_parentHashRule, _parentHashItem)
	}

	logs, sub, err := _RelaySet.contract.WatchLogs(opts, "InitiateChange", _parentHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelaySetInitiateChange)
				if err := _RelaySet.contract.UnpackLog(event, "InitiateChange", log); err != nil {
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

// ParseInitiateChange is a log parse operation binding the contract event 0x55252fa6eee4741b4e24a74a70e9c11fd2c2281df8d6ea13126ff845f7825c89.
//
// Solidity: event InitiateChange(bytes32 indexed _parentHash, address[] _newSet)
func (_RelaySet *RelaySetFilterer) ParseInitiateChange(log types.Log) (*RelaySetInitiateChange, error) {
	event := new(RelaySetInitiateChange)
	if err := _RelaySet.contract.UnpackLog(event, "InitiateChange", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelaySetNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the RelaySet contract.
type RelaySetNewOwnerIterator struct {
	Event *RelaySetNewOwner // Event containing the contract specifics and raw log

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
func (it *RelaySetNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelaySetNewOwner)
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
		it.Event = new(RelaySetNewOwner)
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
func (it *RelaySetNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelaySetNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelaySetNewOwner represents a NewOwner event raised by the RelaySet contract.
type RelaySetNewOwner struct {
	Old     common.Address
	Current common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed old, address indexed current)
func (_RelaySet *RelaySetFilterer) FilterNewOwner(opts *bind.FilterOpts, old []common.Address, current []common.Address) (*RelaySetNewOwnerIterator, error) {

	var oldRule []interface{}
	for _, oldItem := range old {
		oldRule = append(oldRule, oldItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _RelaySet.contract.FilterLogs(opts, "NewOwner", oldRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &RelaySetNewOwnerIterator{contract: _RelaySet.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed old, address indexed current)
func (_RelaySet *RelaySetFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *RelaySetNewOwner, old []common.Address, current []common.Address) (event.Subscription, error) {

	var oldRule []interface{}
	for _, oldItem := range old {
		oldRule = append(oldRule, oldItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _RelaySet.contract.WatchLogs(opts, "NewOwner", oldRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelaySetNewOwner)
				if err := _RelaySet.contract.UnpackLog(event, "NewOwner", log); err != nil {
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

// ParseNewOwner is a log parse operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed old, address indexed current)
func (_RelaySet *RelaySetFilterer) ParseNewOwner(log types.Log) (*RelaySetNewOwner, error) {
	event := new(RelaySetNewOwner)
	if err := _RelaySet.contract.UnpackLog(event, "NewOwner", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelaySetNewRelayedIterator is returned from FilterNewRelayed and is used to iterate over the raw logs and unpacked data for NewRelayed events raised by the RelaySet contract.
type RelaySetNewRelayedIterator struct {
	Event *RelaySetNewRelayed // Event containing the contract specifics and raw log

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
func (it *RelaySetNewRelayedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelaySetNewRelayed)
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
		it.Event = new(RelaySetNewRelayed)
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
func (it *RelaySetNewRelayedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelaySetNewRelayedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelaySetNewRelayed represents a NewRelayed event raised by the RelaySet contract.
type RelaySetNewRelayed struct {
	Old     common.Address
	Current common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewRelayed is a free log retrieval operation binding the contract event 0x4fea88aaf04c303804bb211ecc32a00ac8e5f0656bb854cad8a4a2e438256b74.
//
// Solidity: event NewRelayed(address indexed old, address indexed current)
func (_RelaySet *RelaySetFilterer) FilterNewRelayed(opts *bind.FilterOpts, old []common.Address, current []common.Address) (*RelaySetNewRelayedIterator, error) {

	var oldRule []interface{}
	for _, oldItem := range old {
		oldRule = append(oldRule, oldItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _RelaySet.contract.FilterLogs(opts, "NewRelayed", oldRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &RelaySetNewRelayedIterator{contract: _RelaySet.contract, event: "NewRelayed", logs: logs, sub: sub}, nil
}

// WatchNewRelayed is a free log subscription operation binding the contract event 0x4fea88aaf04c303804bb211ecc32a00ac8e5f0656bb854cad8a4a2e438256b74.
//
// Solidity: event NewRelayed(address indexed old, address indexed current)
func (_RelaySet *RelaySetFilterer) WatchNewRelayed(opts *bind.WatchOpts, sink chan<- *RelaySetNewRelayed, old []common.Address, current []common.Address) (event.Subscription, error) {

	var oldRule []interface{}
	for _, oldItem := range old {
		oldRule = append(oldRule, oldItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _RelaySet.contract.WatchLogs(opts, "NewRelayed", oldRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelaySetNewRelayed)
				if err := _RelaySet.contract.UnpackLog(event, "NewRelayed", log); err != nil {
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

// ParseNewRelayed is a log parse operation binding the contract event 0x4fea88aaf04c303804bb211ecc32a00ac8e5f0656bb854cad8a4a2e438256b74.
//
// Solidity: event NewRelayed(address indexed old, address indexed current)
func (_RelaySet *RelaySetFilterer) ParseNewRelayed(log types.Log) (*RelaySetNewRelayed, error) {
	event := new(RelaySetNewRelayed)
	if err := _RelaySet.contract.UnpackLog(event, "NewRelayed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayedOwnedSetABI is the input ABI used to generate the binding from.
const RelayedOwnedSetABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"recentBlocks\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPending\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_new\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"addValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reporter\",\"type\":\"address\"},{\"name\":\"_validator\",\"type\":\"address\"},{\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"relayReportMalicious\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalizeChange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recentBlocks\",\"type\":\"uint256\"}],\"name\":\"setRecentBlocks\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"relaySet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"finalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_relaySet\",\"type\":\"address\"}],\"name\":\"setRelay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reporter\",\"type\":\"address\"},{\"name\":\"_validator\",\"type\":\"address\"},{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"relayReportBenign\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_relaySet\",\"type\":\"address\"},{\"name\":\"_initial\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"currentSet\",\"type\":\"address[]\"}],\"name\":\"ChangeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"reporter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"reported\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"malicious\",\"type\":\"bool\"}],\"name\":\"Report\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"old\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"current\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"}]"

// RelayedOwnedSetFuncSigs maps the 4-byte function signature to its string representation.
var RelayedOwnedSetFuncSigs = map[string]string{
	"4d238c8e": "addValidator(address)",
	"75286211": "finalizeChange()",
	"b3f05b97": "finalized()",
	"11ae9ed2": "getPending()",
	"b7ab4db5": "getValidators()",
	"8da5cb5b": "owner()",
	"0c857805": "recentBlocks()",
	"f981dfca": "relayReportBenign(address,address,uint256)",
	"5518edf8": "relayReportMalicious(address,address,uint256,bytes)",
	"a6940b07": "relaySet()",
	"40a141ff": "removeValidator(address)",
	"13af4035": "setOwner(address)",
	"8c90ac07": "setRecentBlocks(uint256)",
	"c805f68b": "setRelay(address)",
}

// RelayedOwnedSetBin is the compiled bytecode used for deploying new contracts.
var RelayedOwnedSetBin = "0x608060405260008054600160a060020a0319163317905560146001553480156200002857600080fd5b5060405162000fe238038062000fe2833981016040528051602080830151909201805191929091829160009162000066916003919085019062000144565b50600090505b8151811015620001065760016004600084848151811015156200008b57fe5b602090810291909101810151600160a060020a031682528101919091526040016000908120805460ff19169215159290921790915582518291600491859084908110620000d457fe5b6020908102909101810151600160a060020a03168252810191909152604001600020600190810191909155016200006c565b600380546200011891600291620001ae565b505060058054600160a060020a031916600160a060020a039490941693909317909255506200021b9050565b8280548282559060005260206000209081019282156200019c579160200282015b828111156200019c5782518254600160a060020a031916600160a060020a0390911617825560209092019160019091019062000165565b50620001aa929150620001f1565b5090565b8280548282559060005260206000209081019282156200019c5760005260206000209182015b828111156200019c578254825591600101919060010190620001d4565b6200021891905b80821115620001aa578054600160a060020a0319168155600101620001f8565b90565b610db7806200022b6000396000f3006080604052600436106100cf5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630c85780581146100d457806311ae9ed2146100fb57806313af40351461016057806340a141ff146101835780634d238c8e146101a45780635518edf8146101c557806375286211146101fe5780638c90ac07146102135780638da5cb5b1461022b578063a6940b071461025c578063b3f05b9714610271578063b7ab4db51461029a578063c805f68b146102af578063f981dfca146102d0575b600080fd5b3480156100e057600080fd5b506100e96102fa565b60408051918252519081900360200190f35b34801561010757600080fd5b50610110610300565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561014c578181015183820152602001610134565b505050509050019250505060405180910390f35b34801561016c57600080fd5b50610181600160a060020a0360043516610363565b005b34801561018f57600080fd5b50610181600160a060020a03600435166103e2565b3480156101b057600080fd5b50610181600160a060020a03600435166105b1565b3480156101d157600080fd5b50610181600160a060020a0360048035821691602480359091169160443591606435908101910135610673565b34801561020a57600080fd5b506101816106c6565b34801561021f57600080fd5b506101816004356106e7565b34801561023757600080fd5b50610240610703565b60408051600160a060020a039092168252519081900360200190f35b34801561026857600080fd5b50610240610712565b34801561027d57600080fd5b50610286610721565b604080519115158252519081900360200190f35b3480156102a657600080fd5b50610110610742565b3480156102bb57600080fd5b50610181600160a060020a03600435166107a2565b3480156102dc57600080fd5b50610181600160a060020a03600435811690602435166044356107e8565b60015481565b6060600380548060200260200160405190810160405280929190818152602001828054801561035857602002820191906000526020600020905b8154600160a060020a0316815260019091019060200180831161033a575b505050505090505b90565b600054600160a060020a0316331461037a57600080fd5b60008054604051600160a060020a03808516939216917f70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b236491a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60008054600160a060020a031633146103fa57600080fd5b600160a060020a03821660009081526004602052604090208054600190910154839160ff169081801561042e575060025481105b8015610465575082600160a060020a031660028281548110151561044e57fe5b600091825260209091200154600160a060020a0316145b151561047057600080fd5b600160a060020a038516600090815260046020526040902060010154600380549195509060001981019081106104a257fe5b60009182526020909120015460038054600160a060020a0390921691869081106104c857fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550836004600060038781548110151561050a57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190206001015560038054600019810190811061054457fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff19169055600380549061057d906000198301610ccc565b50600160a060020a0385166000908152600460205260408120805460ff19168155600101556105aa61080f565b5050505050565b600054600160a060020a031633146105c857600080fd5b600160a060020a038116600090815260046020526040902054819060ff16156105f057600080fd5b600160a060020a0382166000818152600460205260408120805460ff1916600190811782556003805492820183905590820181559091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b01805473ffffffffffffffffffffffffffffffffffffffff1916909117905561066f61080f565b5050565b600554600160a060020a0316331461068a57600080fd5b6105aa85858585858080601f0160208091040260200160405190810160405280939291908181526020018383808284375061085e945050505050565b600554600160a060020a031633146106dd57600080fd5b6106e56109b2565b565b600054600160a060020a031633146106fe57600080fd5b600155565b600054600160a060020a031681565b600554600160a060020a031681565b60005474010000000000000000000000000000000000000000900460ff1681565b6060600280548060200260200160405190810160405280929190818152602001828054801561035857602002820191906000526020600020908154600160a060020a0316815260019091019060200180831161033a575050505050905090565b600054600160a060020a031633146107b957600080fd5b6005805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600554600160a060020a031633146107ff57600080fd5b61080a838383610aa4565b505050565b60005474010000000000000000000000000000000000000000900460ff16151561083857600080fd5b6000805474ff0000000000000000000000000000000000000000191690556106e5610bf6565b600160a060020a03841660009081526004602052604090208054600190910154859160ff1690818015610892575060025481105b80156108c9575082600160a060020a03166002828154811015156108b257fe5b600091825260209091200154600160a060020a0316145b15156108d457600080fd5b600160a060020a03861660009081526004602052604090208054600190910154879160ff1690818015610908575060025481105b801561093f575082600160a060020a031660028281548110151561092857fe5b600091825260209091200154600160a060020a0316145b151561094a57600080fd5b876001548101431115801561095e57504381105b151561096957600080fd5b604051600190600160a060020a03808d1691908e16907f32c78b6140c46745a46e88cd883707d70dbd2f06d13dd76fe5f499c01290da4f90600090a45050505050505050505050565b60005474010000000000000000000000000000000000000000900460ff16156109da57600080fd5b600380546109ea91600291610cf0565b506000805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055604080516020808252600280549183018290527f8564cd629b15f47dc310d45bcbfc9bcf5420b0d51bf0659a16c67f91d276325393909291829182019084908015610a9457602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610a76575b50509250505060405180910390a1565b600160a060020a03831660009081526004602052604090208054600190910154849160ff1690818015610ad8575060025481105b8015610b0f575082600160a060020a0316600282815481101515610af857fe5b600091825260209091200154600160a060020a0316145b1515610b1a57600080fd5b600160a060020a03851660009081526004602052604090208054600190910154869160ff1690818015610b4e575060025481105b8015610b85575082600160a060020a0316600282815481101515610b6e57fe5b600091825260209091200154600160a060020a0316145b1515610b9057600080fd5b8660015481014311158015610ba457504381105b1515610baf57600080fd5b604051600090600160a060020a03808c1691908d16907f32c78b6140c46745a46e88cd883707d70dbd2f06d13dd76fe5f499c01290da4f908490a450505050505050505050565b600554604080517fa0285c0100000000000000000000000000000000000000000000000000000000815260001943014060048201818152602483019384526003805460448501819052600160a060020a039096169563a0285c019593949193916064019084908015610c9157602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610c73575b50509350505050600060405180830381600087803b158015610cb257600080fd5b505af1158015610cc6573d6000803e3d6000fd5b50505050565b81548183558181111561080a5760008381526020902061080a918101908301610d40565b828054828255906000526020600020908101928215610d305760005260206000209182015b82811115610d30578254825591600101919060010190610d15565b50610d3c929150610d5a565b5090565b61036091905b80821115610d3c5760008155600101610d46565b61036091905b80821115610d3c57805473ffffffffffffffffffffffffffffffffffffffff19168155600101610d605600a165627a7a72305820416e5ee1c2e3aac275f1c9927467be9b2d8d5931deaa3e9b541359b375ad2ad10029"

// DeployRelayedOwnedSet deploys a new Ethereum contract, binding an instance of RelayedOwnedSet to it.
func DeployRelayedOwnedSet(auth *bind.TransactOpts, backend bind.ContractBackend, _relaySet common.Address, _initial []common.Address) (common.Address, *types.Transaction, *RelayedOwnedSet, error) {
	parsed, err := abi.JSON(strings.NewReader(RelayedOwnedSetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RelayedOwnedSetBin), backend, _relaySet, _initial)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RelayedOwnedSet{RelayedOwnedSetCaller: RelayedOwnedSetCaller{contract: contract}, RelayedOwnedSetTransactor: RelayedOwnedSetTransactor{contract: contract}, RelayedOwnedSetFilterer: RelayedOwnedSetFilterer{contract: contract}}, nil
}

// RelayedOwnedSet is an auto generated Go binding around an Ethereum contract.
type RelayedOwnedSet struct {
	RelayedOwnedSetCaller     // Read-only binding to the contract
	RelayedOwnedSetTransactor // Write-only binding to the contract
	RelayedOwnedSetFilterer   // Log filterer for contract events
}

// RelayedOwnedSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelayedOwnedSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayedOwnedSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelayedOwnedSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayedOwnedSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RelayedOwnedSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayedOwnedSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelayedOwnedSetSession struct {
	Contract     *RelayedOwnedSet  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelayedOwnedSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelayedOwnedSetCallerSession struct {
	Contract *RelayedOwnedSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// RelayedOwnedSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelayedOwnedSetTransactorSession struct {
	Contract     *RelayedOwnedSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// RelayedOwnedSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelayedOwnedSetRaw struct {
	Contract *RelayedOwnedSet // Generic contract binding to access the raw methods on
}

// RelayedOwnedSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelayedOwnedSetCallerRaw struct {
	Contract *RelayedOwnedSetCaller // Generic read-only contract binding to access the raw methods on
}

// RelayedOwnedSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelayedOwnedSetTransactorRaw struct {
	Contract *RelayedOwnedSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelayedOwnedSet creates a new instance of RelayedOwnedSet, bound to a specific deployed contract.
func NewRelayedOwnedSet(address common.Address, backend bind.ContractBackend) (*RelayedOwnedSet, error) {
	contract, err := bindRelayedOwnedSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RelayedOwnedSet{RelayedOwnedSetCaller: RelayedOwnedSetCaller{contract: contract}, RelayedOwnedSetTransactor: RelayedOwnedSetTransactor{contract: contract}, RelayedOwnedSetFilterer: RelayedOwnedSetFilterer{contract: contract}}, nil
}

// NewRelayedOwnedSetCaller creates a new read-only instance of RelayedOwnedSet, bound to a specific deployed contract.
func NewRelayedOwnedSetCaller(address common.Address, caller bind.ContractCaller) (*RelayedOwnedSetCaller, error) {
	contract, err := bindRelayedOwnedSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RelayedOwnedSetCaller{contract: contract}, nil
}

// NewRelayedOwnedSetTransactor creates a new write-only instance of RelayedOwnedSet, bound to a specific deployed contract.
func NewRelayedOwnedSetTransactor(address common.Address, transactor bind.ContractTransactor) (*RelayedOwnedSetTransactor, error) {
	contract, err := bindRelayedOwnedSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RelayedOwnedSetTransactor{contract: contract}, nil
}

// NewRelayedOwnedSetFilterer creates a new log filterer instance of RelayedOwnedSet, bound to a specific deployed contract.
func NewRelayedOwnedSetFilterer(address common.Address, filterer bind.ContractFilterer) (*RelayedOwnedSetFilterer, error) {
	contract, err := bindRelayedOwnedSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RelayedOwnedSetFilterer{contract: contract}, nil
}

// bindRelayedOwnedSet binds a generic wrapper to an already deployed contract.
func bindRelayedOwnedSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RelayedOwnedSetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelayedOwnedSet *RelayedOwnedSetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RelayedOwnedSet.Contract.RelayedOwnedSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelayedOwnedSet *RelayedOwnedSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayedOwnedSet.Contract.RelayedOwnedSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelayedOwnedSet *RelayedOwnedSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelayedOwnedSet.Contract.RelayedOwnedSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelayedOwnedSet *RelayedOwnedSetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RelayedOwnedSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelayedOwnedSet *RelayedOwnedSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayedOwnedSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelayedOwnedSet *RelayedOwnedSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelayedOwnedSet.Contract.contract.Transact(opts, method, params...)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_RelayedOwnedSet *RelayedOwnedSetCaller) Finalized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RelayedOwnedSet.contract.Call(opts, out, "finalized")
	return *ret0, err
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_RelayedOwnedSet *RelayedOwnedSetSession) Finalized() (bool, error) {
	return _RelayedOwnedSet.Contract.Finalized(&_RelayedOwnedSet.CallOpts)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_RelayedOwnedSet *RelayedOwnedSetCallerSession) Finalized() (bool, error) {
	return _RelayedOwnedSet.Contract.Finalized(&_RelayedOwnedSet.CallOpts)
}

// GetPending is a free data retrieval call binding the contract method 0x11ae9ed2.
//
// Solidity: function getPending() view returns(address[])
func (_RelayedOwnedSet *RelayedOwnedSetCaller) GetPending(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _RelayedOwnedSet.contract.Call(opts, out, "getPending")
	return *ret0, err
}

// GetPending is a free data retrieval call binding the contract method 0x11ae9ed2.
//
// Solidity: function getPending() view returns(address[])
func (_RelayedOwnedSet *RelayedOwnedSetSession) GetPending() ([]common.Address, error) {
	return _RelayedOwnedSet.Contract.GetPending(&_RelayedOwnedSet.CallOpts)
}

// GetPending is a free data retrieval call binding the contract method 0x11ae9ed2.
//
// Solidity: function getPending() view returns(address[])
func (_RelayedOwnedSet *RelayedOwnedSetCallerSession) GetPending() ([]common.Address, error) {
	return _RelayedOwnedSet.Contract.GetPending(&_RelayedOwnedSet.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_RelayedOwnedSet *RelayedOwnedSetCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _RelayedOwnedSet.contract.Call(opts, out, "getValidators")
	return *ret0, err
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_RelayedOwnedSet *RelayedOwnedSetSession) GetValidators() ([]common.Address, error) {
	return _RelayedOwnedSet.Contract.GetValidators(&_RelayedOwnedSet.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_RelayedOwnedSet *RelayedOwnedSetCallerSession) GetValidators() ([]common.Address, error) {
	return _RelayedOwnedSet.Contract.GetValidators(&_RelayedOwnedSet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RelayedOwnedSet *RelayedOwnedSetCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RelayedOwnedSet.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RelayedOwnedSet *RelayedOwnedSetSession) Owner() (common.Address, error) {
	return _RelayedOwnedSet.Contract.Owner(&_RelayedOwnedSet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RelayedOwnedSet *RelayedOwnedSetCallerSession) Owner() (common.Address, error) {
	return _RelayedOwnedSet.Contract.Owner(&_RelayedOwnedSet.CallOpts)
}

// RecentBlocks is a free data retrieval call binding the contract method 0x0c857805.
//
// Solidity: function recentBlocks() view returns(uint256)
func (_RelayedOwnedSet *RelayedOwnedSetCaller) RecentBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RelayedOwnedSet.contract.Call(opts, out, "recentBlocks")
	return *ret0, err
}

// RecentBlocks is a free data retrieval call binding the contract method 0x0c857805.
//
// Solidity: function recentBlocks() view returns(uint256)
func (_RelayedOwnedSet *RelayedOwnedSetSession) RecentBlocks() (*big.Int, error) {
	return _RelayedOwnedSet.Contract.RecentBlocks(&_RelayedOwnedSet.CallOpts)
}

// RecentBlocks is a free data retrieval call binding the contract method 0x0c857805.
//
// Solidity: function recentBlocks() view returns(uint256)
func (_RelayedOwnedSet *RelayedOwnedSetCallerSession) RecentBlocks() (*big.Int, error) {
	return _RelayedOwnedSet.Contract.RecentBlocks(&_RelayedOwnedSet.CallOpts)
}

// RelaySet is a free data retrieval call binding the contract method 0xa6940b07.
//
// Solidity: function relaySet() view returns(address)
func (_RelayedOwnedSet *RelayedOwnedSetCaller) RelaySet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RelayedOwnedSet.contract.Call(opts, out, "relaySet")
	return *ret0, err
}

// RelaySet is a free data retrieval call binding the contract method 0xa6940b07.
//
// Solidity: function relaySet() view returns(address)
func (_RelayedOwnedSet *RelayedOwnedSetSession) RelaySet() (common.Address, error) {
	return _RelayedOwnedSet.Contract.RelaySet(&_RelayedOwnedSet.CallOpts)
}

// RelaySet is a free data retrieval call binding the contract method 0xa6940b07.
//
// Solidity: function relaySet() view returns(address)
func (_RelayedOwnedSet *RelayedOwnedSetCallerSession) RelaySet() (common.Address, error) {
	return _RelayedOwnedSet.Contract.RelaySet(&_RelayedOwnedSet.CallOpts)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address _validator) returns()
func (_RelayedOwnedSet *RelayedOwnedSetTransactor) AddValidator(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _RelayedOwnedSet.contract.Transact(opts, "addValidator", _validator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address _validator) returns()
func (_RelayedOwnedSet *RelayedOwnedSetSession) AddValidator(_validator common.Address) (*types.Transaction, error) {
	return _RelayedOwnedSet.Contract.AddValidator(&_RelayedOwnedSet.TransactOpts, _validator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address _validator) returns()
func (_RelayedOwnedSet *RelayedOwnedSetTransactorSession) AddValidator(_validator common.Address) (*types.Transaction, error) {
	return _RelayedOwnedSet.Contract.AddValidator(&_RelayedOwnedSet.TransactOpts, _validator)
}

// FinalizeChange is a paid mutator transaction binding the contract method 0x75286211.
//
// Solidity: function finalizeChange() returns()
func (_RelayedOwnedSet *RelayedOwnedSetTransactor) FinalizeChange(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayedOwnedSet.contract.Transact(opts, "finalizeChange")
}

// FinalizeChange is a paid mutator transaction binding the contract method 0x75286211.
//
// Solidity: function finalizeChange() returns()
func (_RelayedOwnedSet *RelayedOwnedSetSession) FinalizeChange() (*types.Transaction, error) {
	return _RelayedOwnedSet.Contract.FinalizeChange(&_RelayedOwnedSet.TransactOpts)
}

// FinalizeChange is a paid mutator transaction binding the contract method 0x75286211.
//
// Solidity: function finalizeChange() returns()
func (_RelayedOwnedSet *RelayedOwnedSetTransactorSession) FinalizeChange() (*types.Transaction, error) {
	return _RelayedOwnedSet.Contract.FinalizeChange(&_RelayedOwnedSet.TransactOpts)
}

// RelayReportBenign is a paid mutator transaction binding the contract method 0xf981dfca.
//
// Solidity: function relayReportBenign(address _reporter, address _validator, uint256 _blockNumber) returns()
func (_RelayedOwnedSet *RelayedOwnedSetTransactor) RelayReportBenign(opts *bind.TransactOpts, _reporter common.Address, _validator common.Address, _blockNumber *big.Int) (*types.Transaction, error) {
	return _RelayedOwnedSet.contract.Transact(opts, "relayReportBenign", _reporter, _validator, _blockNumber)
}

// RelayReportBenign is a paid mutator transaction binding the contract method 0xf981dfca.
//
// Solidity: function relayReportBenign(address _reporter, address _validator, uint256 _blockNumber) returns()
func (_RelayedOwnedSet *RelayedOwnedSetSession) RelayReportBenign(_reporter common.Address, _validator common.Address, _blockNumber *big.Int) (*types.Transaction, error) {
	return _RelayedOwnedSet.Contract.RelayReportBenign(&_RelayedOwnedSet.TransactOpts, _reporter, _validator, _blockNumber)
}

// RelayReportBenign is a paid mutator transaction binding the contract method 0xf981dfca.
//
// Solidity: function relayReportBenign(address _reporter, address _validator, uint256 _blockNumber) returns()
func (_RelayedOwnedSet *RelayedOwnedSetTransactorSession) RelayReportBenign(_reporter common.Address, _validator common.Address, _blockNumber *big.Int) (*types.Transaction, error) {
	return _RelayedOwnedSet.Contract.RelayReportBenign(&_RelayedOwnedSet.TransactOpts, _reporter, _validator, _blockNumber)
}

// RelayReportMalicious is a paid mutator transaction binding the contract method 0x5518edf8.
//
// Solidity: function relayReportMalicious(address _reporter, address _validator, uint256 _blockNumber, bytes _proof) returns()
func (_RelayedOwnedSet *RelayedOwnedSetTransactor) RelayReportMalicious(opts *bind.TransactOpts, _reporter common.Address, _validator common.Address, _blockNumber *big.Int, _proof []byte) (*types.Transaction, error) {
	return _RelayedOwnedSet.contract.Transact(opts, "relayReportMalicious", _reporter, _validator, _blockNumber, _proof)
}

// RelayReportMalicious is a paid mutator transaction binding the contract method 0x5518edf8.
//
// Solidity: function relayReportMalicious(address _reporter, address _validator, uint256 _blockNumber, bytes _proof) returns()
func (_RelayedOwnedSet *RelayedOwnedSetSession) RelayReportMalicious(_reporter common.Address, _validator common.Address, _blockNumber *big.Int, _proof []byte) (*types.Transaction, error) {
	return _RelayedOwnedSet.Contract.RelayReportMalicious(&_RelayedOwnedSet.TransactOpts, _reporter, _validator, _blockNumber, _proof)
}

// RelayReportMalicious is a paid mutator transaction binding the contract method 0x5518edf8.
//
// Solidity: function relayReportMalicious(address _reporter, address _validator, uint256 _blockNumber, bytes _proof) returns()
func (_RelayedOwnedSet *RelayedOwnedSetTransactorSession) RelayReportMalicious(_reporter common.Address, _validator common.Address, _blockNumber *big.Int, _proof []byte) (*types.Transaction, error) {
	return _RelayedOwnedSet.Contract.RelayReportMalicious(&_RelayedOwnedSet.TransactOpts, _reporter, _validator, _blockNumber, _proof)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validator) returns()
func (_RelayedOwnedSet *RelayedOwnedSetTransactor) RemoveValidator(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _RelayedOwnedSet.contract.Transact(opts, "removeValidator", _validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validator) returns()
func (_RelayedOwnedSet *RelayedOwnedSetSession) RemoveValidator(_validator common.Address) (*types.Transaction, error) {
	return _RelayedOwnedSet.Contract.RemoveValidator(&_RelayedOwnedSet.TransactOpts, _validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validator) returns()
func (_RelayedOwnedSet *RelayedOwnedSetTransactorSession) RemoveValidator(_validator common.Address) (*types.Transaction, error) {
	return _RelayedOwnedSet.Contract.RemoveValidator(&_RelayedOwnedSet.TransactOpts, _validator)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _new) returns()
func (_RelayedOwnedSet *RelayedOwnedSetTransactor) SetOwner(opts *bind.TransactOpts, _new common.Address) (*types.Transaction, error) {
	return _RelayedOwnedSet.contract.Transact(opts, "setOwner", _new)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _new) returns()
func (_RelayedOwnedSet *RelayedOwnedSetSession) SetOwner(_new common.Address) (*types.Transaction, error) {
	return _RelayedOwnedSet.Contract.SetOwner(&_RelayedOwnedSet.TransactOpts, _new)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _new) returns()
func (_RelayedOwnedSet *RelayedOwnedSetTransactorSession) SetOwner(_new common.Address) (*types.Transaction, error) {
	return _RelayedOwnedSet.Contract.SetOwner(&_RelayedOwnedSet.TransactOpts, _new)
}

// SetRecentBlocks is a paid mutator transaction binding the contract method 0x8c90ac07.
//
// Solidity: function setRecentBlocks(uint256 _recentBlocks) returns()
func (_RelayedOwnedSet *RelayedOwnedSetTransactor) SetRecentBlocks(opts *bind.TransactOpts, _recentBlocks *big.Int) (*types.Transaction, error) {
	return _RelayedOwnedSet.contract.Transact(opts, "setRecentBlocks", _recentBlocks)
}

// SetRecentBlocks is a paid mutator transaction binding the contract method 0x8c90ac07.
//
// Solidity: function setRecentBlocks(uint256 _recentBlocks) returns()
func (_RelayedOwnedSet *RelayedOwnedSetSession) SetRecentBlocks(_recentBlocks *big.Int) (*types.Transaction, error) {
	return _RelayedOwnedSet.Contract.SetRecentBlocks(&_RelayedOwnedSet.TransactOpts, _recentBlocks)
}

// SetRecentBlocks is a paid mutator transaction binding the contract method 0x8c90ac07.
//
// Solidity: function setRecentBlocks(uint256 _recentBlocks) returns()
func (_RelayedOwnedSet *RelayedOwnedSetTransactorSession) SetRecentBlocks(_recentBlocks *big.Int) (*types.Transaction, error) {
	return _RelayedOwnedSet.Contract.SetRecentBlocks(&_RelayedOwnedSet.TransactOpts, _recentBlocks)
}

// SetRelay is a paid mutator transaction binding the contract method 0xc805f68b.
//
// Solidity: function setRelay(address _relaySet) returns()
func (_RelayedOwnedSet *RelayedOwnedSetTransactor) SetRelay(opts *bind.TransactOpts, _relaySet common.Address) (*types.Transaction, error) {
	return _RelayedOwnedSet.contract.Transact(opts, "setRelay", _relaySet)
}

// SetRelay is a paid mutator transaction binding the contract method 0xc805f68b.
//
// Solidity: function setRelay(address _relaySet) returns()
func (_RelayedOwnedSet *RelayedOwnedSetSession) SetRelay(_relaySet common.Address) (*types.Transaction, error) {
	return _RelayedOwnedSet.Contract.SetRelay(&_RelayedOwnedSet.TransactOpts, _relaySet)
}

// SetRelay is a paid mutator transaction binding the contract method 0xc805f68b.
//
// Solidity: function setRelay(address _relaySet) returns()
func (_RelayedOwnedSet *RelayedOwnedSetTransactorSession) SetRelay(_relaySet common.Address) (*types.Transaction, error) {
	return _RelayedOwnedSet.Contract.SetRelay(&_RelayedOwnedSet.TransactOpts, _relaySet)
}

// RelayedOwnedSetChangeFinalizedIterator is returned from FilterChangeFinalized and is used to iterate over the raw logs and unpacked data for ChangeFinalized events raised by the RelayedOwnedSet contract.
type RelayedOwnedSetChangeFinalizedIterator struct {
	Event *RelayedOwnedSetChangeFinalized // Event containing the contract specifics and raw log

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
func (it *RelayedOwnedSetChangeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayedOwnedSetChangeFinalized)
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
		it.Event = new(RelayedOwnedSetChangeFinalized)
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
func (it *RelayedOwnedSetChangeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayedOwnedSetChangeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayedOwnedSetChangeFinalized represents a ChangeFinalized event raised by the RelayedOwnedSet contract.
type RelayedOwnedSetChangeFinalized struct {
	CurrentSet []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChangeFinalized is a free log retrieval operation binding the contract event 0x8564cd629b15f47dc310d45bcbfc9bcf5420b0d51bf0659a16c67f91d2763253.
//
// Solidity: event ChangeFinalized(address[] currentSet)
func (_RelayedOwnedSet *RelayedOwnedSetFilterer) FilterChangeFinalized(opts *bind.FilterOpts) (*RelayedOwnedSetChangeFinalizedIterator, error) {

	logs, sub, err := _RelayedOwnedSet.contract.FilterLogs(opts, "ChangeFinalized")
	if err != nil {
		return nil, err
	}
	return &RelayedOwnedSetChangeFinalizedIterator{contract: _RelayedOwnedSet.contract, event: "ChangeFinalized", logs: logs, sub: sub}, nil
}

// WatchChangeFinalized is a free log subscription operation binding the contract event 0x8564cd629b15f47dc310d45bcbfc9bcf5420b0d51bf0659a16c67f91d2763253.
//
// Solidity: event ChangeFinalized(address[] currentSet)
func (_RelayedOwnedSet *RelayedOwnedSetFilterer) WatchChangeFinalized(opts *bind.WatchOpts, sink chan<- *RelayedOwnedSetChangeFinalized) (event.Subscription, error) {

	logs, sub, err := _RelayedOwnedSet.contract.WatchLogs(opts, "ChangeFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayedOwnedSetChangeFinalized)
				if err := _RelayedOwnedSet.contract.UnpackLog(event, "ChangeFinalized", log); err != nil {
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

// ParseChangeFinalized is a log parse operation binding the contract event 0x8564cd629b15f47dc310d45bcbfc9bcf5420b0d51bf0659a16c67f91d2763253.
//
// Solidity: event ChangeFinalized(address[] currentSet)
func (_RelayedOwnedSet *RelayedOwnedSetFilterer) ParseChangeFinalized(log types.Log) (*RelayedOwnedSetChangeFinalized, error) {
	event := new(RelayedOwnedSetChangeFinalized)
	if err := _RelayedOwnedSet.contract.UnpackLog(event, "ChangeFinalized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayedOwnedSetNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the RelayedOwnedSet contract.
type RelayedOwnedSetNewOwnerIterator struct {
	Event *RelayedOwnedSetNewOwner // Event containing the contract specifics and raw log

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
func (it *RelayedOwnedSetNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayedOwnedSetNewOwner)
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
		it.Event = new(RelayedOwnedSetNewOwner)
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
func (it *RelayedOwnedSetNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayedOwnedSetNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayedOwnedSetNewOwner represents a NewOwner event raised by the RelayedOwnedSet contract.
type RelayedOwnedSetNewOwner struct {
	Old     common.Address
	Current common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed old, address indexed current)
func (_RelayedOwnedSet *RelayedOwnedSetFilterer) FilterNewOwner(opts *bind.FilterOpts, old []common.Address, current []common.Address) (*RelayedOwnedSetNewOwnerIterator, error) {

	var oldRule []interface{}
	for _, oldItem := range old {
		oldRule = append(oldRule, oldItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _RelayedOwnedSet.contract.FilterLogs(opts, "NewOwner", oldRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &RelayedOwnedSetNewOwnerIterator{contract: _RelayedOwnedSet.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed old, address indexed current)
func (_RelayedOwnedSet *RelayedOwnedSetFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *RelayedOwnedSetNewOwner, old []common.Address, current []common.Address) (event.Subscription, error) {

	var oldRule []interface{}
	for _, oldItem := range old {
		oldRule = append(oldRule, oldItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _RelayedOwnedSet.contract.WatchLogs(opts, "NewOwner", oldRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayedOwnedSetNewOwner)
				if err := _RelayedOwnedSet.contract.UnpackLog(event, "NewOwner", log); err != nil {
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

// ParseNewOwner is a log parse operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed old, address indexed current)
func (_RelayedOwnedSet *RelayedOwnedSetFilterer) ParseNewOwner(log types.Log) (*RelayedOwnedSetNewOwner, error) {
	event := new(RelayedOwnedSetNewOwner)
	if err := _RelayedOwnedSet.contract.UnpackLog(event, "NewOwner", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayedOwnedSetReportIterator is returned from FilterReport and is used to iterate over the raw logs and unpacked data for Report events raised by the RelayedOwnedSet contract.
type RelayedOwnedSetReportIterator struct {
	Event *RelayedOwnedSetReport // Event containing the contract specifics and raw log

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
func (it *RelayedOwnedSetReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayedOwnedSetReport)
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
		it.Event = new(RelayedOwnedSetReport)
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
func (it *RelayedOwnedSetReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayedOwnedSetReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayedOwnedSetReport represents a Report event raised by the RelayedOwnedSet contract.
type RelayedOwnedSetReport struct {
	Reporter  common.Address
	Reported  common.Address
	Malicious bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReport is a free log retrieval operation binding the contract event 0x32c78b6140c46745a46e88cd883707d70dbd2f06d13dd76fe5f499c01290da4f.
//
// Solidity: event Report(address indexed reporter, address indexed reported, bool indexed malicious)
func (_RelayedOwnedSet *RelayedOwnedSetFilterer) FilterReport(opts *bind.FilterOpts, reporter []common.Address, reported []common.Address, malicious []bool) (*RelayedOwnedSetReportIterator, error) {

	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}
	var reportedRule []interface{}
	for _, reportedItem := range reported {
		reportedRule = append(reportedRule, reportedItem)
	}
	var maliciousRule []interface{}
	for _, maliciousItem := range malicious {
		maliciousRule = append(maliciousRule, maliciousItem)
	}

	logs, sub, err := _RelayedOwnedSet.contract.FilterLogs(opts, "Report", reporterRule, reportedRule, maliciousRule)
	if err != nil {
		return nil, err
	}
	return &RelayedOwnedSetReportIterator{contract: _RelayedOwnedSet.contract, event: "Report", logs: logs, sub: sub}, nil
}

// WatchReport is a free log subscription operation binding the contract event 0x32c78b6140c46745a46e88cd883707d70dbd2f06d13dd76fe5f499c01290da4f.
//
// Solidity: event Report(address indexed reporter, address indexed reported, bool indexed malicious)
func (_RelayedOwnedSet *RelayedOwnedSetFilterer) WatchReport(opts *bind.WatchOpts, sink chan<- *RelayedOwnedSetReport, reporter []common.Address, reported []common.Address, malicious []bool) (event.Subscription, error) {

	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}
	var reportedRule []interface{}
	for _, reportedItem := range reported {
		reportedRule = append(reportedRule, reportedItem)
	}
	var maliciousRule []interface{}
	for _, maliciousItem := range malicious {
		maliciousRule = append(maliciousRule, maliciousItem)
	}

	logs, sub, err := _RelayedOwnedSet.contract.WatchLogs(opts, "Report", reporterRule, reportedRule, maliciousRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayedOwnedSetReport)
				if err := _RelayedOwnedSet.contract.UnpackLog(event, "Report", log); err != nil {
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

// ParseReport is a log parse operation binding the contract event 0x32c78b6140c46745a46e88cd883707d70dbd2f06d13dd76fe5f499c01290da4f.
//
// Solidity: event Report(address indexed reporter, address indexed reported, bool indexed malicious)
func (_RelayedOwnedSet *RelayedOwnedSetFilterer) ParseReport(log types.Log) (*RelayedOwnedSetReport, error) {
	event := new(RelayedOwnedSetReport)
	if err := _RelayedOwnedSet.contract.UnpackLog(event, "Report", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorSetABI is the input ABI used to generate the binding from.
const ValidatorSetABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"finalizeChange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"validator\",\"type\":\"address\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"reportMalicious\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"validator\",\"type\":\"address\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"reportBenign\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_parentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_newSet\",\"type\":\"address[]\"}],\"name\":\"InitiateChange\",\"type\":\"event\"}]"

// ValidatorSetFuncSigs maps the 4-byte function signature to its string representation.
var ValidatorSetFuncSigs = map[string]string{
	"75286211": "finalizeChange()",
	"b7ab4db5": "getValidators()",
	"d69f13bb": "reportBenign(address,uint256)",
	"c476dd40": "reportMalicious(address,uint256,bytes)",
}

// ValidatorSet is an auto generated Go binding around an Ethereum contract.
type ValidatorSet struct {
	ValidatorSetCaller     // Read-only binding to the contract
	ValidatorSetTransactor // Write-only binding to the contract
	ValidatorSetFilterer   // Log filterer for contract events
}

// ValidatorSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorSetSession struct {
	Contract     *ValidatorSet     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorSetCallerSession struct {
	Contract *ValidatorSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ValidatorSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorSetTransactorSession struct {
	Contract     *ValidatorSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ValidatorSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorSetRaw struct {
	Contract *ValidatorSet // Generic contract binding to access the raw methods on
}

// ValidatorSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorSetCallerRaw struct {
	Contract *ValidatorSetCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorSetTransactorRaw struct {
	Contract *ValidatorSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorSet creates a new instance of ValidatorSet, bound to a specific deployed contract.
func NewValidatorSet(address common.Address, backend bind.ContractBackend) (*ValidatorSet, error) {
	contract, err := bindValidatorSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorSet{ValidatorSetCaller: ValidatorSetCaller{contract: contract}, ValidatorSetTransactor: ValidatorSetTransactor{contract: contract}, ValidatorSetFilterer: ValidatorSetFilterer{contract: contract}}, nil
}

// NewValidatorSetCaller creates a new read-only instance of ValidatorSet, bound to a specific deployed contract.
func NewValidatorSetCaller(address common.Address, caller bind.ContractCaller) (*ValidatorSetCaller, error) {
	contract, err := bindValidatorSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetCaller{contract: contract}, nil
}

// NewValidatorSetTransactor creates a new write-only instance of ValidatorSet, bound to a specific deployed contract.
func NewValidatorSetTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorSetTransactor, error) {
	contract, err := bindValidatorSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetTransactor{contract: contract}, nil
}

// NewValidatorSetFilterer creates a new log filterer instance of ValidatorSet, bound to a specific deployed contract.
func NewValidatorSetFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorSetFilterer, error) {
	contract, err := bindValidatorSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetFilterer{contract: contract}, nil
}

// bindValidatorSet binds a generic wrapper to an already deployed contract.
func bindValidatorSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorSetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorSet *ValidatorSetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ValidatorSet.Contract.ValidatorSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorSet *ValidatorSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorSet.Contract.ValidatorSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorSet *ValidatorSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorSet.Contract.ValidatorSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorSet *ValidatorSetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ValidatorSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorSet *ValidatorSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorSet *ValidatorSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorSet.Contract.contract.Transact(opts, method, params...)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_ValidatorSet *ValidatorSetCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ValidatorSet.contract.Call(opts, out, "getValidators")
	return *ret0, err
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_ValidatorSet *ValidatorSetSession) GetValidators() ([]common.Address, error) {
	return _ValidatorSet.Contract.GetValidators(&_ValidatorSet.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_ValidatorSet *ValidatorSetCallerSession) GetValidators() ([]common.Address, error) {
	return _ValidatorSet.Contract.GetValidators(&_ValidatorSet.CallOpts)
}

// FinalizeChange is a paid mutator transaction binding the contract method 0x75286211.
//
// Solidity: function finalizeChange() returns()
func (_ValidatorSet *ValidatorSetTransactor) FinalizeChange(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorSet.contract.Transact(opts, "finalizeChange")
}

// FinalizeChange is a paid mutator transaction binding the contract method 0x75286211.
//
// Solidity: function finalizeChange() returns()
func (_ValidatorSet *ValidatorSetSession) FinalizeChange() (*types.Transaction, error) {
	return _ValidatorSet.Contract.FinalizeChange(&_ValidatorSet.TransactOpts)
}

// FinalizeChange is a paid mutator transaction binding the contract method 0x75286211.
//
// Solidity: function finalizeChange() returns()
func (_ValidatorSet *ValidatorSetTransactorSession) FinalizeChange() (*types.Transaction, error) {
	return _ValidatorSet.Contract.FinalizeChange(&_ValidatorSet.TransactOpts)
}

// ReportBenign is a paid mutator transaction binding the contract method 0xd69f13bb.
//
// Solidity: function reportBenign(address validator, uint256 blockNumber) returns()
func (_ValidatorSet *ValidatorSetTransactor) ReportBenign(opts *bind.TransactOpts, validator common.Address, blockNumber *big.Int) (*types.Transaction, error) {
	return _ValidatorSet.contract.Transact(opts, "reportBenign", validator, blockNumber)
}

// ReportBenign is a paid mutator transaction binding the contract method 0xd69f13bb.
//
// Solidity: function reportBenign(address validator, uint256 blockNumber) returns()
func (_ValidatorSet *ValidatorSetSession) ReportBenign(validator common.Address, blockNumber *big.Int) (*types.Transaction, error) {
	return _ValidatorSet.Contract.ReportBenign(&_ValidatorSet.TransactOpts, validator, blockNumber)
}

// ReportBenign is a paid mutator transaction binding the contract method 0xd69f13bb.
//
// Solidity: function reportBenign(address validator, uint256 blockNumber) returns()
func (_ValidatorSet *ValidatorSetTransactorSession) ReportBenign(validator common.Address, blockNumber *big.Int) (*types.Transaction, error) {
	return _ValidatorSet.Contract.ReportBenign(&_ValidatorSet.TransactOpts, validator, blockNumber)
}

// ReportMalicious is a paid mutator transaction binding the contract method 0xc476dd40.
//
// Solidity: function reportMalicious(address validator, uint256 blockNumber, bytes proof) returns()
func (_ValidatorSet *ValidatorSetTransactor) ReportMalicious(opts *bind.TransactOpts, validator common.Address, blockNumber *big.Int, proof []byte) (*types.Transaction, error) {
	return _ValidatorSet.contract.Transact(opts, "reportMalicious", validator, blockNumber, proof)
}

// ReportMalicious is a paid mutator transaction binding the contract method 0xc476dd40.
//
// Solidity: function reportMalicious(address validator, uint256 blockNumber, bytes proof) returns()
func (_ValidatorSet *ValidatorSetSession) ReportMalicious(validator common.Address, blockNumber *big.Int, proof []byte) (*types.Transaction, error) {
	return _ValidatorSet.Contract.ReportMalicious(&_ValidatorSet.TransactOpts, validator, blockNumber, proof)
}

// ReportMalicious is a paid mutator transaction binding the contract method 0xc476dd40.
//
// Solidity: function reportMalicious(address validator, uint256 blockNumber, bytes proof) returns()
func (_ValidatorSet *ValidatorSetTransactorSession) ReportMalicious(validator common.Address, blockNumber *big.Int, proof []byte) (*types.Transaction, error) {
	return _ValidatorSet.Contract.ReportMalicious(&_ValidatorSet.TransactOpts, validator, blockNumber, proof)
}

// ValidatorSetInitiateChangeIterator is returned from FilterInitiateChange and is used to iterate over the raw logs and unpacked data for InitiateChange events raised by the ValidatorSet contract.
type ValidatorSetInitiateChangeIterator struct {
	Event *ValidatorSetInitiateChange // Event containing the contract specifics and raw log

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
func (it *ValidatorSetInitiateChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorSetInitiateChange)
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
		it.Event = new(ValidatorSetInitiateChange)
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
func (it *ValidatorSetInitiateChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorSetInitiateChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorSetInitiateChange represents a InitiateChange event raised by the ValidatorSet contract.
type ValidatorSetInitiateChange struct {
	ParentHash [32]byte
	NewSet     []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInitiateChange is a free log retrieval operation binding the contract event 0x55252fa6eee4741b4e24a74a70e9c11fd2c2281df8d6ea13126ff845f7825c89.
//
// Solidity: event InitiateChange(bytes32 indexed _parentHash, address[] _newSet)
func (_ValidatorSet *ValidatorSetFilterer) FilterInitiateChange(opts *bind.FilterOpts, _parentHash [][32]byte) (*ValidatorSetInitiateChangeIterator, error) {

	var _parentHashRule []interface{}
	for _, _parentHashItem := range _parentHash {
		_parentHashRule = append(_parentHashRule, _parentHashItem)
	}

	logs, sub, err := _ValidatorSet.contract.FilterLogs(opts, "InitiateChange", _parentHashRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetInitiateChangeIterator{contract: _ValidatorSet.contract, event: "InitiateChange", logs: logs, sub: sub}, nil
}

// WatchInitiateChange is a free log subscription operation binding the contract event 0x55252fa6eee4741b4e24a74a70e9c11fd2c2281df8d6ea13126ff845f7825c89.
//
// Solidity: event InitiateChange(bytes32 indexed _parentHash, address[] _newSet)
func (_ValidatorSet *ValidatorSetFilterer) WatchInitiateChange(opts *bind.WatchOpts, sink chan<- *ValidatorSetInitiateChange, _parentHash [][32]byte) (event.Subscription, error) {

	var _parentHashRule []interface{}
	for _, _parentHashItem := range _parentHash {
		_parentHashRule = append(_parentHashRule, _parentHashItem)
	}

	logs, sub, err := _ValidatorSet.contract.WatchLogs(opts, "InitiateChange", _parentHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorSetInitiateChange)
				if err := _ValidatorSet.contract.UnpackLog(event, "InitiateChange", log); err != nil {
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

// ParseInitiateChange is a log parse operation binding the contract event 0x55252fa6eee4741b4e24a74a70e9c11fd2c2281df8d6ea13126ff845f7825c89.
//
// Solidity: event InitiateChange(bytes32 indexed _parentHash, address[] _newSet)
func (_ValidatorSet *ValidatorSetFilterer) ParseInitiateChange(log types.Log) (*ValidatorSetInitiateChange, error) {
	event := new(ValidatorSetInitiateChange)
	if err := _ValidatorSet.contract.UnpackLog(event, "InitiateChange", log); err != nil {
		return nil, err
	}
	return event, nil
}
