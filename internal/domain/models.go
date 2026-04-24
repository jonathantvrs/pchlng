package domain

import "time"

type Account struct {
	ID             int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

type OperationType struct {
	ID          int
	Description string
	Multiplier  int64
}

type Transaction struct {
	ID              int       `json:"transaction_id"`
	AccountID       int       `json:"account_id"`
	OperationTypeID int       `json:"operation_type_id"`
	Amount          int64     `json:"amount"`
	EventDate       time.Time `json:"event_date"`
}

type AccountRepository interface {
	Create(account *Account) error
	GetByID(id int) (*Account, error)
	GetByDocument(doc string) (*Account, error)
}

type OperationTypeRepository interface {
	GetByID(id int) (*OperationType, error)
}

type TransactionRepository interface {
	Create(transaction *Transaction) error
}
