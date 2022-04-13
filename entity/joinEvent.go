package entity

import (
	"gorm.io/gorm"
)

type JoinEvent struct {
	gorm.Model
	UserID  int
	EventID int
}
