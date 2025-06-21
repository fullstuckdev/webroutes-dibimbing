package main

import (
	"golangapi/config"
	"golangapi/models"
	"golangapi/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitializeApp() *gin.Engine {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading ENV")
	}

	r := gin.Default()

	db := config.ConnectDatabase()

	// auto migrate
	db.AutoMigrate(&models.User{}, &models.Profile{})

	routes.SetupRoutes(r, db)

	return r
}

func main() {
	app := InitializeApp()
	app.Run(":8080")
}