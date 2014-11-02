package controllers

import (
	"github.com/adred/wiki-player/app/models"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/tracelog"
)

type SongController struct {
	SM *models.SongModel
}

func (sc *SongController) Index(c *gin.Context) {
	songs, err := sc.SM.GetAll()
	if err != nil {
		tracelog.CompletedError(err, "SongController", "Index")
		c.JSON(500, gin.H{"message": "Something went wrong.", "status": 500})

		return
	}

	c.JSON(200, gin.H{"songs": songs})
}

func (sc *SongController) Get(c *gin.Context) {
}
