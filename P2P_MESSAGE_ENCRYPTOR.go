package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func AESEncrypt(plainText, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, 12)
	stream := cipher.NewGCM(block)
	cipherText := stream.Seal(nil, nonce, plainText, nil)
	return hex.EncodeToString(cipherText), nil
}

func main() {
	key := []byte("12345678901234567890123456789012")
	msg := []byte("BaseChainP2PMessage")
	enc, _ := AESEncrypt(msg, key)
	fmt.Printf("加密后消息: %s\n", enc)
}
