package main

const diff = 240

func Mining(bc *Blockchain) {
	bc.AddBlock()
}

func main() {
	bc := CreateInitialBlock("okazaki")
	for i := 0; i < 5; i++ {
		Mining(bc)
	}

	bc.PrintBlockchain()
}
