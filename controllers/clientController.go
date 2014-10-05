package controllers

import (
	"github.com/adred/wiki-player/models"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/tracelog"
)

type ClientController struct {
	CM *models.ClientModel
}

func (cc *ClientController) Index(c *gin.Context) {
	keys, err := cc.CM.Register("1", "ownClient")
	if err != nil {
		tracelog.CompletedError(err, "ClientController", "Post")
		c.JSON(500, gin.H{"message": "Something went wrong.", "status": 500})

		return
	}

	c.JSON(200, gin.H{"keys": keys})
}

func (cc *ClientController) Post(c *gin.Context) {
	keys, err := cc.CM.Register("1", "ownClient")
	if err != nil {
		tracelog.CompletedError(err, "ClientController", "Post")
		c.JSON(500, gin.H{"message": "Something went wrong.", "status": 500})

		return
	}

	c.JSON(200, gin.H{"keys": keys})
}

func (cc *ClientController) Get(c *gin.Context) {
}
