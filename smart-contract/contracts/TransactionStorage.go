// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TransactionStorage

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
	_ = abi.ConvertType
)

// TransactionStorageTransaction is an auto generated low-level Go binding around an user-defined struct.
type TransactionStorageTransaction struct {
	Id        *big.Int
	From      string
	To        string
	Amount    *big.Int
	Timestamp *big.Int
}

// TransactionStorageMetaData contains all meta data concerning the TransactionStorage contract.
var TransactionStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TransactionAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_from\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"addTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getTransactionDetail\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionStorage.Transaction\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransactions\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionStorage.Transaction[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506040518060400160405280600181526020017f31000000000000000000000000000000000000000000000000000000000000008152505f90816100539190610293565b50610362565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806100d457607f821691505b6020821081036100e7576100e6610090565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026101497fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261010e565b610153868361010e565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61019761019261018d8461016b565b610174565b61016b565b9050919050565b5f819050919050565b6101b08361017d565b6101c46101bc8261019e565b84845461011a565b825550505050565b5f90565b6101d86101cc565b6101e38184846101a7565b505050565b5b81811015610206576101fb5f826101d0565b6001810190506101e9565b5050565b601f82111561024b5761021c816100ed565b610225846100ff565b81016020851015610234578190505b610248610240856100ff565b8301826101e8565b50505b505050565b5f82821c905092915050565b5f61026b5f1984600802610250565b1980831691505092915050565b5f610283838361025c565b9150826002028217905092915050565b61029c82610059565b67ffffffffffffffff8111156102b5576102b4610063565b5b6102bf82546100bd565b6102ca82828561020a565b5f60209050601f8311600181146102fb575f84156102e9578287015190505b6102f38582610278565b86555061035a565b601f198416610309866100ed565b5f5b828110156103305784890151825560018201915060208501945060208101905061030b565b8683101561034d5784890151610349601f89168261025c565b8355505b6001600288020188555050505b505050505050565b6110008061036f5f395ff3fe608060405234801561000f575f80fd5b5060043610610055575f3560e01c80630d8e6e2c1461005957806354fd4d501461007757806383920e90146100955780639ff512fa146100b3578063a76f4937146100e3575b5f80fd5b6100616100ff565b60405161006e9190610701565b60405180910390f35b61007f61018e565b60405161008c9190610701565b60405180910390f35b61009d610219565b6040516100aa91906108b6565b60405180910390f35b6100cd60048036038101906100c89190610911565b6103b2565b6040516100da91906109b6565b60405180910390f35b6100fd60048036038101906100f89190610b02565b61056d565b005b60605f805461010d90610bb7565b80601f016020809104026020016040519081016040528092919081815260200182805461013990610bb7565b80156101845780601f1061015b57610100808354040283529160200191610184565b820191905f5260205f20905b81548152906001019060200180831161016757829003601f168201915b5050505050905090565b5f805461019a90610bb7565b80601f01602080910402602001604051908101604052809291908181526020018280546101c690610bb7565b80156102115780601f106101e857610100808354040283529160200191610211565b820191905f5260205f20905b8154815290600101906020018083116101f457829003601f168201915b505050505081565b60606001805480602002602001604051908101604052809291908181526020015f905b828210156103a9578382905f5260205f2090600502016040518060a00160405290815f820154815260200160018201805461027690610bb7565b80601f01602080910402602001604051908101604052809291908181526020018280546102a290610bb7565b80156102ed5780601f106102c4576101008083540402835291602001916102ed565b820191905f5260205f20905b8154815290600101906020018083116102d057829003601f168201915b5050505050815260200160028201805461030690610bb7565b80601f016020809104026020016040519081016040528092919081815260200182805461033290610bb7565b801561037d5780601f106103545761010080835404028352916020019161037d565b820191905f5260205f20905b81548152906001019060200180831161036057829003601f168201915b50505050508152602001600382015481526020016004820154815250508152602001906001019061023c565b50505050905090565b6103ba610665565b60025482106103fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103f590610c31565b60405180910390fd5b6001828154811061041257610411610c4f565b5b905f5260205f2090600502016040518060a00160405290815f820154815260200160018201805461044290610bb7565b80601f016020809104026020016040519081016040528092919081815260200182805461046e90610bb7565b80156104b95780601f10610490576101008083540402835291602001916104b9565b820191905f5260205f20905b81548152906001019060200180831161049c57829003601f168201915b505050505081526020016002820180546104d290610bb7565b80601f01602080910402602001604051908101604052809291908181526020018280546104fe90610bb7565b80156105495780601f1061052057610100808354040283529160200191610549565b820191905f5260205f20905b81548152906001019060200180831161052c57829003601f168201915b50505050508152602001600382015481526020016004820154815250509050919050565b60016040518060a00160405280600254815260200185815260200184815260200183815260200142815250908060018154018082558091505060019003905f5260205f2090600502015f909190919091505f820151815f015560208201518160010190816105db9190610e19565b5060408201518160020190816105f19190610e19565b50606082015181600301556080820151816004015550507ffeed241432bfaeea3502dce6d5d036e4c6449696af15966aebd7fe8df15ee29260025484848442604051610641959493929190610ef7565b60405180910390a160025f81548092919061065b90610f83565b9190505550505050565b6040518060a001604052805f815260200160608152602001606081526020015f81526020015f81525090565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6106d382610691565b6106dd818561069b565b93506106ed8185602086016106ab565b6106f6816106b9565b840191505092915050565b5f6020820190508181035f83015261071981846106c9565b905092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f819050919050565b61075c8161074a565b82525050565b5f82825260208201905092915050565b5f61077c82610691565b6107868185610762565b93506107968185602086016106ab565b61079f816106b9565b840191505092915050565b5f60a083015f8301516107bf5f860182610753565b50602083015184820360208601526107d78282610772565b915050604083015184820360408601526107f18282610772565b91505060608301516108066060860182610753565b5060808301516108196080860182610753565b508091505092915050565b5f61082f83836107aa565b905092915050565b5f602082019050919050565b5f61084d82610721565b610857818561072b565b9350836020820285016108698561073b565b805f5b858110156108a457848403895281516108858582610824565b945061089083610837565b925060208a0199505060018101905061086c565b50829750879550505050505092915050565b5f6020820190508181035f8301526108ce8184610843565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b6108f08161074a565b81146108fa575f80fd5b50565b5f8135905061090b816108e7565b92915050565b5f60208284031215610926576109256108df565b5b5f610933848285016108fd565b91505092915050565b5f60a083015f8301516109515f860182610753565b50602083015184820360208601526109698282610772565b915050604083015184820360408601526109838282610772565b91505060608301516109986060860182610753565b5060808301516109ab6080860182610753565b508091505092915050565b5f6020820190508181035f8301526109ce818461093c565b905092915050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610a14826106b9565b810181811067ffffffffffffffff82111715610a3357610a326109de565b5b80604052505050565b5f610a456108d6565b9050610a518282610a0b565b919050565b5f67ffffffffffffffff821115610a7057610a6f6109de565b5b610a79826106b9565b9050602081019050919050565b828183375f83830152505050565b5f610aa6610aa184610a56565b610a3c565b905082815260208101848484011115610ac257610ac16109da565b5b610acd848285610a86565b509392505050565b5f82601f830112610ae957610ae86109d6565b5b8135610af9848260208601610a94565b91505092915050565b5f805f60608486031215610b1957610b186108df565b5b5f84013567ffffffffffffffff811115610b3657610b356108e3565b5b610b4286828701610ad5565b935050602084013567ffffffffffffffff811115610b6357610b626108e3565b5b610b6f86828701610ad5565b9250506040610b80868287016108fd565b9150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610bce57607f821691505b602082108103610be157610be0610b8a565b5b50919050565b7f5472616e73616374696f6e20494420646f6573206e6f742065786973742e00005f82015250565b5f610c1b601e8361069b565b9150610c2682610be7565b602082019050919050565b5f6020820190508181035f830152610c4881610c0f565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610cd87fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610c9d565b610ce28683610c9d565b95508019841693508086168417925050509392505050565b5f819050919050565b5f610d1d610d18610d138461074a565b610cfa565b61074a565b9050919050565b5f819050919050565b610d3683610d03565b610d4a610d4282610d24565b848454610ca9565b825550505050565b5f90565b610d5e610d52565b610d69818484610d2d565b505050565b5b81811015610d8c57610d815f82610d56565b600181019050610d6f565b5050565b601f821115610dd157610da281610c7c565b610dab84610c8e565b81016020851015610dba578190505b610dce610dc685610c8e565b830182610d6e565b50505b505050565b5f82821c905092915050565b5f610df15f1984600802610dd6565b1980831691505092915050565b5f610e098383610de2565b9150826002028217905092915050565b610e2282610691565b67ffffffffffffffff811115610e3b57610e3a6109de565b5b610e458254610bb7565b610e50828285610d90565b5f60209050601f831160018114610e81575f8415610e6f578287015190505b610e798582610dfe565b865550610ee0565b601f198416610e8f86610c7c565b5f5b82811015610eb657848901518255600182019150602085019450602081019050610e91565b86831015610ed35784890151610ecf601f891682610de2565b8355505b6001600288020188555050505b505050505050565b610ef18161074a565b82525050565b5f60a082019050610f0a5f830188610ee8565b8181036020830152610f1c81876106c9565b90508181036040830152610f3081866106c9565b9050610f3f6060830185610ee8565b610f4c6080830184610ee8565b9695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610f8d8261074a565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610fbf57610fbe610f56565b5b60018201905091905056fea26469706673582212200713367e6a9bc7510b8a8ac57b2c663f682e4e33642c9c39b9bd2021b011b5ba64736f6c63430008190033",
}

// TransactionStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use TransactionStorageMetaData.ABI instead.
var TransactionStorageABI = TransactionStorageMetaData.ABI

// TransactionStorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TransactionStorageMetaData.Bin instead.
var TransactionStorageBin = TransactionStorageMetaData.Bin

// DeployTransactionStorage deploys a new Ethereum contract, binding an instance of TransactionStorage to it.
func DeployTransactionStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TransactionStorage, error) {
	parsed, err := TransactionStorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TransactionStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TransactionStorage{TransactionStorageCaller: TransactionStorageCaller{contract: contract}, TransactionStorageTransactor: TransactionStorageTransactor{contract: contract}, TransactionStorageFilterer: TransactionStorageFilterer{contract: contract}}, nil
}

// TransactionStorage is an auto generated Go binding around an Ethereum contract.
type TransactionStorage struct {
	TransactionStorageCaller     // Read-only binding to the contract
	TransactionStorageTransactor // Write-only binding to the contract
	TransactionStorageFilterer   // Log filterer for contract events
}

// TransactionStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransactionStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransactionStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransactionStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransactionStorageSession struct {
	Contract     *TransactionStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TransactionStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransactionStorageCallerSession struct {
	Contract *TransactionStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TransactionStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransactionStorageTransactorSession struct {
	Contract     *TransactionStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TransactionStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransactionStorageRaw struct {
	Contract *TransactionStorage // Generic contract binding to access the raw methods on
}

// TransactionStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransactionStorageCallerRaw struct {
	Contract *TransactionStorageCaller // Generic read-only contract binding to access the raw methods on
}

// TransactionStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransactionStorageTransactorRaw struct {
	Contract *TransactionStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransactionStorage creates a new instance of TransactionStorage, bound to a specific deployed contract.
func NewTransactionStorage(address common.Address, backend bind.ContractBackend) (*TransactionStorage, error) {
	contract, err := bindTransactionStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransactionStorage{TransactionStorageCaller: TransactionStorageCaller{contract: contract}, TransactionStorageTransactor: TransactionStorageTransactor{contract: contract}, TransactionStorageFilterer: TransactionStorageFilterer{contract: contract}}, nil
}

// NewTransactionStorageCaller creates a new read-only instance of TransactionStorage, bound to a specific deployed contract.
func NewTransactionStorageCaller(address common.Address, caller bind.ContractCaller) (*TransactionStorageCaller, error) {
	contract, err := bindTransactionStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransactionStorageCaller{contract: contract}, nil
}

// NewTransactionStorageTransactor creates a new write-only instance of TransactionStorage, bound to a specific deployed contract.
func NewTransactionStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*TransactionStorageTransactor, error) {
	contract, err := bindTransactionStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransactionStorageTransactor{contract: contract}, nil
}

// NewTransactionStorageFilterer creates a new log filterer instance of TransactionStorage, bound to a specific deployed contract.
func NewTransactionStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*TransactionStorageFilterer, error) {
	contract, err := bindTransactionStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransactionStorageFilterer{contract: contract}, nil
}

// bindTransactionStorage binds a generic wrapper to an already deployed contract.
func bindTransactionStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TransactionStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransactionStorage *TransactionStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransactionStorage.Contract.TransactionStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransactionStorage *TransactionStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransactionStorage.Contract.TransactionStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransactionStorage *TransactionStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransactionStorage.Contract.TransactionStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransactionStorage *TransactionStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransactionStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransactionStorage *TransactionStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransactionStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransactionStorage *TransactionStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransactionStorage.Contract.contract.Transact(opts, method, params...)
}

// GetTransactionDetail is a free data retrieval call binding the contract method 0x9ff512fa.
//
// Solidity: function getTransactionDetail(uint256 _id) view returns((uint256,string,string,uint256,uint256))
func (_TransactionStorage *TransactionStorageCaller) GetTransactionDetail(opts *bind.CallOpts, _id *big.Int) (TransactionStorageTransaction, error) {
	var out []interface{}
	err := _TransactionStorage.contract.Call(opts, &out, "getTransactionDetail", _id)

	if err != nil {
		return *new(TransactionStorageTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new(TransactionStorageTransaction)).(*TransactionStorageTransaction)

	return out0, err

}

// GetTransactionDetail is a free data retrieval call binding the contract method 0x9ff512fa.
//
// Solidity: function getTransactionDetail(uint256 _id) view returns((uint256,string,string,uint256,uint256))
func (_TransactionStorage *TransactionStorageSession) GetTransactionDetail(_id *big.Int) (TransactionStorageTransaction, error) {
	return _TransactionStorage.Contract.GetTransactionDetail(&_TransactionStorage.CallOpts, _id)
}

// GetTransactionDetail is a free data retrieval call binding the contract method 0x9ff512fa.
//
// Solidity: function getTransactionDetail(uint256 _id) view returns((uint256,string,string,uint256,uint256))
func (_TransactionStorage *TransactionStorageCallerSession) GetTransactionDetail(_id *big.Int) (TransactionStorageTransaction, error) {
	return _TransactionStorage.Contract.GetTransactionDetail(&_TransactionStorage.CallOpts, _id)
}

// GetTransactions is a free data retrieval call binding the contract method 0x83920e90.
//
// Solidity: function getTransactions() view returns((uint256,string,string,uint256,uint256)[])
func (_TransactionStorage *TransactionStorageCaller) GetTransactions(opts *bind.CallOpts) ([]TransactionStorageTransaction, error) {
	var out []interface{}
	err := _TransactionStorage.contract.Call(opts, &out, "getTransactions")

	if err != nil {
		return *new([]TransactionStorageTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new([]TransactionStorageTransaction)).(*[]TransactionStorageTransaction)

	return out0, err

}

// GetTransactions is a free data retrieval call binding the contract method 0x83920e90.
//
// Solidity: function getTransactions() view returns((uint256,string,string,uint256,uint256)[])
func (_TransactionStorage *TransactionStorageSession) GetTransactions() ([]TransactionStorageTransaction, error) {
	return _TransactionStorage.Contract.GetTransactions(&_TransactionStorage.CallOpts)
}

// GetTransactions is a free data retrieval call binding the contract method 0x83920e90.
//
// Solidity: function getTransactions() view returns((uint256,string,string,uint256,uint256)[])
func (_TransactionStorage *TransactionStorageCallerSession) GetTransactions() ([]TransactionStorageTransaction, error) {
	return _TransactionStorage.Contract.GetTransactions(&_TransactionStorage.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string)
func (_TransactionStorage *TransactionStorageCaller) GetVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TransactionStorage.contract.Call(opts, &out, "getVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string)
func (_TransactionStorage *TransactionStorageSession) GetVersion() (string, error) {
	return _TransactionStorage.Contract.GetVersion(&_TransactionStorage.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string)
func (_TransactionStorage *TransactionStorageCallerSession) GetVersion() (string, error) {
	return _TransactionStorage.Contract.GetVersion(&_TransactionStorage.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_TransactionStorage *TransactionStorageCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TransactionStorage.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_TransactionStorage *TransactionStorageSession) Version() (string, error) {
	return _TransactionStorage.Contract.Version(&_TransactionStorage.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_TransactionStorage *TransactionStorageCallerSession) Version() (string, error) {
	return _TransactionStorage.Contract.Version(&_TransactionStorage.CallOpts)
}

// AddTransaction is a paid mutator transaction binding the contract method 0xa76f4937.
//
// Solidity: function addTransaction(string _from, string _to, uint256 _amount) returns()
func (_TransactionStorage *TransactionStorageTransactor) AddTransaction(opts *bind.TransactOpts, _from string, _to string, _amount *big.Int) (*types.Transaction, error) {
	return _TransactionStorage.contract.Transact(opts, "addTransaction", _from, _to, _amount)
}

// AddTransaction is a paid mutator transaction binding the contract method 0xa76f4937.
//
// Solidity: function addTransaction(string _from, string _to, uint256 _amount) returns()
func (_TransactionStorage *TransactionStorageSession) AddTransaction(_from string, _to string, _amount *big.Int) (*types.Transaction, error) {
	return _TransactionStorage.Contract.AddTransaction(&_TransactionStorage.TransactOpts, _from, _to, _amount)
}

// AddTransaction is a paid mutator transaction binding the contract method 0xa76f4937.
//
// Solidity: function addTransaction(string _from, string _to, uint256 _amount) returns()
func (_TransactionStorage *TransactionStorageTransactorSession) AddTransaction(_from string, _to string, _amount *big.Int) (*types.Transaction, error) {
	return _TransactionStorage.Contract.AddTransaction(&_TransactionStorage.TransactOpts, _from, _to, _amount)
}

// TransactionStorageTransactionAddedIterator is returned from FilterTransactionAdded and is used to iterate over the raw logs and unpacked data for TransactionAdded events raised by the TransactionStorage contract.
type TransactionStorageTransactionAddedIterator struct {
	Event *TransactionStorageTransactionAdded // Event containing the contract specifics and raw log

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
func (it *TransactionStorageTransactionAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransactionStorageTransactionAdded)
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
		it.Event = new(TransactionStorageTransactionAdded)
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
func (it *TransactionStorageTransactionAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransactionStorageTransactionAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransactionStorageTransactionAdded represents a TransactionAdded event raised by the TransactionStorage contract.
type TransactionStorageTransactionAdded struct {
	Id        *big.Int
	From      string
	To        string
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransactionAdded is a free log retrieval operation binding the contract event 0xfeed241432bfaeea3502dce6d5d036e4c6449696af15966aebd7fe8df15ee292.
//
// Solidity: event TransactionAdded(uint256 id, string from, string to, uint256 amount, uint256 timestamp)
func (_TransactionStorage *TransactionStorageFilterer) FilterTransactionAdded(opts *bind.FilterOpts) (*TransactionStorageTransactionAddedIterator, error) {

	logs, sub, err := _TransactionStorage.contract.FilterLogs(opts, "TransactionAdded")
	if err != nil {
		return nil, err
	}
	return &TransactionStorageTransactionAddedIterator{contract: _TransactionStorage.contract, event: "TransactionAdded", logs: logs, sub: sub}, nil
}

// WatchTransactionAdded is a free log subscription operation binding the contract event 0xfeed241432bfaeea3502dce6d5d036e4c6449696af15966aebd7fe8df15ee292.
//
// Solidity: event TransactionAdded(uint256 id, string from, string to, uint256 amount, uint256 timestamp)
func (_TransactionStorage *TransactionStorageFilterer) WatchTransactionAdded(opts *bind.WatchOpts, sink chan<- *TransactionStorageTransactionAdded) (event.Subscription, error) {

	logs, sub, err := _TransactionStorage.contract.WatchLogs(opts, "TransactionAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransactionStorageTransactionAdded)
				if err := _TransactionStorage.contract.UnpackLog(event, "TransactionAdded", log); err != nil {
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

// ParseTransactionAdded is a log parse operation binding the contract event 0xfeed241432bfaeea3502dce6d5d036e4c6449696af15966aebd7fe8df15ee292.
//
// Solidity: event TransactionAdded(uint256 id, string from, string to, uint256 amount, uint256 timestamp)
func (_TransactionStorage *TransactionStorageFilterer) ParseTransactionAdded(log types.Log) (*TransactionStorageTransactionAdded, error) {
	event := new(TransactionStorageTransactionAdded)
	if err := _TransactionStorage.contract.UnpackLog(event, "TransactionAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
