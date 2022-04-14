package main

import (
	_configs "event-planner/configs"
	_middleware "event-planner/delivery/middleware"
	"event-planner/delivery/router"
	"event-planner/driver"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_echoMiddleware "github.com/labstack/echo/v4/middleware"

	userhandler "event-planner/delivery/handler/user"
	userrepo "event-planner/repository/user"
	userusecase "event-planner/service/user"

	authhandler "event-planner/delivery/handler/auth"
	authrepo "event-planner/repository/auth"
	authusecase "event-planner/service/auth"

	eventhandler "event-planner/delivery/handler/event"
	eventrepository "event-planner/repository/event"
	eventusecase "event-planner/service/event"
)

func main() {
	configs := _configs.GetConfig()
	db := driver.InitDB(configs)

	userRepo := userrepo.NewUserRepository(db)
	userUseCase := userusecase.NewUserUseCase(userRepo)
	userHandler := userhandler.NewUserHandler(userUseCase)

	authRepo := authrepo.NewAuthRepository(db)
	authUseCase := authusecase.NewAuthUseCase(authRepo)
	authHandler := authhandler.NewAuthHandler(authUseCase)

	eventRepo := eventrepository.NewEventRepository(db)
	eventUseCase := eventusecase.NewEventUseCase(eventRepo)
	eventHandler := eventhandler.NewEventHandler(eventUseCase)

	e := echo.New()

	e.Use(_echoMiddleware.RemoveTrailingSlash())
	e.Use(_middleware.CustomLogger())
	e.Use(_echoMiddleware.CORSWithConfig(_echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	router.RegisterAuthPath(e, authHandler)
	router.RegisterPath(e, userHandler, eventHandler)
	log.Fatal(e.Start(fmt.Sprintf(":%v", configs.Port)))
}
