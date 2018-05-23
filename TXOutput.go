package main

/*
Transaction Output
*/
type TXOutput struct {
	Value        int
	ScriptPubKey string // https://en.bitcoin.it/wiki/Script
}

func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool  {
	return out.ScriptPubKey == unlockingData
}