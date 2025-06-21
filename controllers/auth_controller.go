package controllers

import (
	"golangapi/models"
	"golangapi/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	AuthService *services.AuthService
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		AuthService: services.NewAuthService(db),
	}
}

func (ac *AuthController) Register(c *gin.Context) {
    var user models.User
	
	// harus JSON
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Use service to register user
    token, err := ac.AuthService.Register(&user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "User registered successfully",
        "token":   token,
    })
}

func (ac *AuthController) Login(c *gin.Context) {
    var loginReq models.LoginRequest

	// harus bentuknya JSON
    if err := c.ShouldBindJSON(&loginReq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Use service to login user
    // testing123
    token, err := ac.AuthService.Login(&loginReq)
    if err != nil {
        if err.Error() == "invalid email or password" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Login successful",
        "token":   token,
    })
}