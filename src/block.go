package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"reflect"
	"time"
)

type Block struct {
	index        int
	name         string
	timestamp    int64
	transaction  []Transaction
  	nonce        int
	previousHash [32]byte
}

func PrintBlock(b Block) {
	fmt.Printf("	Block %d{\n", b.index)
	fmt.Printf("		author       : %s\n", b.name)
	fmt.Printf("		timestamp    : %d\n", b.timestamp)
	for i := 0; i < len(b.transaction); i++ {
		fmt.Printf("		transaction %d{\n", i)
		PrintTransaction(b.transaction[i])
		fmt.Printf("		}\n")
	}
	fmt.Printf("		nonce        : %d\n", b.nonce)
	fmt.Printf("	 	previousHash : %x\n", b.previousHash)
	fmt.Printf("	}\n")
	fmt.Printf("	Hash    	     : %x\n", b.Hash())
}

func BlockToString(b *Block) string {
	transactionStr := TransactionToString(b)
	blockStr := fmt.Sprintf("Block %d {\n"+
		"	Miner        : %s\n"+
		"	Timestamp    : %d\n"+
		"%s"+
		"	Nonce        : %d\n"+
		"	previousHash : %x\n"+
		"}\n",
		b.index, b.name, b.timestamp, transactionStr, b.nonce, b.previousHash)
	hash := fmt.Sprintf("Hash : %x\n", b.Hash())
	blockStr += hash

	return blockStr
}

func CreateInitialBlock(miner string) *Block {
	var initialBlock = &Block{
		index:        0,
		name:         miner,
		timestamp:    time.Now().UnixNano(),
		transaction:  nil,
		nonce:        0,
		previousHash: [32]byte{},
	}
	return initialBlock
}

func CreateNewBlock(previousBlock *Block, miner string) *Block {
	newBlock := &Block{
		index:        previousBlock.index + 1,
		name:         miner,
		timestamp:    time.Now().UnixNano(),
		previousHash: previousBlock.Hash(),
	}
	for {
		newBlock.nonce = rand.Int()
		newBlock.transaction = currentTransaction
		if Compare(newBlock.Hash(), limit) == 1 {
			break
		}
	}

	return newBlock
}

func AddBlock(chainName string, previous *Block, new *Block) bool {
	if !reflect.DeepEqual(previous.Hash(), new.previousHash) {
		fmt.Println("previous hash is not correct")
		return false
	}
	if Compare(new.Hash(), limit) != 1 {
		fmt.Println("new hash is invalid")
		return false
	}

	fileName := "Blockchain/" + chainName
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("file open error")
		return false
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	_, err = fmt.Fprintf(file, "%s\n", BlockToString(new))
	if err != nil {
		fmt.Println("file write error")
		return false
	}

	return true
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
		Name         string        `json:"name"`
		Timestamp    int64         `json:"timestamp"`
		Transaction  []Transaction `json:"transaction"`
		Nonce        int           `json:"nonce"`
		PreviousHash [32]byte      `json:"previous_hash"`
	}{
		Index:        b.index,
		Name:         b.name,
		Timestamp:    b.timestamp,
		Transaction:  b.transaction,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
	})
}

func Compare(a [32]byte, b []byte) int {
	check := make([]byte, 32)
	for i := 0; i < 32; i++ {
		check[i] = a[i]
	}
	return bytes.Compare(b, check)
}

func makeBytePoW(diff int) []byte {
	ans := make([]byte, 32)
	if diff >= 256 {
		ans[0] = byte(255)
		return ans
	}
	remind := diff % 8
	quot := diff / 8
	check := int(math.Pow(2, float64(remind)))
	ans[31-quot] = byte(check)

	return ans
}
