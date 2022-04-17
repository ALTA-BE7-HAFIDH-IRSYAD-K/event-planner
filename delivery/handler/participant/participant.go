package participant

import (
	"event-planner/delivery/middleware"
	"event-planner/delivery/response"
	"event-planner/entity"
	"event-planner/service/participant"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ParticipantHandler struct {
	participantService participant.ServiceInterface
}

func NewParticipantHandler(participantService participant.ServiceInterface) *ParticipantHandler {
	return &ParticipantHandler{
		participantService: participantService,
	}
}

func (ph *ParticipantHandler) CreateParticipantHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var dataParticipant entity.JoinEvent

		id := middleware.ExtractToken(c)

		dataParticipant.UserID = uint(id)

		errBind := c.Bind(&dataParticipant)

		fmt.Println("data", dataParticipant)

		if errBind != nil {
			return errBind
		}

		err := ph.participantService.CreateParticipantion(dataParticipant)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed to create participant"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccessWithoutData("success to create participant"))
	}
}
