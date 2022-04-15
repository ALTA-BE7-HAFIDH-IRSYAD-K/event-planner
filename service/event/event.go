package event

import (
	"event-planner/entity"
	"event-planner/repository/event"
	"event-planner/utils"
	"github.com/jinzhu/copier"
	"mime/multipart"
)

type EventUseCase struct {
	EventRepository event.EventRepositoryInterface
}

func NewEventUseCase(eventRepo event.EventRepositoryInterface) EventUseCaseInterface {
	return &EventUseCase{
		EventRepository: eventRepo,
	}

}

func (euc *EventUseCase) GetAll() ([]entity.EventResponseGetAll, error) {
	events, err := euc.EventRepository.GetAll()
	var eventsResp []entity.EventResponseGetAll
	copier.Copy(&eventsResp, &events)
	return eventsResp, err
}

func (euc *EventUseCase) GetEventById(id int) (entity.EventResponse, error) {
	Product, err := euc.EventRepository.GetEventById(id)
	var eventsResp entity.EventResponse
	copier.Copy(&eventsResp, &Product)
	return eventsResp, err
}

func (euc *EventUseCase) GetEventByIdUser(id int) ([]entity.Event, error) {
	Products, err := euc.EventRepository.GetEventByIdUser(id)
	return Products, err
}

func (euc *EventUseCase) CreateEvent(file *multipart.FileHeader, event entity.Event) error {

	url, err := utils.UploadFile(file, event.Images)

	if err != nil {
		return err
	}

	event.Images = url
	err = euc.EventRepository.CreateEvent(event)
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
