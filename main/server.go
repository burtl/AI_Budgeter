// Package main defines the server setup and entry point for the AI_Budgeter application.
package main

import (
	"AI_Budgeter/handlers" // Adjust import path based on your project
	"AI_Budgeter/services" // Adjust import path based on your project
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // PostgreSQL driver
	"log"
)

// SetupServer initializes the Gin engine and routes.
func SetupServer(db *sql.DB) *gin.Engine {
	r := gin.Default()

	// Initialize the AuthService with the DB connection.
	authService := &services.AuthService{DB: db}

	// Create a user handler for registration.
	userHandler := &handlers.UserHandler{AuthService: authService}
	r.POST("/register", userHandler.RegisterEndpoint)

	// Create a login handler for authentication.
	loginHandler := &handlers.LoginHandler{AuthService: authService}
	r.POST("/login", loginHandler.LoginEndpoint)

	// Additional routes can be defined here (e.g., /transactions, /balance, etc.)

	return r
}

func main() {
	// Connect to PostgreSQL (adjust credentials as needed).
	connStr := "postgres://dev.user:dev.password@localhost:5432/dev.database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Verify DB connection.
	if pingErr := db.Ping(); pingErr != nil {
		log.Fatal("Cannot ping database:", pingErr)
	}

	fmt.Println("Database connection successful.")

	// Create the server and run it on port 8080.
	server := SetupServer(db)
	if httpErr := server.Run(":8080"); httpErr != nil {
		log.Fatal("Server failed to start:", httpErr)
	}
}
