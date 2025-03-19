// Package handlers defines HTTP handlers for user-related endpoints.
package handlers

import (
	"net/http"

	"AI_Budgeter/services" // Adjust the import path based on your project structure
	"github.com/gin-gonic/gin"
)

// LoginHandler handles authentication-related HTTP requests.
type LoginHandler struct {
	AuthService *services.AuthService
}

// LoginEndpoint handles the /login POST request.
// It expects JSON input with "email" and "password" fields and returns a JWT token on success.
func (l *LoginHandler) LoginEndpoint(ctx *gin.Context) {
	// Define the expected request body structure.
	type requestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body requestBody

	// Bind JSON from the request to our struct.
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Attempt to authenticate the user via the AuthService.
	token, err := l.AuthService.LoginUser(body.Email, body.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Return the generated JWT token.
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
