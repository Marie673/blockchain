package main

import "fmt"

type Chain struct {
	author        string
	chainName     string
	previousBlock *Block
	newBlock      *Block
}

func Mining(chainName string, author string) {
	// 原初のブロック作成
	chain := &Chain{
		author:        author,
		chainName:     chainName,
		previousBlock: nil,
		newBlock:      CreateInitialBlock(author),
	}
	// それに連なるブロック作成
	for {
		chain.previousBlock = chain.newBlock
		chain.newBlock = CreateNewBlock(chain.previousBlock, chain.author)
		res := AddBlock(chain.chainName, chain.previousBlock, chain.newBlock)
		if res {
			fmt.Println("new block is added")
		} else {
			fmt.Println("failed add block")
		}
	}
}
