package controller

import (
	"Trainify/database"
	"Trainify/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// SaveWorkoutHistory saves a workout to the user's history
func SaveWorkoutHistory(c *gin.Context) {
	var history model.WorkoutHistory

	if err := c.ShouldBindJSON(&history); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Set the workout_date to the current date if not provided
	if history.WorkoutDate.IsZero() {
		history.WorkoutDate = time.Now()
	}

	// Save the history to the database
	if err := database.DBConn.Create(&history).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save workout history"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Workout history saved successfully"})
}


// GetWorkoutsByDate retrieves all workouts for a specific user on a specific date
func GetWorkoutsByDate(c *gin.Context) {
	userID := c.Param("user_id")
	date := c.Query("date") // Date passed as a query parameter, e.g., ?date=2024-11-25

	var workouts []model.WorkoutHistory
	if err := database.DBConn.Where("user_id = ? AND DATE(workout_date) = ?", userID, date).Find(&workouts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve workouts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"workouts": workouts})
}

