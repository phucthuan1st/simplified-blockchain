package main

import (
	"flag"
	"fmt"
	"os"
	"simplified-blockchain/cmd"
)

func main() {
	// Parse command-line arguments
	helpFlag := flag.Bool("help", false, "Display command outline")
	showFlag := flag.Bool("show", false, "Display all blocks from the current chain database")
	addBlockFlag := flag.Bool("addblock", false, "Create a new block and add it to the chain")

	dbPathFlag := flag.String("db", "", "Path to the database file")
	flag.Parse()

	if *helpFlag {
		cmd.PrintCommandOutline()
		return
	}

	// Use the specified database path or the default if not provided
	dbPath := *dbPathFlag
	if dbPath == "" {
		dbPath = os.Getenv("MYCHAINPATH")
		if dbPath == "" {
			dbPath = "./data/mychain.json"
		}
	}

	// Execute the appropriate command based on user input
	if *showFlag {
		bc, err := cmd.LoadChainData(dbPath)

		if err != nil {
			fmt.Printf("Error loading chain data from %s: %s", dbPath, err.Error())
		}

		cmd.DisplayBlockchain(bc)
	} else if *addBlockFlag {
		bc, err := cmd.LoadChainData(dbPath)

		if err != nil {
			fmt.Printf("Error loading chain data from %s: %s", dbPath, err.Error())
		}

		block, err := cmd.CreateNewBlock()
		if err != nil {
			fmt.Printf("Create new block failed: %v\n", err)
			return
		}

		err = bc.Add(block)

		if err != nil {
			fmt.Printf("Add failed: %v\n", err)
			return
		}

		err = bc.WriteToFile(dbPath)

		if err != nil {
			fmt.Printf("Write chain to db failed: %v\n", err)
		}
	} else {
		// If no valid command is provided, display the main menu
		for {
			cmd.PrintMainMenu()

			var choice int
			fmt.Print("Select an option: ")
			_, err := fmt.Scanf("%d\n", &choice)
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
}
