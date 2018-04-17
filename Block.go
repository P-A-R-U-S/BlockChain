package main

import (
	"strconv"
	"bytes"
	"crypto/sha256"
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
}

/*
	Create Hash for existing Block
*/
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PreBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

