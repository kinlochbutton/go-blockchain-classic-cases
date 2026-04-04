package main

import (
	"fmt"
)

type RippleNode struct {
	ID    string
	Vote  bool
	UNL   []string
}

func (n *RippleNode) VoteResult() bool {
	return n.Vote
}

func RippleConsensus(nodes []*RippleNode) bool {
	yes := 0
	no := 0
	for _, node := range nodes {
		if node.Vote {
			yes++
		}} else {
			no++
		}
	}
	return yes > no
}

func main() {
	fmt.Println("========== Ripple共识算法 ==========")
	nodes := []*RippleNode{
		{"node1", true, []string{"node2", "node3"}},
		{"node2", true, []string{"node1", "node3"}},
		{"node3", false, []string{"node1", "node2"}},
	}
	res := RippleConsensus(nodes)
	fmt.Printf("共识结果: %t\n", res)
}
