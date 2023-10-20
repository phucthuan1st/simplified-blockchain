package pkg

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"log"
	"time"
)

type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
}

func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	if transactions == nil {
		return nil
	}

	block := &Block{
		Timestamp:     time.Now().Unix(),
		Transactions:  transactions,
		PrevBlockHash: prevBlockHash,
	}

	// Calculate and set the block's hash
	block.Hash = block.calculateHash()

	return block
}

func (b *Block) calculateHash() []byte {
	var preparedHash []byte

	preparedHash = append(preparedHash, int64ToBytes(b.Timestamp)...)

	for _, transaction := range b.Transactions {
		preparedHash = append(preparedHash, transaction.Data...)
	}

	preparedHash = append(preparedHash, b.PrevBlockHash...)

	hasher := sha256.New()
	_, err := hasher.Write(preparedHash)

	if err != nil {
		log.Printf("error writing hash: %v", err)
	}

	return hasher.Sum(nil)
}

func int64ToBytes(n int64) []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, uint64(n))
	return bytes
}

func DisplayBlock(b *Block) {
	fmt.Println("--------------- Block ----------------")
	fmt.Printf("Timestamp: %d \n", b.Timestamp)

	fmt.Println("Transactions: ")
	for i, tx := range b.Transactions {
		transactionDetails := string(tx.Data)
		fmt.Printf("\tTransaction %d: %v\n", i, transactionDetails)
	}

	hash := base64.StdEncoding.EncodeToString(b.Hash)
	fmt.Printf("Hash: %v \n", hash)

	prevHash := base64.StdEncoding.EncodeToString(b.PrevBlockHash)
	fmt.Printf("Previous Block Hash: %v \n", prevHash)
	fmt.Println("------------- End Block --------------")
}
