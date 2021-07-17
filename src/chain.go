package main

import (
	"fmt"
	"os"
	"time"
)

type Chain struct {
	author        string
	chainName     string
	previousBlock *Block
	newBlock      *Block
}

func CreateInitialBlock(chainName string, author string) *Chain {
	initialBlock := &Block{
		Index:        0,
		Name:         author,
		Timestamp:    time.Now().UnixNano(),
		Transaction:  nil,
		Nonce:        0,
		PreviousHash: [32]byte{},
	}
	chain := &Chain{
		author:        author,
		chainName:     chainName,
		previousBlock: nil,
		newBlock:      initialBlock,
	}
	fileName := "Blockchain/" + chainName
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("file open error")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	_, err = fmt.Fprintf(file, "%s\n", BlockToString(initialBlock))
	if err != nil {
		fmt.Println("file write error")
	}

	return chain
}

func Mining(chainName string, author string) {
	// 原初のブロック作成
	chain := CreateInitialBlock(chainName, author)
	// それに連なるブロック作成
	for {
		preChain := chain
		chain.previousBlock = chain.newBlock
		chain.newBlock = CreateNewBlock(chain.previousBlock, chain.author)
		res := AddBlock(chain.chainName, chain.previousBlock, chain.newBlock)
		if res {
			fmt.Println("new block &d is added", chain.newBlock.Index)
		} else {
			fmt.Println("failed add block")
			chain = preChain
		}
	}
}
