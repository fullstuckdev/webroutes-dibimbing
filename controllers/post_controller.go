package controllers

import (
	"golangapi/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostController struct {
	DB *gorm.DB
}

func NewPostController(db *gorm.DB) *PostController {
	return &PostController{DB: db}
}

func (pc *PostController) CreatePost(c *gin.Context) {
		userID, _ := strconv.Atoi(c.Param("id")) // ID dari si table users (string)

		var req models.CreatePostRequest

		if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid data",
		})
		return
	}

	post := models.Post{
		UserID: uint(userID),
		Title: req.Title,
		Content: req.Content,
	}

	// setter method
	post.SetExpiration()

	pc.DB.Create(&post)

	c.JSON(http.StatusCreated, models.APIResponse {
		Success: true,
		Message: "Posts created!",
		Data: post,
	})

}