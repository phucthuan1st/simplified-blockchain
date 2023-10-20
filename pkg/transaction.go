package pkg

type Transaction struct {
	Data []byte
}

func NewTransaction(data []byte) *Transaction {
	return &Transaction{Data: data}
}
