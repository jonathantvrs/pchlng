package repository

import (
	"database/sql"
	"errors"

	"github.com/jonathantvrs/pismo/internal/domain"
)

type SQLiteAccountRepo struct {
	db *sql.DB
}

func NewSQLiteAccountRepo(db *sql.DB) *SQLiteAccountRepo {
	return &SQLiteAccountRepo{db: db}
}

func (r *SQLiteAccountRepo) Create(acc *domain.Account) error {
	query := `INSERT INTO accounts (document_number) VALUES (?) RETURNING id`
	return r.db.QueryRow(query, acc.DocumentNumber).Scan(&acc.ID)
}

func (r *SQLiteAccountRepo) GetByID(id int) (*domain.Account, error) {
	acc := &domain.Account{}
	query := `SELECT id, document_number FROM accounts WHERE id = ?`
	err := r.db.QueryRow(query, id).Scan(&acc.ID, &acc.DocumentNumber)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrAccountNotFound
	}
	return acc, err
}

func (r *SQLiteAccountRepo) GetByDocument(doc string) (*domain.Account, error) {
	acc := &domain.Account{}
	query := `SELECT id, document_number FROM accounts WHERE document_number = ?`
	err := r.db.QueryRow(query, doc).Scan(&acc.ID, &acc.DocumentNumber)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrAccountNotFound
	}
	return acc, err
}

type SQLiteOperationRepo struct {
	db *sql.DB
}

func NewSQLiteOperationRepo(db *sql.DB) *SQLiteOperationRepo {
	return &SQLiteOperationRepo{db: db}
}

func (r *SQLiteOperationRepo) GetByID(id int) (*domain.OperationType, error) {
	op := &domain.OperationType{}
	query := `SELECT id, description, multiplier FROM operation_types WHERE id = ?`
	err := r.db.QueryRow(query, id).Scan(&op.ID, &op.Description, &op.Multiplier)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrInvalidOperation
	}
	return op, err
}

type SQLiteTransactionRepo struct {
	db *sql.DB
}

func NewSQLiteTransactionRepo(db *sql.DB) *SQLiteTransactionRepo {
	return &SQLiteTransactionRepo{db: db}
}

func (r *SQLiteTransactionRepo) Create(tx *domain.Transaction) error {
	query := `INSERT INTO transactions (account_id, operation_type_id, amount, event_date) 
	          VALUES (?, ?, ?, ?) RETURNING id`
	return r.db.QueryRow(query, tx.AccountID, tx.OperationTypeID, tx.Amount, tx.EventDate).Scan(&tx.ID)
}
