package participant

import (
	"event-planner/delivery/middleware"
	"event-planner/delivery/response"
	"event-planner/entity"
	"event-planner/service/participant"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ParticipantHandler struct {
	participantService participant.ServiceInterface
}

func NewParticipantHandler(participantService participant.ServiceInterface) *ParticipantHandler {
	return &ParticipantHandler{
		participantService: participantService,
	}
}

func (ph *ParticipantHandler) GetAllParticipantHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		participants, err := ph.participantService.GetAllParticipantions()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed to get all data participant"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("success to get all participant", participants))
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
