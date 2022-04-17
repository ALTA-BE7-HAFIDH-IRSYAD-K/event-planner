package participant

import (
	"event-planner/entity"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	t.Run("TestCreateParticipantSuccess", func(t *testing.T) {
		participantService := NewParticipatService(mockParticipantRepository{})
		err := participantService.CreateParticipantion(entity.JoinEvent{})
		assert.Nil(t, err)
	})

	t.Run("TestCreateParticipantError", func(t *testing.T) {
		participantService := NewParticipatService(mockParticipantRepositoryError{})
		err := participantService.CreateParticipantion(entity.JoinEvent{})
		assert.NotNil(t, err)
	})
}

type mockParticipantRepository struct{}

type mockParticipantRepositoryError struct{}

func (m mockParticipantRepository) CreateParticipantion(participant entity.JoinEvent) error {
	return nil
}

func (m mockParticipantRepositoryError) CreateParticipantion(participant entity.JoinEvent) error {
	return fmt.Errorf("error create participant")
}
