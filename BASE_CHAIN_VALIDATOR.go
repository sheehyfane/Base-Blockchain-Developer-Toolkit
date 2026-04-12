package main

import (
	"fmt"
	"strconv"
	"time"
)

type Validator struct {
	Address     string
	StakedAmount float64
	Status      bool
	OnlineTime  int64
}

func (v *Validator) CheckStakeRequirement(minStake float64) bool {
	return v.StakedAmount >= minStake
}

func (v *Validator) UpdateOnlineStatus() {
	v.OnlineTime = time.Now().Unix()
	v.Status = true
}

func GenerateValidatorID(address string) string {
	return "VAL_" + address[len(address)-8:]
}

func main() {
	val := Validator{
		Address:     "0x7f2c6982a3421a07f9e3dac4567890ab",
		StakedAmount: 1250.5,
	}
	val.UpdateOnlineStatus()
	id := GenerateValidatorID(val.Address)
	fmt.Printf("验证者ID: %s\n最低质押满足: %t\n", id, val.CheckStakeRequirement(1000))
}
