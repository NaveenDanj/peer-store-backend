package handlers

import (
	"net/http"
	"peer-store/dto"
	"peer-store/models"
	"peer-store/service/storage"

	"github.com/gin-gonic/gin"
)

func GenerateLink(c *gin.Context) {
	var requestJson dto.FileShareRequestDTO

	if err := c.ShouldBindJSON(&requestJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := c.Get("currentUser")
	userObj, _ := user.(models.User)

	link, err := storage.GenerateShareLink(&requestJson, &userObj)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Link": link,
	})
	return
}

func RevokeLink(c *gin.Context) {

	var requestJson dto.LinkRevokeRequestDTO

	if err := c.ShouldBindJSON(&requestJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
