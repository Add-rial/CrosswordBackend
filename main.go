package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"CrosswordBackend/handlers"
	"CrosswordBackend/middleware"
	"CrosswordBackend/services"
	"CrosswordBackend/config"
)

func main(){
	config.InitEnv()
	config.InitDB()
	services.JsonGenerator()

	router := gin.Default()

	router.POST("/users/register", handlers.RegisterUser)
	router.POST("/users/login", handlers.LoginUser)
	router.GET("/leaderboard", handlers.GetLeaderboard)
	router.GET("/crossword", handlers.GetCrossword)
	auth := router.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/submitcrossword", handlers.SubmitCrossword)
	}

	log.Fatal(router.Run(":8080"))
}