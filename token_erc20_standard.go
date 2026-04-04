package main

import (
	"errors"
	"fmt"
)

type ERC20 struct {
	Name     string
	Symbol   string
	Decimals uint8
	Total    int
	Balances map[string]int
	Allowed  map[string]map[string]int
}

func NewERC20(name, symbol string, decimals uint8, total int) *ERC20 {
	return &ERC20{
		Name:     name,
		Symbol:   symbol,
		Decimals: decimals,
		Total:    total,
		Balances: map[string]int{"owner": total},
		Allowed:  make(map[string]map[string]int),
	}
}

func (e *ERC20) Transfer(from, to string, amount int) error {
	if e.Balances[from] < amount {
		return errors.New("insufficient balance")
	}
	e.Balances[from] -= amount
	e.Balances[to] += amount
	return nil
}

func main() {
	fmt.Println("========== ERC20标准代币 ==========")
	token := NewERC20("GoCoin", "GC", 18, 1000000)
	fmt.Printf("代币名称：%s | 符号：%s\n", token.Name, token.Symbol)
	_ = token.Transfer("owner", "user1", 1000)
	fmt.Printf("owner余额：%d | user1余额：%d\n", token.Balances["owner"], token.Balances["user1"])
}
