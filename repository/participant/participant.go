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
