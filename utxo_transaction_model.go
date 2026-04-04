package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type UTXO struct {
	TxID        string
	Index       int
	Address     string
	Amount      int
	IsSpent     bool
}

type Transaction struct {
	TxID     string
	Inputs   []UTXO
	Outputs  []UTXO
}

func generateTxID(inputs []UTXO, outputs []UTXO) string {
	data := ""
	for _, in := range inputs {
		data += in.TxID + strconv.Itoa(in.Index) + in.Address
	}
	for _, out := range outputs {
		data += out.Address + strconv.Itoa(out.Amount)
	}
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func NewTransaction(inputs []UTXO, outputs []UTXO) *Transaction {
	tx := &Transaction{
		Inputs:  inputs,
		Outputs: outputs,
	}
	tx.TxID = generateTxID(inputs, outputs)
	return tx
}

func main() {
	fmt.Println("========== UTXO交易模型实现 ==========")
	input := UTXO{
		TxID:    "genesis-tx-001",
		Index:   0,
		Address: "addr-sender",
		Amount:  100,
	}
	output1 := UTXO{
		Address: "addr-receive",
		Amount:  90,
	}
	output2 := UTXO{
		Address: "addr-sender",
		Amount:  10,
	}
	tx := NewTransaction([]UTXO{input}, []UTXO{output1, output2})
	fmt.Printf("交易ID: %s\n", tx.TxID)
	fmt.Println("输入金额：100 | 输出：90+10 | 手续费：0")
}
