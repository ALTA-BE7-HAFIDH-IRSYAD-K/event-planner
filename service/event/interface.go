package event

import (
	"event-planner/entity"
	"mime/multipart"
)

type EventUseCaseInterface interface {
	GetAll() ([]entity.EventResponseGetAll, error)
	GetEventById(id int) (entity.EventResponse, error)
	GetEventByIdUser(id int) ([]entity.Event, error)
	CreateEvent(file *multipart.FileHeader, event entity.Event) error
	DeleteEvent(id, userID int) error
	UpdateEvent(id int, Product entity.Event) error
}
