package comment

import (
	"event-planner/entity"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	t.Run("TestCreateCommentSuccess", func(t *testing.T) {
		commentService := NewCommentService(mockRepositoryComment{})
		err := commentService.CreateComment(entity.Comment{})
		assert.Nil(t, err)
	})

	t.Run("TestCreateCommentError", func(t *testing.T) {
		commentService := NewCommentService(mockRepositoryCommentError{})
		err := commentService.CreateComment(entity.Comment{})
		assert.NotNil(t, err)
	})
}

type mockRepositoryComment struct{}

type mockRepositoryCommentError struct{}

func (m mockRepositoryComment) CreateComment(comment entity.Comment) error {
	return nil
}

func (m mockRepositoryCommentError) CreateComment(comment entity.Comment) error {
	return fmt.Errorf("error create comment")
}
