// Package main defines the server setup and entry point for the AI_Budgeter application.
package main

import (
	"AI_Budgeter/handlers"
	"AI_Budgeter/services"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// SetupServer initializes the Gin engine and routes.
func SetupServer(db *sql.DB) *gin.Engine {
	r := gin.Default()

	// Initialize the AuthService with the DB connection.
	authService := &services.AuthService{DB: db}

	// Existing user-related handlers
	userHandler := &handlers.UserHandler{AuthService: authService}
	r.POST("/register", userHandler.RegisterEndpoint)

	loginHandler := &handlers.LoginHandler{AuthService: authService}
	r.POST("/login", loginHandler.LoginEndpoint)

	// ---------------------------------------
	// 1) Initialize the TransactionService
	transactionService := &services.TransactionService{DB: db}

	// 2) Create the TransactionsHandler
	transactionsHandler := &handlers.TransactionsHandler{
		TransactionService: transactionService,
	}

	// 3) Register the /transactions routes
	//    (Assuming you'd want JWT middleware on these routes)
	r.GET("/transactions", transactionsHandler.GetTransactionsEndpoint)
	r.POST("/transactions", transactionsHandler.CreateTransactionEndpoint)

	return r
}

func main() {
	connStr := "postgres://dev.user:dev.password@localhost:5432/dev.database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	if pingErr := db.Ping(); pingErr != nil {
		log.Fatal("Cannot ping database:", pingErr)
	}
	fmt.Println("Database connection successful.")

	server := SetupServer(db)
	if httpErr := server.Run(":8080"); httpErr != nil {
		log.Fatal("Server failed to start:", httpErr)
	}
}
