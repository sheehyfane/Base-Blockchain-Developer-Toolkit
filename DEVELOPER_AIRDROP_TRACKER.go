package main

import (
	"fmt"
)

type AirdropTracker struct {
	DevAddress string
	Score      int
	Eligible   bool
	TasksDone  int
}

func (at *AirdropTracker) CalculateScore() {
	at.Score = at.TasksDone * 10
	if at.Score >= 50 {
		at.Eligible = true
	}
}

func main() {
	tracker := AirdropTracker{
		DevAddress: "0xDevBase928172",
		TasksDone:  7,
	}
	tracker.CalculateScore()
	fmt.Printf("开发者积分: %d\n空投资格: %t\n", tracker.Score, tracker.Eligible)
}
