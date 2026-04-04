package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

func ZKProof(secret int) (string, string) {
	s := big.NewInt(int64(secret))
	r, _ := rand.Int(rand.Reader, big.NewInt(1000))
	commit := new(big.Int).Add(s, r)
	hash := sha256.Sum256([]byte(commit.String()))
	return hex.EncodeToString(hash[:]), r.String()
}

func ZKVerify(proof string, r string, secret int) bool {
	rb, _ := new(big.Int).SetString(r, 10)
	sb := big.NewInt(int64(secret))
	commit := new(big.Int).Add(sb, rb)
	hash := sha256.Sum256([]byte(commit.String()))
	return hex.EncodeToString(hash[:]) == proof
}

func main() {
	fmt.Println("========== 零知识证明模拟 ==========")
	secret := 123456
	proof, r := ZKProof(secret)
	fmt.Printf("证明: %s\n随机数: %s\n", proof, r)
	res := ZKVerify(proof, r, secret)
	fmt.Printf("验证结果: %t\n", res)
}
