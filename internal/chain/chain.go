package chain

import "github.com/rchelbii/tock/internal/block"

type BlockChain struct {
	blocks []*block.Block
}

func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := block.NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewGenesisBlock() *block.Block {
	return block.NewBlock("Genesis Block", []byte{})
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*block.Block{NewGenesisBlock()}}
}

func (bc BlockChain) Blocks() []*block.Block {
    return bc.blocks
}
