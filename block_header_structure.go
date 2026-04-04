package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type BlockHeader struct {
	Version    int
	PrevHash   string
	MerkleRoot string
	Timestamp  int64
	Difficulty int
	Nonce      int
}

func (bh *BlockHeader) Hash() string {
	data := strconv.Itoa(bh.Version) + bh.PrevHash + bh.MerkleRoot +
		strconv.FormatInt(bh.Timestamp, 10) + strconv.Itoa(bh.Difficulty) + strconv.Itoa(bh.Nonce)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func main() {
	fmt.Println("========== 区块头标准结构 ==========")
	header := BlockHeader{
		Version:    1,
		PrevHash:   "00000000000000000005a3d8",
		MerkleRoot: "abcdef1234567890",
		Timestamp:  time.Now().Unix(),
		Difficulty: 5,
		Nonce:      123456,
	}
	fmt.Printf("区块头哈希: %s\n", header.Hash())
	fmt.Println("字段：版本+前哈希+默克尔根+时间+难度+随机数")
}
