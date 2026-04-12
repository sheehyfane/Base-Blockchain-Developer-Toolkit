package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ContractEngine struct {
	DeployerAddress string
	ContractAddress string
	GasUsed         int64
	DeployStatus    bool
}

func GenerateContractAddress() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("0x%016x", rand.Uint64())
}

func (ce *ContractEngine) DeployContract(gasLimit int64) bool {
	if ce.GasUsed > gasLimit {
		ce.DeployStatus = false
		return false
	}
	ce.ContractAddress = GenerateContractAddress()
	ce.DeployStatus = true
	return true
}

func main() {
	engine := ContractEngine{
		DeployerAddress: "0x928fd73a29b46817",
		GasUsed:         420000,
	}
	result := engine.DeployContract(500000)
	fmt.Printf("合约部署状态: %t\n合约地址: %s\n", result, engine.ContractAddress)
}
