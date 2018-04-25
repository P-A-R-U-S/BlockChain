package main

import (
	"time"
	"bytes"
	"encoding/gob"
	"log"
)

/*
Base type for the BlockChain contains initial information for the block
(version, date or timestamp, hash of current and previous blocks.

Note: 	By bitcoin spec extract Timestamp, PrevBlockHash, Hash into separate struct.
		For simplification we keep it as part of the current struct.
*/
type Block struct {
	Timestamp     	int64  // date of creation
	Transactions	[]*Transaction // transactions
	PrevBlockHash 	[]byte // previous block hash
	Hash          	[]byte // current block hash
	Nonce         	int    // counter
}


/*
Constructor of Blocks
*/
func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{
		time.Now().Unix(),
		transactions,
		prevBlockHash,
		[]byte{},
		0 }

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

/*
Add Genesis-block (first block) into BlockChain
*/
func NewGenesisBlock(coinbase *Transaction) *Block {
	return  NewBlock([]*Transaction{coinbase}, []byte{})
}

/*
Serialization Block to byte array
*/
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Printf("Error serializing block: %s", err)
	}
	return result.Bytes()
}

/*
Deserialization byte array to block
*/
func Deserialize(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))

	err := decoder.Decode(&block)
	if err != nil {
		log.Printf("Error serializing block: %s", err)
	}
	return &block
}