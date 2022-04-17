package participant

import _entity "event-planner/entity"

type ParticipantRepositoryInterface interface {
	CreateParticipantion(participation _entity.JoinEvent) error
}
