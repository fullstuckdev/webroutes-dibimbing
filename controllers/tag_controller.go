package controllers

import (
	"golangapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TagController struct {
	DB *gorm.DB
}

func NewTagontroller(db *gorm.DB) *TagController {
	return &TagController{DB: db}
}

func (tc *TagController) CreateTag(c *gin.Context) {
	
	var req models.CreateTagRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid data",
		})
		return
	}

	tag := models.Tag {
		Name: req.Name,
	}

	tc.DB.Create(&tag)

	c.JSON(http.StatusCreated, models.APIResponse{
		Success: true,
		Message: "Tag created",
		Data: tag,
	})
}