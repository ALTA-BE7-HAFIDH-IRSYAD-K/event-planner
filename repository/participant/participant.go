package participant

import (
	"event-planner/entity"
	"fmt"

	"gorm.io/gorm"
)

type ParticipationRepository struct {
	database *gorm.DB
}

func NewParticipationRepository(database *gorm.DB) *ParticipationRepository {
	return &ParticipationRepository{
		database: database,
	}
}

func (pr *ParticipationRepository) CreateParticipantion(participant entity.JoinEvent) error {
	fmt.Println("participationRepo", participant)
	tx := pr.database.Save(&participant)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (pr *ParticipationRepository) GetAllParticipantions() ([]entity.JoinEvent, error) {
	var participant []entity.JoinEvent

	tx := pr.database.Find(&participant)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return participant, nil
}
