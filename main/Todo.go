// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// TodoTask is an auto generated low-level Go binding around an user-defined struct.
type TodoTask struct {
	Data      string
	Completed bool
}

// TodoMetaData contains all meta data concerning the Todo contract.
var TodoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"addTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_id\",\"type\":\"uint32\"}],\"name\":\"getTask\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"}],\"internalType\":\"structTodo.Task\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTaskList\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"}],\"internalType\":\"structTodo.Task[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_id\",\"type\":\"uint32\"}],\"name\":\"removeTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_id\",\"type\":\"uint32\"}],\"name\":\"toggleCompleted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_id\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"updateTaskData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506200003c60405180606001604052806021815260200162000f76602191396200004260201b60201c565b620001e3565b6000604051806040016040528083815260200160001515815250905060008190806001815401808255809150506001900390600052602060002090600202016000909190919091506000820151816000019080519060200190620000a8929190620000cf565b5060208201518160010160006101000a81548160ff02191690831515021790555050505050565b828054620000dd90620001ae565b90600052602060002090601f0160209004810192826200010157600085556200014d565b82601f106200011c57805160ff19168380011785556200014d565b828001600101855582156200014d579182015b828111156200014c5782518255916020019190600101906200012f565b5b5090506200015c919062000160565b5090565b5b808211156200017b57600081600090555060010162000161565b5090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620001c757607f821691505b602082108103620001dd57620001dc6200017f565b5b50919050565b610d8380620001f36000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806338940fae146100675780633ad68c761461008557806357342d74146100a157806367238562146100d1578063714ff71f146100ed57806378c01ecc14610109575b600080fd5b61006f610125565b60405161007c91906108e4565b60405180910390f35b61009f600480360381019061009a9190610956565b610231565b005b6100bb60048036038101906100b69190610956565b6102b0565b6040516100c891906109c0565b60405180910390f35b6100eb60048036038101906100e69190610b17565b61039f565b005b61010760048036038101906101029190610b60565b61042a565b005b610123600480360381019061011e9190610956565b61046e565b005b60606000805480602002602001604051908101604052809291908181526020016000905b82821015610228578382906000526020600020906002020160405180604001604052908160008201805461017c90610beb565b80601f01602080910402602001604051908101604052809291908181526020018280546101a890610beb565b80156101f55780601f106101ca576101008083540402835291602001916101f5565b820191906000526020600020905b8154815290600101906020018083116101d857829003601f168201915b505050505081526020016001820160009054906101000a900460ff16151515158152505081526020019060010190610149565b50505050905090565b60008163ffffffff168154811061024b5761024a610c1c565b5b906000526020600020906002020160010160009054906101000a900460ff161560008263ffffffff168154811061028557610284610c1c565b5b906000526020600020906002020160010160006101000a81548160ff02191690831515021790555050565b6102b86105a5565b60008263ffffffff16815481106102d2576102d1610c1c565b5b90600052602060002090600202016040518060400160405290816000820180546102fb90610beb565b80601f016020809104026020016040519081016040528092919081815260200182805461032790610beb565b80156103745780601f1061034957610100808354040283529160200191610374565b820191906000526020600020905b81548152906001019060200180831161035757829003601f168201915b505050505081526020016001820160009054906101000a900460ff1615151515815250509050919050565b60006040518060400160405280838152602001600015158152509050600081908060018154018082558091505060019003906000526020600020906002020160009091909190915060008201518160000190805190602001906104039291906105c1565b5060208201518160010160006101000a81548160ff02191690831515021790555050505050565b8060008363ffffffff168154811061044557610444610c1c565b5b906000526020600020906002020160000190805190602001906104699291906105c1565b505050565b60008190505b60016000805490506104869190610c84565b8163ffffffff1610156105525760006001826104a29190610cb8565b63ffffffff16815481106104b9576104b8610c1c565b5b906000526020600020906002020160008263ffffffff16815481106104e1576104e0610c1c565b5b9060005260206000209060020201600082018160000190805461050390610beb565b61050e929190610647565b506001820160009054906101000a900460ff168160010160006101000a81548160ff021916908315150217905550905050808061054a90610cf2565b915050610474565b50600080548061056557610564610d1e565b5b60019003818190600052602060002090600202016000808201600061058a91906106d4565b6001820160006101000a81549060ff02191690555050905550565b6040518060400160405280606081526020016000151581525090565b8280546105cd90610beb565b90600052602060002090601f0160209004810192826105ef5760008555610636565b82601f1061060857805160ff1916838001178555610636565b82800160010185558215610636579182015b8281111561063557825182559160200191906001019061061a565b5b5090506106439190610714565b5090565b82805461065390610beb565b90600052602060002090601f01602090048101928261067557600085556106c3565b82601f1061068657805485556106c3565b828001600101855582156106c357600052602060002091601f016020900482015b828111156106c25782548255916001019190600101906106a7565b5b5090506106d09190610714565b5090565b5080546106e090610beb565b6000825580601f106106f25750610711565b601f0160209004906000526020600020908101906107109190610714565b5b50565b5b8082111561072d576000816000905550600101610715565b5090565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561079757808201518184015260208101905061077c565b838111156107a6576000848401525b50505050565b6000601f19601f8301169050919050565b60006107c88261075d565b6107d28185610768565b93506107e2818560208601610779565b6107eb816107ac565b840191505092915050565b60008115159050919050565b61080b816107f6565b82525050565b6000604083016000830151848203600086015261082e82826107bd565b91505060208301516108436020860182610802565b508091505092915050565b600061085a8383610811565b905092915050565b6000602082019050919050565b600061087a82610731565b610884818561073c565b9350836020820285016108968561074d565b8060005b858110156108d257848403895281516108b3858261084e565b94506108be83610862565b925060208a0199505060018101905061089a565b50829750879550505050505092915050565b600060208201905081810360008301526108fe818461086f565b905092915050565b6000604051905090565b600080fd5b600080fd5b600063ffffffff82169050919050565b6109338161091a565b811461093e57600080fd5b50565b6000813590506109508161092a565b92915050565b60006020828403121561096c5761096b610910565b5b600061097a84828501610941565b91505092915050565b600060408301600083015184820360008601526109a082826107bd565b91505060208301516109b56020860182610802565b508091505092915050565b600060208201905081810360008301526109da8184610983565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610a24826107ac565b810181811067ffffffffffffffff82111715610a4357610a426109ec565b5b80604052505050565b6000610a56610906565b9050610a628282610a1b565b919050565b600067ffffffffffffffff821115610a8257610a816109ec565b5b610a8b826107ac565b9050602081019050919050565b82818337600083830152505050565b6000610aba610ab584610a67565b610a4c565b905082815260208101848484011115610ad657610ad56109e7565b5b610ae1848285610a98565b509392505050565b600082601f830112610afe57610afd6109e2565b5b8135610b0e848260208601610aa7565b91505092915050565b600060208284031215610b2d57610b2c610910565b5b600082013567ffffffffffffffff811115610b4b57610b4a610915565b5b610b5784828501610ae9565b91505092915050565b60008060408385031215610b7757610b76610910565b5b6000610b8585828601610941565b925050602083013567ffffffffffffffff811115610ba657610ba5610915565b5b610bb285828601610ae9565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610c0357607f821691505b602082108103610c1657610c15610bbc565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610c8f82610c4b565b9150610c9a83610c4b565b925082821015610cad57610cac610c55565b5b828203905092915050565b6000610cc38261091a565b9150610cce8361091a565b92508263ffffffff03821115610ce757610ce6610c55565b5b828201905092915050565b6000610cfd8261091a565b915063ffffffff8203610d1357610d12610c55565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea2646970667358221220a1d15b5f4860d413a3d13a5ea0d86e2efc0e38e1299322a33875b9bac2419db864736f6c634300080e0033636f6d706c6574652061737369676e6d656e742062792032326e64204d61792e21",
}

// TodoABI is the input ABI used to generate the binding from.
// Deprecated: Use TodoMetaData.ABI instead.
var TodoABI = TodoMetaData.ABI

// TodoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TodoMetaData.Bin instead.
var TodoBin = TodoMetaData.Bin

// DeployTodo deploys a new Ethereum contract, binding an instance of Todo to it.
func DeployTodo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Todo, error) {
	parsed, err := TodoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TodoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// Todo is an auto generated Go binding around an Ethereum contract.
type Todo struct {
	TodoCaller     // Read-only binding to the contract
	TodoTransactor // Write-only binding to the contract
	TodoFilterer   // Log filterer for contract events
}

// TodoCaller is an auto generated read-only Go binding around an Ethereum contract.
type TodoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TodoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TodoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TodoSession struct {
	Contract     *Todo             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TodoCallerSession struct {
	Contract *TodoCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TodoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TodoTransactorSession struct {
	Contract     *TodoTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoRaw is an auto generated low-level Go binding around an Ethereum contract.
type TodoRaw struct {
	Contract *Todo // Generic contract binding to access the raw methods on
}

// TodoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TodoCallerRaw struct {
	Contract *TodoCaller // Generic read-only contract binding to access the raw methods on
}

// TodoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TodoTransactorRaw struct {
	Contract *TodoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTodo creates a new instance of Todo, bound to a specific deployed contract.
func NewTodo(address common.Address, backend bind.ContractBackend) (*Todo, error) {
	contract, err := bindTodo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// NewTodoCaller creates a new read-only instance of Todo, bound to a specific deployed contract.
func NewTodoCaller(address common.Address, caller bind.ContractCaller) (*TodoCaller, error) {
	contract, err := bindTodo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TodoCaller{contract: contract}, nil
}

// NewTodoTransactor creates a new write-only instance of Todo, bound to a specific deployed contract.
func NewTodoTransactor(address common.Address, transactor bind.ContractTransactor) (*TodoTransactor, error) {
	contract, err := bindTodo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TodoTransactor{contract: contract}, nil
}

// NewTodoFilterer creates a new log filterer instance of Todo, bound to a specific deployed contract.
func NewTodoFilterer(address common.Address, filterer bind.ContractFilterer) (*TodoFilterer, error) {
	contract, err := bindTodo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TodoFilterer{contract: contract}, nil
}

// bindTodo binds a generic wrapper to an already deployed contract.
func bindTodo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TodoABI))
	if err != nil {
		return nil, err
	}

	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.TodoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transact(opts, method, params...)
}

// GetTask is a free data retrieval call binding the contract method 0x57342d74.
//
// Solidity: function getTask(uint32 _id) view returns((string,bool))
func (_Todo *TodoCaller) GetTask(opts *bind.CallOpts, _id uint32) (TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "getTask", _id)

	if err != nil {
		return *new(TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new(TodoTask)).(*TodoTask)

	return out0, err

}

// GetTask is a free data retrieval call binding the contract method 0x57342d74.
//
// Solidity: function getTask(uint32 _id) view returns((string,bool))
func (_Todo *TodoSession) GetTask(_id uint32) (TodoTask, error) {
	return _Todo.Contract.GetTask(&_Todo.CallOpts, _id)
}

// GetTask is a free data retrieval call binding the contract method 0x57342d74.
//
// Solidity: function getTask(uint32 _id) view returns((string,bool))
func (_Todo *TodoCallerSession) GetTask(_id uint32) (TodoTask, error) {
	return _Todo.Contract.GetTask(&_Todo.CallOpts, _id)
}

// GetTaskList is a free data retrieval call binding the contract method 0x38940fae.
//
// Solidity: function getTaskList() view returns((string,bool)[])
func (_Todo *TodoCaller) GetTaskList(opts *bind.CallOpts) ([]TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "getTaskList")

	if err != nil {
		return *new([]TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new([]TodoTask)).(*[]TodoTask)

	return out0, err

}

// GetTaskList is a free data retrieval call binding the contract method 0x38940fae.
//
// Solidity: function getTaskList() view returns((string,bool)[])
func (_Todo *TodoSession) GetTaskList() ([]TodoTask, error) {
	return _Todo.Contract.GetTaskList(&_Todo.CallOpts)
}

// GetTaskList is a free data retrieval call binding the contract method 0x38940fae.
//
// Solidity: function getTaskList() view returns((string,bool)[])
func (_Todo *TodoCallerSession) GetTaskList() ([]TodoTask, error) {
	return _Todo.Contract.GetTaskList(&_Todo.CallOpts)
}

// AddTask is a paid mutator transaction binding the contract method 0x67238562.
//
// Solidity: function addTask(string _data) returns()
func (_Todo *TodoTransactor) AddTask(opts *bind.TransactOpts, _data string) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "addTask", _data)
}

// AddTask is a paid mutator transaction binding the contract method 0x67238562.
//
// Solidity: function addTask(string _data) returns()
func (_Todo *TodoSession) AddTask(_data string) (*types.Transaction, error) {
	return _Todo.Contract.AddTask(&_Todo.TransactOpts, _data)
}

// AddTask is a paid mutator transaction binding the contract method 0x67238562.
//
// Solidity: function addTask(string _data) returns()
func (_Todo *TodoTransactorSession) AddTask(_data string) (*types.Transaction, error) {
	return _Todo.Contract.AddTask(&_Todo.TransactOpts, _data)
}

// RemoveTask is a paid mutator transaction binding the contract method 0x78c01ecc.
//
// Solidity: function removeTask(uint32 _id) returns()
func (_Todo *TodoTransactor) RemoveTask(opts *bind.TransactOpts, _id uint32) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "removeTask", _id)
}

// RemoveTask is a paid mutator transaction binding the contract method 0x78c01ecc.
//
// Solidity: function removeTask(uint32 _id) returns()
func (_Todo *TodoSession) RemoveTask(_id uint32) (*types.Transaction, error) {
	return _Todo.Contract.RemoveTask(&_Todo.TransactOpts, _id)
}

// RemoveTask is a paid mutator transaction binding the contract method 0x78c01ecc.
//
// Solidity: function removeTask(uint32 _id) returns()
func (_Todo *TodoTransactorSession) RemoveTask(_id uint32) (*types.Transaction, error) {
	return _Todo.Contract.RemoveTask(&_Todo.TransactOpts, _id)
}

// ToggleCompleted is a paid mutator transaction binding the contract method 0x3ad68c76.
//
// Solidity: function toggleCompleted(uint32 _id) returns()
func (_Todo *TodoTransactor) ToggleCompleted(opts *bind.TransactOpts, _id uint32) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "toggleCompleted", _id)
}

// ToggleCompleted is a paid mutator transaction binding the contract method 0x3ad68c76.
//
// Solidity: function toggleCompleted(uint32 _id) returns()
func (_Todo *TodoSession) ToggleCompleted(_id uint32) (*types.Transaction, error) {
	return _Todo.Contract.ToggleCompleted(&_Todo.TransactOpts, _id)
}

// ToggleCompleted is a paid mutator transaction binding the contract method 0x3ad68c76.
//
// Solidity: function toggleCompleted(uint32 _id) returns()
func (_Todo *TodoTransactorSession) ToggleCompleted(_id uint32) (*types.Transaction, error) {
	return _Todo.Contract.ToggleCompleted(&_Todo.TransactOpts, _id)
}

// UpdateTaskData is a paid mutator transaction binding the contract method 0x714ff71f.
//
// Solidity: function updateTaskData(uint32 _id, string _data) returns()
func (_Todo *TodoTransactor) UpdateTaskData(opts *bind.TransactOpts, _id uint32, _data string) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "updateTaskData", _id, _data)
}

// UpdateTaskData is a paid mutator transaction binding the contract method 0x714ff71f.
//
// Solidity: function updateTaskData(uint32 _id, string _data) returns()
func (_Todo *TodoSession) UpdateTaskData(_id uint32, _data string) (*types.Transaction, error) {
	return _Todo.Contract.UpdateTaskData(&_Todo.TransactOpts, _id, _data)
}

// UpdateTaskData is a paid mutator transaction binding the contract method 0x714ff71f.
//
// Solidity: function updateTaskData(uint32 _id, string _data) returns()
func (_Todo *TodoTransactorSession) UpdateTaskData(_id uint32, _data string) (*types.Transaction, error) {
	return _Todo.Contract.UpdateTaskData(&_Todo.TransactOpts, _id, _data)
}
