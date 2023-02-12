package main

import (
	"fmt"
	"golang/blockchain"
	"time"
)

func main() {
	blockchain := blockchain.CreateBlockChain()

	time.Sleep(time.Second) //休息個幾秒再新增區塊
	blockchain.AddBlock("第一筆資料")

	for _, block := range blockchain.Blocks {
		fmt.Println("==============================================================================")
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("Nonce: %d\n", block.Nonce)

	}

}
