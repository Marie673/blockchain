package main

import (
	"fmt"
	"reflect"
	"time"
)

type Blockchain struct {
	author string
	block  []*Block
}

func (bc *Blockchain) PrintBlockchain() {
	fmt.Printf("Blockchain author : %s{\n", bc.author)
	for i := 0; i < len(bc.block); i++ {
		PrintBlock(*bc.block[i])
		// fmt.Printf("	Hashed : %x\n", bc.block[i].Hash())
	}
	fmt.Printf("}\n\n")
}

func CreateInitialBlock(name string) *Blockchain {
	InitializeTransaction()
	b := &Block{
		index:        0,
		name:         name,
		timestamp:    time.Now().UnixNano(),
		transaction:  currentTransaction,
		proof:        1,
		previousHash: [32]byte{},
	}
	bc := &Blockchain{}
	bc.author = name
	bc.block = append(bc.block, b)

	return bc
}

func (bc *Blockchain) AddBlock() {
	newBlock := bc.ProofOfWork()
	newBlock.index = len(bc.block)
	bc.block = append(bc.block, newBlock)
}

func (bc *Blockchain) CheckHash() bool {
	for i := 0; i < len(bc.block)-1; i++ {
		before := bc.block[i].Hash()
		after := bc.block[i+1].previousHash
		if !reflect.DeepEqual(after, before) {
			// if after != before {
			return false
		}
	}

	for i := 0; i < len(bc.block); i++ {
		limit := makeByte(diff)
		nowBlockHash := bc.block[i].Hash()
		if compare(nowBlockHash, limit) != 1 {
			return false
		}
	}
	return true
}

func Consensus() []*Block {
	chainLength := make([]int, len(nodes))
	for i := 0; i < len(nodes); i++ {
		if nodes[i].CheckHash() {
			chainLength[i] = len(nodes[i].block)
		} else {
			chainLength[i] = -1
		}
	}
	maxLength := 0
	maxIndex := 0
	for i := 0; i < len(chainLength); i++ {
		if chainLength[i] > maxLength {
			maxLength = chainLength[i]
			maxIndex = i
		}
	}
	ansChain := make([]*Block, maxLength)
	ansChain = nodes[maxIndex].block

	// consensusがとれる度に出力
	consensusChain.block = nodes[maxIndex].block
	consensusChain.PrintBlockchain()

	return ansChain
}
