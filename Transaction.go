package main

import (
	"fmt"
	"crypto/sha256"
	"bytes"
	"encoding/gob"
	"log"
	"encoding/hex"
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

func (tx *Transaction) Serialize() []byte {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	return encoded.Bytes()
}

/*
Check if transactions is coin base.
*/
func (tx *Transaction) IsCoinbase() bool  {
	return len(tx.Vin) == 1 && len(tx.Vin[0].Txid) == 0 && tx.Vin[0].Vout == -1
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

/*
Create new transaction
*/
func NewUTXOTransactions(from, to string, amount int, bc *BlockChain) *Transaction  {

	var inputs []TXInput
	var outputs []TXOutput

	acc, validOutputs := bc.FindSpendableOutputs(from, amount)

	if acc < amount {
		log.Panic("Error: Not enough funds.")
	}

	//build list of inputs
	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)

		if err != nil {
			log.Panicf("Not able to DecodeString:%s",txid)
		}

		for _, out := range outs {
			input := TXInput{ txID, out, from}
			inputs = append(inputs, input)
		}
	}

	//Build list of outputs
	outputs = append(outputs, TXOutput{amount, to})
	if acc > amount {
		outputs = append(outputs, TXOutput{acc - amount, from})
	}

	tx := Transaction{nil, inputs, outputs}
	tx.ID = tx.Hash()

	return &tx
}

