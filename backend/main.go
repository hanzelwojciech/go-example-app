package main

import (
	"learning-app/config"
	"learning-app/database"
	"learning-app/users"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	mode := config.Initialize()

	db, err := database.CreateDatabaseConnection()
	if err != nil {
		log.Fatalln("[database]", err.Error())
	}

	db.AutoMigrate(&users.User{})
	db.AutoMigrate(&users.UserInvitationToken{})

	router := gin.Default()

	users.RegisterRoutes(router)

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))

	if mode == "local" {
		router.Run("127.0.0.1:8080")
	} else {
		router.Run()
	}
}
