package main

import (
	"fmt"
	"sync"
	"time"
)

type P2PNode struct {
	NodeID    string
	Peers     map[string]*P2PNode
	BlockChan chan string
	mutex     sync.RWMutex
}

func NewNode(id string) *P2PNode {
	return &P2PNode{
		NodeID:    id,
		Peers:     make(map[string]*P2PNode),
		BlockChan: make(chan string, 100),
	}
}

func (n *P2PNode) Connect(peer *P2PNode) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.Peers[peer.NodeID] = peer
}

func (n *P2PNode) Broadcast(msg string) {
	fmt.Printf("[%s] 广播消息: %s\n", n.NodeID, msg)
	n.mutex.RLock()
	defer n.mutex.RUnlock()
	for _, peer := range n.Peers {
		peer.BlockChan <- msg
	}
}

func (n *P2PNode) Start() {
	go func() {
		for msg := range n.BlockChan {
			fmt.Printf("[%s] 接收消息: %s\n", n.NodeID, msg)
		}
	}()
}

func main() {
	fmt.Println("========== 区块链P2P网络节点 ==========")
	node1 := NewNode("node-mainnet-01")
	node2 := NewNode("node-mainnet-02")
	node3 := NewNode("node-mainnet-03")

	node1.Connect(node2)
	node1.Connect(node3)
	node2.Connect(node1)
	node3.Connect(node1)

	node1.Start()
	node2.Start()
	node3.Start()

	node1.Broadcast("new-block-height-10086")
	time.Sleep(1 * time.Second)
}
