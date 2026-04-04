package main

import (
	"fmt"
	"time"
)

type Bridge struct {
	ChainA []string
	ChainB []string
}

func (b *Bridge) Lock(chain string, hash string) {
	if chain == "A" {
		b.ChainA = append(b.ChainA, hash)
	}
}

func (b *Bridge) Mint(chain string, hash string) {
	if chain == "B" {
		b.ChainB = append(b.ChainB, hash)
	}
}

func (b *Bridge) CrossTransfer(hash string) {
	fmt.Println("跨链转账中...锁定A链资产")
	b.Lock("A", hash)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("发行B链资产")
	b.Mint("B", hash)
	fmt.Println("跨链完成")
}

func main() {
	fmt.Println("========== 跨链桥模拟 ==========")
	bridge := &Bridge{}
	bridge.CrossTransfer("cross-tx-001")
	fmt.Printf("A链锁定: %d | B链发行: %d\n", len(bridge.ChainA), len(bridge.ChainB))
}
