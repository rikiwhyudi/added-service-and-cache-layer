package models

import "time"

type User struct {
	ID          int           `json:"id"`
	Name        string        `json:"name" gorm:"type: varchar(255)"`
	Email       string        `json:"email" gorm:"type: varchar(255);unique"`
	Password    string        `json:"password" gorm:"type: varchar(255)"`
	Status      string        `json:"status" gorm:"type: varchar(255)"`
	Transaction []Transaction `json:"transaction" gorm:"foreignkey:UserID"` //has many
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

type UserResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

func (UserResponse) TableName() string {
	return "users"
}
