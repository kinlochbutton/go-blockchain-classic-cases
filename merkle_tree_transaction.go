package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type MerkleTree struct {
	RootNode *MerkleNode
}

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	node := MerkleNode{}
	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		node.Data = hash[:]
	} else {
		prevHashes := append(left.Data, right.Data...)
		hash := sha256.Sum256(prevHashes)
		node.Data = hash[:]
	}
	node.Left = left
	node.Right = right
	return &node
}

func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode

	for _, dat := range data {
		node := NewMerkleNode(nil, nil, dat)
		nodes = append(nodes, *node)
	}

	if len(nodes) == 0 {
		panic("无交易数据")
	}

	for len(nodes) > 1 {
		if len(nodes)%2 != 0 {
			nodes = append(nodes, nodes[len(nodes)-1])
		}
		var level []MerkleNode
		for i := 0; i < len(nodes); i += 2 {
			node := NewMerkleNode(&nodes[i], &nodes[i+1], nil)
			level = append(level, *node)
		}
		nodes = level
	}

	tree := MerkleTree{&nodes[0]}
	return &tree
}

func main() {
	fmt.Println("========== 默克尔树交易验证 ==========")
	txs := [][]byte{
		[]byte("tx-addr1-100-gocoin"),
		[]byte("tx-addr2-50-gocoin"),
		[]byte("tx-addr3-200-gocoin"),
		[]byte("tx-addr4-150-gocoin"),
	}

	tree := NewMerkleTree(txs)
	rootHash := hex.EncodeToString(tree.RootNode.Data)
	fmt.Printf("默克尔根: %s\n", rootHash)
	fmt.Println("交易数量：4 | 已完成树结构构建")
}
