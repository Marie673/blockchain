package main

import (
	"fmt"
)

type Transaction struct {
	sender    string
	recipient string
	amount    int
}

func PrintTransaction(t Transaction) {
	fmt.Printf("			sender    : %s\n", t.sender)
	fmt.Printf("			recipient : %s\n", t.recipient)
	fmt.Printf("			amount    : %d\n", t.amount)
}

var currentTransaction []Transaction

func InitializeTransaction() {
	currentTransaction = []Transaction{}
}

func AddTransaction(new Transaction) {
	currentTransaction = append(currentTransaction, new)
}
