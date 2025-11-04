package handlers

import (
	"context"
	"net/http"

	"cloud.google.com/go/auth/credentials/idtoken"
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

// AuthWithGoogle godoc
// @Summary Authenticate with Google
// @Description Verifies Google token, auto-registers user if new, returns JWT
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param token body map[string]string true "Google ID Token"
// @Success 200 {object} map[string]interface{} "Authenticated successfully"
// @Failure 400 {object} map[string]string "Invalid request format"
// @Failure 401 {object} map[string]string "Invalid Google token"
// @Router /auth/google [post]
func AuthWithGoogle(c *gin.Context) {
    var body struct {
        Token string `json:"token" binding:"required"`
    }
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    }

	payload, err := idtoken.Validate(context.Background(), body.Token, config.GoogleClientID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Google token"})
		return
	}

	email, _ := payload.Claims["email"].(string)
	name, _ := payload.Claims["name"].(string)

    var userInDB model.User
	result := config.DB.Where("email = ?", email).FirstOrCreate(&userInDB, model.User{
		Username: name, 
		Score: 0, 
		Email: email,
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
	c.JSON(http.StatusOK, gin.H{
        "token": token,
        "user":  gin.H{
			"email": userInDB.Email,
            "username": userInDB.Username,
            "score": userInDB.Score,
		},
    })
}
