package entity

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID  uint   `json:"user_id" form:"user_id"`
	EventID uint   `json:"event_id" form:"event_id"`
	Comment string `json:"comment" form:"comment"`
	User    User   `json:"user" gorm:"foreignKey:UserID;references:ID"`
}

type CommentResponse struct {
	gorm.Model
	UserID  uint         `json:"user_id" form:"user_id"`
	EventID uint         `json:"event_id" form:"event_id"`
	Comment string       `json:"comment" form:"comment"`
	User    UserResponse `json:"user"`
}
