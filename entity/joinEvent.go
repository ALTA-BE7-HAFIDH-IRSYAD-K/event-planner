package entity

import (
	"gorm.io/gorm"
)

type JoinEvent struct {
	gorm.Model
	UserID  uint `json:"user_id" form:"user_id"`
	EventID uint `json:"event_id" form:"event_id"`
	User    User `json:"user" gorm:"foreignKey:UserID;references:ID"`
}

type JoinEventResponse struct {
	UserID  uint         `json:"user_id" form:"user_id"`
	EventID uint         `json:"event_id" form:"event_id"`
	User    UserResponse `json:"user"`
}
