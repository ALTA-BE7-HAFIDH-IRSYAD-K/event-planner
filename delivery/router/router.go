package router

import (
	"event-planner/delivery/handler/auth"
	"event-planner/delivery/handler/comment"
	"event-planner/delivery/handler/event"
	"event-planner/delivery/handler/participant"
	"event-planner/delivery/handler/user"
	"event-planner/delivery/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterAuthPath(e *echo.Echo, ah *auth.AuthHandler) {
	e.POST("/auth", ah.LoginHandler())
}

func RegisterPath(e *echo.Echo, uh *user.UserHandler, eh *event.EventHandler, ph *participant.ParticipantHandler, ch *comment.HandlerComment) {

	e.GET("/users", uh.GetAllHandler(), middleware.JWTMiddleware())
	e.GET("/users/profile", uh.GetUserById(), middleware.JWTMiddleware())
	e.POST("/users", uh.CreateUser())
	e.DELETE("/users", uh.DeleteUser(), middleware.JWTMiddleware())
	e.PUT("/users", uh.UpdateUser(), middleware.JWTMiddleware())

	e.GET("/events", eh.GetAllHandler())
	e.GET("/events/:id", eh.GetEventById())
	e.GET("/events/profile", eh.GetEventByIdUser(), middleware.JWTMiddleware())
	e.POST("/events", eh.CreateEvent(), middleware.JWTMiddleware())
	e.DELETE("/events/:id", eh.DeleteEvent(), middleware.JWTMiddleware())
	e.PUT("/events/:id", eh.UpdateEvent(), middleware.JWTMiddleware())

	e.POST("/events/participation", ph.CreateParticipantHandler(), middleware.JWTMiddleware())

	e.POST("/events/comment", ch.CreateParticipantHandler(), middleware.JWTMiddleware())

}
