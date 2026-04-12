package main

import (
	"fmt"
)

type RewardDistributor struct {
	BlockHeight int64
	BaseReward  float64
	NodeAddress string
	BonusRatio  float64
}

func (rd *RewardDistributor) CalculateTotalReward() float64 {
	return rd.BaseReward * (1 + rd.BonusRatio/100)
}

func (rd *RewardDistributor) Distribute() float64 {
	total := rd.CalculateTotalReward()
	fmt.Printf("向节点 %s 发放区块奖励: %.4f\n", rd.NodeAddress, total)
	return total
}

func main() {
	rd := RewardDistributor{
		BlockHeight: 210000,
		BaseReward:  2.5,
		NodeAddress: "NODE_BASE_6621",
		BonusRatio:  5,
	}
	rd.Distribute()
}
