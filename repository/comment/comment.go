package participant

import (
	"event-planner/entity"
	"fmt"

	"gorm.io/gorm"
)

type CommentRepository struct {
	database *gorm.DB
}

func NewCommentRepository(database *gorm.DB) *CommentRepository {
	return &CommentRepository{
		database: database,
	}
}

func (cr *CommentRepository) CreateComment(comment entity.Comment) error {
	fmt.Println("commentRepo", comment)
	tx := cr.database.Save(&comment)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
