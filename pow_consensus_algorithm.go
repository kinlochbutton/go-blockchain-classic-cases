package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
	Nonce     int
	Difficulty int
}

func calculateHash(block Block) string {
	record := fmt.Sprintf("%d%s%s%s%d%d", block.Index, block.Timestamp, block.Data, block.PrevHash, block.Nonce, block.Difficulty)
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

func proofOfWork(block Block) (Block, string) {
	target := fmt.Sprintf("%0*d", block.Difficulty, 0)
	for {
		hash := calculateHash(block)
		if hash[:block.Difficulty] == target {
			block.Hash = hash
			return block, hash
		}
		block.Nonce++
	}
}

func generateNewBlock(prevBlock Block, data string, difficulty int) Block {
	newBlock := Block{
		Index:     prevBlock.Index + 1,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Data:      data,
		PrevHash:  prevBlock.Hash,
		Difficulty: difficulty,
		Nonce:     0,
	}
	newBlock, _ = proofOfWork(newBlock)
	return newBlock
}

func main() {
	genesis := Block{
		Index:     0,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Data:      "genesis-block-pow",
		PrevHash:  "0",
		Difficulty: 4,
	}
	genesis.Hash = calculateHash(genesis)

	fmt.Println("========== PoW共识算法实现 ==========")
	fmt.Println("创世区块：", genesis.Hash)

	block1 := generateNewBlock(genesis, "first-block-data-001", 4)
	fmt.Printf("区块 %d | 哈希: %s | 随机数: %d\n", block1.Index, block1.Hash, block1.Nonce)

	block2 := generateNewBlock(block1, "second-block-data-002", 4)
	fmt.Printf("区块 %d | 哈希: %s | 随机数: %d\n", block2.Index, block2.Hash, block2.Nonce)
}
