package participant

import _entity "event-planner/entity"

type ServiceInterface interface {
	CreateParticipantion(participation _entity.JoinEvent) error
}
