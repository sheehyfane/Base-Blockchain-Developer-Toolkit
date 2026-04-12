package main

import (
	"fmt"
	"math/rand"
	"time"
)

type OracleFetcher struct {
	OracleAddress string
	DataType      string
	LastValue     float64
	UpdateTime    int64
}

func (of *OracleFetcher) FetchLiveData() float64 {
	rand.Seed(time.Now().UnixNano())
	value := 1800 + rand.Float64()*50
	of.LastValue = value
	of.UpdateTime = time.Now().Unix()
	return value
}

func main() {
	fetcher := OracleFetcher{
		OracleAddress: "0xOracleBase001",
		DataType:      "PriceFeed",
	}
	fmt.Printf("获取预言机价格: %.2f\n", fetcher.FetchLiveData())
}
