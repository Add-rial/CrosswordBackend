package main

// @title Crossword API
// @version 1.0
// @description Backend API for Crossword App

// @host https://crosswordbackend.onrender.com/
// @BasePath /

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"CrosswordBackend/docs"

	"CrosswordBackend/config"
	"CrosswordBackend/handlers"
	"CrosswordBackend/middleware"
)

func main(){
	config.InitEnv()
	config.InitDB()
	config.InitCORS()

	docs.SwaggerInfo.BasePath = "/"
	router := gin.Default()
	router.Use(cors.New(config.CORSconfig))

	router.POST("/users/register", handlers.RegisterUser)
	router.POST("/users/login", handlers.LoginUser)
	router.POST("/auth/google", handlers.AuthWithGoogle)
	router.GET("/leaderboard", handlers.GetLeaderboard)
	router.GET("/crossword", handlers.GetCrossword)
	submitcrossword := router.Group("/")
	submitcrossword.Use(middleware.AuthMiddleware())
	{
		submitcrossword.POST("/submitcrossword", handlers.SubmitCrossword)
	}
	admin := router.Group("/admin")
	admin.Use(middleware.AdminMiddleware())
	{
		admin.POST("/upload-crossword", handlers.UpdateCrossword)
		admin.POST("/update-scores", handlers.UpdateScore)
		admin.POST("/update-solution", handlers.UpdateSolution)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/", ginSwagger.WrapHandler(swaggerFiles.Handler))


	log.Fatal(router.Run(":8080"))
}