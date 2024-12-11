package controller

import (
	"Trainify/database"
	"Trainify/helper"
	"Trainify/model"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type FormData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login Handler
func Login(c *gin.Context) {
	returnObject := gin.H{
		"status":  "ERROR",
		"message": "Invalid email or password",
	}

	var formData FormData
	if err := c.ShouldBindJSON(&formData); err != nil {
		log.Println("Form binding error:", err)
		c.JSON(http.StatusBadRequest, returnObject)
		return
	}

	var user model.User
	if err := database.DBConn.First(&user, "email = ?", formData.Email).Error; err != nil {
		log.Println("User not found")
		c.JSON(http.StatusUnauthorized, returnObject)
		return
	}

	// Validate password
	if err := helper.CheckPasswordHash(formData.Password, user.Password); err != nil {
		log.Println("Password mismatch")
		c.JSON(http.StatusUnauthorized, returnObject)
		return
	}

	token, err := helper.GenerateToken(user)
	if err != nil {
		returnObject["message"] = "Error in generating token"
		c.JSON(http.StatusInternalServerError, returnObject)
		return
	}

	returnObject["token"] = token
	returnObject["status"] = "OK"
	returnObject["message"] = "User authenticated"
	c.JSON(http.StatusOK, returnObject)
}

// Register Handler
func Register(c *gin.Context) {
	returnObject := gin.H{
		"status":  "ERROR",
		"message": "Registration failed",
	}

	// Collect form data
	var formData FormData
	if err := c.ShouldBindJSON(&formData); err != nil {
		log.Print("Error in JSON binding:", err)
		c.JSON(400, returnObject)
		return
	}

	// Add form data to user model
	var user model.User
	user.Email = formData.Email
	user.Password = helper.HashPassword(formData.Password) // Ensure bcrypt is used in helper function

	result := database.DBConn.Create(&user)

	if result.Error != nil {
		log.Println("Error creating user:", result.Error)
		returnObject["message"] = "User already exists or database error"
		c.JSON(400, returnObject)
		return
	}

	returnObject["status"] = "OK"
	returnObject["message"] = "User registered successfully"
	c.JSON(201, returnObject)
}

// Logout Handler
func Logout(c *gin.Context) {
	returnObject := gin.H{
		"status":  "OK",
		"message": "Logged out successfully",
	}
	c.JSON(200, returnObject)
}

func TokenAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get the Authorization header
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Authorization header is missing",
            })
            c.Abort()
            return
        }

        // Remove "Bearer " prefix if present
        tokenString = strings.TrimPrefix(tokenString, "Bearer ")

        // Validate the token
        claims, err := helper.ValidateToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Invalid or expired token",
            })
            c.Abort()
            return
        }

        // Check token expiration
        if time.Now().After(time.Unix(claims.ExpiresAt.Unix(), 0)) {
            c.JSON(http.StatusUnauthorized, gin.H{
               "error": "Invalid or expired token",
            })
            c.Abort()
            return
        }

        // Set user ID in context for further use
        c.Set("userID", claims.UserId)
        c.Next()
    }
}

func RefreshToken(c *gin.Context) {
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{
            "status":  "ERROR",
            "message": "Unauthorized - userID not found in context",
        })
        return
    }

    var user model.User
    if err := database.DBConn.First(&user, "id = ?", userID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "status":  "ERROR",
            "message": "User not found",
        })
        return
    }

    // Generate a new token
    newToken, err := helper.GenerateToken(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  "ERROR",
            "message": "Failed to generate new token",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status": "OK",
        "token":  newToken,
        "user":   user,
    })
}



