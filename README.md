# Simple Blockchain Implementation in Go

## Overview

This project is a simplified blockchain implementation in Go. It includes basic structures for Blocks, Transactions, and a Blockchain.

## Block Structure

```go
type Block struct {
    Transactions  []*Transaction
    PrevBlockHash []byte
    Hash          []byte
    Timestamp     int64
}
```

## Transaction Structure

```go
type Transaction struct {
    Data []byte
}
```

## Transaction Data Structure

```go
type TransactionData struct {
    Sender    string `json:"sender"`
    Receiver  string `json:"receiver"`
    Signature string `json:"signature"`
}
```

## Blockchain Structure

```go
type Blockchain struct {
    Blocks []*Block
}
```

## Merkle Tree for Block Integrity

In this implementation, a Merkle tree is used to ensure the integrity of the transactions within a block. The Merkle tree is constructed using a bottom-up approach, where leaf nodes represent individual transactions, and internal nodes represent the hash concatenation of their child nodes.

### Block Integrity Validation

To validate the integrity of a block, the Merkle root checksum is calculated and compared with the one stored in the block:

```go
func ValidateBlockIntegrity(b *Block) bool {
    merkleRootChecksum := NewMerkleTree(b.Transactions).GetMerkleRootChecksum()

    return bytes.Equal(merkleRootChecksum, b.MerkleRootChecksum)
}
```

## Usage

You can use this blockchain implementation for educational purposes, understanding the basic concepts of blockchains, and experimenting with Go programming.

Feel free to explore the code and modify it according to your learning needs.

## How to Run

To run the project, use the following commands:

```bash
git clone https://github.com/phucthuan1st/simplified-blockchain
cd simplified-blockchain
go run main.go
```

### Command-Line Interface

The blockchain provides a command-line interface with the following usage:

```plaintext

Usage:
No flag        --> Interactive mode
--help         --> Display command outline
--show         --> Display all blocks from the current chain database
--addblock     --> Create a new block and add it to the chain

Optional Flags:
--db pathToDBFile  --> Specify the path to the database file (default: MYCHAINPATH or ./data/mychain.json)
```

## Contributions

Contributions are welcome. If you find any issues or want to enhance the project, please create a pull request.

### License

This project is licensed under the [GNU General Public License v3.0](LICENSE). Feel free to use, modify, and distribute the code in accordance with the terms of the license.

Feel free to explore the code and modify it according to your learning needs.

Happy learning and happy coding!
