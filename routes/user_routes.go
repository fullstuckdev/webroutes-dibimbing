package routes

import (
	"golangapi/controllers"
	"golangapi/middleware"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func SetupUserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	userController := controllers.NewUserController(db)
	profileController := controllers.NewProfileController(db)

	protected := router.Group("/")
	nonProtected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/users", userController.GetUsers)
		protected.POST("/users/:id/profile", profileController.CreateProfile)
		protected.GET("/users/:id/profile", profileController.GetProfile)
	}
	
	nonProtected.GET("/users/:id", userController.GetUserByID)
}