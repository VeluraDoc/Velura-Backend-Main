package model

import (
	"errors"

	"github.com/VeluraDoc/Velura-Backend-Main/config"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) HashPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return errors.New("failed to hash password")
	}

	u.Password = string(hash)

	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}

func (u *User) Save() error {
	db := config.ConnectToDB()

	result := db.Save(u)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *User) Delete() error {
	db := config.ConnectToDB()

	result := db.Delete(&User{}, "id = ?", u.ID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
