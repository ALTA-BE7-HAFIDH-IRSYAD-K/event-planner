package main

import (
	_configs "event-planner/configs"
	"event-planner/driver"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	configs := _configs.GetConfig()
	db := driver.InitDB(configs)

	e := echo.New()

	fmt.Println("db", db)
	log.Fatal(e.Start(fmt.Sprintf(":%v", configs.Port)))
}
