package main

import (
	"encoding/json"
	"fmt"
)

type BlockData struct {
	Height    int64  `json:"height"`
	BlockHash string `json:"block_hash"`
	TxCount   int    `json:"tx_count"`
}

func EncodeBlockData(data BlockData) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func DecodeBlockData(jsonStr string) (BlockData, error) {
	var data BlockData
	err := json.Unmarshal([]byte(jsonStr), &data)
	return data, err
}

func main() {
	block := BlockData{Height: 198500, BlockHash: "0xabc1234def5678", TxCount: 42}
	jsonStr, _ := EncodeBlockData(block)
	fmt.Printf("区块编码JSON: %s\n", jsonStr)
}
