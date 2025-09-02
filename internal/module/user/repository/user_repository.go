package user_repository

import (
	user_model "github.com/VeluraOpenSource/Velura_Documents_Service/internal/module/user/model"
)

type UserRepository interface {
	Save(user *user_model.User) error
	Delete(id string) error
	FindByEmail(email string) (*user_model.User, error)
	FindByID(id string) (*user_model.User, error)
}
