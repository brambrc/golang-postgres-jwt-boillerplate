package Model

import (
	"liamelior-api/Database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `form:"username" json:"username" binding:"required" gorm:"unique"`
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required" gorm:"unique"`
	Name     string `form:"name" json:"name" binding:"required"`
	Role     string `form:"role" json:"role" binding:"required"`
}

func (u *User) BeforeSave(*gorm.DB) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Save() (*User, error) {
	var err error

	// check duplicate username
	err = Database.Database.Create(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByUsername(username string) (User, error) {
	var user User
	err := Database.Database.Where("username = ?", username).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
