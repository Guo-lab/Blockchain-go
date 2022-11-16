package main

// using the bytes library
import (
	"bytes"
    "crypto/sha256"
	"fmt"
)

// var identifier []type
type Block struct {
	Hash []byte
    Data []byte
	previousHash []byte
}

type blockChain struct {
	blocks []*Block
}


// func function_name( [parameter list] ) [return_types] {}

// Join 
//// https://www.geeksforgeeks.org/how-to-join-the-elements-of-the-byte-slice-in-golang/#:~:text=In%20the%20Go%20slice%20of%20bytes%2C%20you%20are,these%20joined%20elements%20separated%20by%20the%20given%20separator.
// Sum256
//// https://cloud.tencent.com/developer/section/1140766
func (tmpBlock *Block) generateHash() { // pointer passing
    info := bytes.Join([][]byte{tmpBlock.Data, tmpBlock.previousHash}, []byte{})
    sha  := sha256.Sum256(info) // return [Size]byte
	tmpBlock.Hash = sha[:] 
}
func generateBlock (data string, previousHash []byte) *Block {
	new := &Block{[]byte{}, []byte(data), previousHash} // allocate
	new.generateHash()
	return new
}
func (originalChain *blockChain) addBlock(data string) {
    new := generateBlock(data, originalChain.blocks[len(originalChain.blocks)-1].Hash)
	originalChain.blocks = append(originalChain.blocks, new) // append for []*Block
}

// for Initialization
func Genesis() *Block {
	return generateBlock("Block Head", []byte{})
}
func InitChain() *blockChain {
    return &blockChain{ []*Block{ Genesis() } } // Generate Chain
}



func main () {
    GlobalChain := InitChain()
	GlobalChain.addBlock("First Block after Genesis")
	GlobalChain.addBlock("Second Block after First")
	GlobalChain.addBlock("Third Block after Second")
    // Range
	//// https://www.runoob.com/go/go-range.html
	for _, block := range GlobalChain.blocks {
		fmt.Printf("###################################\n",)
        //fmt.Printf("%-19s: %x\n", "Last Block Hash", block.previousHash)
		fmt.Printf("%-19s: %s\n", "Block Data", block.Data) // s stands for //// https://www.cnblogs.com/xwxz/p/14498237.html
		fmt.Printf("%-19s: %x\n", "Current Block Hash", block.Hash)
		fmt.Printf("-----------------------------------\n\n",)		
	}
}

// Output
/*
###################################
Last Block Hash    : 
Block Data         : Block Head
Current Block Hash : 164047fe99c4a1f534ef73528533b11118de7bd77897751cf638dbc72824002c
-----------------------------------
###################################
Last Block Hash    : 164047fe99c4a1f534ef73528533b11118de7bd77897751cf638dbc72824002c
Block Data         : First Block after Genesis
Current Block Hash : 87167d6fc72aad7296baf65c0d430742c2e9a8cd371064eb372962635ab31680
-----------------------------------
###################################
Last Block Hash    : 87167d6fc72aad7296baf65c0d430742c2e9a8cd371064eb372962635ab31680
Block Data         : Second Block after First
Current Block Hash : b7d74545d0187d5ca92582021371ace01ede34f5821fca4cd98f208e299944b5
-----------------------------------
###################################
Last Block Hash    : b7d74545d0187d5ca92582021371ace01ede34f5821fca4cd98f208e299944b5
Block Data         : Third Block after Second
Current Block Hash : 18cb77f41cb3853014012c150ea262e52fd806cc45a578ffd964363b261a80a0
-----------------------------------
*/