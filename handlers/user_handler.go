// Package handlers defines HTTP handlers for user-related endpoints.
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"your_project/services" // Adjust import path based on your project structure
)

// UserHandler handles user-related HTTP requests.
type UserHandler struct {
	AuthService *services.AuthService
}

// RegisterEndpoint handles the /register POST request.
// It expects JSON input with "username", "email", and "password" fields.
func (u *UserHandler) RegisterEndpoint(ctx *gin.Context) {
	type requestBody struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body requestBody

	// Single-line comment: bind JSON from the request to our struct.
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Single-line comment: check for duplicate email before registering.
	emailTaken, err := u.AuthService.IsEmailTaken(body.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	if emailTaken {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Email already in use"})
		return
	}

	// Single-line comment: hash the password using bcrypt for security.
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if hashErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	// Single-line comment: attempt to register the user via the AuthService.
	registerErr := u.AuthService.RegisterUser(body.Username, body.Email, string(hashedPassword))
	if registerErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}