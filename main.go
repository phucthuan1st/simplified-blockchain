package main

import (
	"fmt"
	"os"
	"simplified-blockchain/cmd"
)

func main() {
	for {
		cmd.PrintMainMenu()

		var choice int
		fmt.Print("Select an option: ")
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			cmd.CreateNewChainMenu()
		case 2:
			cmd.LoadChainMenu()
		case 3:
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Please select a valid option.")
		}
	}
}
