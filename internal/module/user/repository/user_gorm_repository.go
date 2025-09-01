package user_repository

import (
	"errors"
	"sync"

	"github.com/VeluraDoc/Velura-Backend-Main/internal/config"
	user_model "github.com/VeluraDoc/Velura-Backend-Main/internal/module/user/model"
)

type GormUserRepository struct{}

var (
	userRepoInstance *GormUserRepository
	once             sync.Once
)

func GetRepository() *GormUserRepository {
	once.Do(func() {
		userRepoInstance = &GormUserRepository{}
	})
	return userRepoInstance
}

func (r *GormUserRepository) Save(user *user_model.User) error {
	return config.DB.Save(user).Error
}

func (r *GormUserRepository) Delete(id string) error {
	return config.DB.Delete(&user_model.User{}, "id = ?", id).Error
}

func (r *GormUserRepository) FindByEmail(email string) (*user_model.User, error) {
	var user user_model.User
	result := config.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, errors.New("user with email " + email + " not found")
	}
	return &user, nil
}

func (r *GormUserRepository) FindByID(id string) (*user_model.User, error) {
	var user user_model.User
	result := config.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, errors.New("user with id " + id + " not found")
	}
	return &user, nil
}
