package main

type Node struct {
	ip    []string
	chain *Chain
}

func (c *Chain) GetBlockChain(name string, ip string) bool {
	return false
}
