package main

import (
    "Trainify/database"
    "Trainify/routes"
    "log"
    "os"

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

    // Add CORS middleware
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"}, // Frontend URL
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Auth-Token", "Token", "Content-Type"},
        AllowCredentials: true,
    }))

    routes.SetupRoutes(router)

    log.Printf("Server running on port %s", port)
    log.Fatal(router.Run(":" + port))
}
