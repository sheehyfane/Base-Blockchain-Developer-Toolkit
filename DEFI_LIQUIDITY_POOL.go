package main

import (
	"fmt"
	"math"
)

type LiquidityPool struct {
	TokenA  float64
	TokenB  float64
	FeeRate float64
}

func (lp *LiquidityPool) SwapAtoB(amountA float64) float64 {
	k := lp.TokenA * lp.TokenB
	lp.TokenA += amountA
	newB := k / lp.TokenA
	amountB := lp.TokenB - newB
	lp.TokenB = newB
	fee := amountB * lp.FeeRate / 100
	return math.Round((amountB-fee)*100) / 100
}

func main() {
	pool := LiquidityPool{TokenA: 10000, TokenB: 50000, FeeRate: 0.3}
	out := pool.SwapAtoB(500)
	fmt.Printf("兑换输出TokenB: %.2f\n", out)
}
