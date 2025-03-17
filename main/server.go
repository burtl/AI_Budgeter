// Package main defines the server setup and entry point for the AI_Budgeter application.
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // PostgreSQL driver
	"your_project/handlers" // Adjust import path based on your project
	"your_project/services" // Adjust import path based on your project
)

// SetupServer initializes the Gin engine and routes.
func SetupServer(db *sql.DB) *gin.Engine {
	// Single-line comment: create a new router.
	r := gin.Default()

	// Single-line comment: initialize the AuthService with our DB connection.
	authService := &services.AuthService{DB: db}

	// Single-line comment: create a user handler with the AuthService.
	userHandler := &handlers.UserHandler{AuthService: authService}

	// Single-line comment: define the /register route.
	r.POST("/register", userHandler.RegisterEndpoint)

	// You can define other routes here (e.g., /login, /transactions, etc.)

	return r
}

func main() {
	// Single-line comment: connect to PostgreSQL (adjust credentials as needed).
	connStr := "postgres://dev.user:dev.password@localhost:5432/dev.database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Single-line comment: verify DB connection.
	if pingErr := db.Ping(); pingErr != nil {
		log.Fatal("Cannot ping database:", pingErr)
	}

	fmt.Println("Database connection successful.")

	// Single-line comment: create the server and run it on port 8080.
	server := SetupServer(db)
	httpErr := server.Run(":8080")
	if httpErr != nil {
		log.Fatal("Server failed to start:", httpErr)
	}
}