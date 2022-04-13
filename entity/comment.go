package entity

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID  int
	EventID int
	Comment string `json:"comment" form:"comment"`
}
