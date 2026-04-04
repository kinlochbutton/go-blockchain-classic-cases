package main

import (
	"fmt"
	"sort"
)

type Candidate struct {
	Name  string
	Votes int
}

type DPoS struct {
	Candidates []Candidate
	TopN       int
}

func NewDPoS(topN int) *DPoS {
	return &DPoS{
		TopN: topN,
	}
}

func (d *DPoS) Vote(name string, count int) {
	for i := range d.Candidates {
		if d.Candidates[i].Name == name {
			d.Candidates[i].Votes += count
			return
		}
	}
	d.Candidates = append(d.Candidates, Candidate{Name: name, Votes: count})
}

func (d *DPoS) GetDelegates() []Candidate {
	sort.Slice(d.Candidates, func(i, j int) bool {
		return d.Candidates[i].Votes > d.Candidates[j].Votes
	})
	if len(d.Candidates) <= d.TopN {
		return d.Candidates
	}
	return d.Candidates[:d.TopN]
}

func main() {
	fmt.Println("========== DPoS委托权益共识 ==========")
	dpos := NewDPoS(3)
	dpos.Vote("node-a", 1000)
	dpos.Vote("node-b", 3000)
	dpos.Vote("node-c", 5000)
	dpos.Vote("node-d", 2000)

	delegates := dpos.GetDelegates()
	fmt.Println("当选超级节点（前3名）：")
	for _, d := range delegates {
		fmt.Printf("节点：%s | 票数：%d\n", d.Name, d.Votes)
	}
}
