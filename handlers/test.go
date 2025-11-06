package handlers

import (
	"CrosswordBackend/config"
	"CrosswordBackend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnScore(c *gin.Context){
	var b struct{
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&b); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"messsage": "Invalid format"})
		return
	}

	var userInDB model.User
	if err := config.DB.Where("email = ?", b.Email).First(&userInDB).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error": "User not registered"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Score": userInDB.Score})
}