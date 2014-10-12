package controllers

import (
	"github.com/adred/wiki-player/models"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/tracelog"
	"github.com/gorilla/sessions"
)

type (
	// NonceController is the type of this class
	NonceController struct {
		NM *models.NonceModel
	}
)

// Create creates a nonce
func (nc *NonceController) Create(c *gin.Context) {
	// Get session
	session := c.MustGet("session").(*sessions.Session)

	nonce, err := nc.NM.Create(session.Values["uid"].(string))
	if err != nil {
		tracelog.CompletedError(err, "NonceController", "Create")
		c.JSON(500, gin.H{"message": "Something went wrong.", "status": 500})

		return
	}

	c.JSON(200, gin.H{"nonce": nonce})
}
