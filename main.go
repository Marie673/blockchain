package main

import (
	"time"
)

const diff = 235

const nodeNUM = 5

var names = [nodeNUM]string{"okazaki", "yude", "yamakasu", "str4w", "mimimi"}
var nodes [nodeNUM]*Blockchain

func InitialNodes() {
	for i := 0; i < nodeNUM; i++ {
		nodes[i] = CreateInitialBlock(names[i])
		// nodes[i].PrintBlockchain()
	}
}

func Mining(index int) {
	nodes[index].AddBlock()
	nodes[index].block = Consensus()
}

func main() {
	InitialNodes()
	go func() {
		for {
			Mining(0)
		}
	}()
	go func() {
		for {
			Mining(1)
		}
	}()
	go func() {
		for {
			Mining(2)
		}
	}()
	go func() {
		for {
			Mining(3)
		}
	}()
	go func() {
		for {
			Mining(4)
		}
	}()

	go func() {
		for {
			time.Sleep(5 * time.Second)
			for i := 0; i < nodeNUM; i++ {
				nodes[i].PrintBlockchain()
			}
		}
	}()

	time.Sleep(300 * time.Second)
	/*
		file, err := os.Create(`result\\mining.txt`)
		if err != nil {
			panic(err)
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)

		ls := Consensus()
		for i := 0; i < len(ls); i++ {
			m, err := json.Marshal(ls[i])
			if err != nil {
				panic(err)
			}
			output := string(m)
			output += "\r\n"
			_, err = file.Write(([]byte)(output))
			if err != nil {
				return
			}
		}
	*/
}
