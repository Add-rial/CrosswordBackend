package services

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context, updatingWhat string){
	jsonData, err := io.ReadAll(c.Request.Body)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

	fileName := updatingWhat + "JSON.json"
	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil{
		errmsg := "Error updataing the " + updatingWhat
		c.JSON(http.StatusInternalServerError, gin.H{"message": errmsg})
	}
	
	successmsg := updatingWhat + " uploaded successfully"
	c.JSON(http.StatusOK, gin.H{"message": successmsg})
}