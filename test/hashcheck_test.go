packege main

import {
  "testing"
}

func HashcheckOfAfterAddingTransaction(t *testing.T) {
  AddTransaction()
  block := &Block{
    transaction: currentTransaction
  }
  hash1 := block.Hash()
  InitializeTransaction()
  AddTransaction()
  block.transaction = currentTransaction
  hash2 := block.Hash()
  
  if hash1 == hash2 {
    t.Fatal("faled :hash is not changed")
  }
  
}
