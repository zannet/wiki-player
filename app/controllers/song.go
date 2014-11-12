package controllers

import (
	"github.com/adred/wiki-player/app/models"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/tracelog"
)

type Song struct {
	SM *models.Song
}

func (sc *Song) Index(c *gin.Context) {
	songs, err := sc.SM.GetAll()
	if err != nil {
		tracelog.CompletedError(err, "Song", "Index")
		c.JSON(500, gin.H{"message": "Something went wrong.", "status": 500})

		return
	}

	c.JSON(200, gin.H{"songs": songs})
}

func (sc *Song) Get(c *gin.Context) {
}
