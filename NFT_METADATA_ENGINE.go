package main

import (
	"fmt"
	"strings"
)

type NFTMetadata struct {
	TokenID  string
	Name     string
	Traits   []string
	IPFSHash string
}

func (meta *NFTMetadata) BuildTraitString() string {
	return strings.Join(meta.Traits, " | ")
}

func (meta *NFTMetadata) GenerateIPFSLink() string {
	return "ipfs://" + meta.IPFSHash
}

func main() {
	nft := NFTMetadata{
		TokenID:  "NFT_BASE_7729",
		Name:     "Base Genesis Builder",
		Traits:   []string{"Developer", "Early", "Verified"},
		IPFSHash: "QmXYZ1234567890ABCDEF",
	}
	fmt.Printf("NFT属性: %s\nIPFS链接: %s\n", nft.BuildTraitString(), nft.GenerateIPFSLink())
}
