package user

import (
	"event-planner/entity"
)

type UserUseCaseInterface interface {
	GetAll() ([]entity.UserResponse, error)
	GetUserById(id int) (entity.UserResponse, error)
	CreateUser(user entity.User) error
	DeleteUser(id int) error
	UpdateUser(id int, user entity.User) error
}
