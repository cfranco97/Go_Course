package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFile = "balance.txt"

func main() {

	var balance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR:")
		fmt.Println(err)
		fmt.Println("-------")
	}
	fmt.Println("Hello user, welcome to the Go bank!")
	//loop
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")
		var choice int

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your balance: %.2f\n", balance)
		case 2:
			var deposit float64
			fmt.Print("Enter the amount you wish to deposit:")
			fmt.Scan(&deposit)
			balance += deposit
			writeBalanceToFile(balance)
			fmt.Printf("Your updated balance: %.2f\n", balance)
			continue
		case 3:
			var withdraw float64
			fmt.Print("Enter the amount you wish to withdraw:")
			fmt.Scan(&withdraw)
			if (balance - withdraw) < 0 {
				fmt.Println("Error, balance cannot be negative!")
				continue
			}
			balance -= withdraw
			writeBalanceToFile(balance)
			fmt.Printf("Your updated balance: %.2f\n", balance)
			continue
		default:
			fmt.Println("goodbye!")
			return
		}

	}
}

func getBalanceFromFile() (float64, error) {
	file, err := os.ReadFile("balance.txt")
	if err != nil {
		return 0, errors.New("Failed to find balance file.")
	}

	balanceText := string(file)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 0, errors.New("failed to parse stored balance value.")
	}

	// returns balance value and error, in this case error should be null when everything goes right.
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(balanceFile, []byte(balanceText), 0644)
}
