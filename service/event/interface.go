package event

import (
	"event-planner/entity"
)

type EventUseCaseInterface interface {
	GetAll() ([]entity.Event, error)
	GetEventById(id int) (entity.Event, error)
	GetEventByIdUser(id int) ([]entity.Event, error)
	CreateEvent(Product entity.Event) error
	DeleteEvent(id, userID int) error
	UpdateEvent(id int, Product entity.Event) error
}
