package main

import (
	"fmt"
	"os"
)

func main() {

	for {
		printMainMenu()

		var choice int
		fmt.Print("Select an option: ")
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			createNewChainMenu()
		case 2:
			loadChainMenu()
		case 3:
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Please select a valid option.")
		}
	}
}
