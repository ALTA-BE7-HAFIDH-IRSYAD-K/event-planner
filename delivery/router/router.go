package router

import (
	"event-planner/delivery/handler/user"
	"event-planner/delivery/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, uh *user.UserHandler) {

	e.GET("/users", uh.GetAllHandler(), middleware.JWTMiddleware())
	e.GET("/users/profile", uh.GetUserById(), middleware.JWTMiddleware())
	e.POST("/users", uh.CreateUser())
	e.DELETE("/users", uh.DeleteUser(), middleware.JWTMiddleware())
	e.PUT("/users", uh.UpdateUser(), middleware.JWTMiddleware())

}
