package main

// @title Crossword API
// @version 1.0
// @description Backend API for Crossword App 
// @termsOfService http://swagger.io/terms/


// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"

	"CrosswordBackend/docs"

	"CrosswordBackend/handlers"
	"CrosswordBackend/middleware"
	"CrosswordBackend/services"
	"CrosswordBackend/config"
)

func main(){
	config.InitEnv()
	config.InitDB()
	services.JsonGenerator()

	docs.SwaggerInfo.BasePath = "/"
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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	log.Fatal(router.Run(":8080"))
}