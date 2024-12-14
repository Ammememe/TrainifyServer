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
    
    // Add logging middleware
    router.Use(gin.Logger())
    router.Use(gin.Recovery())

    // CORS configuration
    router.Use(cors.New(cors.Config{
        AllowOrigins: []string{
            "http://localhost:3000",
            "http://localhost:8001",
            "http://swipetofit.com",
            "http://api.swipetofit.com",
            "http://login.swipetofit.com",
        },
        AllowMethods: []string{
            "GET",
            "POST",
            "PUT",
            "PATCH",
            "DELETE",
            "HEAD",
            "OPTIONS",
        },
        AllowHeaders: []string{
            "Origin",
            "Content-Length",
            "Content-Type",
            "Authorization",
            "Accept",
            "X-Requested-With",
            "Access-Control-Allow-Origin",
            "Access-Control-Allow-Headers",
            "Access-Control-Allow-Methods",
            "Access-Control-Allow-Credentials",
        },
        ExposeHeaders: []string{
            "Content-Length",
            "Access-Control-Allow-Origin",
            "Access-Control-Allow-Headers",
            "Access-Control-Allow-Methods",
            "Access-Control-Allow-Credentials",
        },
        AllowCredentials: true,
        AllowWildcard:   false,
        MaxAge:          12 * time.Hour,
    }))

    // Add OPTIONS handler for preflight requests
    router.OPTIONS("/*path", func(c *gin.Context) {
        c.Status(204)
        c.Done()
    })

    routes.SetupRoutes(router)

    log.Printf("Server running on port %s", port)
    log.Fatal(router.Run(":" + port))
}