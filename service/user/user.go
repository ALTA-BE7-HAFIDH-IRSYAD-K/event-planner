package user

import (
	password2 "event-planner/delivery/password"
	"event-planner/entity"
	"event-planner/repository/user"

	"github.com/jinzhu/copier"
)

type UserUseCase struct {
	UserRepository user.UserRepositoryInterface
}

func NewUserUseCase(userRepo user.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		UserRepository: userRepo,
	}
}

func (uuc *UserUseCase) GetAll() ([]entity.UserResponse, error) {
	users, err := uuc.UserRepository.GetAll()
	var userRes []entity.UserResponse
	copier.Copy(&userRes, &users)
	return userRes, err
}

func (uuc *UserUseCase) GetUserById(id int) (entity.UserResponse, error) {
	user, err := uuc.UserRepository.GetUserById(id)
	userRes := entity.UserResponse{}
	copier.Copy(&userRes, &user)
	return userRes, err
}

func (uuc *UserUseCase) CreateUser(user entity.User) error {
	password, _ := password2.HashPassword(user.Password)
	user.Password = password
	err := uuc.UserRepository.CreateUser(user)
	return err
}

func (uuc *UserUseCase) DeleteUser(id int) error {
	err := uuc.UserRepository.DeleteUser(id)
	return err
}

func (uuc *UserUseCase) UpdateUser(id int, user entity.User) error {
	password, _ := password2.HashPassword(user.Password)
	user.Password = password
	err := uuc.UserRepository.UpdateUser(id, user)
	return err
}
