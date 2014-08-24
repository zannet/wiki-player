package main

import (
	"os"

	"bitbucket.org/adred/wiki-player/config"
	"bitbucket.org/adred/wiki-player/controllers"
	"bitbucket.org/adred/wiki-player/models"
	"bitbucket.org/adred/wiki-player/utils"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/tracelog"
)

// Our main go routine
func main() {
	// Start logger
	tracelog.StartFile(1, config.Entry("LogDir"), 1)

	// Create new DB
	db, err := utils.NewDB()
	if err != nil {
		os.Exit(1)
	}
	// Close DB
	defer db.Handle.Close()

	// Init SongModel
	sm := &models.SongModel{DB: db}
	// Init SongController
	sc := &controllers.SongController{SM: sm}

	mux := gin.Default()
	mux.GET("/", sc.Index)

	// Listen and server on 0.0.0.0:8080
	mux.Run(":8080")

	tracelog.Stop()
}
