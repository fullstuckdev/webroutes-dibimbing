package services

import (
	"errors"
	"golangapi/models"
	"golangapi/utils"

	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{DB: db}
}

func (as *AuthService) Register(user *models.User) (string, error) {
	// Hash password
	if err := user.HashPassword(user.Password); err != nil {
		return "", errors.New("error hashing password")
	}

	// Create user in database
	if err := as.DB.Create(user).Error; err != nil {
		return "", errors.New("error creating user")
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", errors.New("error generating token")
	}

	return token, nil
}

func (as *AuthService) Login(loginReq *models.LoginRequest) (string, error) {
	var user models.User

	// Check if user exists
	if err := as.DB.Where("email = ?", loginReq.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("invalid email or password")
		}
		return "", err
	}

	// Check password
	if err := user.CheckPassword(loginReq.Password); err != nil {
		return "", errors.New("invalid email or password")
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", errors.New("error generating token")
	}

	return token, nil
} 