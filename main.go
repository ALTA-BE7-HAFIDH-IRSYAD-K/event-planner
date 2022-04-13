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
)

func main() {
	configs := _configs.GetConfig()
	db := driver.InitDB(configs)

	userRepo := userrepo.NewUserRepository(db)
	userUseCase := userusecase.NewUserUseCase(userRepo)
	userHandler := userhandler.NewUserHandler(userUseCase)

	e := echo.New()

	router.RegisterPath(e, userHandler)
	log.Fatal(e.Start(fmt.Sprintf(":%v", configs.Port)))
}
