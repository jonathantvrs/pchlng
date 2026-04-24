package usecase

import (
	"time"

	"github.com/jonathantvrs/pismo/internal/domain"
	"github.com/jonathantvrs/pismo/pkg/abs"
)

type TransactionUseCase struct {
	txRepo  domain.TransactionRepository
	accRepo domain.AccountRepository
	opRepo  domain.OperationTypeRepository
}

func NewTransactionUseCase(
	tr domain.TransactionRepository,
	ar domain.AccountRepository,
	or domain.OperationTypeRepository) *TransactionUseCase {
	return &TransactionUseCase{txRepo: tr, accRepo: ar, opRepo: or}
}

func (uc *TransactionUseCase) Create(accountID, opTypeID int, amount int64) (*domain.Transaction, error) {
	if _, err := uc.accRepo.GetByID(accountID); err != nil {
		return nil, domain.ErrAccountNotFound
	}

	opType, err := uc.opRepo.GetByID(opTypeID)
	if err != nil {
		return nil, domain.ErrInvalidOperation
	}

	finalAmount := abs.Abs(amount) * opType.Multiplier

	tx := &domain.Transaction{
		AccountID:       accountID,
		OperationTypeID: opTypeID,
		Amount:          finalAmount,
		EventDate:       time.Now(),
	}

	if err := uc.txRepo.Create(tx); err != nil {
		return nil, err
	}

	return tx, nil
}
