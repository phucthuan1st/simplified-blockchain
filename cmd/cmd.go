package cmd

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"simplified-blockchain/pkg"
	"strconv"
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

// TODO: working with a new chain
func CreateNewChainMenu() {
	fmt.Print("Enter an identifier for the new chain: ")
	var identifier string
	fmt.Scanln(&identifier)

	// save current blockchain to json file
	path := ""
	fmt.Print("Enter path to save: ")
	fmt.Scanf("%s", &path)

	blockchain, err := pkg.NewBlockchain(identifier)
	if err != nil {
		log.Printf("Error creating new blockchain: %v", err)
	}

	interactiveMenu(blockchain, path)
}

// TODO: load current chain from database
func LoadChainMenu() {
	fmt.Print("Enter the JSON file path to load the chain from: ")
	var filePath string
	fmt.Scanln(&filePath)

	blockchain, err := pkg.ReadFromFile(filePath)
	if err != nil {
		fmt.Printf("Error loading blockchain from file: %v\n", err)
		return
	}

	interactiveMenu(blockchain, filePath)
}

// TODO: this is client side forge for a block, not a backend forge
func CreateNewBlock() (*pkg.Block, error) {
	var transactions []*pkg.Transaction

	fmt.Print("Enter the number of transactions: ")
	scanner := bufio.NewReader(os.Stdin)
	numTransactionsInput, _ := scanner.ReadString('\n')
	numTransactionsInput = strings.TrimSpace(numTransactionsInput)

	numTransactions, err := strconv.Atoi(numTransactionsInput)
	if err != nil {
		return nil, fmt.Errorf("invalid input for number of transactions: %v", err)
	}

	fmt.Println("------------ Transactions in Block ------------")

	for i := 1; i <= numTransactions; i++ {
		fmt.Printf("-- Transaction %d --\n", i)

		fmt.Print("Enter sender: ")
		sender, _ := scanner.ReadString('\n')
		sender = strings.TrimSpace(sender)

		fmt.Print("Enter receiver: ")
		receiver, _ := scanner.ReadString('\n')
		receiver = strings.TrimSpace(receiver)

		fmt.Print("Enter signature: ")
		signature, _ := scanner.ReadString('\n')
		signature = strings.TrimSpace(signature)

		transaction, err := pkg.NewTransaction(sender, receiver, signature)
		if err != nil {
			return nil, fmt.Errorf("error creating transaction: %v", err)
		}

		transactions = append(transactions, transaction)
	}

	fmt.Println("---------- End Transactions in Block ----------")

	//block, err := pkg.NewBlock(transactions, nil)
	block, err := pkg.NewBlock(transactions, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating block: %v", err)
	}

	return block, nil
}

// TODO: client-side for interact with backend of a blockchain, with a place to store chain data
func interactiveMenu(blockchain *pkg.Blockchain, filePath string) {
	running := true

	for running {
		fmt.Printf("\x1bc")
		PrintSubMenu("New Chain")

		fmt.Print("Select an option: ")
		var choice int
		_, err := fmt.Scanf("%d\n", &choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			// Add a block to the blockchain
			block, err := CreateNewBlock()

			if err != nil {
				log.Printf("Error creating a new block: %v", err)
			}
			block.PrevBlockHash = blockchain.GetLastBlock().Hash

			err = blockchain.Add(block)
			if err != nil {
				log.Printf("Error adding block: %v", err)
			}
		case 2:
			DisplayBlockchain(blockchain)
		case 3:
			err := blockchain.WriteToFile(filePath)
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

// PrintCommandOutline prints an outline of available commands and their descriptions.
func PrintCommandOutline() {
	fmt.Println("Usage:")
	fmt.Println("No flag        --> Interactive mode")
	fmt.Println("--help         --> Display command outline")
	fmt.Println("--show         --> Display all blocks from the current chain database")
	fmt.Println("--addblock     --> Create a new block and add it to the chain")
	fmt.Println("\nOptional Flags:")
	fmt.Println("--db pathToDBFile  --> Specify the path to the database file (default: ./data/mychain.json)")
}

func DisplayBlockchain(bc *pkg.Blockchain) {
	for i, b := range bc.Blocks {
		fmt.Println("--------------- Block ----------------")
		DisplayBlockData(b)

		if i == 0 || (i != 0 && pkg.ValidateBlockIntegrity(b)) {
			fmt.Printf("Block has GOOD integrity!\n")
		} else {
			fmt.Printf("Block has BAD integrity!\n")
		}

		fmt.Println("------------- End Block --------------")
	}
}

func DisplayBlockData(b *pkg.Block) {
	fmt.Printf("Timestamp: %d \n", b.Timestamp)

	prevHash := base64.StdEncoding.EncodeToString(b.PrevBlockHash)
	fmt.Printf("Previous Block Hash: %v \n", prevHash)

	fmt.Println("Transactions: ")
	for i, tx := range b.Transactions {
		var transactionDetails pkg.TransactionData
		err := json.Unmarshal(tx.Data, &transactionDetails)

		if err != nil {
			fmt.Printf("\tGenesis Block: %v \n", string(tx.Data))
		} else {
			fmt.Printf("\tTransaction %d: %v\n", i, transactionDetails)
		}
	}

	hash := base64.StdEncoding.EncodeToString(b.Hash)
	fmt.Printf("Hash: %v \n", hash)

	merkleRootChecksum := base64.StdEncoding.EncodeToString(b.MerkleRootChecksum)
	fmt.Printf("MerkleRootChecksum: %v\n", merkleRootChecksum)
}

func LoadChainData(filePath string) (*pkg.Blockchain, error) {
	fmt.Printf("Try loading the blockchain data from %v\n", filePath)
	bc, err := pkg.ReadFromFile(filePath)

	if err != nil {
		fmt.Printf("Error while reading the database: %v\n", err)
		fmt.Printf("Create a new database in path: %v\n", filePath)

		bc, err = pkg.NewBlockchain(filePath)

		if err != nil {
			fmt.Printf("Error creating new chain database: %v\n", err)
			return nil, err
		}

		err := bc.WriteToFile(filePath)

		if err != nil {
			return nil, err
		}
	}

	return bc, nil
}
