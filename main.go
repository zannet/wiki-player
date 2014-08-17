package main

import (
	"net/http"
	"os"

	"bitbucket.org/adred/cloud-music-player/config"
	"bitbucket.org/adred/cloud-music-player/controllers"
	"bitbucket.org/adred/cloud-music-player/models"
	"bitbucket.org/adred/cloud-music-player/utils"
	"bitbucket.org/adred/cloud-music-player/utils/helpers"
	"github.com/goinggo/tracelog"
	"github.com/unrolled/render"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

// Our custom handler type
type appHandler func(web.C, http.ResponseWriter, *http.Request) *helpers.AppError

// We need to implement net/http's ServeHTTP too
func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

// Handler
func (fn appHandler) ServeHTTPC(c web.C, w http.ResponseWriter, r *http.Request) {
	if err := fn(c, w, r); err != nil {
		switch err.Code {
		case http.StatusOK:
			// Do nothing
		case http.StatusNotFound:
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		case http.StatusInternalServerError:
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		default:
			// Catch any other errors we haven't explicitly handled
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}

// Our main go routine
func main() {
	// Start logger
	tracelog.StartFile(1, config.Entry("LogDir"), 15)

	// Open DB
	db := &utils.DB{}
	err := db.Open()
	if err != nil {
		os.Exit(1)
	}
	// Close DB
	defer db.Close()
	// Init render
	rdr := render.New(render.Options{Directory: "views"})
	// Init SongModel
	sm := &models.SongModel{DB: db.Conn}
	// Init SongController
	sc := &controllers.SongController{SM: sm, Rdr: rdr}

	// Routes
	goji.Get("/", appHandler(sc.Index))

	// Serve app
	goji.Serve()

	tracelog.Stop()
}
