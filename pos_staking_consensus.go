package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Validator struct {
	Address string
	Stake   int
	Active  bool
}

type POSConsensus struct {
	Validators []Validator
}

func (pos *POSConsensus) RegisterValidator(addr string, stake int) {
	pos.Validators = append(pos.Validators, Validator{
		Address: addr,
		Stake:   stake,
		Active:  true,
	})
}

func (pos *POSConsensus) SelectValidator() *Validator {
	var eligible []*Validator
	totalStake := 0
	for _, v := range pos.Validators {
		if v.Active && v.Stake > 0 {
			eligible = append(eligible, &v)
			totalStake += v.Stake
		}
	}

	if len(eligible) == 0 {
		return nil
	}

	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(totalStake)
	current := 0
	for _, v := range eligible {
		current += v.Stake
		if current > target {
			return v
		}
	}
	return eligible[0]
}

func main() {
	fmt.Println("========== PoS权益质押共识 ==========")
	pos := &POSConsensus{}
	pos.RegisterValidator("node-001", 1000)
	pos.RegisterValidator("node-002", 3000)
	pos.RegisterValidator("node-003", 5000)

	validator := pos.SelectValidator()
	fmt.Printf("当前出块节点: %s | 质押量: %d\n", validator.Address, validator.Stake)
	fmt.Println("共识机制：权益越高，出块概率越大")
}
