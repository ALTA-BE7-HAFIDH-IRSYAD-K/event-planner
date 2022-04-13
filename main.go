package main

import (
	_configs "event-planner/configs"
	"event-planner/delivery/router"
	"event-planner/driver"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"

	userhandler "event-planner/delivery/handler/user"
	userrepo "event-planner/repository/user"
	userusecase "event-planner/service/user"

	authhandler "event-planner/delivery/handler/auth"
	authrepo "event-planner/repository/auth"
	authusecase "event-planner/service/auth"
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

	e := echo.New()
	
	router.RegisterAuthPath(e, authHandler)
	router.RegisterPath(e, userHandler)
	log.Fatal(e.Start(fmt.Sprintf(":%v", configs.Port)))
}
