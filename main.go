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
            "http://swipetofit.com",
            "http://api.swipetofit.com",
            "http://login.swipetofit.com",
        },
        AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowHeaders: []string{
            "Origin",
            "Content-Type",
            "Accept",
            "Authorization",
            "X-Requested-With",
        },
        ExposeHeaders: []string{"Content-Length"},
        AllowCredentials: true,
        AllowOriginFunc: func(origin string) bool {
            return true // Be careful with this in production
        },
        MaxAge: 12 * time.Hour,
    }))
    
    // Add a custom middleware to handle OPTIONS requests
    router.Use(func(c *gin.Context) {
        if c.Request.Method == "OPTIONS" {
            c.Header("Access-Control-Allow-Origin", c.GetHeader("Origin"))
            c.Header("Access-Control-Allow-Credentials", "true")
            c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
            c.Header("Access-Control-Allow-Headers", "Authorization, Origin, Content-Type, Accept")
            c.Status(204)
            c.Abort()
            return
        }
        c.Next()
    })

    // Add OPTIONS handler for preflight requests
    router.OPTIONS("/*path", func(c *gin.Context) {
        c.Status(204)
        c.Done()
    })

    routes.SetupRoutes(router)

    log.Printf("Server running on port %s", port)
    log.Fatal(router.Run(":" + port))
}