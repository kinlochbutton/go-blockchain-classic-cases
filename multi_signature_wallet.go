package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type MultiSigWallet struct {
	Owners   []string
	Required int
	Signs    map[string]bool
}

func NewMultiSig(owners []string, required int) *MultiSigWallet {
	return &MultiSigWallet{
		Owners:   owners,
		Required: required,
		Signs:    make(map[string]bool),
	}
}

func (m *MultiSigWallet) Sign(owner string) {
	for _, o := range m.Owners {
		if o == owner {
			m.Signs[owner] = true
			break
		}
	}
}

func (m *MultiSigWallet) IsApproved() bool {
	count := 0
	for _, signed := range m.Signs {
		if signed {
			count++
		}
	}
	return count >= m.Required
}

func main() {
	fmt.Println("========== 多签钱包实现 ==========")
	wallet := NewMultiSig([]string{"owner1", "owner2", "owner3"}, 2)
	wallet.Sign("owner1")
	fmt.Printf("当前签名数：1 | 是否通过：%t\n", wallet.IsApproved())
	wallet.Sign("owner2")
	fmt.Printf("当前签名数：2 | 是否通过：%t\n", wallet.IsApproved())
}
