package controllers

import (
	"fmt"
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

func (pc *PostController) UpdatePost(c *gin.Context) {
	postID := c.Param("id")

	var post models.Post

	if err := pc.DB.First(&post, postID).Error; err != nil {
		c.JSON(http.StatusNotFound, models.APIResponse{
			Success: false,
			Message: "Post not found",
		})	
		return
	}

	var req models.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusNotFound, models.APIResponse{
			Success: false,
			Message: "Update kurang lengkap",
		})	
		return
	}
    
	// ini data dari si update post  {golang api updated update golang API}
	fmt.Println("ini data dari si update post ", req)

	post.Title = req.Title
	post.Content = req.Content

	//  &{{2 2025-06-21 15:00:15.927 +0700 WIB 2025-06-21 15:00:15.927 +0700 WIB {0001-01-01 00:00:00 +0000 UTC false}} 1 golang api updated update golang API 2025-07-21 15:00:15.919 +0700 WIB []}
	fmt.Println("ini postingan (&) ", &post)

	// {{2 2025-06-21 15:00:15.927 +0700 WIB 2025-06-21 15:00:15.927 +0700 WIB {0001-01-01 00:00:00 +0000 UTC false}} 1 golang api updated update golang API 2025-07-21 15:00:15.919 +0700 WIB []}
	fmt.Println("ini postingan ", post)

	pc.DB.Save(&post)

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Post updated",
		Data: post,
	})
}

func (pc *PostController) DeletePost(c *gin.Context) {
	postID := c.Param("id")

	var post models.Post

	// pencarian postingan
	if err := pc.DB.First(&post, postID).Error; err != nil {
		c.JSON(http.StatusNotFound, models.APIResponse{
			Success: false,
			Message: "Post not found",
		})	
		return
	}

	// sql builder untuk delete
	pc.DB.Delete(&post)

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Post deleted",
	})
}