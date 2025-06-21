package controllers

import (
	"fmt"
	"golangapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TagController struct {
	DB *gorm.DB // ngarahnya ke GORM, database MYSQL
	// DATABASE 2 => DB2
	// DATABASE 3 => DB3
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

func (tc *TagController) UpdateTag(c *gin.Context) {
	tagID := c.Param("id")

	var tag models.Tag

	if err := tc.DB.First(&tag, tagID).Error; err != nil {
		c.JSON(http.StatusNotFound, models.APIResponse{
			Success: false,
			Message: "Tag not found",
		})
	}

	var req models.UpdateTagRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusNotFound, models.APIResponse{
			Success: false,
			Message: "Data tidak lengkap",
		})	
		return
	}

	tag.Name = req.Name

	tc.DB.Save(&tag) // save artinya insert / create

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Tag updated",
		Data: tag,
	})
}

func (tc *TagController) DeleteTag(c *gin.Context) {
	tagID := c.Param("id")

	var tag models.Tag

	if err := tc.DB.First(&tag, tagID).Error; err != nil {
		c.JSON(http.StatusNotFound, models.APIResponse{
			Success: false,
			Message: "Tag not found",
		})
	}
	deleteTagQuery := "DELETE from tags where id = ?"
	result := tc.DB.Exec(deleteTagQuery, tagID)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to delete tag",
		})
		return
	}
	fmt.Println("hasil 0 ", result.Error)
	fmt.Println("hasil 1, ", result)
	fmt.Println("hasil 2, ", &result)


	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Tag deleted",
	})
}