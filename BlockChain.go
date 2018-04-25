package main

import (
	"github.com/boltdb/bolt"
	"log"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"
const genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

/*
Base type to to contains block
*/
type BlockChain struct {
	tip []byte
	db  *bolt.DB
}

/*
Creates a new BlockChain DB
*/
func CreateBlockChain(address string) *BlockChain {

	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	err := db.Update(func(tx *bolt.Tx) error {
		cbtx := NewCoinbaseTX(address, genesisCoinbaseData)
		genesis := NewGenesisBlock(cbtx)

		b, err := tx.CreateBucket([]byte{blocksBucket})
		err = b.Put(genesis.Hash, genesis.Serialize())

		return err
	})
}

/*
Add new Blocks into BlockChain
*/
func (bc *BlockChain) AddBlock(data string) {

	var lastHash []byte

	err := bc.db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	newBlock 	:= NewBlock(data, lastHash)


	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Fatal(err)
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Fatal(err)
		}

		bc.tip = newBlock.Hash

		return nil
	})
}

/*
Create BlockChain with Genesis-Blocks
*/
func NewBlockChain() *BlockChain  {

	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}

			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panic(err)
			}

			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				log.Panic(err)
			}
			tip  = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}
	bc := BlockChain{tip, db }

	return &bc
}


type BlockChainIterator struct {
	currentHash []byte
	db *bolt.DB
}

func(bc *BlockChain) Iterator() *BlockChainIterator {
	bci := &BlockChainIterator{bc.tip, bc.db}

	return bci
}

func(i *BlockChainIterator) Next() *Block {

	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = Deserialize(encodedBlock)

		return nil
	})
	if err != nil {
		log.Panic(err)
	}


	i.currentHash = block.PrevBlockHash

	return block
}