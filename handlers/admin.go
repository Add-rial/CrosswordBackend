package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"CrosswordBackend/config"
	"CrosswordBackend/model"
	"CrosswordBackend/services"
)


// UpdateCrossword godoc
// @Summary Upload or update the crossword puzzle
// @Description Allows admin to upload or replace the crossword JSON file for the current day.
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param X-Admin-Key header string true "Admin authentication key"
// @Param crossword body object true "Crossword JSON data"
// @Success 200 {object} map[string]string "Crossword uploaded successfully"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 401 {object} map[string]string "Unauthorized (invalid admin key)"
// @Failure 500 {object} map[string]string "Error updating crossword"
// @Router /admin/updateCrossword [post]
func UpdateCrossword(c *gin.Context){
	services.Update(c, "crossword")
}

// UpdateSolution godoc
// @Summary Upload or update the crossword solution
// @Description Allows admin to upload the official solution JSON file after the puzzle day ends.
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param X-Admin-Key header string true "Admin authentication key"
// @Param solution body object true "Solution JSON data"
// @Success 200 {object} map[string]string "Solution uploaded successfully"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 401 {object} map[string]string "Unauthorized (invalid admin key)"
// @Failure 500 {object} map[string]string "Error updating solution"
// @Router /admin/updateSolution [post]
func UpdateSolution(c *gin.Context){
	services.Update(c, "solution")
}

// UpdateScore godoc
// @Summary Compare user answers and update scores
// @Description Runs the daily crossword scoring process. Compares every user's stored crossword answers with the official solution and updates their total scores accordingly.
// @Tags Admin
// @Produce  json
// @Param X-Admin-Key header string true "Admin authentication key"
// @Success 200 {object} map[string]string "Scores updated successfully"
// @Failure 401 {object} map[string]string "Unauthorized (invalid admin key)"
// @Failure 500 {object} map[string]string "Error loading solutions or updating scores"
// @Router /admin/updateScore [post]
func UpdateScore(c *gin.Context){
	log.Println("Running daily crossword scoring task...")

	solution, crosswordid, err := services.LoadOfficialSolution()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error loadind solutions"})
	}

	var answers []model.CrosswordAnswer
	config.DB.Where("crossword_id = ?", crosswordid).Find(&answers)

	for _, ans := range answers{
		score := services.CompareAnswer(ans.Answers, solution)
		config.DB.Model(&model.User{}).Where("id = ?", ans.UserID).
			Update("score", gorm.Expr("score + ?", score))
	}

	c.JSON(http.StatusOK, gin.H{"message": "Scores updated successfully"})
}