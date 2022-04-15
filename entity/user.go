package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber"`
	Address     string `json:"address" form:"address"`
	Images      string `json:"images" form:"images"`
	Event       []Event
	JoinEvent   []JoinEvent
	Comment     []Comment
}

type UserResponse struct {
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber"`
	Address     string `json:"address" form:"address"`
	Images      string `json:"images" form:"images"`
}
