package auth

import (
	"errors"
	"event-planner/delivery/middleware"
	password2 "event-planner/delivery/password"
	"event-planner/entity"
	"fmt"
	"gorm.io/gorm"
)

type AuthRepository struct {
	database *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		database: db,
	}
}

func (ar *AuthRepository) Login(email string, password string) (string, error) {
	var user entity.User
	tx := ar.database.Where("email = ?", email).Find(&user)
	if tx.Error != nil {
		return "failed", tx.Error
	}

	//jika data user dengan email tsb tidak ada
	if tx.RowsAffected == 0 {
		return "user not found", errors.New("user not found")
	}

	fmt.Println("data user", user.Password)
	fmt.Println("data rows", tx.RowsAffected)

	//jika ada, maka cek

	if password2.CheckPasswordHash(password, user.Password) != true {
		return "password incorrect", errors.New("password incorrect")
	}

	//jika password sama
	token, err := middleware.CreateToken(int(user.ID), user.Name)
	if err != nil {
		return "create token failed", err
	}

	return token, nil

}
