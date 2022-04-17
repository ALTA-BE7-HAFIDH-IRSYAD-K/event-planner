package entity

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	UserID      int
	Name        string    `json:"name" form:"name"`
	Location    string    `json:"location" form:"location"`
	Status      string    `json:"status" form:"status"`
	Images      string    `json:"images" form:"images"`
	Description string    `json:"description" form:"description"`
	Category    string    `json:"category" form:"category"`
	Date        time.Time `json:"date" form:"date"`
	JoinEvent   []JoinEvent
	Comment     []Comment
	User        User
}

type EventResponseGetAll struct {
	gorm.Model
	UserID      int
	Name        string       `json:"name" form:"name"`
	Location    string       `json:"location" form:"location"`
	Status      string       `json:"status" form:"status"`
	Images      string       `json:"images" form:"images"`
	Description string       `json:"description" form:"description"`
	Category    string       `json:"category" form:"category"`
	Date        time.Time    `json:"date" form:"date"`
	User        UserResponse `json:"user"`
}

type EventResponse struct {
	gorm.Model
	UserID      int
	Name        string              `json:"name" form:"name"`
	Location    string              `json:"location" form:"location"`
	Status      string              `json:"status" form:"status"`
	Images      string              `json:"images" form:"images"`
	Description string              `json:"description" form:"description"`
	Category    string              `json:"category" form:"category"`
	Date        time.Time           `json:"date" form:"date"`
	JoinEvent   []JoinEventResponse `json:"join_event"`
	Comment     []CommentResponse   `json:"comment"`
}
