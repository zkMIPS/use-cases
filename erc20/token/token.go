// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

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

// PairingG1Point is an auto generated low-level Go binding around an user-defined struct.
type PairingG1Point struct {
	X *big.Int
	Y *big.Int
}

// PairingG2Point is an auto generated low-level Go binding around an user-defined struct.
type PairingG2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// Proof is an auto generated low-level Go binding around an user-defined struct.
type Proof struct {
	A PairingG1Point
	B PairingG2Point
	C PairingG1Point
}

// TestTokenMetaData contains all meta data concerning the TestToken contract.
var TestTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structPairing.G1Point\",\"name\":\"a\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structPairing.G2Point\",\"name\":\"b\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structPairing.G1Point\",\"name\":\"c\",\"type\":\"tuple\"}],\"internalType\":\"structProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"uint256[65]\",\"name\":\"input\",\"type\":\"uint256[65]\"},{\"internalType\":\"uint256[2]\",\"name\":\"proof_commitment\",\"type\":\"uint256[2]\"}],\"name\":\"mintWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506040518060400160405280600a81526020017f5465737420546f6b656e000000000000000000000000000000000000000000008152505f90816100539190610359565b506040518060400160405280600281526020017f5454000000000000000000000000000000000000000000000000000000000000815250600190816100989190610359565b50601260025f6101000a81548160ff021916908360ff1602179055506a084595161401484a00000060038190555073012ef3e31ba2664163bd039535889ae7be9e7e8660045f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610428565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061019a57607f821691505b6020821081036101ad576101ac610156565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261020f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826101d4565b61021986836101d4565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61025d61025861025384610231565b61023a565b610231565b9050919050565b5f819050919050565b61027683610243565b61028a61028282610264565b8484546101e0565b825550505050565b5f90565b61029e610292565b6102a981848461026d565b505050565b5b818110156102cc576102c15f82610296565b6001810190506102af565b5050565b601f821115610311576102e2816101b3565b6102eb846101c5565b810160208510156102fa578190505b61030e610306856101c5565b8301826102ae565b50505b505050565b5f82821c905092915050565b5f6103315f1984600802610316565b1980831691505092915050565b5f6103498383610322565b9150826002028217905092915050565b6103628261011f565b67ffffffffffffffff81111561037b5761037a610129565b5b6103858254610183565b6103908282856102d0565b5f60209050601f8311600181146103c1575f84156103af578287015190505b6103b9858261033e565b865550610420565b601f1984166103cf866101b3565b5f5b828110156103f6578489015182556001820191506020850194506020810190506103d1565b86831015610413578489015161040f601f891682610322565b8355505b6001600288020188555050505b505050505050565b6113ce806104355f395ff3fe608060405234801561000f575f80fd5b506004361061009c575f3560e01c8063313ce56711610064578063313ce5671461015857806370a082311461017657806395d89b41146101a6578063a9059cbb146101c4578063dd62ed3e146101f45761009c565b806306fdde03146100a0578063095ea7b3146100be5780630acf8131146100ee57806318160ddd1461010a57806323b872dd14610128575b5f80fd5b6100a8610224565b6040516100b591906109fe565b60405180910390f35b6100d860048036038101906100d39190610ab8565b6102b3565b6040516100e59190610b10565b60405180910390f35b61010860048036038101906101039190610e05565b6103d7565b005b610112610526565b60405161011f9190610e8e565b60405180910390f35b610142600480360381019061013d9190610ea7565b61052f565b60405161014f9190610b10565b60405180910390f35b6101606106d4565b60405161016d9190610f12565b60405180910390f35b610190600480360381019061018b9190610f2b565b6106e9565b60405161019d9190610e8e565b60405180910390f35b6101ae61072f565b6040516101bb91906109fe565b60405180910390f35b6101de60048036038101906101d99190610ab8565b6107bf565b6040516101eb9190610b10565b60405180910390f35b61020e60048036038101906102099190610f56565b61090c565b60405161021b9190610e8e565b60405180910390f35b60605f805461023290610fc1565b80601f016020809104026020016040519081016040528092919081815260200182805461025e90610fc1565b80156102a95780601f10610280576101008083540402835291602001916102a9565b820191905f5260205f20905b81548152906001019060200180831161028c57829003601f168201915b5050505050905090565b5f8073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036102eb575f80fd5b8160065f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516103c59190610e8e565b60405180910390a36001905092915050565b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bc232a008484846040518463ffffffff1660e01b815260040161043593929190611210565b6020604051808303815f875af1158015610451573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104759190611272565b6104b4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104ab906112e7565b60405180910390fd5b8360035f8282546104c59190611332565b925050819055508360055f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546105189190611332565b925050819055505050505050565b5f600354905090565b5f8160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461057c9190611365565b925050819055508160065f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461060a9190611365565b925050819055508160055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461065d9190611332565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516106c19190610e8e565b60405180910390a3600190509392505050565b5f60025f9054906101000a900460ff16905090565b5f60055f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b60606001805461073e90610fc1565b80601f016020809104026020016040519081016040528092919081815260200182805461076a90610fc1565b80156107b55780601f1061078c576101008083540402835291602001916107b5565b820191905f5260205f20905b81548152906001019060200180831161079857829003601f168201915b5050505050905090565b5f8073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036107f7575f80fd5b8160055f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546108439190611365565b925050819055508160055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546108969190611332565b925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516108fa9190610e8e565b60405180910390a36001905092915050565b5f60065f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6109d08261098e565b6109da8185610998565b93506109ea8185602086016109a8565b6109f3816109b6565b840191505092915050565b5f6020820190508181035f830152610a1681846109c6565b905092915050565b5f604051905090565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610a5482610a2b565b9050919050565b610a6481610a4a565b8114610a6e575f80fd5b50565b5f81359050610a7f81610a5b565b92915050565b5f819050919050565b610a9781610a85565b8114610aa1575f80fd5b50565b5f81359050610ab281610a8e565b92915050565b5f8060408385031215610ace57610acd610a27565b5b5f610adb85828601610a71565b9250506020610aec85828601610aa4565b9150509250929050565b5f8115159050919050565b610b0a81610af6565b82525050565b5f602082019050610b235f830184610b01565b92915050565b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610b63826109b6565b810181811067ffffffffffffffff82111715610b8257610b81610b2d565b5b80604052505050565b5f610b94610a1e565b9050610ba08282610b5a565b919050565b5f60408284031215610bba57610bb9610b29565b5b610bc46040610b8b565b90505f610bd384828501610aa4565b5f830152506020610be684828501610aa4565b60208301525092915050565b5f80fd5b5f67ffffffffffffffff821115610c1057610c0f610b2d565b5b602082029050919050565b5f80fd5b5f610c31610c2c84610bf6565b610b8b565b90508060208402830185811115610c4b57610c4a610c1b565b5b835b81811015610c745780610c608882610aa4565b845260208401935050602081019050610c4d565b5050509392505050565b5f82601f830112610c9257610c91610bf2565b5b6002610c9f848285610c1f565b91505092915050565b5f60808284031215610cbd57610cbc610b29565b5b610cc76040610b8b565b90505f610cd684828501610c7e565b5f830152506040610ce984828501610c7e565b60208301525092915050565b5f6101008284031215610d0b57610d0a610b29565b5b610d156060610b8b565b90505f610d2484828501610ba5565b5f830152506040610d3784828501610ca8565b60208301525060c0610d4b84828501610ba5565b60408301525092915050565b5f67ffffffffffffffff821115610d7157610d70610b2d565b5b602082029050919050565b5f610d8e610d8984610d57565b610b8b565b90508060208402830185811115610da857610da7610c1b565b5b835b81811015610dd15780610dbd8882610aa4565b845260208401935050602081019050610daa565b5050509392505050565b5f82601f830112610def57610dee610bf2565b5b6041610dfc848285610d7c565b91505092915050565b5f805f805f6109a08688031215610e1f57610e1e610a27565b5b5f610e2c88828901610a71565b9550506020610e3d88828901610aa4565b9450506040610e4e88828901610cf5565b935050610140610e6088828901610ddb565b925050610960610e7288828901610c7e565b9150509295509295909350565b610e8881610a85565b82525050565b5f602082019050610ea15f830184610e7f565b92915050565b5f805f60608486031215610ebe57610ebd610a27565b5b5f610ecb86828701610a71565b9350506020610edc86828701610a71565b9250506040610eed86828701610aa4565b9150509250925092565b5f60ff82169050919050565b610f0c81610ef7565b82525050565b5f602082019050610f255f830184610f03565b92915050565b5f60208284031215610f4057610f3f610a27565b5b5f610f4d84828501610a71565b91505092915050565b5f8060408385031215610f6c57610f6b610a27565b5b5f610f7985828601610a71565b9250506020610f8a85828601610a71565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610fd857607f821691505b602082108103610feb57610fea610f94565b5b50919050565b610ffa81610a85565b82525050565b604082015f8201516110145f850182610ff1565b5060208201516110276020850182610ff1565b50505050565b5f60029050919050565b5f81905092915050565b5f819050919050565b5f6110558383610ff1565b60208301905092915050565b5f602082019050919050565b6110768161102d565b6110808184611037565b925061108b82611041565b805f5b838110156110bb5781516110a2878261104a565b96506110ad83611061565b92505060018101905061108e565b505050505050565b608082015f8201516110d75f85018261106d565b5060208201516110ea604085018261106d565b50505050565b61010082015f8201516111055f850182611000565b50602082015161111860408501826110c3565b50604082015161112b60c0850182611000565b50505050565b5f60419050919050565b5f81905092915050565b5f819050919050565b5f602082019050919050565b61116381611131565b61116d818461113b565b925061117882611145565b805f5b838110156111a857815161118f878261104a565b965061119a8361114e565b92505060018101905061117b565b505050505050565b5f81905092915050565b6111c38161102d565b6111cd81846111b0565b92506111d882611041565b805f5b838110156112085781516111ef878261104a565b96506111fa83611061565b9250506001810190506111db565b505050505050565b5f610960820190506112245f8301866110f0565b61123261010083018561115a565b6112406109208301846111ba565b949350505050565b61125181610af6565b811461125b575f80fd5b50565b5f8151905061126c81611248565b92915050565b5f6020828403121561128757611286610a27565b5b5f6112948482850161125e565b91505092915050565b7f50726f6f6620697320696e76616c6964000000000000000000000000000000005f82015250565b5f6112d1601083610998565b91506112dc8261129d565b602082019050919050565b5f6020820190508181035f8301526112fe816112c5565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61133c82610a85565b915061134783610a85565b925082820190508082111561135f5761135e611305565b5b92915050565b5f61136f82610a85565b915061137a83610a85565b925082820390508181111561139257611391611305565b5b9291505056fea2646970667358221220599c00a58c6c0d0bca9c08aebb057c97cfb301b486a2ae2c9fb398f5160e6e5364736f6c63430008190033",
}

// TestTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TestTokenMetaData.ABI instead.
var TestTokenABI = TestTokenMetaData.ABI

// TestTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestTokenMetaData.Bin instead.
var TestTokenBin = TestTokenMetaData.Bin

// DeployTestToken deploys a new Ethereum contract, binding an instance of TestToken to it.
func DeployTestToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestToken, error) {
	parsed, err := TestTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestToken{TestTokenCaller: TestTokenCaller{contract: contract}, TestTokenTransactor: TestTokenTransactor{contract: contract}, TestTokenFilterer: TestTokenFilterer{contract: contract}}, nil
}

// TestToken is an auto generated Go binding around an Ethereum contract.
type TestToken struct {
	TestTokenCaller     // Read-only binding to the contract
	TestTokenTransactor // Write-only binding to the contract
	TestTokenFilterer   // Log filterer for contract events
}

// TestTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestTokenSession struct {
	Contract     *TestToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestTokenCallerSession struct {
	Contract *TestTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TestTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestTokenTransactorSession struct {
	Contract     *TestTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TestTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestTokenRaw struct {
	Contract *TestToken // Generic contract binding to access the raw methods on
}

// TestTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestTokenCallerRaw struct {
	Contract *TestTokenCaller // Generic read-only contract binding to access the raw methods on
}

// TestTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestTokenTransactorRaw struct {
	Contract *TestTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestToken creates a new instance of TestToken, bound to a specific deployed contract.
func NewTestToken(address common.Address, backend bind.ContractBackend) (*TestToken, error) {
	contract, err := bindTestToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestToken{TestTokenCaller: TestTokenCaller{contract: contract}, TestTokenTransactor: TestTokenTransactor{contract: contract}, TestTokenFilterer: TestTokenFilterer{contract: contract}}, nil
}

// NewTestTokenCaller creates a new read-only instance of TestToken, bound to a specific deployed contract.
func NewTestTokenCaller(address common.Address, caller bind.ContractCaller) (*TestTokenCaller, error) {
	contract, err := bindTestToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestTokenCaller{contract: contract}, nil
}

// NewTestTokenTransactor creates a new write-only instance of TestToken, bound to a specific deployed contract.
func NewTestTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TestTokenTransactor, error) {
	contract, err := bindTestToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestTokenTransactor{contract: contract}, nil
}

// NewTestTokenFilterer creates a new log filterer instance of TestToken, bound to a specific deployed contract.
func NewTestTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TestTokenFilterer, error) {
	contract, err := bindTestToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestTokenFilterer{contract: contract}, nil
}

// bindTestToken binds a generic wrapper to an already deployed contract.
func bindTestToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestToken *TestTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestToken.Contract.TestTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestToken *TestTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestToken.Contract.TestTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestToken *TestTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestToken.Contract.TestTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestToken *TestTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestToken *TestTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestToken *TestTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) view returns(uint256 remaining)
func (_TestToken *TestTokenCaller) Allowance(opts *bind.CallOpts, tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestToken.contract.Call(opts, &out, "allowance", tokenOwner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) view returns(uint256 remaining)
func (_TestToken *TestTokenSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _TestToken.Contract.Allowance(&_TestToken.CallOpts, tokenOwner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) view returns(uint256 remaining)
func (_TestToken *TestTokenCallerSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _TestToken.Contract.Allowance(&_TestToken.CallOpts, tokenOwner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256 balance)
func (_TestToken *TestTokenCaller) BalanceOf(opts *bind.CallOpts, tokenOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestToken.contract.Call(opts, &out, "balanceOf", tokenOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256 balance)
func (_TestToken *TestTokenSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _TestToken.Contract.BalanceOf(&_TestToken.CallOpts, tokenOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256 balance)
func (_TestToken *TestTokenCallerSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _TestToken.Contract.BalanceOf(&_TestToken.CallOpts, tokenOwner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestToken *TestTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TestToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestToken *TestTokenSession) Decimals() (uint8, error) {
	return _TestToken.Contract.Decimals(&_TestToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestToken *TestTokenCallerSession) Decimals() (uint8, error) {
	return _TestToken.Contract.Decimals(&_TestToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestToken *TestTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestToken *TestTokenSession) Name() (string, error) {
	return _TestToken.Contract.Name(&_TestToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestToken *TestTokenCallerSession) Name() (string, error) {
	return _TestToken.Contract.Name(&_TestToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestToken *TestTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestToken *TestTokenSession) Symbol() (string, error) {
	return _TestToken.Contract.Symbol(&_TestToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestToken *TestTokenCallerSession) Symbol() (string, error) {
	return _TestToken.Contract.Symbol(&_TestToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestToken *TestTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestToken *TestTokenSession) TotalSupply() (*big.Int, error) {
	return _TestToken.Contract.TotalSupply(&_TestToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestToken *TestTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _TestToken.Contract.TotalSupply(&_TestToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 tokens) returns(bool success)
func (_TestToken *TestTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _TestToken.contract.Transact(opts, "approve", spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 tokens) returns(bool success)
func (_TestToken *TestTokenSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _TestToken.Contract.Approve(&_TestToken.TransactOpts, spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 tokens) returns(bool success)
func (_TestToken *TestTokenTransactorSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _TestToken.Contract.Approve(&_TestToken.TransactOpts, spender, tokens)
}

// MintWithProof is a paid mutator transaction binding the contract method 0x0acf8131.
//
// Solidity: function mintWithProof(address to, uint256 value, ((uint256,uint256),(uint256[2],uint256[2]),(uint256,uint256)) proof, uint256[65] input, uint256[2] proof_commitment) returns()
func (_TestToken *TestTokenTransactor) MintWithProof(opts *bind.TransactOpts, to common.Address, value *big.Int, proof Proof, input [65]*big.Int, proof_commitment [2]*big.Int) (*types.Transaction, error) {
	return _TestToken.contract.Transact(opts, "mintWithProof", to, value, proof, input, proof_commitment)
}

// MintWithProof is a paid mutator transaction binding the contract method 0x0acf8131.
//
// Solidity: function mintWithProof(address to, uint256 value, ((uint256,uint256),(uint256[2],uint256[2]),(uint256,uint256)) proof, uint256[65] input, uint256[2] proof_commitment) returns()
func (_TestToken *TestTokenSession) MintWithProof(to common.Address, value *big.Int, proof Proof, input [65]*big.Int, proof_commitment [2]*big.Int) (*types.Transaction, error) {
	return _TestToken.Contract.MintWithProof(&_TestToken.TransactOpts, to, value, proof, input, proof_commitment)
}

// MintWithProof is a paid mutator transaction binding the contract method 0x0acf8131.
//
// Solidity: function mintWithProof(address to, uint256 value, ((uint256,uint256),(uint256[2],uint256[2]),(uint256,uint256)) proof, uint256[65] input, uint256[2] proof_commitment) returns()
func (_TestToken *TestTokenTransactorSession) MintWithProof(to common.Address, value *big.Int, proof Proof, input [65]*big.Int, proof_commitment [2]*big.Int) (*types.Transaction, error) {
	return _TestToken.Contract.MintWithProof(&_TestToken.TransactOpts, to, value, proof, input, proof_commitment)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 tokens) returns(bool success)
func (_TestToken *TestTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _TestToken.contract.Transact(opts, "transfer", to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 tokens) returns(bool success)
func (_TestToken *TestTokenSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _TestToken.Contract.Transfer(&_TestToken.TransactOpts, to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 tokens) returns(bool success)
func (_TestToken *TestTokenTransactorSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _TestToken.Contract.Transfer(&_TestToken.TransactOpts, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokens) returns(bool success)
func (_TestToken *TestTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _TestToken.contract.Transact(opts, "transferFrom", from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokens) returns(bool success)
func (_TestToken *TestTokenSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _TestToken.Contract.TransferFrom(&_TestToken.TransactOpts, from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokens) returns(bool success)
func (_TestToken *TestTokenTransactorSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _TestToken.Contract.TransferFrom(&_TestToken.TransactOpts, from, to, tokens)
}

// TestTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TestToken contract.
type TestTokenApprovalIterator struct {
	Event *TestTokenApproval // Event containing the contract specifics and raw log

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
func (it *TestTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestTokenApproval)
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
		it.Event = new(TestTokenApproval)
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
func (it *TestTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestTokenApproval represents a Approval event raised by the TestToken contract.
type TestTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestToken *TestTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TestTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TestToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TestTokenApprovalIterator{contract: _TestToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestToken *TestTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TestTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TestToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestTokenApproval)
				if err := _TestToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestToken *TestTokenFilterer) ParseApproval(log types.Log) (*TestTokenApproval, error) {
	event := new(TestTokenApproval)
	if err := _TestToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TestToken contract.
type TestTokenTransferIterator struct {
	Event *TestTokenTransfer // Event containing the contract specifics and raw log

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
func (it *TestTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestTokenTransfer)
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
		it.Event = new(TestTokenTransfer)
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
func (it *TestTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestTokenTransfer represents a Transfer event raised by the TestToken contract.
type TestTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestToken *TestTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TestTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TestTokenTransferIterator{contract: _TestToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestToken *TestTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TestTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestTokenTransfer)
				if err := _TestToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestToken *TestTokenFilterer) ParseTransfer(log types.Log) (*TestTokenTransfer, error) {
	event := new(TestTokenTransfer)
	if err := _TestToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
