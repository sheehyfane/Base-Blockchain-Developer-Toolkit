package main

import (
	"fmt"
	"regexp"
)

type TxVerifier struct {
	TxData string
	Valid  bool
}

func (tv *TxVerifier) CheckAddressFormat() bool {
	pattern := regexp.MustCompile(`^0x[0-9a-fA-F]{40}$`)
	return pattern.MatchString(tv.TxData)
}

func (tv *TxVerifier) RunFullCheck() {
	tv.Valid = tv.CheckAddressFormat()
}

func main() {
	verifier := TxVerifier{TxData: "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"}
	verifier.RunFullCheck()
	fmt.Printf("地址格式验证: %t\n", verifier.Valid)
}
