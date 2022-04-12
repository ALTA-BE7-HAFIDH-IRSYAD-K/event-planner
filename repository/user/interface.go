package user

import (
	"event-planner/entity"
)

type UserRepositoryInterface interface {
	GetAll() ([]entity.User, error)
	GetUserById(id int) (entity.User, error)
	CreateUser(user entity.User) error
	DeleteUser(id int) error
	UpdateUser(id int, user entity.User) error
}
