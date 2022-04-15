package event

import (
	"event-planner/entity"
	"fmt"
	"gorm.io/gorm"
)

type EventRepository struct {
	database *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{
		database: db,
	}
}

func (er *EventRepository) GetAll() ([]entity.Event, error) {
	var events []entity.Event
	tx := er.database.Find(&events)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return events, nil

}

func (er *EventRepository) GetEventById(id int) (entity.Event, error) {
	var events entity.Event
	tx := er.database.Preload("JoinEvent").Preload("JoinEvent.User").Preload("Comment").Preload("Comment.User").Where("id = ? ", id).First(&events)
	//tx := er.database.Where("id = ?", id).First(&events)
	if tx.Error != nil {
		return events, tx.Error
	}
	return events, nil

}

func (er *EventRepository) GetEventByIdUser(id int) ([]entity.Event, error) {
	var events []entity.Event
	tx := er.database.Where("user_id = ?", id).Find(&events)
	if tx.Error != nil {
		return events, tx.Error
	}
	return events, nil

}

func (er *EventRepository) CreateEvent(event entity.Event) error {

	tx := er.database.Save(&event)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

func (er *EventRepository) DeleteEvent(id, userID int) error {
	var Event entity.Event
	err := er.database.Where("id =? and user_id = ?", id, userID).First(&Event).Error
	if err != nil {
		return err
	}
	er.database.Delete(&Event)

	return nil

}

func (er *EventRepository) UpdateEvent(id int, event entity.Event) error {

	tx := er.database.Where("id = ?", id).Updates(&event)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {

		return fmt.Errorf("%s", "error")
	}
	return nil

}
