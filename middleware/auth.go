package middleware

import (
	"Trainify/helper"
	"log"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {

	token := c.GetHeader("token")

	if token == "" {
		c.JSON(401, gin.H{
			"message": "Unauthorized access",
		})
		c.Abort()
		return
	}

	claims, msg := helper.ValidateToken(token)

	log.Println(claims)

	if msg != ""{
		c.JSON(401, gin.H{
			"error": msg,
		})
		c.Abort()
		return
	}

	c.Next()

}