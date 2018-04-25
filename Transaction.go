package main

import "fmt"

const subsidy = 10

/*
BlockChain transaction
*/
type Transaction struct {
	ID 		[]byte
	Vin		[]TXInput
	Vout	[]TXOutput
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
	tx.SetID()

	return &tx
}