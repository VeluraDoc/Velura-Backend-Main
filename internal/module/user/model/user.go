package user_model

import (
	"errors"

	"github.com/VeluraDoc/Velura-Backend-Main/internal/config"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var validate = validator.New()

type User struct {
	ID       string `json:"id" gorm:"primaryKey;type:uuid"`
	Email    string `json:"email" gorm:"unique;index;not null" validate:"required,email"`
	Password string `json:"-" gorm:"not null" validate:"required,min=8"`
	Role     Role   `json:"role" gorm:"type:int;not null;default:2"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return nil
}

func (u *User) Validate() error {
	if err := validate.Struct(u); err != nil {
		return err
	}

	return nil
}

func (u *User) GetRole() string {
	return u.Role.String()
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
	result := config.DB.Save(u)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *User) Delete() error {
	result := config.DB.Delete(&User{}, "id = ?", u.ID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetUserByEmail(email string) (User, error) {
	var user User

	result := config.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return User{}, errors.New("user with email " + email + " not found")
	}

	return user, nil
}

func GetUserByID(id string) (User, error) {
	var user User

	result := config.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return User{}, errors.New("user with id " + id + " not found")
	}

	return user, nil
}
