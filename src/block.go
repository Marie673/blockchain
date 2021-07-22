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
	Index        int           `json:"index"`
	Name         string        `json:"name"`
	Timestamp    int64         `json:"timestamp"`
	Transaction  []Transaction `json:"transaction"`
	Nonce        int           `json:"nonce"`
	PreviousHash [32]byte      `json:"previous_hash"`
}

func PrintBlock(b Block) {
	fmt.Printf("	Block %d{\n", b.Index)
	fmt.Printf("		author       : %s\n", b.Name)
	fmt.Printf("		Timestamp    : %d\n", b.Timestamp)
	for i := 0; i < len(b.Transaction); i++ {
		fmt.Printf("		Transaction %d{\n", i)
		PrintTransaction(b.Transaction[i])
		fmt.Printf("		}\n")
	}
	fmt.Printf("		Nonce        : %d\n", b.Nonce)
	fmt.Printf("	 	PreviousHash : %x\n", b.PreviousHash)
	fmt.Printf("	}\n")
	fmt.Printf("	Hash    	     : %x\n", b.Hash())
}

func BlockToString(b *Block) string {
	transactionStr := TransactionToString(b)
	blockStr := fmt.Sprintf("{\n"+
		"    \"Index\"        : %d,\n"+
		"    \"Miner\"        : \"%s\",\n"+
		"    \"Timestamp\"    : %d,\n"+
		"    \"Transaction\"  : [\n"+
		"%s"+
		"    ]\n"+
		"    \"Nonce\"        : %d,\n"+
		"    \"PreviousHash\" : %x,\n"+
		"}\n",
		b.Index, b.Name, b.Timestamp, transactionStr, b.Nonce, b.PreviousHash)

	return blockStr
}

func CreateNewBlock(previousBlock *Block, miner string) *Block {
	newBlock := &Block{
		Index:        previousBlock.Index + 1,
		Name:         miner,
		Timestamp:    time.Now().UnixNano(),
		PreviousHash: previousBlock.Hash(),
	}
	// Proof of Workに基づくブロック生成
	// TODO いい感じのアルゴリズム考えて
	for {
		newBlock.Nonce = rand.Int()
		newBlock.Transaction = currentTransaction
		if Compare(newBlock.Hash(), limit) == 1 {
			break
		}
	}

	return newBlock
}

func AddBlock(chainName string, previous *Block, new *Block) bool {
	if !reflect.DeepEqual(previous.Hash(), new.PreviousHash) {
		fmt.Println("previous hash is not correct")
		return false
	}
	if Compare(new.Hash(), limit) != 1 {
		fmt.Println("new hash is invalid")
		return false
	}

	fileName := "Blockchain/" + chainName
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0666)
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
		Index        int           `json:"Index"`
		Name         string        `json:"Name"`
		Timestamp    int64         `json:"Timestamp"`
		Transaction  []Transaction `json:"Transaction"`
		Nonce        int           `json:"Nonce"`
		PreviousHash [32]byte      `json:"previous_hash"`
	}{
		Index:        b.Index,
		Name:         b.Name,
		Timestamp:    b.Timestamp,
		Transaction:  b.Transaction,
		Nonce:        b.Nonce,
		PreviousHash: b.PreviousHash,
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
