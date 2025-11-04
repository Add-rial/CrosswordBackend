package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"CrosswordBackend/config"
	"CrosswordBackend/model"
	"CrosswordBackend/utils"
)

func RegisterUser(c *gin.Context){
	var input model.User
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	var userInDB model.User
	
	result := config.DB.Where("email = ?", input.Email).FirstOrCreate(&userInDB, model.User{
		Username: input.Username, 
		Score: 0, 
		Email: input.Email,
	})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error during registration"})
		return
	}

	token, err := utils.GenerateJWT(userInDB.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"email": userInDB.Email,
			"username": userInDB.Username,
			"score": userInDB.Score,
			"message": "User already registered",
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"email": userInDB.Email,
			"username": userInDB.Username,
			"score": userInDB.Score,
			"message": "User successfully registered.",
			"token": token,
		})
	}
}

func LoginUser(c *gin.Context){
	var input model.User

	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	var userInDB model.User
	result := config.DB.Where("email = ?", input.Email).First(&userInDB)

	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not registered"})
		return
	}

	token, err := utils.GenerateJWT(userInDB.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}	
	c.JSON(http.StatusOK, gin.H{
        "token": token,
        "user":  userInDB,
    })
}

func GetLeaderboard(c *gin.Context){
	var leaderboard []model.User
	
	result := config.DB.Order("score desc").Find(&leaderboard)
	
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error while fetching leaderboard"})
		return
	}

	c.JSON(http.StatusOK, leaderboard)
}

func GetCrossword(c *gin.Context){
	file, err := os.ReadFile("crosswordOfTheDay")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Crossword not found"})
		return
	}
	var crossword model.Crossword
	err = json.Unmarshal(file, &crossword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error extracting crossword"})
		return
	}
	c.JSON(http.StatusOK, crossword)
}

func SubmitCrossword(c *gin.Context){
	userIDVal, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
        return
    }
    userID := userIDVal.(uint)

	var input model.CrosswordAnswer
	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
	}
	input.UserID = userID

	var existingAnswer model.CrosswordAnswer
	result := config.DB.Where("user_id = ?", userID).Assign(input).FirstOrCreate(&existingAnswer)

    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
        return
    }

	c.JSON(http.StatusOK, gin.H{"message": "Answer stored"})
}