package services

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context, updatingWhat string, jsonData []byte, fileName string){
	err := os.WriteFile(fileName, jsonData, 0644)
	if err != nil{
		errmsg := "Error updataing the " + updatingWhat
		c.JSON(http.StatusInternalServerError, gin.H{"message": errmsg})
	}
	
	successmsg := updatingWhat + " uploaded successfully"
	c.JSON(http.StatusOK, gin.H{"message": successmsg})
}