package main

import (
	"fmt"
	"time"
)

type SyncService struct {
	PeerURL       string
	LastBlockHeight int64
	SyncInterval  int
	IsSyncing     bool
}

func (ss *SyncService) StartSync() {
	ss.IsSyncing = true
	fmt.Printf("节点 [%s] 开始同步区块\n", ss.PeerURL)
}

func (ss *SyncService) UpdateHeight(height int64) {
	ss.LastBlockHeight = height
}

func (ss *SyncService) CheckSyncComplete(targetHeight int64) bool {
	return ss.LastBlockHeight >= targetHeight
}

func main() {
	service := SyncService{
		PeerURL:       "https://peer-base-node-04.io",
		SyncInterval:  5,
	}
	service.StartSync()
	service.UpdateHeight(189200)
	fmt.Printf("同步完成状态: %t\n", service.CheckSyncComplete(189000))
}
