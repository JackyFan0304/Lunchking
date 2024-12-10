// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package voting

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

// VotingCandidate is an auto generated low-level Go binding around an user-defined struct.
type VotingCandidate struct {
	Id        *big.Int
	Name      string
	VoteCount *big.Int
}

// VotingMetaData contains all meta data concerning the Voting contract.
var VotingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"CandidateAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"candidateId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"addCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"candidates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"candidatesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllCandidates\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"}],\"internalType\":\"structVoting.Candidate[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_candidateId\",\"type\":\"uint256\"}],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_candidateId\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506111bb8061005b5f395ff3fe608060405234801561000f575f5ffd5b5060043610610086575f3560e01c8063462e91ec11610059578063462e91ec146101145780638da5cb5b14610130578063a3ec138d1461014e578063ff9810991461017e57610086565b80630121b93f1461008a5780632d35a8a2146100a65780632e6997fe146100c45780633477ee2e146100e2575b5f5ffd5b6100a4600480360381019061009f9190610761565b6101ae565b005b6100ae610342565b6040516100bb919061079b565b60405180910390f35b6100cc610348565b6040516100d9919061093b565b60405180910390f35b6100fc60048036038101906100f79190610761565b6104b6565b60405161010b939291906109a3565b60405180910390f35b61012e60048036038101906101299190610b0b565b610562565b005b6101386106a0565b6040516101459190610b91565b60405180910390f35b61016860048036038101906101639190610bd4565b6106c4565b6040516101759190610c19565b60405180910390f35b61019860048036038101906101939190610761565b6106e1565b6040516101a5919061079b565b60405180910390f35b60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1615610238576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161022f90610c7c565b60405180910390fd5b5f8111801561024957506003548111155b610288576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161027f90610ce4565b60405180910390fd5b600160025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555060015f8281526020019081526020015f206002015f81548092919061030190610d2f565b91905055507f030b0f8dcd86a031eddb071f91882edeac8173663ba775713b677b42b51be44b8133604051610337929190610d76565b60405180910390a150565b60035481565b60605f60035467ffffffffffffffff811115610367576103666109e7565b5b6040519080825280602002602001820160405280156103a057816020015b61038d6106fe565b8152602001906001900390816103855790505b5090505f600190505b60035481116104ae5760015f8281526020019081526020015f206040518060600160405290815f82015481526020016001820180546103e790610dca565b80601f016020809104026020016040519081016040528092919081815260200182805461041390610dca565b801561045e5780601f106104355761010080835404028352916020019161045e565b820191905f5260205f20905b81548152906001019060200180831161044157829003601f168201915b505050505081526020016002820154815250508260018361047f9190610dfa565b815181106104905761048f610e2d565b5b602002602001018190525080806104a690610d2f565b9150506103a9565b508091505090565b6001602052805f5260405f205f91509050805f0154908060010180546104db90610dca565b80601f016020809104026020016040519081016040528092919081815260200182805461050790610dca565b80156105525780601f1061052957610100808354040283529160200191610552565b820191905f5260205f20905b81548152906001019060200180831161053557829003601f168201915b5050505050908060020154905083565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105e790610eca565b60405180910390fd5b60035f81548092919061060290610d2f565b9190505550604051806060016040528060035481526020018281526020015f81525060015f60035481526020019081526020015f205f820151815f015560208201518160010190816106549190611088565b50604082015181600201559050507fe83b2a43e7e82d975c8a0a6d2f045153c869e111136a34d1889ab7b598e396a360035482604051610695929190611157565b60405180910390a150565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6002602052805f5260405f205f915054906101000a900460ff1681565b5f60015f8381526020019081526020015f20600201549050919050565b60405180606001604052805f8152602001606081526020015f81525090565b5f604051905090565b5f5ffd5b5f5ffd5b5f819050919050565b6107408161072e565b811461074a575f5ffd5b50565b5f8135905061075b81610737565b92915050565b5f6020828403121561077657610775610726565b5b5f6107838482850161074d565b91505092915050565b6107958161072e565b82525050565b5f6020820190506107ae5f83018461078c565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6107e68161072e565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61082e826107ec565b61083881856107f6565b9350610848818560208601610806565b61085181610814565b840191505092915050565b5f606083015f8301516108715f8601826107dd565b50602083015184820360208601526108898282610824565b915050604083015161089e60408601826107dd565b508091505092915050565b5f6108b4838361085c565b905092915050565b5f602082019050919050565b5f6108d2826107b4565b6108dc81856107be565b9350836020820285016108ee856107ce565b805f5b85811015610929578484038952815161090a85826108a9565b9450610915836108bc565b925060208a019950506001810190506108f1565b50829750879550505050505092915050565b5f6020820190508181035f83015261095381846108c8565b905092915050565b5f82825260208201905092915050565b5f610975826107ec565b61097f818561095b565b935061098f818560208601610806565b61099881610814565b840191505092915050565b5f6060820190506109b65f83018661078c565b81810360208301526109c8818561096b565b90506109d7604083018461078c565b949350505050565b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610a1d82610814565b810181811067ffffffffffffffff82111715610a3c57610a3b6109e7565b5b80604052505050565b5f610a4e61071d565b9050610a5a8282610a14565b919050565b5f67ffffffffffffffff821115610a7957610a786109e7565b5b610a8282610814565b9050602081019050919050565b828183375f83830152505050565b5f610aaf610aaa84610a5f565b610a45565b905082815260208101848484011115610acb57610aca6109e3565b5b610ad6848285610a8f565b509392505050565b5f82601f830112610af257610af16109df565b5b8135610b02848260208601610a9d565b91505092915050565b5f60208284031215610b2057610b1f610726565b5b5f82013567ffffffffffffffff811115610b3d57610b3c61072a565b5b610b4984828501610ade565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610b7b82610b52565b9050919050565b610b8b81610b71565b82525050565b5f602082019050610ba45f830184610b82565b92915050565b610bb381610b71565b8114610bbd575f5ffd5b50565b5f81359050610bce81610baa565b92915050565b5f60208284031215610be957610be8610726565b5b5f610bf684828501610bc0565b91505092915050565b5f8115159050919050565b610c1381610bff565b82525050565b5f602082019050610c2c5f830184610c0a565b92915050565b7f596f75206861766520616c726561647920766f7465642e0000000000000000005f82015250565b5f610c6660178361095b565b9150610c7182610c32565b602082019050919050565b5f6020820190508181035f830152610c9381610c5a565b9050919050565b7f496e76616c69642063616e6469646174652e00000000000000000000000000005f82015250565b5f610cce60128361095b565b9150610cd982610c9a565b602082019050919050565b5f6020820190508181035f830152610cfb81610cc2565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610d398261072e565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610d6b57610d6a610d02565b5b600182019050919050565b5f604082019050610d895f83018561078c565b610d966020830184610b82565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610de157607f821691505b602082108103610df457610df3610d9d565b5b50919050565b5f610e048261072e565b9150610e0f8361072e565b9250828203905081811115610e2757610e26610d02565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4f6e6c7920746865206f776e65722063616e20706572666f726d2074686973205f8201527f616374696f6e2e00000000000000000000000000000000000000000000000000602082015250565b5f610eb460278361095b565b9150610ebf82610e5a565b604082019050919050565b5f6020820190508181035f830152610ee181610ea8565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610f447fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610f09565b610f4e8683610f09565b95508019841693508086168417925050509392505050565b5f819050919050565b5f610f89610f84610f7f8461072e565b610f66565b61072e565b9050919050565b5f819050919050565b610fa283610f6f565b610fb6610fae82610f90565b848454610f15565b825550505050565b5f5f905090565b610fcd610fbe565b610fd8818484610f99565b505050565b5b81811015610ffb57610ff05f82610fc5565b600181019050610fde565b5050565b601f8211156110405761101181610ee8565b61101a84610efa565b81016020851015611029578190505b61103d61103585610efa565b830182610fdd565b50505b505050565b5f82821c905092915050565b5f6110605f1984600802611045565b1980831691505092915050565b5f6110788383611051565b9150826002028217905092915050565b611091826107ec565b67ffffffffffffffff8111156110aa576110a96109e7565b5b6110b48254610dca565b6110bf828285610fff565b5f60209050601f8311600181146110f0575f84156110de578287015190505b6110e8858261106d565b86555061114f565b601f1984166110fe86610ee8565b5f5b8281101561112557848901518255600182019150602085019450602081019050611100565b86831015611142578489015161113e601f891682611051565b8355505b6001600288020188555050505b505050505050565b5f60408201905061116a5f83018561078c565b818103602083015261117c818461096b565b9050939250505056fea26469706673582212202e939714eaf55b6adfa86e7c87a3c9e37119f2271577483d5abec79600f6a6d264736f6c634300081c0033",
}

// VotingABI is the input ABI used to generate the binding from.
// Deprecated: Use VotingMetaData.ABI instead.
var VotingABI = VotingMetaData.ABI

// VotingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VotingMetaData.Bin instead.
var VotingBin = VotingMetaData.Bin

// DeployVoting deploys a new Ethereum contract, binding an instance of Voting to it.
func DeployVoting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Voting, error) {
	parsed, err := VotingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VotingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Voting{VotingCaller: VotingCaller{contract: contract}, VotingTransactor: VotingTransactor{contract: contract}, VotingFilterer: VotingFilterer{contract: contract}}, nil
}

// Voting is an auto generated Go binding around an Ethereum contract.
type Voting struct {
	VotingCaller     // Read-only binding to the contract
	VotingTransactor // Write-only binding to the contract
	VotingFilterer   // Log filterer for contract events
}

// VotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotingSession struct {
	Contract     *Voting           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotingCallerSession struct {
	Contract *VotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotingTransactorSession struct {
	Contract     *VotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotingRaw struct {
	Contract *Voting // Generic contract binding to access the raw methods on
}

// VotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotingCallerRaw struct {
	Contract *VotingCaller // Generic read-only contract binding to access the raw methods on
}

// VotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotingTransactorRaw struct {
	Contract *VotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoting creates a new instance of Voting, bound to a specific deployed contract.
func NewVoting(address common.Address, backend bind.ContractBackend) (*Voting, error) {
	contract, err := bindVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Voting{VotingCaller: VotingCaller{contract: contract}, VotingTransactor: VotingTransactor{contract: contract}, VotingFilterer: VotingFilterer{contract: contract}}, nil
}

// NewVotingCaller creates a new read-only instance of Voting, bound to a specific deployed contract.
func NewVotingCaller(address common.Address, caller bind.ContractCaller) (*VotingCaller, error) {
	contract, err := bindVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotingCaller{contract: contract}, nil
}

// NewVotingTransactor creates a new write-only instance of Voting, bound to a specific deployed contract.
func NewVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*VotingTransactor, error) {
	contract, err := bindVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotingTransactor{contract: contract}, nil
}

// NewVotingFilterer creates a new log filterer instance of Voting, bound to a specific deployed contract.
func NewVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*VotingFilterer, error) {
	contract, err := bindVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotingFilterer{contract: contract}, nil
}

// bindVoting binds a generic wrapper to an already deployed contract.
func bindVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VotingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.VotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transact(opts, method, params...)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(uint256 id, string name, uint256 voteCount)
func (_Voting *VotingCaller) Candidates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id        *big.Int
	Name      string
	VoteCount *big.Int
}, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "candidates", arg0)

	outstruct := new(struct {
		Id        *big.Int
		Name      string
		VoteCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.VoteCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(uint256 id, string name, uint256 voteCount)
func (_Voting *VotingSession) Candidates(arg0 *big.Int) (struct {
	Id        *big.Int
	Name      string
	VoteCount *big.Int
}, error) {
	return _Voting.Contract.Candidates(&_Voting.CallOpts, arg0)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(uint256 id, string name, uint256 voteCount)
func (_Voting *VotingCallerSession) Candidates(arg0 *big.Int) (struct {
	Id        *big.Int
	Name      string
	VoteCount *big.Int
}, error) {
	return _Voting.Contract.Candidates(&_Voting.CallOpts, arg0)
}

// CandidatesCount is a free data retrieval call binding the contract method 0x2d35a8a2.
//
// Solidity: function candidatesCount() view returns(uint256)
func (_Voting *VotingCaller) CandidatesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "candidatesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CandidatesCount is a free data retrieval call binding the contract method 0x2d35a8a2.
//
// Solidity: function candidatesCount() view returns(uint256)
func (_Voting *VotingSession) CandidatesCount() (*big.Int, error) {
	return _Voting.Contract.CandidatesCount(&_Voting.CallOpts)
}

// CandidatesCount is a free data retrieval call binding the contract method 0x2d35a8a2.
//
// Solidity: function candidatesCount() view returns(uint256)
func (_Voting *VotingCallerSession) CandidatesCount() (*big.Int, error) {
	return _Voting.Contract.CandidatesCount(&_Voting.CallOpts)
}

// GetAllCandidates is a free data retrieval call binding the contract method 0x2e6997fe.
//
// Solidity: function getAllCandidates() view returns((uint256,string,uint256)[])
func (_Voting *VotingCaller) GetAllCandidates(opts *bind.CallOpts) ([]VotingCandidate, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getAllCandidates")

	if err != nil {
		return *new([]VotingCandidate), err
	}

	out0 := *abi.ConvertType(out[0], new([]VotingCandidate)).(*[]VotingCandidate)

	return out0, err

}

// GetAllCandidates is a free data retrieval call binding the contract method 0x2e6997fe.
//
// Solidity: function getAllCandidates() view returns((uint256,string,uint256)[])
func (_Voting *VotingSession) GetAllCandidates() ([]VotingCandidate, error) {
	return _Voting.Contract.GetAllCandidates(&_Voting.CallOpts)
}

// GetAllCandidates is a free data retrieval call binding the contract method 0x2e6997fe.
//
// Solidity: function getAllCandidates() view returns((uint256,string,uint256)[])
func (_Voting *VotingCallerSession) GetAllCandidates() ([]VotingCandidate, error) {
	return _Voting.Contract.GetAllCandidates(&_Voting.CallOpts)
}

// GetVotes is a free data retrieval call binding the contract method 0xff981099.
//
// Solidity: function getVotes(uint256 _candidateId) view returns(uint256)
func (_Voting *VotingCaller) GetVotes(opts *bind.CallOpts, _candidateId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getVotes", _candidateId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotes is a free data retrieval call binding the contract method 0xff981099.
//
// Solidity: function getVotes(uint256 _candidateId) view returns(uint256)
func (_Voting *VotingSession) GetVotes(_candidateId *big.Int) (*big.Int, error) {
	return _Voting.Contract.GetVotes(&_Voting.CallOpts, _candidateId)
}

// GetVotes is a free data retrieval call binding the contract method 0xff981099.
//
// Solidity: function getVotes(uint256 _candidateId) view returns(uint256)
func (_Voting *VotingCallerSession) GetVotes(_candidateId *big.Int) (*big.Int, error) {
	return _Voting.Contract.GetVotes(&_Voting.CallOpts, _candidateId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Voting *VotingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Voting *VotingSession) Owner() (common.Address, error) {
	return _Voting.Contract.Owner(&_Voting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Voting *VotingCallerSession) Owner() (common.Address, error) {
	return _Voting.Contract.Owner(&_Voting.CallOpts)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(bool)
func (_Voting *VotingCaller) Voters(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "voters", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(bool)
func (_Voting *VotingSession) Voters(arg0 common.Address) (bool, error) {
	return _Voting.Contract.Voters(&_Voting.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(bool)
func (_Voting *VotingCallerSession) Voters(arg0 common.Address) (bool, error) {
	return _Voting.Contract.Voters(&_Voting.CallOpts, arg0)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x462e91ec.
//
// Solidity: function addCandidate(string _name) returns()
func (_Voting *VotingTransactor) AddCandidate(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "addCandidate", _name)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x462e91ec.
//
// Solidity: function addCandidate(string _name) returns()
func (_Voting *VotingSession) AddCandidate(_name string) (*types.Transaction, error) {
	return _Voting.Contract.AddCandidate(&_Voting.TransactOpts, _name)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x462e91ec.
//
// Solidity: function addCandidate(string _name) returns()
func (_Voting *VotingTransactorSession) AddCandidate(_name string) (*types.Transaction, error) {
	return _Voting.Contract.AddCandidate(&_Voting.TransactOpts, _name)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _candidateId) returns()
func (_Voting *VotingTransactor) Vote(opts *bind.TransactOpts, _candidateId *big.Int) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "vote", _candidateId)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _candidateId) returns()
func (_Voting *VotingSession) Vote(_candidateId *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.Vote(&_Voting.TransactOpts, _candidateId)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _candidateId) returns()
func (_Voting *VotingTransactorSession) Vote(_candidateId *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.Vote(&_Voting.TransactOpts, _candidateId)
}

// VotingCandidateAddedIterator is returned from FilterCandidateAdded and is used to iterate over the raw logs and unpacked data for CandidateAdded events raised by the Voting contract.
type VotingCandidateAddedIterator struct {
	Event *VotingCandidateAdded // Event containing the contract specifics and raw log

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
func (it *VotingCandidateAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingCandidateAdded)
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
		it.Event = new(VotingCandidateAdded)
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
func (it *VotingCandidateAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingCandidateAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingCandidateAdded represents a CandidateAdded event raised by the Voting contract.
type VotingCandidateAdded struct {
	Id   *big.Int
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCandidateAdded is a free log retrieval operation binding the contract event 0xe83b2a43e7e82d975c8a0a6d2f045153c869e111136a34d1889ab7b598e396a3.
//
// Solidity: event CandidateAdded(uint256 id, string name)
func (_Voting *VotingFilterer) FilterCandidateAdded(opts *bind.FilterOpts) (*VotingCandidateAddedIterator, error) {

	logs, sub, err := _Voting.contract.FilterLogs(opts, "CandidateAdded")
	if err != nil {
		return nil, err
	}
	return &VotingCandidateAddedIterator{contract: _Voting.contract, event: "CandidateAdded", logs: logs, sub: sub}, nil
}

// WatchCandidateAdded is a free log subscription operation binding the contract event 0xe83b2a43e7e82d975c8a0a6d2f045153c869e111136a34d1889ab7b598e396a3.
//
// Solidity: event CandidateAdded(uint256 id, string name)
func (_Voting *VotingFilterer) WatchCandidateAdded(opts *bind.WatchOpts, sink chan<- *VotingCandidateAdded) (event.Subscription, error) {

	logs, sub, err := _Voting.contract.WatchLogs(opts, "CandidateAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingCandidateAdded)
				if err := _Voting.contract.UnpackLog(event, "CandidateAdded", log); err != nil {
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

// ParseCandidateAdded is a log parse operation binding the contract event 0xe83b2a43e7e82d975c8a0a6d2f045153c869e111136a34d1889ab7b598e396a3.
//
// Solidity: event CandidateAdded(uint256 id, string name)
func (_Voting *VotingFilterer) ParseCandidateAdded(log types.Log) (*VotingCandidateAdded, error) {
	event := new(VotingCandidateAdded)
	if err := _Voting.contract.UnpackLog(event, "CandidateAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the Voting contract.
type VotingVotedIterator struct {
	Event *VotingVoted // Event containing the contract specifics and raw log

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
func (it *VotingVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingVoted)
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
		it.Event = new(VotingVoted)
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
func (it *VotingVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingVoted represents a Voted event raised by the Voting contract.
type VotingVoted struct {
	CandidateId *big.Int
	Voter       common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x030b0f8dcd86a031eddb071f91882edeac8173663ba775713b677b42b51be44b.
//
// Solidity: event Voted(uint256 candidateId, address voter)
func (_Voting *VotingFilterer) FilterVoted(opts *bind.FilterOpts) (*VotingVotedIterator, error) {

	logs, sub, err := _Voting.contract.FilterLogs(opts, "Voted")
	if err != nil {
		return nil, err
	}
	return &VotingVotedIterator{contract: _Voting.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x030b0f8dcd86a031eddb071f91882edeac8173663ba775713b677b42b51be44b.
//
// Solidity: event Voted(uint256 candidateId, address voter)
func (_Voting *VotingFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *VotingVoted) (event.Subscription, error) {

	logs, sub, err := _Voting.contract.WatchLogs(opts, "Voted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingVoted)
				if err := _Voting.contract.UnpackLog(event, "Voted", log); err != nil {
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

// ParseVoted is a log parse operation binding the contract event 0x030b0f8dcd86a031eddb071f91882edeac8173663ba775713b677b42b51be44b.
//
// Solidity: event Voted(uint256 candidateId, address voter)
func (_Voting *VotingFilterer) ParseVoted(log types.Log) (*VotingVoted, error) {
	event := new(VotingVoted)
	if err := _Voting.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
