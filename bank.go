package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFile = "balance.txt"

func readFromFile() (float64, error) {
	data, err := os.ReadFile(balanceFile)
	if err != nil {
		return 0.0, errors.New("could not read balance file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 0.0, errors.New("could not parse balance")
	}
	return balance, nil
}
func writeToFile(balance float64) {
	balanceText := fmt.Sprintf("%.2f", balance)
	os.WriteFile(balanceFile, []byte(balanceText), 0644)
}

func main() {
	accountBalance, err := readFromFile()
	if err != nil {
		fmt.Println("Error reading balance:", err)
		fmt.Println("__________________")
	}

	fmt.Println("Wellcome to GO BANK")

	for {
		fmt.Println()
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")
		fmt.Println()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)

		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New Amount:", accountBalance)
			writeToFile(accountBalance)

		case 3:
			fmt.Print("Your withdraw amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient balance.")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! New Amount:", accountBalance)
			writeToFile(accountBalance)

		case 4:
			fmt.Println("Thanks for visiting!")
			return

		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}

		/*if Choice == 1 {
					fmt.Println("Your balnce is: ", accountBalance)
				} else if Choice == 2 {
					fmt.Print("Your deposit: ")
					var depositAmount float64
					fmt.Scan(&depositAmount)
					if depositAmount <= 0 {
						fmt.Println("Invalid amount. Must be greater than 0.")
						continue
					}
					accountBalance += depositAmount
					fmt.Println("Balance updated! New Amount: ", accountBalance)
					writeToFile(accountBalance)
				} else if Choice == 3 {
					fmt.Print("Your withdraw amount: ")
					var withdrawAmount float64
					fmt.Scan(&withdrawAmount)
					if withdrawAmount > accountBalance {
						fmt.Println("Insufficient balance.")
						continue
					}
					if withdrawAmount <= 0 {
		            fmt.Println("Invalid amount.")
		   			 continue
					}
					accountBalance -= withdrawAmount
					fmt.Println("Balance updated! New Amount: ", accountBalance)
					writeToFile(accountBalance)
				} else {
					break
				}*/
	}
	fmt.Println("Thanks for Visiting!")

}
