package routes

import (
	"Trainify/controller"
	"Trainify/middleware"

	"github.com/gin-gonic/gin"
)

import "github.com/gin-contrib/cors"

// In your main setup function



func SetupRoutes(r *gin.Engine) {
	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)
	r.POST("/logout", controller.Logout)
	r.Use(cors.Default())
	r.POST("/workout-history", controller.SaveWorkoutHistory)
    r.GET("/workout-history/:user_id", controller.GetWorkoutsByDate)


	

	private := r.Group("/private")

	private.Use(middleware.Authenticate)

	private.GET("refreshtoken", controller.RefreshToken)

}