package main

import (
	"fmt"
)

func main() {
	bc := NewBlockchain()
	bc.AddBlock("send 1 BTC to Tang")
	bc.AddBlock("send 2 more BTC to Tang")
	for _, block :=range bc.blocks {
		fmt.Println("Prev. Hash :%x",block.PrevBlockHash)
		fmt.Println("Data: %s", block.Data)
		fmt.Println("Hash: %x",block.Hash)
		fmt.Println()

	}

}
