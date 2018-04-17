package main

/*
Base type to to contains block
*/
type BlockChain struct {
	Blocks []*Block
}

/*
Add new Blocks into BlockChain
*/
func (bc *BlockChain) AddBlock(data string) {
	prevBlock 	:= bc.Blocks[len(bc.Blocks) - 1]
	newBlock 	:= NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

/*
Add Genesis-block (first block) into BlockChain
*/
func NewGenesisBlock() *Block {
	return  NewBlock("Genesis Blocks", []byte{})
}

/*
Create BlockChain with Genesis-Blocks
*/
func NewBlockChain() *BlockChain  {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}