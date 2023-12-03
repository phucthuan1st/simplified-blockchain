package pkg

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Timestamp     int64
}

func (b *Block) SetHash() error {
	timestamp_bytes := []byte(strconv.FormatInt(b.Timestamp, 15))
	transactions_bytes, err := json.Marshal(b.Transactions)

	if err != nil {
		return err
	}

	// TODO: Feed the PrevBlockHash, Transactions, and Timestamp into the hash in this order
	header := bytes.Join([][]byte{b.PrevBlockHash, transactions_bytes, timestamp_bytes}, []byte{})

	// TODO: calculate the hash of this block
	hash := sha256.Sum256(header)

	b.Hash = hash[:]

	return nil
}

func NewBlock(transactions []*Transaction, prevBlockHash []byte) (*Block, error) {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Transactions:  transactions,
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
	}

	err := block.SetHash()

	return block, err
}

func DisplayBlock(b *Block) {
	fmt.Println("--------------- Block ----------------")
	fmt.Printf("Timestamp: %d \n", b.Timestamp)

	prevHash := base64.StdEncoding.EncodeToString(b.PrevBlockHash)
	fmt.Printf("Previous Block Hash: %v \n", prevHash)

	fmt.Println("Transactions: ")
	for i, tx := range b.Transactions {
		var transactionDetails TransactionData
		err := json.Unmarshal(tx.Data, &transactionDetails)
		if err != nil {
			fmt.Printf("Genesis Block: %v \n", string(tx.Data))
		} else {
			fmt.Printf("\tTransaction %d: %v\n", i, transactionDetails)
		}
	}

	hash := base64.StdEncoding.EncodeToString(b.Hash)
	fmt.Printf("Hash: %v \n", hash)
	fmt.Println("------------- End Block --------------")
}
