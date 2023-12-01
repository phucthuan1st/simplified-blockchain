package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"simplified-blockchain/pkg"
	"strings"
)

func PrintMainMenu() {
	fmt.Printf("\x1bc")
	fmt.Println("Main Menu:")
	fmt.Println("--------------- Menu ----------------")
	fmt.Println("1. Create a new chain")
	fmt.Println("2. Load a chain from file")
	fmt.Println("3. Exit")
	fmt.Println("-------------------------------------")
}

func PrintSubMenu(menuTitle string) {
	fmt.Printf("%s:\n", menuTitle)
	fmt.Println("--------------- Menu ----------------")
	fmt.Println("1. Add a Block")
	fmt.Println("2. Display Chain")
	fmt.Println("3. Save current chain")
	fmt.Println("4. Back to Main Menu")
	fmt.Println("-------------------------------------")
}

func CreateNewChainMenu() {
	running := true
	fmt.Print("Enter an identifier for the new chain: ")
	var identifier string
	fmt.Scanln(&identifier)

	blockchain := pkg.NewBlockchain(identifier)

	for running {
		fmt.Printf("\x1bc")
		PrintSubMenu("New Chain")

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
			block := CreateNewBlock()
			block.PrevBlockHash = blockchain.GetLastBlock().Hash
			err := blockchain.Add(block)
			if err != nil {
				log.Printf("Error adding block: %v", err)
			}
		case 2:
			pkg.DisplayBlockchain(blockchain)
		case 3:
			// save current blockchain to json file
			path := ""
			fmt.Print("Enter path to save: ")
			fmt.Scanf("%s", &path)
			err := blockchain.WriteToFile(path + identifier + ".json")
			if err != nil {
				log.Printf("Error writing blockchain to file: %s\n", err.Error())
			}
		case 4:
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

func CreateNewBlock() *pkg.Block {
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

func LoadChainMenu() {
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
		PrintSubMenu("Chain from File")

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
			block := CreateNewBlock()
			block.PrevBlockHash = blockchain.GetLastBlock().Hash
			err := blockchain.Add(block)
			if err != nil {
				log.Printf("Error adding block: %v", err)
			}
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
