package main

import (
	_configs "event-planner/configs"
	_commentHandler "event-planner/delivery/handler/comment"
	_participantHandler "event-planner/delivery/handler/participant"
	_middleware "event-planner/delivery/middleware"
	"event-planner/delivery/router"
	"event-planner/driver"
	_commentRepo "event-planner/repository/comment"
	_participantRepo "event-planner/repository/participant"
	_commentService "event-planner/service/comment"
	_participantService "event-planner/service/participant"
	"fmt"
	"github.com/labstack/echo/v4"
	_echoMiddleware "github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"

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

	participantRepo := _participantRepo.NewParticipationRepository(db)
	participantService := _participantService.NewParticipatService(participantRepo)
	participantHandler := _participantHandler.NewParticipantHandler(participantService)

	commentRepo := _commentRepo.NewCommentRepository(db)
	commentService := _commentService.NewCommentService(commentRepo)
	commentHandler := _commentHandler.NewCommentHandler(commentService)

	e := echo.New()

	e.Use(_echoMiddleware.RemoveTrailingSlash())
	e.Use(_middleware.CustomLogger())
	e.Use(_echoMiddleware.CORSWithConfig(_echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	router.RegisterAuthPath(e, authHandler)
	router.RegisterPath(e, userHandler, eventHandler, participantHandler, commentHandler)
	log.Fatal(e.Start(fmt.Sprintf(":%v", configs.Port)))
}
