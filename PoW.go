package main

import (
	"bytes"
	"math"
	"math/rand"
)

func (b *Block) ProofOfWork() *Block {
	var newHash [32]byte
	limit := makeByte(diff)
	newBlock := &Block{}
	for {
		newBlock = NewBlock(b.previousHash, rand.Int(), b.name)
		newHash = newBlock.Hash()
		if compare(newHash, limit) == 1 {
			break
		}
	}
	InitializeTransaction()

	return newBlock
}

func compare(a [32]byte, b []byte) int {
	check := make([]byte, 32)
	for i := 0; i < 32; i++ {
		check[i] = a[i]
	}
	return bytes.Compare(b, check)
}

func makeByte(diff int) []byte {
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
