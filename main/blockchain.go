package main

type BlockChain struct {
	blocks []*Block
}

/**
 * 向BlockChain中增加Block
 */
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks) - 1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

/**
 * 向空的BlockChain中增加Genesis Block
 */
 func NewGenesisBlock() *Block {
 	return NewBlock("Genesis Block", []byte{})
 }

 func NewBlockChain() *BlockChain {
 	blockChain := BlockChain{[]*Block{NewGenesisBlock()}}
 	return &blockChain
 }



