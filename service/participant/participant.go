package participant

import (
	"event-planner/entity"
	_participant "event-planner/repository/participant"
	"fmt"
)

type ParticipantService struct {
	participantRepository _participant.ParticipantRepositoryInterface
}

func NewParticipatService(participantRepo _participant.ParticipantRepositoryInterface) ServiceInterface {
	return &ParticipantService{
		participantRepository: participantRepo,
	}

}

func (p ParticipantService) GetAllParticipantions() ([]entity.JoinEvent, error) {
	//TODO implement me
	participant, err := p.participantRepository.GetAllParticipantions()

	if err != nil {
		return nil, err
	}

	return participant, nil
}

func (p ParticipantService) CreateParticipantion(participation entity.JoinEvent) error {
	//TODO implement me
	fmt.Println("participationService", participation)
	err := p.participantRepository.CreateParticipantion(participation)

	return err
}
