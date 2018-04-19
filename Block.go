package main

import (
	"time"
)

/*
Base type for the BlockChain contains initial information for the block
(version, date or timestamp, hash of current and previous blocks.

Note: 	By bitcoin spec extract Timestamp, PreBlockHash, Hash into separate struct.
		For simplification we keep it as part of the current struct.
*/
type Block struct {
	Timestamp 		int64 	// date of creation
	Data 			[]byte	// version
	PreBlockHash	[]byte	// previous block hash
	Hash			[]byte	// current block hash
	Nonce			int 	// counter
}


/*
Constructor of Blocks
*/
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		time.Now().Unix(),
		[]byte(data),
		prevBlockHash,
		[]byte{},
		0 }

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}