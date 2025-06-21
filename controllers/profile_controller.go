package controllers

import (
	"golangapi/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProfileController struct {
	DB *gorm.DB
}

func NewProfileController(db *gorm.DB) *ProfileController {
	return &ProfileController{DB: db}
}

func (pc *ProfileController) CreateProfile(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id")) // ID dari si table users (string)

	var req models.CreateProfileRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid data",
		})
		return
	}

	profile := models.Profile{
		UserID: uint(userID), // convert to integer
		FirstName: req.FirstName,
		LastName: req.LastName,
		Bio: req.Bio,
	}

	// buat create ke dalam table
	pc.DB.Create(&profile)

	c.JSON(http.StatusCreated, models.APIResponse{
		Success: true,
		Message: "Profile Created",
		Data: profile,
	})
}

func(pc *ProfileController) GetProfile(c *gin.Context) {
	userID := c.Param("id")

	var profile models.Profile

	// builder engine dari ORM
	pc.DB.Where("user_id = ?", userID).First(&profile)

	// Response
	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Profile found",
		Data: profile,
	}) 
}