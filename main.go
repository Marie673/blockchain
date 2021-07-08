package main

func main() {
	bc := CreateInitialBlock("okazaki")
	{
		t := Transaction{
			sender:    "sender_test1",
			recipient: "recipient_test1",
			amount:    5,
		}
		AddTransaction(t)
		bc.AddBlock()
	}

	{
		t := Transaction{
			sender:    "sender_test2",
			recipient: "recipient_test2",
			amount:    6,
		}
		AddTransaction(t)
		bc.AddBlock()
	}

	bc.PrintBlockchain()
}
