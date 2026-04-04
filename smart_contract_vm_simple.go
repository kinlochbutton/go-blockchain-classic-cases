package main

import (
	"errors"
	"fmt"
	"strconv"
)

type ContractVM struct {
	Storage map[string]int
}

func NewVM() *ContractVM {
	return &ContractVM{
		Storage: make(map[string]int),
	}
}

func (vm *ContractVM) Set(key string, value int) {
	vm.Storage[key] = value
}

func (vm *ContractVM) Get(key string) (int, error) {
	val, ok := vm.Storage[key]
	if !ok {
		return 0, errors.New("key not found")
	}
	return val, nil
}

func (vm *ContractVM) Add(key string, num int) int {
	val := vm.Storage[key]
	val += num
	vm.Storage[key] = val
	return val
}

func main() {
	fmt.Println("========== 简易智能合约虚拟机 ==========")
	vm := NewVM()
	vm.Set("token-total", 1000000)
	vm.Add("token-total", 50000)
	total, _ := vm.Get("token-total")
	fmt.Printf("代币总量: %d\n", total)
	fmt.Println("VM状态：运行正常 | 支持存储/加减/查询操作")
}
