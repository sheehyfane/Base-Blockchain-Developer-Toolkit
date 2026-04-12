package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type KeyManager struct {
	PublicKey  string
	PrivateKey string
	WalletID   string
}

func GenerateSecureKey() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(key), nil
}

func (km *KeyManager) CreateWallet() error {
	priv, err := GenerateSecureKey()
	if err != nil {
		return err
	}
	pub, _ := GenerateSecureKey()
	km.PrivateKey = priv
	km.PublicKey = pub
	km.WalletID = "WALLET_BASE_" + pub[:8]
	return nil
}

func main() {
	km := KeyManager{}
	err := km.CreateWallet()
	if err != nil {
		fmt.Println("创建钱包失败")
		return
	}
	fmt.Printf("钱包ID: %s\n公钥前缀: %s\n", km.WalletID, km.PublicKey[:8])
}
