package controllers

import (
	"github.com/adred/wiki-player/app/models"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/tracelog"
	"github.com/gorilla/sessions"
)

type (
	// Nonce is the type of this class
	Nonce struct {
		NM *models.Nonce
	}
)

// Create creates a nonce
func (nc *Nonce) Create(c *gin.Context) {
	// Get session
	session := c.MustGet("session").(*sessions.Session)

	nonce, err := nc.NM.Create(session.Values["uid"].(string))
	if err != nil {
		tracelog.CompletedError(err, "Nonce", "Create")
		c.JSON(500, gin.H{"message": "Something went wrong.", "status": 500})

		return
	}

	c.JSON(200, gin.H{"nonce": nonce})
}
