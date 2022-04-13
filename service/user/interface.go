package user

import (
	"event-planner/entity"
)

type UserUseCaseInterface interface {
	GetAll() ([]entity.User, error)
	GetUserById(id int) (entity.User, error)
	CreateUser(user entity.User) error
	DeleteUser(id int) error
	UpdateUser(id int, user entity.User) error
}
