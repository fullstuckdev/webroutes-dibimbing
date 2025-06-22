package routes

import (
	"golangapi/controllers"
	"golangapi/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func SetupSysRoutes(router *gin.RouterGroup, db *gorm.DB) {
	sysController := controllers.NewSysController(db)

	protected := router.Group("/sys")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/directory", sysController.CreateDirectory)
		protected.POST("/file", sysController.CreateFile)
	}
}