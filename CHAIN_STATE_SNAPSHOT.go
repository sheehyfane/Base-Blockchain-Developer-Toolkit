package main

import (
	"fmt"
	"time"
)

type StateSnapshot struct {
	SnapshotID  string
	BlockHeight int64
	StateRoot   string
	CreatedAt   int64
}

func CreateSnapshot(height int64, root string) StateSnapshot {
	return StateSnapshot{
		SnapshotID:  fmt.Sprintf("SNAP_%d", time.Now().Unix()),
		BlockHeight: height,
		StateRoot:   root,
		CreatedAt:   time.Now().Unix(),
	}
}

func main() {
	snap := CreateSnapshot(215600, "ROOT_HASH_987654321")
	fmt.Printf("快照ID: %s\n状态根: %s\n", snap.SnapshotID, snap.StateRoot)
}
