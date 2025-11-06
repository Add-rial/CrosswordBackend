package handlers

import (
	"encoding/json"
	"fmt"
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
	filePath := "data/day1/crosswordJSON.json"
	file, err := os.ReadFile(filePath)
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

	var existing model.CrosswordAnswer
    if err := config.DB.
        Where("user_id = ? AND crossword_id = ?", userID, input.CrosswordID).
        First(&existing).Error; err == nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "You have already submitted this crossword.",
        })
        return
    }

    if err := config.DB.Create(&input).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
        return
    }

	c.JSON(http.StatusOK, gin.H{"message": "Answer stored"})
}

// @Summary      Get Crossword and Solution
// @Description  Fetches both the crossword question (grid, clues, etc.) and the corresponding solution for a given crossword ID.
// @Tags         Crossword
// @Accept       json
// @Produce      json
// @Param        crossword_id  body  int  true  "ID of the crossword to fetch"
// @Success      200  {object}  map[string]interface{}  "Crossword and solution successfully fetched"
// @Failure      400  {object}  map[string]string        "Bad request or solution not found"
// @Failure      500  {object}  map[string]string        "Internal server error while reading files"
// @Router       /getsolution [post]
func GetSolution(c *gin.Context){
	var b struct{
		CrosswordID uint `json:"crossword_id"`
	}

	if err:= c.ShouldBindJSON(&b); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format, provide crossword_id"})
	}

	filePath := fmt.Sprintf("data/day%d/solutionJSON.json", b.CrosswordID)
	file, err := os.ReadFile(filePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Solution not found"})
		return
	}

	var solution model.CrosswordSolution
	err = json.Unmarshal(file, &solution)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error extracting crossword solution"})
		return
	}

	filePath = fmt.Sprintf("data/day%d/crosswordJSON.json", b.CrosswordID)
	file, err = os.ReadFile(filePath)
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

	response := gin.H{
		"crossword": crossword,
		"solution":  solution,
	}

	c.JSON(http.StatusOK, response)
}