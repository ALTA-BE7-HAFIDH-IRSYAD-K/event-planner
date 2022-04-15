package participant

import (
	"event-planner/entity"
	_participant "event-planner/repository/participant"
)

type ParticipantService struct {
	participantRepository _participant.ParticipantRepositoryInterface
}

func NewParticipatService(participantRepo _participant.ParticipantRepositoryInterface) ServiceInterface {
	return &ParticipantService{
		participantRepository: participantRepo,
	}

}

func (p ParticipantService) CreateParticipantion(participation entity.JoinEvent) error {
	//TODO implement me
	err := p.participantRepository.CreateParticipantion(participation)

	return err
}
