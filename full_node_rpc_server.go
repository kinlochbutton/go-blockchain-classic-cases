package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type ChainRPC int

type BlockArgs struct {
	Height int
}

type BlockReply struct {
	Hash string
	Data string
}

func (c *ChainRPC) GetBlock(args *BlockArgs, reply *BlockReply) error {
	reply.Hash = fmt.Sprintf("hash-height-%d", args.Height)
	reply.Data = fmt.Sprintf("block-data-height-%d", args.Height)
	return nil
}

func main() {
	fmt.Println("========== 全节点RPC服务 ==========")
	chain := new(ChainRPC)
	_ = rpc.Register(chain)
	listener, _ := net.Listen("tcp", ":1234")
	fmt.Println("RPC服务已启动 :1234 | 支持GetBlock接口")
	go func() {
		for {
			conn, _ := listener.Accept()
			go rpc.ServeConn(conn)
		}
	}()
	select {}
}
