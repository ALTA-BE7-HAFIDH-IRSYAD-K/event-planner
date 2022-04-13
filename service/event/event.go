package event

import (
	"event-planner/entity"
	"event-planner/repository/event"
)

type EventUseCase struct {
	EventRepository event.EventRepositoryInterface
}

func NewEventUseCase(eventRepo event.EventRepositoryInterface) EventUseCaseInterface {
	return &EventUseCase{
		EventRepository: eventRepo,
	}

}

func (euc *EventUseCase) GetAll() ([]entity.Event, error) {
	Products, err := euc.EventRepository.GetAll()
	return Products, err
}

func (euc *EventUseCase) GetEventById(id int) (entity.Event, error) {
	Product, err := euc.EventRepository.GetEventById(id)
	return Product, err
}

func (euc *EventUseCase) GetEventByIdUser(id int) ([]entity.Event, error) {
	Products, err := euc.EventRepository.GetEventByIdUser(id)
	return Products, err
}

func (euc *EventUseCase) CreateEvent(Product entity.Event) error {
	err := euc.EventRepository.CreateEvent(Product)
	return err
}

func (euc *EventUseCase) DeleteEvent(id, userID int) error {
	err := euc.EventRepository.DeleteEvent(id, userID)
	return err
}

func (euc *EventUseCase) UpdateEvent(id int, Product entity.Event) error {
	err := euc.EventRepository.UpdateEvent(id, Product)
	return err
}
