package main

import (
	"strings"
	"testing"
	"time"
)

func TestCalculateHash(t *testing.T) {
	block := Block{
		Index:     1,
		Timestamp: time.Now().String(),
		Data:      "Test Block",
		PrevHash:  "previous-hash",
	}

	hash := calculateHash(block)

	if len(hash) != 64 {
		t.Errorf("Expected hash length of 64, got %d", len(hash))
	}

	// Test that the same block always produces the same hash
	hash2 := calculateHash(block)
	if hash != hash2 {
		t.Errorf("Hash calculation is not deterministic")
	}
}

func TestCreateBlock(t *testing.T) {
	index := 1
	data := "Test Data"
	prevHash := "previous-hash"

	block := createBlock(index, data, prevHash)

	if block.Index != index {
		t.Errorf("Expected Index %d, got %d", index, block.Index)
	}

	if block.Data != data {
		t.Errorf("Expected Data %s, got %s", data, block.Data)
	}

	if block.PrevHash != prevHash {
		t.Errorf("Expected PrevHash %s, got %s", prevHash, block.PrevHash)
	}

	if len(block.Hash) != 64 {
		t.Errorf("Expected hash length of 64, got %d", len(block.Hash))
	}

	if !strings.Contains(block.Timestamp, time.Now().Format("2006")) {
		t.Errorf("Timestamp %s doesn't contain current year", block.Timestamp)
	}
}

func TestBlockchain(t *testing.T) {
	blockchain := []Block{
		createBlock(0, "Genesis Block", ""),
	}

	if len(blockchain) != 1 {
		t.Errorf("Expected blockchain length of 1, got %d", len(blockchain))
	}

	blockchain = append(blockchain, createBlock(1, "Second Block", blockchain[0].Hash))

	if len(blockchain) != 2 {
		t.Errorf("Expected blockchain length of 2, got %d", len(blockchain))
	}

	if blockchain[1].PrevHash != blockchain[0].Hash {
		t.Errorf("Second block's PrevHash doesn't match first block's Hash")
	}
}
