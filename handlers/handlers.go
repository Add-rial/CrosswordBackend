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

// RegisterUser godoc
// @Summary Register a new user
// @Description Creates a new user if not registered and returns a JWT token. If already registered, returns the user(no JWT token)
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body model.User true "User info"
// @Success 201 {object} map[string]interface{} "User successfully registered"
// @Failure 400 {object} map[string]string "Invalid request format"
// @Failure 500 {object} map[string]string "Database or token error"
// @Router /users/register [post]
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

// LoginUser godoc
// @Summary Log in existing user
// @Description Returns JWT token if user exists
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body model.User true "User login info"
// @Success 200 {object} map[string]interface{} "Successful login"
// @Failure 400 {object} map[string]string "Invalid request format"
// @Failure 401 {object} map[string]string "User not registered"
// @Router /users/login [post]
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

// GetLeaderboard godoc
// @Summary Get leaderboard
// @Description Returns a list of users sorted by score in descending order
// @Tags Leaderboard
// @Produce  json
// @Success 200 {array} model.User
// @Failure 500 {object} map[string]string "Database error"
// @Router /leaderboard [get]
func GetLeaderboard(c *gin.Context){
	var leaderboard []model.User
	
	result := config.DB.Order("score desc").Find(&leaderboard)
	
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error while fetching leaderboard"})
		return
	}

	c.JSON(http.StatusOK, leaderboard)
}

// GetCrossword godoc
// @Summary Get today's crossword
// @Description Returns crossword grid data and clues from JSON file
// @Tags Crossword
// @Produce  json
// @Success 200 {object} model.Crossword
// @Failure 500 {object} map[string]string "Crossword not found or parse error"
// @Router /crossword [get]
func GetCrossword(c *gin.Context){
	file, err := os.ReadFile("crosswordJSON.json")
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

// SubmitCrossword godoc
// @Summary Submit crossword answers
// @Description Stores or updates a user’s crossword answers. Requires JWT authentication.
// @Tags Crossword
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param crossword body model.CrosswordAnswer true "Crossword answer data"
// @Success 200 {object} map[string]string "Answer stored"
// @Failure 400 {object} map[string]string "Invalid request format"
// @Failure 401 {object} map[string]string "Unauthorized or missing token"
// @Failure 500 {object} map[string]string "Database error"
// @Router /submitcrossword [post]
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
		return
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