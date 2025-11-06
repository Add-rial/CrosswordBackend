package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"CrosswordBackend/config"
	"CrosswordBackend/model"
)

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
	result := config.DB.Where("user_id = ? AND crossword_id = ?", userID, input.CrosswordID).Assign(input).FirstOrCreate(&existingAnswer)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
        return
    }

	c.JSON(http.StatusOK, gin.H{"message": "Answer stored"})
}