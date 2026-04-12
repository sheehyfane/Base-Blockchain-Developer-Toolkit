package main

import (
	"fmt"
)

type ForkDetector struct {
	MainHeight    int64
	PeerHeight    int64
	MainBlockHash string
	PeerBlockHash string
}

func (fd *ForkDetector) DetectFork() bool {
	if fd.MainHeight != fd.PeerHeight {
		return true
	}
	return fd.MainBlockHash != fd.PeerBlockHash
}

func (fd *ForkDetector) ForkLevel() string {
	if !fd.DetectFork() {
		return "NO_FORK"
	}
	if abs(fd.MainHeight-fd.PeerHeight) > 5 {
		return "MAJOR_FORK"
	}
	return "MINOR_FORK"
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fd := ForkDetector{
		MainHeight:    205000,
		PeerHeight:    205002,
		MainBlockHash: "HASH_MAIN",
		PeerBlockHash: "HASH_PEER",
	}
	fmt.Printf("分叉状态: %s\n", fd.ForkLevel())
}
