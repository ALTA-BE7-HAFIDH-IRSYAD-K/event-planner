package event

import (
	"event-planner/delivery/middleware"
	"event-planner/delivery/response"
	"event-planner/entity"
	"event-planner/service/event"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

type EventHandler struct {
	eventUseCase event.EventUseCaseInterface
}

func NewEventHandler(eventUseCase event.EventUseCaseInterface) *EventHandler {
	return &EventHandler{
		eventUseCase: eventUseCase,
	}
}

func (eh *EventHandler) GetAllHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		events, err := eh.eventUseCase.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed to get all event"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("success get all event", events))
	}
}

func (eh *EventHandler) GetEventById() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id, _ = strconv.Atoi(c.Param("id"))

		event, err := eh.eventUseCase.GetEventById(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("Event not found"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("success get event by id", event))
	}
}

func (eh *EventHandler) GetEventByIdUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id, _ = strconv.Atoi(c.Param("id"))
		id = middleware.ExtractToken(c)
		events, err := eh.eventUseCase.GetEventByIdUser(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("success get events by id user", events))
	}
}

func (eh *EventHandler) CreateEvent() echo.HandlerFunc {
	return func(c echo.Context) error {
		var events EventRequest
		c.Bind(&events)

		file, _ := c.FormFile("images")
		id := middleware.ExtractToken(c)
		now := time.Now()
		year := now.Year()
		month := now.Month()
		day := now.Day()
		hour := now.Hour()
		minute := now.Minute()
		second := now.Second()

		fileName := fmt.Sprint(id, ".", year, month, day, hour, minute, second)
		fmt.Println(fileName)
		file.Filename = fileName + ".png"

		date, err := time.Parse("2006-01-02", events.Date)

		var eventData = entity.Event{

			Name:        events.Name,
			Date:        date,
			Location:    events.Location,
			Images:      fileName,
			Description: events.Description,
			Category:    events.Category,
			UserID:      id,
		}
		fmt.Println(events)
		err = eh.eventUseCase.CreateEvent(file, eventData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed to create product"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccessWithoutData("success Create event"))
	}
}

func (eh *EventHandler) DeleteEvent() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id, _ = strconv.Atoi(c.Param("id"))
		var userID = middleware.ExtractToken(c)

		err := eh.eventUseCase.DeleteEvent(id, userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccessWithoutData("success delete product by id"))
	}
}

func (eh *EventHandler) UpdateEvent() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id, _ = strconv.Atoi(c.Param("id"))
		var event entity.Event
		c.Bind(&event)

		err := eh.eventUseCase.UpdateEvent(id, event)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("success update product by id", event))
	}
}
