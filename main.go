package main

//@ https://blog.csdn.net/SUZHOUTONY/article/details/90049626
//@ https://blog.csdn.net/wohu1104/article/details/123209272
import (
	"fmt"
    //@ https://newbedev.com/relative-imports-in-go
	//@ so, with git init
	"github.com/Guo-lab/Blockchain-go/blockchain" //@ import directory and call package
	"strconv"

	"os"
    "runtime"
	"flag"
)



// func main () {
//  GlobalChain := blockchain.InitChain()
// 	GlobalChain.AddBlock("First Block after Genesis")
// 	GlobalChain.AddBlock("Second Block after First")
// 	GlobalChain.AddBlock("Third Block after Second")
    
// 	//@ Range
// 	//@ https://www.runoob.com/go/go-range.html
// 	for _, block := range GlobalChain.Blocks {

// 		fmt.Printf("###################################\n",)
//      fmt.Printf("%-19s: %x\n", "Last Block Hash", block.PreviousHash)
// 		fmt.Printf("%-19s: %s\n", "Block Data", block.Data) // s stands for //@ https://www.cnblogs.com/xwxz/p/14498237.html
// 		fmt.Printf("%-19s: %x\n", "Current Block Hash", block.Hash)
// 		fmt.Printf("-----------------------------------\n",)

// 		fmt.Printf("***\n",)	
// 		proofofwork := blockchain.NewProof(block)
// 		fmt.Printf("%-19s: %s\n", "Proof of Work", strconv.FormatBool(proofofwork.ValidateHash()))
// 		fmt.Printf("***\n\n",)	
// 	}
// }




//////////// Version-Update CommandLine ////////
type CommandLine struct {
	blockchain *blockchain.BlockChain //@ in blockchain.go
}

//@ three function: print help, add blocks, print the chain
func (commandline *CommandLine) printHelp() {
	fmt.Printf("%-28s%s%28s\n", "--------------------------", "H E L P", "--------------------------") 
	fmt.Printf("%-25s %-s\n", " add -block INPUT_DATA", "[add one block]")
	fmt.Printf("%-25s %-s\n", " print", "[display all blocks in the blockchain]")
	fmt.Printf("%-28s%s%28s\n", "--------------------------", "-------", "--------------------------")
}
func (commandline *CommandLine) addBlock(data string) {
    commandline.blockchain.AddBlock(data)
	fmt.Printf("Block added successfully.\n")
}
func (commandline *CommandLine) printChain() {
	iterator := commandline.blockchain.GenerateIterator()
	for {
		displayBlock := iterator.Back()
        fmt.Printf("-------------- Block Information ----------\n",)
		fmt.Printf("%-19s: %x\n", "Last Block Hash", displayBlock.PreviousHash)
		fmt.Printf("%-19s: %s\n", "Block Data", displayBlock.Data) // s stands for //@ https://www.cnblogs.com/xwxz/p/14498237.html
		fmt.Printf("%-19s: %x\n", "Current Block Hash", displayBlock.Hash)
		proofofwork := blockchain.NewProof(displayBlock)
		fmt.Printf("%-19s: %s\n", "Proof of Work", strconv.FormatBool(proofofwork.ValidateHash()))
		fmt.Printf("-------------------------------------------\n")	
		if len(displayBlock.PreviousHash) == 0 {
			break
		}
	}
}

func (commandline *CommandLine) run() {
	addCmd   := flag.NewFlagSet("ADD", flag.ExitOnError) //@ https://pkg.go.dev/flag#NewFlagSet and https://blog.csdn.net/lin_strong/article/details/120660753
	printCmd := flag.NewFlagSet("PRINT", flag.ExitOnError)
    flagData := addCmd.String("block", "", "Input Data")
	switch os.Args[1] { //@ https://pkg.go.dev/flag#Parse
	case "add":
        addCmd.Parse(os.Args[2:])
	case "print":
		printCmd.Parse(os.Args[2:])
	case "help":
		commandline.printHelp()
		runtime.Goexit()
	default:
		commandline.printHelp()
		runtime.Goexit()
	}

	//@ flags are accessed
	if addCmd.Parsed() {
		if *flagData == "" {
			addCmd.Usage()
			runtime.Goexit()
		}
		commandline.addBlock(*flagData)
		commandline.printChain()
	}
	if printCmd.Parsed() {
		commandline.printChain()
	}
}



///////////////// Main ////////////////////
func main () {
    //@ validate args
	if len(os.Args) < 2 {
		fmt.Println("Variables not enough.")
        os.Exit(-1) //@https://www.zhihu.com/question/392413832#
	}

	defer os.Exit(0) //@ https://blog.csdn.net/u013161278/article/details/117661652
	GlobalChain := blockchain.InitChain()
	defer GlobalChain.Database.Close()

    session := CommandLine{GlobalChain}
    session.run()
}