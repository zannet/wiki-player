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

	// Create DB connection
	db, err := utils.NewDB()
	if err != nil {
		os.Exit(1)
	}
	// Close DB
	defer db.Handle.Close()

	// Init Models
	sm := &models.SongModel{DB: db}
	nm := &models.NonceModel{DB: db}

	// Init Controllers
	sc := &controllers.SongController{SM: sm}
	nc := &controllers.NonceController{NM: nm}

	// Init Gin
	mux := gin.Default()

	// Setup routes
	mux.GET("/", sc.Index)
	mux.GET("/nonce", nc.Create)

	// Listen and serve on 0.0.0.0:8080
	mux.Run(":8080")

	tracelog.Stop()
}
