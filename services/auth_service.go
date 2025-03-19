// Package services provides the business logic for user authentication and registration.
package services

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"          // JWT library for token generation
	"golang.org/x/crypto/bcrypt"             // Package for hashing and comparing passwords
)

const jwtSecret = "mysecretkey" // In production, use a secure environment variable

// AuthService is responsible for handling user authentication and registration.
type AuthService struct {
	DB *sql.DB
}

// RegisterUser creates a new user record in the database.
// It expects the caller to provide a hashed password (for security best practices).
func (a *AuthService) RegisterUser(username, email, hashedPassword string) error {
	stmt, err := a.DB.Prepare("INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)")
	if err != nil {
		log.Println("Error preparing INSERT statement:", err)
		return err
	}
	defer stmt.Close()

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

// LoginUser validates the user credentials and returns a JWT token if successful.
func (a *AuthService) LoginUser(email, password string) (string, error) {
	// Retrieve the user record based on the provided email.
	var userID int
	var username, hashedPassword string
	query := "SELECT id, username, password_hash FROM users WHERE email = $1"
	err := a.DB.QueryRow(query, email).Scan(&userID, &username, &hashedPassword)
	if err != nil {
		log.Println("Error retrieving user:", err)
		return "", errors.New("invalid credentials")
	}

	// Compare the provided password with the stored hashed password.
	if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		log.Println("Password mismatch:", err)
		return "", errors.New("invalid credentials")
	}

	// Create a new JWT token with the user's id and username as claims.
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // token expires in 1 hour
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using the secret key.
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Println("Error signing JWT token:", err)
		return "", err
	}

	return tokenString, nil
}
