package main

import (
	"fmt"
	"crypto/sha256"
	"bytes"
	"encoding/gob"
	"log"
)

const subsidy = 10

/*
BlockChain transaction
*/
type Transaction struct {
	ID 		[]byte
	Vin		[]TXInput
	Vout	[]TXOutput
}


// Hash returns the hash of the Transaction
func (tx *Transaction) Hash() []byte {
	var hash [32]byte

	txCopy := *tx
	txCopy.ID = []byte{}

	hash = sha256.Sum256(txCopy.Serialize())

	return hash[:]
}


func (tx Transaction) Serialize() []byte {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	return encoded.Bytes()
}

/*
Transaction Input
*/
type TXInput struct {
	Txid 		[]byte
	Vout		int
	ScriptSig	string
}

/*
Transaction Output
*/
type TXOutput struct {
	Value			int
	ScriptPubKey	string // https://en.bitcoin.it/wiki/Script
}

/*
Coinbase Transaction constructor
*/
func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to %s", to)
	}

	txin := TXInput{[]byte{}, -1, data}
	txout := TXOutput{subsidy, to}
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.ID = tx.Hash()

	return &tx
}

