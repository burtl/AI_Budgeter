// Package services provides the business logic for user authentication and registration.
package services

import (
	"database/sql"
	"errors"
	"log"
)

// AuthService is responsible for handling user authentication and registration.
type AuthService struct {
	DB *sql.DB
}

// RegisterUser creates a new user record in the database.
// It expects the caller to provide a hashed password (for security best practices).
func (a *AuthService) RegisterUser(username, email, hashedPassword string) error {
	// Single-line comment: prepare an INSERT statement.
	stmt, err := a.DB.Prepare("INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)")
	if err != nil {
		log.Println("Error preparing INSERT statement:", err)
		return err
	}
	defer stmt.Close()

	// Single-line comment: execute the INSERT with the provided user details.
	_, execErr := stmt.Exec(username, email, hashedPassword)
	if execErr != nil {
		log.Println("Error executing INSERT for new user:", execErr)
		return execErr
	}

	return nil
}

// IsEmailTaken checks if the given email already exists in the users table.
// Returns true if the email is already in use, false otherwise.
func (a *AuthService) IsEmailTaken(email string) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM users WHERE email = $1"
	err := a.DB.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}