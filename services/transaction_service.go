// Package services provides the business logic for transaction operations.
package services

import (
	"AI_Budgeter/models" // Import to use the transaction type constants
	"database/sql"
	"errors"
	"log"
	"time"
)

// Transaction holds the data for a single transaction record.
type Transaction struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	Amount          float64   `json:"amount"`
	TransactionType string    `json:"transaction_type"`
	CategoryID      *int      `json:"category_id,omitempty"`
	Timestamp       time.Time `json:"timestamp"`
}

// TransactionService manages transaction-related database operations.
type TransactionService struct {
	DB *sql.DB
}

// GetUserTransactions retrieves paginated transactions for a specific user.
// It expects page and limit to determine offset and row count.
func (t *TransactionService) GetUserTransactions(userID, page, limit int) ([]Transaction, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	query := `
		SELECT id, user_id, amount, transaction_type, category_id, timestamp
		FROM transactions
		WHERE user_id = $1
		ORDER BY timestamp DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := t.DB.Query(query, userID, limit, offset)
	if err != nil {
		log.Println("Error querying transactions:", err)
		return nil, err
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var tx Transaction
		if scanErr := rows.Scan(
			&tx.ID,
			&tx.UserID,
			&tx.Amount,
			&tx.TransactionType,
			&tx.CategoryID,
			&tx.Timestamp,
		); scanErr != nil {
			log.Println("Error scanning transaction row:", scanErr)
			return nil, scanErr
		}
		transactions = append(transactions, tx)
	}

	return transactions, rows.Err()
}

// CreateTransaction inserts a new transaction record for the given user.
// It validates the transaction type (credit/debit) before inserting.
func (t *TransactionService) CreateTransaction(userID int, amount float64, transactionType string, categoryID *int) (*Transaction, error) {
	// Validate transaction type
	if transactionType != models.TransactionTypeCredit && transactionType != models.TransactionTypeDebit {
		return nil, errors.New("invalid transaction type")
	}
	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}

	insertQuery := `
		INSERT INTO transactions (user_id, amount, transaction_type, category_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id, timestamp
	`

	var txID int
	var createdTime time.Time
	err := t.DB.QueryRow(insertQuery, userID, amount, transactionType, categoryID).Scan(&txID, &createdTime)
	if err != nil {
		log.Println("Error inserting transaction:", err)
		return nil, err
	}

	// Construct the Transaction object to return
	newTx := &Transaction{
		ID:              txID,
		UserID:          userID,
		Amount:          amount,
		TransactionType: transactionType,
		CategoryID:      categoryID,
		Timestamp:       createdTime,
	}

	return newTx, nil
}
