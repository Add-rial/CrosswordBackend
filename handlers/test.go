package handlers

import (
	"CrosswordBackend/config"
	"CrosswordBackend/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	if err := config.DB.Model(&model.User{}).Where("email = ?", b.Email).Update("score", gorm.Expr("score + ?", 5)).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Where("email = ?", b.Email).First(&userInDB).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch updated user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Score": userInDB.Score})
}

func SeeSubmittedCrossword(c *gin.Context){
	var userInDB model.CrosswordAnswer
	if err := config.DB.Model(model.CrosswordAnswer{}).
				Where("crossword_id = ?", 1).First(&userInDB); err != nil{
					c.JSON(http.StatusInternalServerError, gin.H{"error": "coulnd't fetch the crossword"})
					return
				}
	
	c.JSON(http.StatusOK, userInDB)
}