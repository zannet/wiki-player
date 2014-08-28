package controllers

import (
	"bitbucket.org/adred/wiki-player/models"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/tracelog"
)

type NonceController struct {
	NM *models.NonceModel
}

func (nc *NonceController) Create(c *gin.Context) {
	nonce, err := nc.NM.Create()
	if err != nil {
		tracelog.CompletedError(err, "NonceController", "Create")
		c.JSON(500, gin.H{"message": "Something went wrong.", "status": 500})

		return
	}

	c.JSON(200, gin.H{"nonce": nonce})
}
