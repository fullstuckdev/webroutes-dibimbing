package controllers

import (
	"golangapi/models"
	"golangapi/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SysController struct {
	SysService *services.SysServices
}

func NewSysController(db *gorm.DB) *SysController {
	return &SysController{
		SysService: services.NewSysService(db),
	}
}

func (sc *SysController) CreateDirectory(c *gin.Context) {
	var req models.CreateDirectoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
		Success: false,
		Message: "Invalid data",
		})
	return
	}

	// dilakukan pemanggilan sebuah service
	err := sc.SysService.CreateDirectory(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Folder berhasil dibuat"})
}

func (sc *SysController) CreateFile(c *gin.Context) {
	var req models.CreateFileRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filePath, err := sc.SysService.CreateFile(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "File berhasil dibuat dan ditulis",
		"path": filePath,
	})
}

func (sc *SysController) UploadFile(c *gin.Context) {

	form, err := c.MultipartForm()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal membaca form data"})
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "File tidak boleh kosong",
		})
		return
	}

	var uploadFiles []any

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Gagal membuka file",
		})
		return
		}

		response, err := sc.SysService.UploadFile(file.Filename, src)
		
		if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Gagal upload file",
		})
		return
		}

		uploadFiles = append(uploadFiles, response)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": uploadFiles,
	})

}