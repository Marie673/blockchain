package main

import (
	"fmt"
)

type Transaction struct {
	Id        int    `json:"id"`
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Amount    int    `json:"amount"`
}

func PrintTransaction(t Transaction) {
	fmt.Printf("			id        : %d\n", t.Id)
	fmt.Printf("			sender    : %s\n", t.Sender)
	fmt.Printf("			recipient : %s\n", t.Recipient)
	fmt.Printf("			amount    : %d\n", t.Amount)
}

func TransactionToString(b *Block) string {
	var str string
	for i := 0; i < len(b.Transaction); i++ {
		strBuf := fmt.Sprintf("    {\n"+
			"        \"Id\"        : %d,\n"+
			"        \"Sender\"    : \"%s\",\n"+
			"        \"Recipient\" : \"%s\",\n"+
			"        \"Amount\"    : %d,\n"+
			"	}\n",
			b.Transaction[i].Id, b.Transaction[i].Sender, b.Transaction[i].Recipient, b.Transaction[i].Amount)
		str = str + strBuf
	}

	return str
}

var currentTransaction []Transaction

func InitializeTransaction() {
	currentTransaction = []Transaction{}
}

func AddTransaction(new Transaction) {
	currentTransaction = append(currentTransaction, new)
}
