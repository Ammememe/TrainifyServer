package main

import (
	"Trainify/database"
	"Trainify/routes"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
    if err := godotenv.Load(); err != nil {
        log.Println("Error loading .env file")
    }

    database.Connect()
}

func main() {
    // Close the db connection using defer clause
    sqlDb, err := database.DBConn.DB()
    if err != nil {
        log.Fatalf("Error getting DB connection: %v", err)
    }
    defer sqlDb.Close()

    port := os.Getenv("PORT")
    if port == "" {
        port = "8001"
    }

    gin.SetMode(gin.ReleaseMode)
    router := gin.New()
    router.Use(func(c *gin.Context) {
        log.Printf("Incoming request from Origin: %s", c.Request.Header.Get("Origin"))
        log.Printf("Request Method: %s", c.Request.Method)
        log.Printf("Request Headers: %v", c.Request.Header)
        c.Next()
    })


    router.Use(cors.New(cors.Config{
        AllowOrigins: []string{
            "http://172.232.130.101",    // Frontend IP
            "http://172.232.130.217",    // API IP
            "http://172.232.131.139",    // Auth IP
            "http://swipetofit.com",
            "http://api.swipetofit.com",
            "http://login.swipetofit.com",
        },
        AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowHeaders: []string{
            "Origin",
            "Auth-Token",
            "Token",
            "Content-Type",
            "Authorization",
            "Access-Control-Allow-Credentials",  // Add this
            "Access-Control-Allow-Headers",      // Add this
            "Access-Control-Allow-Origin",       // Add this
            "Accept",                           // Add this
            "X-Requested-With",            
        },
        ExposeHeaders:    []string{"Content-Length", "Authorization"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,  // Cache preflight requests
    }))
   
    
    routes.SetupRoutes(router)

    log.Printf("Server running on port %s", port)
    log.Fatal(router.Run(":" + port))
}
