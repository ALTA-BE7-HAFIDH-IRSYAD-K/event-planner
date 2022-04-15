package participant

import _entity "event-planner/entity"

type RepositoryCommentInterface interface {
	CreateComment(comment _entity.Comment) error
}
