// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oracle

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

// OracleMetaData contains all meta data concerning the Oracle contract.
var OracleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousPriceMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestedPriceMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPriceMantissa\",\"type\":\"uint256\"}],\"name\":\"PricePosted\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"assetPrices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"getUnderlyingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isPriceOracle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setDirectPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"underlyingPriceMantissa\",\"type\":\"uint256\"}],\"name\":\"setUnderlyingPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061056e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806309a8acb01461005c578063127ffda01461008a5780635e9a523c146100b657806366331bba146100ee578063fc57d4df1461010a575b600080fd5b6100886004803603604081101561007257600080fd5b506001600160a01b038135169060200135610130565b005b610088600480360360408110156100a057600080fd5b506001600160a01b0381351690602001356101a8565b6100dc600480360360208110156100cc57600080fd5b50356001600160a01b031661022e565b60408051918252519081900360200190f35b6100f6610249565b604080519115158252519081900360200190f35b6100dc6004803603602081101561012057600080fd5b50356001600160a01b031661024e565b6001600160a01b038216600081815260208181526040918290205482519384529083015281810183905260608201839052517fdd71a1d19fcba687442a1d5c58578f1e409af71a79d10fd95a4d66efd8fa9ae79181900360800190a16001600160a01b03909116600090815260208190526040902055565b60006101b38361027c565b6001600160a01b038116600081815260208181526040918290205482519384529083015281810185905260608201859052519192507fdd71a1d19fcba687442a1d5c58578f1e409af71a79d10fd95a4d66efd8fa9ae7919081900360800190a16001600160a01b031660009081526020819052604090205550565b6001600160a01b031660009081526020819052604090205490565b600181565b600080600061025c8461027c565b6001600160a01b0316815260208101919091526040016000205492915050565b6000806103c4836001600160a01b03166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b1580156102bb57600080fd5b505afa1580156102cf573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156102f857600080fd5b810190808051604051939291908464010000000082111561031857600080fd5b90830190602082018581111561032d57600080fd5b825164010000000081118282018810171561034757600080fd5b82525081516020918201929091019080838360005b8381101561037457818101518382015260200161035c565b50505050905090810190601f1680156103a15780820380516001836020036101000a031916815260200191505b506040818101905260048152630c68aa8960e31b60208201529250610452915050565b156103e4575073eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee61044c565b826001600160a01b0316636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b15801561041d57600080fd5b505afa158015610431573d6000803e3d6000fd5b505050506040513d602081101561044757600080fd5b505190505b92915050565b6000816040516020018082805190602001908083835b602083106104875780518252601f199092019160209182019101610468565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405280519060200120836040516020018082805190602001908083835b602083106104f55780518252601f1990920191602091820191016104d6565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051602081830303815290604052805190602001201490509291505056fea265627a7a72315820edab8aa61fedaba5437614cbe1b7751597cd88abcbe3ab5ece689a5dac42d4e864736f6c63430005100032",
}

// OracleABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleMetaData.ABI instead.
var OracleABI = OracleMetaData.ABI

// OracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OracleMetaData.Bin instead.
var OracleBin = OracleMetaData.Bin

// DeployOracle deploys a new Ethereum contract, binding an instance of Oracle to it.
func DeployOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Oracle, error) {
	parsed, err := OracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// Oracle is an auto generated Go binding around an Ethereum contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// AssetPrices is a free data retrieval call binding the contract method 0x5e9a523c.
//
// Solidity: function assetPrices(address asset) view returns(uint256)
func (_Oracle *OracleCaller) AssetPrices(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "assetPrices", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssetPrices is a free data retrieval call binding the contract method 0x5e9a523c.
//
// Solidity: function assetPrices(address asset) view returns(uint256)
func (_Oracle *OracleSession) AssetPrices(asset common.Address) (*big.Int, error) {
	return _Oracle.Contract.AssetPrices(&_Oracle.CallOpts, asset)
}

// AssetPrices is a free data retrieval call binding the contract method 0x5e9a523c.
//
// Solidity: function assetPrices(address asset) view returns(uint256)
func (_Oracle *OracleCallerSession) AssetPrices(asset common.Address) (*big.Int, error) {
	return _Oracle.Contract.AssetPrices(&_Oracle.CallOpts, asset)
}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xfc57d4df.
//
// Solidity: function getUnderlyingPrice(address cToken) view returns(uint256)
func (_Oracle *OracleCaller) GetUnderlyingPrice(opts *bind.CallOpts, cToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getUnderlyingPrice", cToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xfc57d4df.
//
// Solidity: function getUnderlyingPrice(address cToken) view returns(uint256)
func (_Oracle *OracleSession) GetUnderlyingPrice(cToken common.Address) (*big.Int, error) {
	return _Oracle.Contract.GetUnderlyingPrice(&_Oracle.CallOpts, cToken)
}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xfc57d4df.
//
// Solidity: function getUnderlyingPrice(address cToken) view returns(uint256)
func (_Oracle *OracleCallerSession) GetUnderlyingPrice(cToken common.Address) (*big.Int, error) {
	return _Oracle.Contract.GetUnderlyingPrice(&_Oracle.CallOpts, cToken)
}

// IsPriceOracle is a free data retrieval call binding the contract method 0x66331bba.
//
// Solidity: function isPriceOracle() view returns(bool)
func (_Oracle *OracleCaller) IsPriceOracle(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "isPriceOracle")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPriceOracle is a free data retrieval call binding the contract method 0x66331bba.
//
// Solidity: function isPriceOracle() view returns(bool)
func (_Oracle *OracleSession) IsPriceOracle() (bool, error) {
	return _Oracle.Contract.IsPriceOracle(&_Oracle.CallOpts)
}

// IsPriceOracle is a free data retrieval call binding the contract method 0x66331bba.
//
// Solidity: function isPriceOracle() view returns(bool)
func (_Oracle *OracleCallerSession) IsPriceOracle() (bool, error) {
	return _Oracle.Contract.IsPriceOracle(&_Oracle.CallOpts)
}

// SetDirectPrice is a paid mutator transaction binding the contract method 0x09a8acb0.
//
// Solidity: function setDirectPrice(address asset, uint256 price) returns()
func (_Oracle *OracleTransactor) SetDirectPrice(opts *bind.TransactOpts, asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setDirectPrice", asset, price)
}

// SetDirectPrice is a paid mutator transaction binding the contract method 0x09a8acb0.
//
// Solidity: function setDirectPrice(address asset, uint256 price) returns()
func (_Oracle *OracleSession) SetDirectPrice(asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SetDirectPrice(&_Oracle.TransactOpts, asset, price)
}

// SetDirectPrice is a paid mutator transaction binding the contract method 0x09a8acb0.
//
// Solidity: function setDirectPrice(address asset, uint256 price) returns()
func (_Oracle *OracleTransactorSession) SetDirectPrice(asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SetDirectPrice(&_Oracle.TransactOpts, asset, price)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x127ffda0.
//
// Solidity: function setUnderlyingPrice(address cToken, uint256 underlyingPriceMantissa) returns()
func (_Oracle *OracleTransactor) SetUnderlyingPrice(opts *bind.TransactOpts, cToken common.Address, underlyingPriceMantissa *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setUnderlyingPrice", cToken, underlyingPriceMantissa)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x127ffda0.
//
// Solidity: function setUnderlyingPrice(address cToken, uint256 underlyingPriceMantissa) returns()
func (_Oracle *OracleSession) SetUnderlyingPrice(cToken common.Address, underlyingPriceMantissa *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SetUnderlyingPrice(&_Oracle.TransactOpts, cToken, underlyingPriceMantissa)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x127ffda0.
//
// Solidity: function setUnderlyingPrice(address cToken, uint256 underlyingPriceMantissa) returns()
func (_Oracle *OracleTransactorSession) SetUnderlyingPrice(cToken common.Address, underlyingPriceMantissa *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SetUnderlyingPrice(&_Oracle.TransactOpts, cToken, underlyingPriceMantissa)
}

// OraclePricePostedIterator is returned from FilterPricePosted and is used to iterate over the raw logs and unpacked data for PricePosted events raised by the Oracle contract.
type OraclePricePostedIterator struct {
	Event *OraclePricePosted // Event containing the contract specifics and raw log

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
func (it *OraclePricePostedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OraclePricePosted)
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
		it.Event = new(OraclePricePosted)
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
func (it *OraclePricePostedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OraclePricePostedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OraclePricePosted represents a PricePosted event raised by the Oracle contract.
type OraclePricePosted struct {
	Asset                  common.Address
	PreviousPriceMantissa  *big.Int
	RequestedPriceMantissa *big.Int
	NewPriceMantissa       *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterPricePosted is a free log retrieval operation binding the contract event 0xdd71a1d19fcba687442a1d5c58578f1e409af71a79d10fd95a4d66efd8fa9ae7.
//
// Solidity: event PricePosted(address asset, uint256 previousPriceMantissa, uint256 requestedPriceMantissa, uint256 newPriceMantissa)
func (_Oracle *OracleFilterer) FilterPricePosted(opts *bind.FilterOpts) (*OraclePricePostedIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "PricePosted")
	if err != nil {
		return nil, err
	}
	return &OraclePricePostedIterator{contract: _Oracle.contract, event: "PricePosted", logs: logs, sub: sub}, nil
}

// WatchPricePosted is a free log subscription operation binding the contract event 0xdd71a1d19fcba687442a1d5c58578f1e409af71a79d10fd95a4d66efd8fa9ae7.
//
// Solidity: event PricePosted(address asset, uint256 previousPriceMantissa, uint256 requestedPriceMantissa, uint256 newPriceMantissa)
func (_Oracle *OracleFilterer) WatchPricePosted(opts *bind.WatchOpts, sink chan<- *OraclePricePosted) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "PricePosted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OraclePricePosted)
				if err := _Oracle.contract.UnpackLog(event, "PricePosted", log); err != nil {
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

// ParsePricePosted is a log parse operation binding the contract event 0xdd71a1d19fcba687442a1d5c58578f1e409af71a79d10fd95a4d66efd8fa9ae7.
//
// Solidity: event PricePosted(address asset, uint256 previousPriceMantissa, uint256 requestedPriceMantissa, uint256 newPriceMantissa)
func (_Oracle *OracleFilterer) ParsePricePosted(log types.Log) (*OraclePricePosted, error) {
	event := new(OraclePricePosted)
	if err := _Oracle.contract.UnpackLog(event, "PricePosted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
