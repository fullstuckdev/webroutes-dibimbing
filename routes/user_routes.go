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
	postController := controllers.NewPostController(db)
	tagController := controllers.NewTagontroller(db)

	protected := router.Group("/")
	nonProtected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/users", userController.GetUsers)
		protected.POST("/users/:id/profile", profileController.CreateProfile)
		protected.GET("/users/:id/profile", profileController.GetProfile)
		protected.POST("/users/:id/posts", postController.CreatePost)
		protected.PUT("/users/:id/posts", postController.UpdatePost)
		protected.DELETE("/users/:id/posts", postController.DeletePost)
		protected.POST("/tags", tagController.CreateTag)
		protected.PATCH("/tags/:id", tagController.UpdateTag) // parsial
		protected.DELETE("/tags/:id", tagController.DeleteTag)
	}
	
	nonProtected.GET("/users/:id", userController.GetUserByID)
}