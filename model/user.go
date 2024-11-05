package model

import "time"

type User struct {
	ID        uint   `json:"id" gorm:"primary_key;auto_increment;NOT NULL"`
	FirstName string `json:"first_name" gorm:"column:first_name;NOT NULL"`
	LastName  string `json:"last_name" gorm:"column:last_name;NOT NULL"`
	Email     string `json:"email" gorm:"column:email;unique;NOT NULL"`
	Password  string `json:"_" gorm:"column:password;hashed"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;NOT NULL;autoCreateTime"`
}