package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	
    "encoding/binary"
	"math/big"
	"math"
)


const Difficulty = 18
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}


func NewProof(tmpBlock *Block) *ProofOfWork {
    target := big.NewInt(1)
	//// https://docs.studygolang.com/pkg/math/big/#Int.Lsh
	target.Lsh(target, uint(256 - Difficulty)) // The bits not reach would be 0s
	return &ProofOfWork{tmpBlock, target}
}


// 1. data 2. hash 3. find target
func (proofofwork *ProofOfWork) GenerateHash() (int, []byte) {
	var intHash  big.Int
	var hash [32]byte // 32 * 8
	count := 0
	
	for count < math.MaxInt64 {

        countBuffer := new(bytes.Buffer)
	    binary.Write(countBuffer, binary.BigEndian, int64(count))
		diffiBuffer := new(bytes.Buffer)
	    binary.Write(diffiBuffer, binary.BigEndian, int64(Difficulty))
		data := bytes.Join([][]byte{
			proofofwork.Block.PreviousHash,
			proofofwork.Block.Data,
            countBuffer.Bytes(),
            diffiBuffer.Bytes(),
		}, []byte{} )
		
		hash = sha256.Sum256(data) //! := Compile Pass 
		fmt.Printf("\r%x", hash)   // Flush
		intHash.SetBytes(hash[:])  // byte to big Int

		if intHash.Cmp(proofofwork.Target) == -1 { // intHash is the combination of data and nonce
			fmt.Println()
			break
		}
		count++ 
	}

	return count, hash[:]
}
func (proofofwork *ProofOfWork) ValidateHash() bool {
    var intHash big.Int

	countBuffer := new(bytes.Buffer)
	binary.Write(countBuffer, binary.BigEndian, int64(proofofwork.Block.Count))
	diffiBuffer := new(bytes.Buffer)
	binary.Write(diffiBuffer, binary.BigEndian, int64(Difficulty))
	data := bytes.Join([][]byte{
		proofofwork.Block.PreviousHash,
		proofofwork.Block.Data,
		countBuffer.Bytes(),
		diffiBuffer.Bytes(),
	}, []byte{} )

	hash := sha256.Sum256(data)
	//fmt.Printf("\r%x", hash) 
	//fmt.Println()
	intHash.SetBytes(hash[:])

	return intHash.Cmp(proofofwork.Target) == -1
}