package model

import "time"

// WorkoutHistory represents a single record of a workout performed by a user
type WorkoutHistory struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	WorkoutID   uint      `json:"workout_id" gorm:"not null"`
	WorkoutDate time.Time `json:"workout_date" gorm:"not null;default:CURRENT_TIMESTAMP"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
}
// TableName overrides the default pluralized table name
func (WorkoutHistory) TableName() string {
	return "workout_history"
}
