package participant

import _entity "event-planner/entity"

type ServiceInterface interface {
	GetAllParticipantions() ([]_entity.JoinEvent, error)
	CreateParticipantion(participation _entity.JoinEvent) error
}
