package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key")

func Login(c *gin.Context) {
    var inputUser User
    if err := c.ShouldBindJSON(&inputUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // Authenticate user (e.g., check username and password)
    user := authenticateUser(inputUser.Username, inputUser.Password)
    if user == nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // Generate JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": user.Username,
        "role":     user.Role,
        "exp":      time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func authenticateUser(username, password string) *User {
    // Replace with actual user lookup logic
    if username == "admin" && password == "admin123" {
        return &User{Username: "admin", Role: "admin"}
    }
    return nil
}