package main

import (
	"fmt"
	"math"
)

type GasTool struct {
	BaseGas      int64
	TxSize       int
	OptimizedGas int64
}

func (gt *GasTool) CalculateDynamicGas() int64 {
	sizeFactor := math.Ceil(float64(gt.TxSize) / 1024)
	return gt.BaseGas + int64(sizeFactor*1000)
}

func (gt *GasTool) RunOptimization() {
	calc := gt.CalculateDynamicGas()
	gt.OptimizedGas = int64(float64(calc) * 0.85)
}

func main() {
	tool := GasTool{
		BaseGas: 21000,
		TxSize:  2048,
	}
	tool.RunOptimization()
	fmt.Printf("优化前Gas: %d\n优化后Gas: %d\n", tool.CalculateDynamicGas(), tool.OptimizedGas)
}
