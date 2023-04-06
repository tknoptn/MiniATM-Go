package main

import (
	"fmt"
)

type Account struct {
	id      string
	name    string
	balance float64
	pin     string
}

func (a *Account) withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("invalid withdrawal amount")
	}
	if amount > a.balance {
		return fmt.Errorf("insufficient funds")
	}
	a.balance -= amount
	if a.balance < 100 {
		fmt.Println("Your balance is low, Maintain minimum balance to avoid charges")
	}
	return nil
}

func (a *Account) deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("invalid deposit amount")
	}
	a.balance += amount
	return nil
}

func (a *Account) checkBalance() float64 {
	return a.balance
}

func (a *Account) changePin(pin string) error {
	if len(pin) != 4 {
		return fmt.Errorf("invalid PIN")
	}
	a.pin = pin
	return nil
}

func (a *Account) authenticate(pin string) bool {
	return a.pin == pin
}

func main() {
	// create an account
	account := &Account{id: "123456789", name: "John Doe", balance: 500.0, pin: "1234"}

	// authenticate the user
	var pin string
	fmt.Print("Enter your PIN: ")
	fmt.Scanln(&pin)
	if !account.authenticate(pin) {
		fmt.Println("Invalid PIN")
		return
	}

	// display the account balance and menu options
	fmt.Printf("Account balance: $%.2f\n", account.checkBalance())
	fmt.Println("Menu:")
	fmt.Println("1. Withdraw")
	fmt.Println("2. Deposit")
	fmt.Println("3. Change PIN")
	fmt.Println("4. Exit")

	// process user input
	var choice int
	for {
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			var amount float64
			fmt.Print("Enter amount to withdraw: $")
			fmt.Scanln(&amount)
			if err := account.withdraw(amount); err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Withdrawal of $%.2f successful\n", amount)
				fmt.Printf("New balance: $%.2f\n", account.checkBalance())
			}
		case 2:
			var amount float64
			fmt.Print("Enter amount to deposit: $")
			fmt.Scanln(&amount)
			if err := account.deposit(amount); err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Deposit of $%.2f successful\n", amount)
				fmt.Printf("New balance: $%.2f\n", account.checkBalance())
			}
		case 3:
			var pin string
			fmt.Print("Enter new PIN (4 digits): ")
			fmt.Scanln(&pin)
			if err := account.changePin(pin); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("PIN changed successfully")
			}
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
