package user

import (
	"event-planner/entity"
	"fmt"
	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (ur *UserRepository) GetAll() ([]entity.User, error) {
	var users []entity.User
	tx := ur.database.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil

}

func (ur *UserRepository) GetUserById(id int) (entity.User, error) {
	var users entity.User
	tx := ur.database.Where("id = ?", id).First(&users)
	if tx.Error != nil {
		return users, tx.Error
	}
	return users, nil

}

func (ur *UserRepository) CreateUser(user entity.User) error {

	tx := ur.database.Create(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

func (ur *UserRepository) DeleteUser(id int) error {
	var users entity.User
	tx := ur.database.Where("id = ?", id).Delete(&users)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

func (ur *UserRepository) UpdateUser(id int, user entity.User) error {

	tx := ur.database.Where("id = ?", id).Updates(&user)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {

		return fmt.Errorf("%s", "error")
	}
	return nil

}
