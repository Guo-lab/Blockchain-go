package blockchain

// using the bytes library
import (
	//"bytes"
    //"crypto/sha256"
	//"fmt"
)

// https://blog.csdn.net/KO_NO_JOJO/article/details/109780090

// var identifier []type
type Block struct {
	Hash []byte
    Data []byte
	PreviousHash []byte

	Count int // For ProofOfWork
}

type BlockChain struct {
	Blocks []*Block
}


// func function_name( [parameter list] ) [return_types] {}



// Join 
//// https://www.geeksforgeeks.org/how-to-join-the-elements-of-the-byte-slice-in-golang/#:~:text=In%20the%20Go%20slice%20of%20bytes%2C%20you%20are,these%20joined%20elements%20separated%20by%20the%20given%20separator.
// Sum256
//// https://cloud.tencent.com/developer/section/1140766
//////////// Version-Update GenerateHash() in proof.go ////
//func (tmpBlock *Block) GenerateHash() { // pointer passing
//    info := bytes.Join([][]byte{tmpBlock.Data, tmpBlock.PreviousHash}, []byte{})
//    sha  := sha256.Sum256(info) // return [Size]byte
//	tmpBlock.Hash = sha[:] 
//}
//
//func GenerateBlock (data string, previousHash []byte) *Block {
//	new := &Block{[]byte{}, []byte(data), previousHash} // allocate
//	new.GenerateHash()
//	return new
//}
////////////// Version-Update GenerateBlock() /////////////
func GenerateBlock (data string, previousHash []byte) *Block {
	newBlock := &Block{[]byte{}, []byte(data), previousHash, 0} // allocate
	proofofwork := NewProof(newBlock)
	
	count, hash := proofofwork.GenerateHash()
	newBlock.Hash  = hash[:]
    newBlock.Count = count
	////fmt.Printf("%x\n", hash)
	return newBlock
}
/////////////////////////////////////////////////////////////
func (originalChain *BlockChain) AddBlock(data string) {
    new := GenerateBlock(data, originalChain.Blocks[len(originalChain.Blocks)-1].Hash)
	originalChain.Blocks = append(originalChain.Blocks, new) // append for []*Block
}

// for Initialization
func Genesis() *Block {
	return GenerateBlock("Block Head", []byte{})
}
func InitChain() *BlockChain {
    return &BlockChain{ []*Block{ Genesis() } } // Generate Chain
}
