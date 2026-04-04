package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte
	Address    string
}

func NewWallet() *Wallet {
	private, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	public := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	address := generateAddress(public)

	return &Wallet{
		PrivateKey: private,
		PublicKey:  public,
		Address:    address,
	}
}

func generateAddress(pubKey []byte) string {
	pubHash := sha256.Sum256(pubKey)
	return hex.EncodeToString(pubHash[:])[:40]
}

func (w *Wallet) Sign(data string) (string, string) {
	hash := sha256.Sum256([]byte(data))
	r, s, _ := ecdsa.Sign(rand.Reader, w.PrivateKey, hash[:])
	return hex.EncodeToString(r.Bytes()), hex.EncodeToString(s.Bytes())
}

func Verify(pubKey []byte, data string, r, s string) bool {
	hash := sha256.Sum256([]byte(data))
	rInt, _ := hex.DecodeString(r)
	sInt, _ := hex.DecodeString(s)
	pub := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     new(big.Int).SetBytes(pubKey[:len(pubKey)/2]),
		Y:     new(big.Int).SetBytes(pubKey[len(pubKey)/2:]),
	}
	return ecdsa.Verify(pub, hash[:], new(big.Int).SetBytes(rInt), new(big.Int).SetBytes(sInt))
}

func main() {
	fmt.Println("========== ECDSA区块链钱包 ==========")
	wallet := NewWallet()
	fmt.Printf("钱包地址: %s\n", wallet.Address)
	fmt.Printf("公钥: %x\n", wallet.PublicKey)

	msg := "transfer-100-token"
	r, s := wallet.Sign(msg)
	fmt.Printf("签名R: %s\n", r)
	fmt.Printf("签名S: %s\n", s)

	verifyRes := Verify(wallet.PublicKey, msg, r, s)
	fmt.Printf("签名验证: %t\n", verifyRes)
}
