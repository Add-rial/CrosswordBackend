package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"CrosswordBackend/config"
	"CrosswordBackend/model"
)

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
