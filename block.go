package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	index        int
	timestamp    int64
	transaction  []Transaction
	proof        int
	previousHash [32]byte
}

func PrintBlock(b Block) {
	fmt.Printf("	Block %d{\n", b.index)
	fmt.Printf("		timestamp    : %d\n", b.timestamp)
	for i := 0; i < len(b.transaction); i++ {
		PrintTransaction(b.transaction[i])
	}
	fmt.Printf("		proof        : %d\n", b.proof)
	fmt.Printf("	 	previousHash : %x\n", b.previousHash)
	fmt.Printf("	}\n")
}

func NewBlock(previousHash [32]byte, nonce int) *Block {
	b := &Block{}
	b.index = -1
	b.timestamp = time.Now().UnixNano()
	b.transaction = currentTransaction
	InitializeTransaction()
	b.proof = nonce
	b.previousHash = previousHash

	return b
}

func (b *Block) Hash() [32]byte {
	m, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}

	return sha256.Sum256(m)
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Index        int           `json:"index"`
		Timestamp    int64         `json:"timestamp"`
		Transaction  []Transaction `json:"transaction"`
		Proof        int           `json:"proof"`
		PreviousHash [32]byte      `json:"previous_hash"`
	}{
		Index:        b.index,
		Timestamp:    b.timestamp,
		Transaction:  b.transaction,
		Proof:        b.proof,
		PreviousHash: b.previousHash,
	})
}
