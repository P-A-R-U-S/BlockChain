package main

import "fmt"

func main() {
	bc := NewBlockChain()

	bc.AddBlock("Sent 1 BTC to John")
	bc.AddBlock("Sent 2 BTC to Peter from John")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PreBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
