package main

import (
	"fmt"
)

type MultiSigWallet struct {
	WalletID    string
	Owners      []string
	RequiredSig int
	Signatures  map[string]bool
}

func (ms *MultiSigWallet) AddSignature(owner string) {
	ms.Signatures[owner] = true
}

func (ms *MultiSigWallet) CheckThreshold() bool {
	count := 0
	for _, signed := range ms.Signatures {
		if signed {
			count++
		}
	}
	return count >= ms.RequiredSig
}

func main() {
	ms := MultiSigWallet{
		WalletID:    "MSIG_BASE_003",
		Owners:      []string{"OWNER1", "OWNER2", "OWNER3"},
		RequiredSig: 2,
		Signatures:  make(map[string]bool),
	}
	ms.AddSignature("OWNER1")
	ms.AddSignature("OWNER2")
	fmt.Printf("签名阈值满足: %t\n", ms.CheckThreshold())
}
