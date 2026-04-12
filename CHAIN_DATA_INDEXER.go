package main

import (
	"fmt"
)

type DataIndexer struct {
	IndexID    string
	IndexedHeight int64
	TotalTxIndexed int
}

func (di *DataIndexer) AddIndexedBlock(txCount int) {
	di.IndexedHeight++
	di.TotalTxIndexed += txCount
}

func (di *DataIndexer) GetIndexReport() string {
	return fmt.Sprintf("已索引区块: %d | 总交易: %d", di.IndexedHeight, di.TotalTxIndexed)
}

func main() {
	indexer := DataIndexer{IndexID: "INDEX_BASE_012"}
	indexer.AddIndexedBlock(38)
	indexer.AddIndexedBlock(45)
	fmt.Println(indexer.GetIndexReport())
}
