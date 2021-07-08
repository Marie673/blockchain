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
	counter := 0
	limit := makeByte(diff)
	var ph [32]byte
	b := NewBlock(ph, counter)
	for {
		b = NewBlock(ph, counter)
		ph = b.Hash()
		fmt.Printf("%x %d\n", ph, counter)
		if compare(ph, limit) == 1 {
			break
		}
		counter++
	}

	b.index = len(bc.block)
	//b.proof =
	bc.block = append(bc.block, b)

}
