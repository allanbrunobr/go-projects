// Copyright 2024 Allan Bruno.

package main

import (
	_ "compress/flate"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

// calculateHash calculates the SHA-256 hash of a given block.
//
// The function takes a Block struct as input and returns its SHA-256 hash as a hexadecimal string.
// The hash is calculated by concatenating the block's index, timestamp, data, and previous hash,
// and then applying the SHA-256 hashing algorithm to the resulting string.
//
// Parameters:
// - block: The Block struct for which the hash is to be calculated.
//
// Returns:
// - string: The hexadecimal representation of the SHA-256 hash of the block.
func calculateHash(block Block) string {
	record := fmt.Sprintf("%d%s%s%s", block.Index, block.Timestamp, block.Data, block.PrevHash)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func createBlock(index int, data string, prevHash string) Block {
	block := Block{
		Index:     index,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevHash,
	}
	block.Hash = calculateHash(block)
	return block
}

func main() {
	// Initialize the blockchain with a genesis block
	blockchain := []Block{
		createBlock(0, "Genesis Block", ""),
	}

	// Add two additional blocks with sample data
	blockchain = append(blockchain, createBlock(1, "Hello", blockchain[0].Hash))
	blockchain = append(blockchain, createBlock(2, "World", blockchain[1].Hash))

	// Print the details of each block in the blockchain
	for _, block := range blockchain {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("PrevHash: %s\n\n", block.PrevHash)
	}
}
