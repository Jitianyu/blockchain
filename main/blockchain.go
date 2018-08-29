package main
import (
	"github.com/boltdb/bolt"
	"log"
)

const dbFile = "blockchain.db"
const blocksBuckets = "blocks"

type Blockchain struct {
	tip []byte
	db *bolt.DB
}

/**
 * 向BlockChain中增加Block
 */
func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte
	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBuckets))
		lastHash = b.Get([]byte("1"))
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	newBlock := NewBlock(data, lastHash)
	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBuckets))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}
		bc.tip = newBlock.Hash
		return nil
	})
}

type BlockchainIterator struct{
	currentHash []byte
	db *bolt.DB
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}
	return bci
}

func (i *BlockchainIterator) next() *Block {
	var block *Block
	err := i.db.View(func(tx *bolt.Tx) error{
		b := tx.Bucket([]byte(blocksBuckets))
		encodedBlock := b.Get(i.currentHash)
		block = DeSerializeBlock(encodedBlock)
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	i.currentHash = block.PrevBlockHash
	return block
}


 func NewGenesisBlock() *Block {
 	return NewBlock("Genesis Block", []byte{})
 }

/**
 * 向空的BlockChain中增加Genesis Block
 */
 func NewBlockChain() *Blockchain {
 	var tip []byte
 	db, err := bolt.Open(dbFile, 0600, nil)
 	if err != nil {
 		log.Panic(err)
	}

 	err = db.Update(func (tx *bolt.Tx) error {
 		b := tx.Bucket([]byte(blocksBuckets))
 		if b == nil {
 			genesis := NewGenesisBlock()
 			b, err := tx.CreateBucket([]byte(blocksBuckets))
 			if err != nil {
 				log.Panic(err)
			}
 			err = b.Put(genesis.Hash, genesis.Serialize())
 			err = b.Put([]byte("1"), genesis.Hash)
 			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("1"))
		}
		return nil
	})
 	bc := Blockchain{tip, db}
 	return &bc
 }
