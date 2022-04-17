package comment

import _entity "event-planner/entity"

type ServiceCommentInterface interface {
	CreateComment(comment _entity.Comment) error
}
