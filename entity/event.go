package entity

import (
	"gorm.io/gorm"
	"time"
)

type Event struct {
	gorm.Model
	Name        string    `json:"name" form:"name"`
	Date        time.Time `json:"date" form:"date"`
	Location    string    `json:"location" form:"location"`
	Status      string    `json:"status" form:"status"`
	Images      string    `json:"images" form:"images"`
	Description string    `json:"description" form:"description"`
	Category    string    `json:"category" form:"category"`
	UserID      int
	JoinEvent   []JoinEvent
	Comment     []Comment
}
