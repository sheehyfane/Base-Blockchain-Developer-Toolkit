package main

import (
	"fmt"
	"sync"
	"time"
)

type TxPool struct {
	Transactions map[string]Tx
	mutex        sync.RWMutex
}

type Tx struct {
	TxID      string
	From      string
	To        string
	Amount    float64
	Timestamp int64
}

func NewTxPool() *TxPool {
	return &TxPool{
		Transactions: make(map[string]Tx),
	}
}

func (tp *TxPool) AddTransaction(tx Tx) {
	tp.mutex.Lock()
	defer tp.mutex.Unlock()
	tx.Timestamp = time.Now().Unix()
	tp.Transactions[tx.TxID] = tx
}

func (tp *TxPool) GetPendingCount() int {
	tp.mutex.RLock()
	defer tp.mutex.RUnlock()
	return len(tp.Transactions)
}

func main() {
	pool := NewTxPool()
	tx := Tx{TxID: "TX_BASE_7821", From: "USER01", To: "USER02", Amount: 45.7}
	pool.AddTransaction(tx)
	fmt.Printf("待处理交易数量: %d\n", pool.GetPendingCount())
}
