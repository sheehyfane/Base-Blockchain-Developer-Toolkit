package main

import (
	"fmt"
	"time"
)

type Monitor struct {
	ChainName        string
	CurrentHeight int64
	NodeHealth     bool
	LastUpdate     int64
}

func (m *Monitor) Refresh() {
	m.LastUpdate = time.Now().Unix()
	m.NodeHealth = true
	fmt.Printf("[%s] 区块高度已刷新: %d\n", m.ChainName, m.CurrentHeight)
}

func (m *Monitor) HealthCheck() string {
	if m.NodeHealth {
		return "HEALTHY"
	}
	return "ABNORMAL"
}

func main() {
	monitor := Monitor{ChainName: "Base", CurrentHeight: 201400}
	monitor.Refresh()
	fmt.Printf("节点状态: %s\n", monitor.HealthCheck())
}
