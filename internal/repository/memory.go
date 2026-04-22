package repository

import (
	"errors"
	"sync"

	"github.com/jonathantvrs/pismo/internal/domain"
)

type AccountMemoryRepo struct {
	accounts map[int]*domain.Account
	docIndex map[string]int
	seq      int
	mu       sync.RWMutex
}

func NewAccountMemoryRepo() *AccountMemoryRepo {
	return &AccountMemoryRepo{
		accounts: make(map[int]*domain.Account),
		docIndex: make(map[string]int),
	}
}

func (r *AccountMemoryRepo) Create(acc *domain.Account) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.seq++
	acc.ID = r.seq
	r.accounts[acc.ID] = acc
	r.docIndex[acc.DocumentNumber] = acc.ID
	return nil
}

func (r *AccountMemoryRepo) GetByDocument(doc string) (*domain.Account, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	id, exists := r.docIndex[doc]
	if !exists {
		return nil, errors.New("account not found")
	}
	return r.accounts[id], nil
}

func (r *AccountMemoryRepo) GetByID(id int) (*domain.Account, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if acc, exists := r.accounts[id]; exists {
		return acc, nil
	}
	return nil, errors.New("account not found")
}

type OperationMemoryRepo struct {
	ops map[int]*domain.OperationType
	mu  sync.RWMutex
}

func NewOperationMemoryRepo() *OperationMemoryRepo {
	return &OperationMemoryRepo{
		ops: map[int]*domain.OperationType{
			1: {ID: 1, Description: "COMPRA A VISTA", Multiplier: -1},
			2: {ID: 2, Description: "COMPRA PARCELADA", Multiplier: -1},
			3: {ID: 3, Description: "SAQUE", Multiplier: -1},
			4: {ID: 4, Description: "PAGAMENTO", Multiplier: 1},
		},
	}
}

func (r *OperationMemoryRepo) GetByID(id int) (*domain.OperationType, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if op, exists := r.ops[id]; exists {
		return op, nil
	}
	return nil, errors.New("operation type not found")
}

type TransactionMemoryRepo struct {
	transactions map[int]*domain.Transaction
	seq          int
	mu           sync.RWMutex
}

func NewTransactionMemoryRepo() *TransactionMemoryRepo {
	return &TransactionMemoryRepo{transactions: make(map[int]*domain.Transaction)}
}

func (r *TransactionMemoryRepo) Create(tx *domain.Transaction) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.seq++
	tx.ID = r.seq
	r.transactions[tx.ID] = tx
	return nil
}
