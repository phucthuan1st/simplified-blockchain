package pkg

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
)

const targetBits = 10 // gia tri co the thay doi duoc so cang lon thi do kho cang cao, thuan tien cho viec kiem tra chay thu nen de so la 10 de chay nhanh hon
var maxNonce = math.MaxInt64

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// ham bat tra ve mot PoW moi voi 2 tham so b va target.
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}
	return pow
}

// HashTransactions hashes the block's transactions
func (b *Block) HashTransactions() []byte {
	var transactions [][]byte
	for _, tx := range b.Transactions {
		txJSON, err := json.Marshal(tx)
		if err != nil {
			panic(err)
		}
		transactions = append(transactions, txJSON)
	}
	merkleTree := NewMerkleTree(transactions)
	RootNode := merkleTree.GetMerkleRootData()
	return RootNode
}

// chuan bi data de hash.

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.HashTransactions(),
			int64ToBytes(pow.block.Timestamp),
			int64ToBytes(int64(targetBits)),
			int64ToBytes(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) MakeProofOfWork() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Minining the block containing \"%s\"\n", pow.block.Transactions[0].Data) // sua lai transaction data cho nay.
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")
	return nonce, hash[:]
}

// ham validate
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1 // ktr nho hon

	return isValid
}
