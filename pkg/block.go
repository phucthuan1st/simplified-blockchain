package pkg

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"strconv"
	"time"
)

type Block struct {
	PrevBlockHash      []byte
	Transactions       []*Transaction
	Hash               []byte
	Timestamp          int64
	MerkleRootChecksum []byte
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
	}

	err := block.SetHash()
	block.MerkleRootChecksum = NewMerkleTree(block.Transactions).GetMerkleRootChecksum()

	return block, err
}
