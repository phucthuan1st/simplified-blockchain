package pkg

import (
	"encoding/json"
	"fmt"
)

type Transaction struct {
	Data []byte
}

type TransactionData struct {
	Sender    string `json:"sender"`
	Receiver  string `json:"receiver"`
	Signature string `json:"signature"`
}

func NewTransaction(sender, receiver, signature string) (*Transaction, error) {
	// Create a Data struct
	data := TransactionData{
		Sender:    sender,
		Receiver:  receiver,
		Signature: signature,
	}

	// Convert struct to JSON byte array
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshalling data: %v", err)
	}

	// Create Transaction with JSON data
	transaction := &Transaction{
		Data: jsonData,
	}

	return transaction, nil
}
