package controllers

import (
	"net/http"

	"bitbucket.org/adred/cloud-music-player/models"
	"bitbucket.org/adred/cloud-music-player/utils/helpers"
	"github.com/goinggo/tracelog"
	"github.com/unrolled/render"
	"github.com/zenazn/goji/web"
)

type SongController struct {
	SM  *models.SongModel
	Rdr *render.Render
}

func (sc *SongController) Index(c web.C, w http.ResponseWriter, r *http.Request) *helpers.AppError {
	songs, err := sc.SM.GetAll()
	if err != nil {
		tracelog.CompletedError(err, "SongController", "models.GetAll")
		return &helpers.AppError{Error: err, Code: 500, Message: "SongController: models.GetAll"}
	}

	sc.Rdr.JSON(w, http.StatusOK, songs)
	return &helpers.AppError{Error: nil, Code: 200}
}

func (sc *SongController) Get(c web.C, w http.ResponseWriter, r *http.Request) *helpers.AppError {
	return &helpers.AppError{Error: nil, Code: 200}
}
