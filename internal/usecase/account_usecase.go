package usecase

import (
	"errors"

	"github.com/jonathantvrs/pismo/internal/domain"
)

type AccountUseCase struct {
	repo domain.AccountRepository
}

func NewAccountUseCase(r domain.AccountRepository) *AccountUseCase {
	return &AccountUseCase{repo: r}
}

func (uc *AccountUseCase) Create(documentNumber string) (*domain.Account, error) {
	existing, err := uc.repo.GetByDocument(documentNumber)
	if err != nil && !errors.Is(err, domain.ErrAccountNotFound) {
		return nil, err
	}

	if existing != nil {
		return nil, domain.ErrDocumentAlreadyExists
	}

	acc := &domain.Account{DocumentNumber: documentNumber}
	if err := uc.repo.Create(acc); err != nil {
		return nil, err
	}

	return acc, nil
}

func (uc *AccountUseCase) GetByID(id int) (*domain.Account, error) {
	return uc.repo.GetByID(id)
}
