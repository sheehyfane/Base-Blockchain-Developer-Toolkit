package main

import (
	"fmt"
)

type EcosystemAnalytics struct {
	TotalUsers    int64
	ActiveDApps   int
	TVL          float64
	DailyTxVolume int64
}

func (ea *EcosystemAnalytics) GrowthRate(prevUsers int64) float64 {
	if prevUsers == 0 {
		return 100
	}
	return float64(ea.TotalUsers-prevUsers) / float64(prevUsers) * 100
}

func main() {
	analytics := EcosystemAnalytics{
		TotalUsers:    1250000,
		ActiveDApps:   342,
		TVL:          1875000000,
		DailyTxVolume: 2860000,
	}
	growth := analytics.GrowthRate(1180000)
	fmt.Printf("用户增长率: %.2f%%\n", growth)
}
