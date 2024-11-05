package controller

import (
	"Trainify/database"
	"Trainify/helper"
	"Trainify/model"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
		c.JSON(400, returnObject)
		return
	}

	var user model.User
	database.DBConn.First(&user, "email = ?", formData.Email)

	if user.ID == 0 {
		log.Println("User not found")
		c.JSON(400, returnObject)
		return
	}

	// Validate password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formData.Password)); err != nil {
		log.Println("Password mismatch")
		c.JSON(401, returnObject)
		return
	}

	token, err := helper.GenerateToken(user)

	if err != nil {
		returnObject["message"] = "Error in generating token"
		c.JSON(401, returnObject)
		return
	}

	returnObject["token"] = token
	returnObject["status"] = "OK"
	returnObject["message"] = "User authenticated"
	c.JSON(200, returnObject)
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

// RefreshToken Handler
func RefreshToken(c *gin.Context) {
	returnObject := gin.H{
		"status":  "OK",
		"message": "Refresh Token Route",
	}
	c.JSON(200, returnObject)
}
