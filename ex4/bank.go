package ex4

import (
	"fmt"
)

type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

func (ba *BankAccount) Deposit(amount float64) {
	ba.Balance += amount
	fmt.Printf("Deposited %.2f. New Balance: %.2f\n", amount, ba.Balance)
}

func (ba *BankAccount) Withdraw(amount float64) {
	if ba.Balance < amount {
		fmt.Println("Insufficient balance.")
		return
	}
	ba.Balance -= amount
	fmt.Printf("Withdrawn %.2f. New Balance: %.2f\n", amount, ba.Balance)
}

func (ba *BankAccount) GetBalance() {
	fmt.Printf("Current Balance: %.2f\n", ba.Balance)
}

func Transaction(account *BankAccount, transactions []float64) {
	for _, t := range transactions {
		if t > 0 {
			account.Deposit(t)
		} else {
			account.Withdraw(-t)
		}
	}
}
