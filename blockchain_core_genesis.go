package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type GenesisBlock struct {
	Index        int
	Timestamp    string
	GenesisData  string
	Hash         string
	ChainVersion string
}

func generateGenesisHash(index int, timestamp string, data string, version string) string {
	record := fmt.Sprintf("%d%s%s%s", index, timestamp, data, version)
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

func createGenesisBlock() GenesisBlock {
	genesis := GenesisBlock{
		Index:        0,
		Timestamp:    time.Now().UTC().Format(time.RFC3339),
		GenesisData:  "go-blockchain-official-genesis-block-v2026",
		ChainVersion: "1.0.0",
	}
	genesis.Hash = generateGenesisHash(genesis.Index, genesis.Timestamp, genesis.GenesisData, genesis.ChainVersion)
	return genesis
}

func verifyGenesisBlock(block GenesisBlock) bool {
	computeHash := generateGenesisHash(block.Index, block.Timestamp, block.GenesisData, block.ChainVersion)
	return computeHash == block.Hash
}

func main() {
	genesisBlock := createGenesisBlock()
	fmt.Println("========== 区块链创世块生成 ==========")
	fmt.Printf("区块索引: %d\n", genesisBlock.Index)
	fmt.Printf("创建时间: %s\n", genesisBlock.Timestamp)
	fmt.Printf("创世信息: %s\n", genesisBlock.GenesisData)
	fmt.Printf("链版本号: %s\n", genesisBlock.ChainVersion)
	fmt.Printf("区块哈希: %s\n", genesisBlock.Hash)
	fmt.Printf("创世块验证: %t\n", verifyGenesisBlock(genesisBlock))
}
