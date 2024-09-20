package transaction

import (
	"net/http"
	"strconv"

	"fintrack/internal/domain/entity"

	"fintrack/internal/service/transaction"

	"github.com/gin-gonic/gin"
)

// TransactionHandler handles transaction-related requests.
type TransactionHandler struct {
	service *transaction.TransactionService
}

// NewTransactionHandler creates a new instance of TransactionHandler.
func NewTransactionHandler(s *transaction.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: s}
}

// PostTransaction godoc
// @Summary Create a new transaction
// @Description Creates a new transaction record for the user
// @Tags Transactions
// @Accept json
// @Produce json
// @Param transaction body entity.Transaction true "Transaction details"
// @Success 200 "Transaction created successfully"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /transactions [post]
func (h *TransactionHandler) PostTransaction(c *gin.Context) {
	var t entity.Transaction
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.SaveTransaction(t); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// GetTransactions godoc
// @Summary Get transactions for a specific user, year, and month
// @Description Retrieves the list of transactions for the specified user, year, and month
// @Tags Transactions
// @Accept json
// @Produce json
// @Param year path int true "Year"
// @Param month path int true "Month"
// @Success 200 {array} entity.Transaction "List of transactions"
// @Failure 400 {object} map[string]string "Invalid year or month"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /transactions/{year}/{month} [get]
func (h *TransactionHandler) GetTransactions(c *gin.Context) {
	userID := c.GetUint64("userID") // Assuming this is set via middleware or is part of the session
	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid year"})
		return
	}
	month, err := strconv.Atoi(c.Param("month"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid month"})
		return
	}

	transactions, err := h.service.GetTransactions(userID, year, month)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}
