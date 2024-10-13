package main

import (
	"fmt"
	"github.com/rchelbii/tock/internal/chain"
)

func main() {
	bc := chain.NewBlockChain()
	bc.AddBlock("Send 1 tock to rchelbi")
	bc.AddBlock("Send 4 tock to alpha")

	for _, block := range bc.Blocks() {
		fmt.Printf("previous hash \t %x \n", block.PrevBlockHash)
		fmt.Printf("current hash \t %x \n", block.Hash)
		fmt.Printf("data \t %s \n", block.Data)
		fmt.Println("-------------------------")
	}
}
