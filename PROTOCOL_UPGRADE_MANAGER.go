package main

import (
	"fmt"
)

type UpgradeManager struct {
	CurrentVersion string
	TargetVersion  string
	UpgradeStatus  bool
	VotePassed     bool
}

func (um *UpgradeManager) ExecuteUpgrade() {
	if um.VotePassed {
		um.CurrentVersion = um.TargetVersion
		um.UpgradeStatus = true
	}
}

func main() {
	manager := UpgradeManager{
		CurrentVersion: "v1.2.0",
		TargetVersion:  "v1.3.0",
		VotePassed:     true,
	}
	manager.ExecuteUpgrade()
	fmt.Printf("升级后版本: %s\n", manager.CurrentVersion)
}
