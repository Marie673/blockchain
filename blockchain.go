package main

import (
	"fmt"
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
	}
	fmt.Printf("}\n")
}

func CreateInitialBlock(name string) *Blockchain {
	InitializeTransaction()
	b := &Block{
		index:        0,
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
	ph := bc.block[len(bc.block)-1].Hash()
	b := NewBlock(ph)
	b.index = len(bc.block)
	//b.proof =
	bc.block = append(bc.block, b)

}
