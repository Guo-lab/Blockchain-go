package main

//// https://blog.csdn.net/SUZHOUTONY/article/details/90049626
//// https://blog.csdn.net/wohu1104/article/details/123209272
import (
	"fmt"
    //// https://newbedev.com/relative-imports-in-go
	// so, with git init
	"github.com/Guo-lab/Blockchain-go/blockchain" // import directory and call package
	"strconv"
)




func main () {
    GlobalChain := blockchain.InitChain()
	GlobalChain.AddBlock("First Block after Genesis")
	GlobalChain.AddBlock("Second Block after First")
	GlobalChain.AddBlock("Third Block after Second")
    
	// Range
	//// https://www.runoob.com/go/go-range.html
	for _, block := range GlobalChain.Blocks {

		fmt.Printf("###################################\n",)
        fmt.Printf("%-19s: %x\n", "Last Block Hash", block.PreviousHash)
		fmt.Printf("%-19s: %s\n", "Block Data", block.Data) // s stands for //// https://www.cnblogs.com/xwxz/p/14498237.html
		fmt.Printf("%-19s: %x\n", "Current Block Hash", block.Hash)
		fmt.Printf("-----------------------------------\n",)

		fmt.Printf("***\n",)	
		proofofwork := blockchain.NewProof(block)
		fmt.Printf("%-19s: %s\n", "Proof of Work", strconv.FormatBool(proofofwork.ValidateHash()))
		fmt.Printf("***\n\n",)	

	}
}