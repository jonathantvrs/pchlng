package handler

import (
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jonathantvrs/pismo/internal/usecase"
)

type TransactionHandler struct {
	useCase *usecase.TransactionUseCase
}

func NewTransactionHandler(uc *usecase.TransactionUseCase) *TransactionHandler {
	return &TransactionHandler{useCase: uc}
}

type CreateTransactionRequest struct {
	AccountID       int     `json:"account_id" binding:"required,gt=0"`
	OperationTypeID int     `json:"operation_type_id" binding:"required,gt=0"`
	Amount          float64 `json:"amount" binding:"required,ne=0"`
}

type TransactionResponse struct {
	ID              int       `json:"transaction_id"`
	AccountID       int       `json:"account_id"`
	OperationTypeID int       `json:"operation_type_id"`
	Amount          float64   `json:"amount"`
	EventDate       time.Time `json:"event_date"`
}

// Create godoc
// @Summary Cria uma nova transação
// @Description Registra uma transação e inverte o sinal baseado no tipo
// @Tags transactions
// @Accept json
// @Produce json
// @Param transaction body CreateTransactionRequest true "Dados da transação"
// @Success 201 {object} domain.Transaction
// @Failure 400 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /transactions [post]
func (h *TransactionHandler) Create(c *gin.Context) {
	var req CreateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data", "details": err.Error()})
		return
	}

	amountInCents := int64(math.Round(req.Amount * 100))

	tx, err := h.useCase.Create(req.AccountID, req.OperationTypeID, amountInCents)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	resp := TransactionResponse{
		ID:              tx.ID,
		AccountID:       tx.AccountID,
		OperationTypeID: tx.OperationTypeID,
		Amount:          float64(tx.Amount) / 100.0,
		EventDate:       tx.EventDate,
	}

	c.JSON(http.StatusCreated, resp)
}
