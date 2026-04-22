package domain

import "errors"

var (
	ErrDocumentAlreadyExists = errors.New("account with this document already exists")
	ErrAccountNotFound       = errors.New("account not found")
	ErrInvalidOperation      = errors.New("invalid operation type")
)
