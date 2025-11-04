package main

// @title Crossword API
// @version 1.0
// @description Backend API for Crossword App 

// @host https://crosswordbackend.onrender.com/
// @BasePath /

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"

	"CrosswordBackend/docs"

	"CrosswordBackend/handlers"
	"CrosswordBackend/middleware"
	"CrosswordBackend/config"
)

func main(){
	config.InitEnv()
	config.InitDB()

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
	admin := router.Group("/admin")
	admin.Use(middleware.AdminMiddleware())
	{
		admin.POST("/upload-crossword", handlers.UpdateCrossword)
		admin.POST("/update-scores", handlers.UpdateScore)
		admin.POST("/update-solution", handlers.UpdateSolution)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	log.Fatal(router.Run(":8080"))
}