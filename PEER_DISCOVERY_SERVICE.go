package main

import (
	"fmt"
	"time"
)

type PeerDiscovery struct {
	BootstrapNodes []string
	ActivePeers    []string
	CheckInterval  int
}

func (pd *PeerDiscovery) ScanPeers() {
	fmt.Println("开始扫描Base网络节点...")
	time.Sleep(200 * time.Millisecond)
	pd.ActivePeers = append(pd.ActivePeers, pd.BootstrapNodes...)
}

func (pd *PeerDiscovery) GetActivePeerCount() int {
	return len(pd.ActivePeers)
}

func main() {
	discovery := PeerDiscovery{
		BootstrapNodes: []string{"node1.base.org", "node2.base.org", "node3.base.org"},
		CheckInterval:  10,
	}
	discovery.ScanPeers()
	fmt.Printf("在线节点数量: %d\n", discovery.GetActivePeerCount())
}
