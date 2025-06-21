package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`

	// One to one with profile
	Profile *Profile `json:"profile,omitempty"`
}

// Profile model (one to one with users)
type Profile struct {
	gorm.Model
	UserID uint `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Bio string `json:"bio"`
}

type LoginRequest struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"Password" binding:"required"`
}

// Base Request
type CreateProfileRequest struct {
	FirstName string `json:"first_name" binding:"required"` // required artinya wajib dimasukan
	LastName string `json:"last_name" binding:"required"`  // required artinya wajib dimasukan
	Bio string `json:"bio"`
}

// Base Response
type APIResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
}


func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}