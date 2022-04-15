package comment

import (
	"event-planner/entity"
	_comment "event-planner/repository/comment"
)

type ServiceComment struct {
	commentRepository _comment.RepositoryCommentInterface
}

func NewCommentService(commentRepo _comment.RepositoryCommentInterface) ServiceCommentInterface {
	return &ServiceComment{
		commentRepository: commentRepo,
	}
}

func (s ServiceComment) CreateComment(comment entity.Comment) error {
	//TODO implement me
	err := s.commentRepository.CreateComment(comment)
	return err
}
