package event

import (
	"event-planner/entity"
)

type EventRepositoryInterface interface {
	GetAll() ([]entity.Event, error)
	GetEventById(id int) (entity.Event, error)
	GetEventByIdUser(id int) ([]entity.Event, error)
	CreateEvent(event entity.Event) error
	DeleteEvent(id, userID int) error
	UpdateEvent(id int, event entity.Event) error
}
