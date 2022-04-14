package entity

import (
	"gorm.io/gorm"
)

type JoinEvent struct {
	gorm.Model
	UserID  uint    `json:"user_id" form:"user_id"`
	EventID uint    `json:"event_id" form:"event_id"`
	Event   []Event `gorm:"foreignKey:ID;references:EventID"`
}
