package controllers

import (
	"net/http"

	"bitbucket.org/adred/wiki-player/models"
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
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, songs)
}

func (sc *SongController) Get(c *gin.Context) {
}
