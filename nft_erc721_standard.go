package main

import (
	"errors"
	"fmt"
)

type ERC721 struct {
	Name     string
	Symbol   string
	Owner    map[int]string
	Balances map[string]int
}

func NewERC721(name, symbol string) *ERC721 {
	return &ERC721{
		Name:     name,
		Symbol:   symbol,
		Owner:    make(map[int]string),
		Balances: make(map[string]int),
	}
}

func (e *ERC721) Mint(to string, tokenID int) {
	e.Owner[tokenID] = to
	e.Balances[to]++
}

func (e *ERC721) Transfer(from, to string, tokenID int) error {
	if e.Owner[tokenID] != from {
		return errors.New("not owner")
	}
	e.Owner[tokenID] = to
	e.Balances[from]--
	e.Balances[to]++
	return nil
}

func main() {
	fmt.Println("========== ERC721 NFT标准 ==========")
	nft := NewERC721("GoArt", "GART")
	nft.Mint("artist", 1001)
	fmt.Printf("NFT 1001 持有者: %s\n", nft.Owner[1001])
	_ = nft.Transfer("artist", "collector", 1001)
	fmt.Printf("NFT 1001 持有者: %s\n", nft.Owner[1001])
}
