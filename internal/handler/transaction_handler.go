package handler

import (
	"net/http"

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

	tx, err := h.useCase.Create(req.AccountID, req.OperationTypeID, req.Amount)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, tx)
}
