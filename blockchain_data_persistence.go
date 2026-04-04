package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type BlockData struct {
	Height int    `json:"height"`
	Hash   string `json:"hash"`
	Data   string `json:"data"`
}

func SaveBlockToFile(block BlockData, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	return encoder.Encode(block)
}

func LoadBlockFromFile(filename string) (BlockData, error) {
	var block BlockData
	file, err := os.Open(filename)
	if err != nil {
		return block, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&block)
	return block, err
}

func main() {
	fmt.Println("========== 区块链数据持久化 ==========")
	block := BlockData{
		Height: 100,
		Hash:   "0xabc123def456",
		Data:   "persistence-test-data",
	}
	err := SaveBlockToFile(block, "block_data.json")
	if err != nil {
		fmt.Println("保存失败：", err)
		return
	}
	fmt.Println("区块已保存至文件")

	loadBlock, _ := LoadBlockFromFile("block_data.json")
	fmt.Printf("读取区块高度：%d | 哈希：%s\n", loadBlock.Height, loadBlock.Hash)
	os.Remove("block_data.json")
}
