package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/phucthuan1st/simplified-blockchain/pkg"
)

func printMainMenu() {
	fmt.Printf("\x1bc")
	fmt.Println("Main Menu:")
	fmt.Println("--------------- Menu ----------------")
	fmt.Println("1. Create a new chain")
	fmt.Println("2. Load a chain from file")
	fmt.Println("3. Exit")
	fmt.Println("-------------------------------------")
}

func printSubMenu(menuTitle string) {
	fmt.Printf("%s:\n", menuTitle)
	fmt.Println("--------------- Menu ----------------")
	fmt.Println("1. Add a Block")
	fmt.Println("2. Display Chain")
	fmt.Println("3. Back to Main Menu")
	fmt.Println("-------------------------------------")
}

func createNewChainMenu() {
	running := true
	fmt.Print("Enter an identifier for the new chain: ")
	var identifier string
	fmt.Scanln(&identifier)

	blockchain := pkg.NewBlockchain(identifier)

	for running {
		fmt.Printf("\x1bc")
		printSubMenu("New Chain")

		fmt.Print("Select an option: ")
		var choice int
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			// Add a block to the blockchain
			block := createNewBlock()
			block.PrevBlockHash = blockchain.GetLastBlock().Hash
			err := blockchain.Add(block)
			if err != nil {
				log.Printf("Error adding block: %v", err)
			}
		case 2:
			pkg.DisplayBlockchain(blockchain)
		case 3:
			// Back to the main menu
			running = false
		default:
			fmt.Println("Invalid option. Please select a valid option.")
		}

		if running {
			fmt.Print("Press Enter to continue!!")
			fmt.Scanln()
		}
	}
}

func createNewBlock() *pkg.Block {
	var transactions []*pkg.Transaction
	var buffer string
	count := 0

	fmt.Println("------------ Transactions in Block ------------")
	scanner := bufio.NewReader(os.Stdin)

	for count == 0 || buffer != "" {
		fmt.Print("Enter a transaction description (press Enter to finish): ")
		buffer, _ = scanner.ReadString('\n')
		buffer = strings.TrimSpace(buffer)

		if buffer != "" {
			transactions = append(transactions, pkg.NewTransaction([]byte(buffer)))
			count++
		}
	}

	fmt.Println("---------- End Transactions in Block ----------")

	block := pkg.NewBlock(transactions, nil)
	return block
}

func loadChainMenu() {
	running := true
	fmt.Print("Enter the JSON file path to load the chain from: ")
	var filePath string
	fmt.Scanln(&filePath)

	blockchain, err := pkg.ReadFromFile(filePath)
	if err != nil {
		fmt.Printf("Error loading blockchain from file: %v\n", err)
		return
	}

	for running {
		fmt.Printf("\x1bc")
		printSubMenu("Chain from File")

		fmt.Print("Select an option: ")
		var choice int
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			// Add a block to the blockchain
			// You can implement this functionality
		case 2:
			// Display the blockchain
			pkg.DisplayBlockchain(blockchain)
		case 3:
			// Back to the main menu
			running = false
		default:
			fmt.Println("Invalid option. Please select a valid option.")
		}

		if running {
			fmt.Print("Press Enter to continue!!")
			fmt.Scanln()
		}
	}
}
