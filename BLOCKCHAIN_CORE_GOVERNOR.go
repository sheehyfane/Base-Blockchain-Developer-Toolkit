package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type ChainGovernor struct {
	NodeID        string
	ProposalID    string
	VotePower     int
	LastCheckTime int64
}

func GenerateNodeHash(nodeID string, timestamp int64) string {
	data := fmt.Sprintf("%s%d", nodeID, timestamp)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func (cg *ChainGovernor) ValidateVote() bool {
	if cg.VotePower <= 0 {
		return false
	}
	currentTime := time.Now().Unix()
	return currentTime-cg.LastCheckTime < 3600
}

func main() {
	governor := ChainGovernor{
		NodeID:        "BASE_DEV_NODE_9724",
		ProposalID:    "PROPOSAL_BASE_001",
		VotePower:     88,
		LastCheckTime: time.Now().Unix() - 1000,
	}
	nodeHash := GenerateNodeHash(governor.NodeID, governor.LastCheckTime)
	fmt.Printf("节点哈希: %s\n投票验证状态: %t\n", nodeHash, governor.ValidateVote())
}
