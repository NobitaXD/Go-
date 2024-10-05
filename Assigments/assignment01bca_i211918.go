package assignment01bca_i211918

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

var Blockchain []Block

func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	block.Hash = CreateHash(block)
	return &block
}

func ListBlocks() {
	for i, block := range Blockchain {
		fmt.Printf("Block: %d\n", i+1)
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Current Hash:  %s\n", block.Hash)
		fmt.Println(strings.Repeat("-", 50))
	}
}

func ChangeBlock(index int, newTransaction string) {
	if index >= 0 && index < len(Blockchain) {
		Blockchain[index].Transaction = newTransaction
		Blockchain[index].Hash = CreateHash(Blockchain[index])
	} else {
		fmt.Println("Invalid Block Index. ")
	}
}

func VerifyChain() bool {
	for i := 1; i < len(Blockchain); i++ {
		if Blockchain[i].PreviousHash != Blockchain[i-1].Hash {
			fmt.Printf("Blockchain invalid at block %d\n", i)
			return false
		}
	}
	fmt.Println("Blockchain is valid. ")
	return true
}

func CalculateHash(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

func CreateHash(block Block) string {
	blockData := block.Transaction + fmt.Sprintf("%d", block.Nonce) + block.PreviousHash
	return CalculateHash(blockData)
}
