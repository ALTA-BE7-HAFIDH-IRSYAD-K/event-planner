package driver

import (
	"event-planner/configs"
	"event-planner/entity"
	"fmt"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(configs *configs.AppConfig) *gorm.DB {

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		configs.Database.Username,
		configs.Database.Password,
		configs.Database.Address,
		configs.Database.Port,
		configs.Database.Name,
	)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Info("failed to connect database :", err)
		panic(err)
	}

	InitialMigration(db)
	return db
}

func InitialMigration(db *gorm.DB) {
	err := db.AutoMigrate(&entity.Comment{})
	err = db.AutoMigrate(&entity.JoinEvent{})
	err = db.AutoMigrate(&entity.Event{})
	err = db.AutoMigrate(&entity.User{})

	if err != nil {
		log.Info("error auto migrate", err)
	}

}
