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

## Contributions

Contributions are welcome. If you find any issues or want to enhance the project, please create a pull request.

## License

This project is licensed under the MIT License. Feel free to use, modify, and distribute the code for educational purposes.

Happy learning and happy coding!
