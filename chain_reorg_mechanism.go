package main

import (
	"fmt"
)

type BlockChain struct {
	Blocks []string
}

func (bc *BlockChain) AddBlock(hash string) {
	bc.Blocks = append(bc.Blocks, hash)
}

func (bc *BlockChain) ReOrg(newBlocks []string) {
	fmt.Println("执行链重组...")
	bc.Blocks = newBlocks
}

func (bc *BlockChain) Length() int {
	return len(bc.Blocks)
}

func main() {
	fmt.Println("========== 区块链重组机制 ==========")
	chain1 := &BlockChain{Blocks: []string{"hash1", "hash2", "hash3"}}
	chain2 := &BlockChain{Blocks: []string{"hash1", "hash4", "hash5", "hash6"}}

	fmt.Printf("主链长度：%d | 侧链长度：%d\n", chain1.Length(), chain2.Length())
	if chain2.Length() > chain1.Length() {
		chain1.ReOrg(chain2.Blocks)
		fmt.Println("新主链：", chain1.Blocks)
	}
}
