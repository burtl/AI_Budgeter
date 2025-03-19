// Package handlers defines HTTP handlers for transaction-related endpoints.
package handlers

import (
	"AI_Budgeter/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// TransactionsHandler handles /transactions endpoints for listing and creating transactions.
type TransactionsHandler struct {
	TransactionService *services.TransactionService
}

// GetTransactionsEndpoint handles GET /transactions.
// It expects a valid JWT to retrieve user_id from the context,
// then optionally reads "page" and "limit" query parameters.
func (h *TransactionsHandler) GetTransactionsEndpoint(ctx *gin.Context) {
	// Example: extracting user_id from JWT claims. Adjust to your actual auth middleware usage.
	userIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID, ok := userIDVal.(int)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

	// Parse pagination query params
	pageStr := ctx.Query("page")
	limitStr := ctx.Query("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	transactions, dbErr := h.TransactionService.GetUserTransactions(userID, page, limit)
	if dbErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transactions"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"transactions": transactions})
}

// CreateTransactionEndpoint handles POST /transactions.
// It expects a valid JWT (to identify user_id) and JSON body with "amount",
// "transaction_type" (credit/debit), and optionally "category_id".
func (h *TransactionsHandler) CreateTransactionEndpoint(ctx *gin.Context) {
	// Example: extracting user_id from JWT claims
	userIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID, ok := userIDVal.(int)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

	// Request body structure
	type requestBody struct {
		Amount          float64 `json:"amount"`
		TransactionType string  `json:"transaction_type"`
		CategoryID      *int    `json:"category_id"`
	}

	var body requestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	newTx, createErr := h.TransactionService.CreateTransaction(
		userID,
		body.Amount,
		body.TransactionType,
		body.CategoryID,
	)
	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": createErr.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"transaction": newTx})
}
