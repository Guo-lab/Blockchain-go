package blockchain


//@ https://pkg.go.dev/github.com/dgraph-io/badger
import (
	"fmt"

	"github.com/dgraph-io/badger" 
) //@ go get github.com/dgraph-io/badger
const (
	dbPath = "./tmp/blocks"
)





type BlockChain struct {
	//Blocks []*Block //@ array blocks in memory
	LastHash []byte
	Database *badger.DB //@ store a pointer to database
}


//@ After BlockChain refactoring
func InitChain() *BlockChain { //@ ignore err not used
//    return &BlockChain{ []*Block{ Genesis() } } 
    //@ init
	var lastHash []byte
	opt := badger.DefaultOptions(dbPath)
	opt.Dir, opt.ValueDir = dbPath, dbPath

	db, _ := badger.Open(opt)

	//@ Check if exists 
	db.Update(func(txn *badger.Txn) error {
        if _, err := txn.Get([]byte("lastHash")); err == badger.ErrKeyNotFound {
			fmt.Printf("BlockChain NOT EXISTS\n")
			newOne := Genesis()
			fmt.Printf("Genesis PROVED\n")	
			
			txn.Set(newOne.Hash, newOne.Serialize())	
			txn.Set([]byte("lastHash"), newOne.Hash)
			lastHash = newOne.Hash
			return nil	
		} else {
			item, _ := txn.Get([]byte("lastHash"))
			lastHash, _ = item.ValueCopy(nil) //@ https://pkg.go.dev/github.com/dgraph-io/badger#section-readme
			return nil
		}
	}) //@ https://dgraph.io/badger

	return &BlockChain{lastHash, db}
}

//@ Add block, traverse all first
func (originalChain *BlockChain) AddBlock(data string) {
//    new := GenerateBlock(data, originalChain.Blocks[ len(originalChain.Blocks) - 1 ].Hash)
//    originalChain.Blocks = append(originalChain.Blocks, new)
    var lastHash []byte
	originalChain.Database.View(func(txn *badger.Txn) error {
		item, _ := txn.Get([]byte("lastHash"))
		lastHash, _ = item.ValueCopy(nil) 
		return nil
	})
    newBlock := GenerateBlock(data, lastHash)    
	
	originalChain.Database.Update(func(txn *badger.Txn) error {
        txn.Set(newBlock.Hash, newBlock.Serialize())
		txn.Set([]byte("lastHash"), newBlock.Hash)
		originalChain.LastHash = newBlock.Hash
		return nil
	})
}




//@ for inverse traversal
type ChainIterator struct {
	CurrentHash []byte
	Database *badger.DB
}


func (chain *BlockChain) GenerateIterator() *ChainIterator {
	return &ChainIterator{chain.LastHash, chain.Database}
}
func (iterator *ChainIterator) Back() *Block {
	var rawData []byte
	var popBlock *Block

	iterator.Database.View(func(txn *badger.Txn) error {
		item, _ := txn.Get(iterator.CurrentHash)
        rawData, _ = item.ValueCopy(nil)
        return nil
	})
	popBlock = DeSerialize(rawData)
	
    iterator.CurrentHash = popBlock.PreviousHash
    return popBlock
}