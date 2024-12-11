// In middleware/auth.go
package middleware

import (
    "Trainify/helper"
    "github.com/gin-gonic/gin"
    "net/http"
    "strings"
    "time"
)

func Authenticate() gin.HandlerFunc {
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
                "error": "Token has expired",
            })
            c.Abort()
            return
        }

        // Set user ID in context for further use
        c.Set("userID", claims.UserId)
        c.Next()
    }
}