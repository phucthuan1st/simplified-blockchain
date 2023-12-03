package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) Add(block *Block) error {
	lastBlock := bc.Blocks[len(bc.Blocks)-1]

	if bytes.Equal(lastBlock.Hash, block.PrevBlockHash) {
		bc.Blocks = append(bc.Blocks, block)
		return nil
	}

	return fmt.Errorf("Invalid new block hash")
}

func (bc *Blockchain) GetLastBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

func NewBlockchain(identifier string) *Blockchain {

	genesisData := []byte(identifier)

	// Create the Genesis block
	genesisBlock := createGenesisBlock(genesisData)

	return &Blockchain{
		Blocks: []*Block{genesisBlock},
	}
}

func createGenesisBlock(genesisData []byte) *Block {
	// The Genesis block has no previous block, so set PrevBlockHash to nil or an empty byte slice
	// You can also use a predefined constant for the Genesis block's PrevBlockHash
	emptyPrevBlockHash := []byte{}

	// Create the Genesis block
	sender := string(genesisData)
	receiver := string(genesisData)
	signature := string(genesisData)

	transaction, err := NewTransaction(sender, receiver, signature)
	if err != nil {
		// Handle the error, for example, print an error message and return nil or handle it as appropriate.
		fmt.Println("Error creating transaction for Genesis block:", err)
		return nil
	}

	genesisBlock := NewBlock([]*Transaction{transaction}, emptyPrevBlockHash)
	return genesisBlock
}

func (bc *Blockchain) WriteToFile(filename string) error {
	jsonData, err := json.MarshalIndent(bc, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, jsonData, 0644)
	return err
}

func ReadFromFile(filename string) (*Blockchain, error) {
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var blockchain Blockchain
	err = json.Unmarshal(jsonData, &blockchain)
	if err != nil {
		return nil, err
	}

	return &blockchain, nil
}

func DisplayBlockchain(bc *Blockchain) {
	for _, b := range bc.Blocks {
		DisplayBlock(b)
	}
}
