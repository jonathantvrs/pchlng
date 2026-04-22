package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jonathantvrs/pismo/internal/domain"
	"github.com/jonathantvrs/pismo/internal/usecase"
)

type AccountHandler struct {
	useCase *usecase.AccountUseCase
}

func NewAccountHandler(uc *usecase.AccountUseCase) *AccountHandler {
	return &AccountHandler{useCase: uc}
}

type CreateAccountRequest struct {
	DocumentNumber string `json:"document_number" binding:"required,len=11,numeric"`
}

// Create godoc
// @Summary Cria uma nova conta
// @Description Cria uma conta baseada no número do documento
// @Tags accounts
// @Accept json
// @Produce json
// @Param account body CreateAccountRequest true "Dados da conta"
// @Success 201 {object} domain.Account
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Router /accounts [post]
func (h *AccountHandler) Create(c *gin.Context) {
	var req CreateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data", "details": err.Error()})
		return
	}

	acc, err := h.useCase.Create(req.DocumentNumber)
	if err != nil {
		if errors.Is(err, domain.ErrDocumentAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, acc)
}

// Get godoc
// @Summary Consulta uma conta
// @Description Retorna os detalhes de uma conta pelo ID
// @Tags accounts
// @Accept json
// @Produce json
// @Param accountId path int true "ID da conta"
// @Success 200 {object} domain.Account
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /accounts/{accountId} [get]
func (h *AccountHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("accountId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid account id"})
		return
	}

	acc, err := h.useCase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		return
	}

	c.JSON(http.StatusOK, acc)
}
