package main

import (
	"fmt"
)

type Transaction struct {
	id        int
	sender    string
	recipient string
	amount    int
}

func PrintTransaction(t Transaction) {
	fmt.Printf("			id        : %d\n", t.id)
	fmt.Printf("			sender    : %s\n", t.sender)
	fmt.Printf("			recipient : %s\n", t.recipient)
	fmt.Printf("			amount    : %d\n", t.amount)
}

func TransactionToString(b *Block) string {
	var str string
	for i := 0; i < len(b.transaction); i++ {
		strBuf := fmt.Sprintf("	Transaction %d {\n"+
			"		Sender    :%s\n"+
			"		Recipient : %s\n"+
			"		Amount    :%d\n"+
			"	}\n",
			b.transaction[i].id, b.transaction[i].sender, b.transaction[i].recipient, b.transaction[i].amount)
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
