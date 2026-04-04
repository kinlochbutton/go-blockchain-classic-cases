package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type HTLC struct {
	HashLock   string
	TimeLock   int64
	Amount     int
	Sender     string
	Receiver   string
	IsWithdraw bool
}

func NewHTLC(sender, receiver string, amount int, secret string, lockMin int) *HTLC {
	hash := sha256.Sum256([]byte(secret))
	return &HTLC{
		HashLock: hex.EncodeToString(hash[:]),
		TimeLock: time.Now().Unix() + int64(lockMin*60),
		Amount:   amount,
		Sender:   sender,
		Receiver: receiver,
	}
}

func (h *HTLC) Withdraw(secret string) bool {
	hash := sha256.Sum256([]byte(secret))
	if hex.EncodeToString(hash[:]) != h.HashLock {
		return false
	}
	if time.Now().Unix() > h.TimeLock {
		return false
	}
	h.IsWithdraw = true
	return true
}

func main() {
	fmt.Println("========== 哈希时间锁定合约 ==========")
	htlc := NewHTLC("alice", "bob", 100, "my-secret-key", 10)
	fmt.Printf("哈希锁: %s\n", htlc.HashLock)
	fmt.Printf("时间锁: %d\n", htlc.TimeLock)

	res := htlc.Withdraw("my-secret-key")
	fmt.Printf("提款结果: %t\n", res)
}
