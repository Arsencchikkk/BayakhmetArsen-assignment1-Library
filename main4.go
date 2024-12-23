package main

import (
	"fmt"
	"github.com/Arsencchikkk/BayakhmetArsen-assignment1-Library/ex4"
)

func main() {
	account := &ex4.BankAccount{
		AccountNumber: "123456789",
		HolderName:    "John Doe",
		Balance:       1000.0,
	}

	for {
		fmt.Println("\nBank Account Menu:")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Get Balance")
		fmt.Println("4. Process Transactions")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var amount float64
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&amount)
			account.Deposit(amount)
		case 2:
			var amount float64
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&amount)
			account.Withdraw(amount)
		case 3:
			account.GetBalance()
		case 4:
			fmt.Println("Enter transaction amounts (positive for deposit, negative for withdrawal). Type '0' to end:")
			var transactions []float64
			for {
				var t float64
				fmt.Scan(&t)
				if t == 0 {
					break
				}
				transactions = append(transactions, t)
			}
			ex4.Transaction(account, transactions)
		case 5:
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
