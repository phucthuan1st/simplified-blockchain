package pkg

import (
	"bytes"
	cryptoRandom "crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	mathRandom "math/rand"
	"os"
	"time"
)

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) Add(block *Block) error {
	n := len(bc.Blocks) - 1
	lastBlock := bc.Blocks[n]

	if bytes.Equal(lastBlock.Hash, block.PrevBlockHash) {
		bc.Blocks = append(bc.Blocks, block)
		return nil
	}

	return fmt.Errorf("invalid new block hash")
}

func (bc *Blockchain) GetLastBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

func NewBlockchain(identifier string) (*Blockchain, error) {
	// Create the Genesis block
	genesisBlock, err := newGenesisBlock([]byte(identifier))

	return &Blockchain{
		Blocks: []*Block{genesisBlock},
	}, err
}

func newGenesisBlock(genesisData []byte) (*Block, error) {
	if len(genesisData) == 0 {
		genesisData, _ = generateRandomByteArray()
	}

	genesisTransaction := Transaction{
		Data: genesisData,
	}

	block := &Block{
		Timestamp:     time.Now().Unix(),
		Transactions:  []*Transaction{&genesisTransaction},
		PrevBlockHash: []byte{},
	}

	// Calculate and set the block's hash
	err := block.SetHash()

	return block, err
}

func generateRandomByteArray() ([]byte, error) {
	// Generate a random length for the byte array
	randomLength := 16 + mathRandom.Int()%17 // minimum 16 bytes, maximum 32 bytes

	// Generate a random byte array with the chosen length
	randomBytes := make([]byte, randomLength)
	_, err := cryptoRandom.Read(randomBytes)
	if err != nil {
		return nil, fmt.Errorf("error generating random bytes: %v", err)
	}

	// Create a SHA-256 hash of the random byte array
	hashedBytes := sha256.Sum256(randomBytes)

	return hashedBytes[:], nil
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
