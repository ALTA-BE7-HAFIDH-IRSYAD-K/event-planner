package user

import (
	"event-planner/entity"
	"event-planner/repository/user"
)

type UserUseCase struct {
	UserRepository user.UserRepositoryInterface
}

func NewUserUseCase(userRepo user.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		UserRepository: userRepo,
	}

}

func (uuc *UserUseCase) GetAll() ([]entity.User, error) {
	users, err := uuc.UserRepository.GetAll()
	return users, err
}

func (uuc *UserUseCase) GetUserById(id int) (entity.User, error) {
	user, err := uuc.UserRepository.GetUserById(id)
	return user, err
}

func (uuc *UserUseCase) CreateUser(user entity.User) error {
	err := uuc.UserRepository.CreateUser(user)
	return err
}

func (uuc *UserUseCase) DeleteUser(id int) error {
	err := uuc.UserRepository.DeleteUser(id)
	return err
}

func (uuc *UserUseCase) UpdateUser(id int, user entity.User) error {
	err := uuc.UserRepository.UpdateUser(id, user)
	return err
}
