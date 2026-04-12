package main

import (
	"fmt"
	"math"
)

type RewardCalc struct {
	StakedAmount float64
	DurationDays int
	APR          float64
}

func (rc *RewardCalc) CalculateReward() float64 {
	rate := rc.APR / 100
	yearRatio := float64(rc.DurationDays) / 365
	reward := rc.StakedAmount * rate * yearRatio
	return math.Round(reward*100) / 100
}

func main() {
	calc := RewardCalc{
		StakedAmount: 5000,
		DurationDays: 180,
		APR:          8.5,
	}
	fmt.Printf("预计质押收益: %.2f\n", calc.CalculateReward())
}
