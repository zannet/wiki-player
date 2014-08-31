package main

import (
	"bitbucket.org/adred/wiki-player/controllers"
	"bitbucket.org/adred/wiki-player/middlewares"
	"bitbucket.org/adred/wiki-player/models"
	"bitbucket.org/adred/wiki-player/utils"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/tracelog"
)

// Our main go routine
func main() {
	// Start logger
	tracelog.StartFile(1, utils.ConfigEntry("LogDir"), 1)

	// Create DB connection
	dbHandle := utils.DBHandle()
	// Close DB
	defer dbHandle.Close()

	// Init Models
	sm := &models.SongModel{DBHandle: dbHandle}
	cm := &models.ClientModel{DBHandle: dbHandle}
	nm := &models.NonceModel{DBHandle: dbHandle}

	// Init Controllers
	sc := &controllers.SongController{SM: sm}
	cc := &controllers.ClientController{CM: cm}
	nc := &controllers.NonceController{NM: nm}

	// Init Gin
	mux := gin.Default()

	// Middlewares
	mux.Use(middlewares.ClientAuth(dbHandle))

	// Setup routes
	mux.GET("/", sc.Index)
	mux.POST("/client", cc.Index)
	mux.GET("/nonce", nc.Create)

	// Listen and serve on 0.0.0.0:8080
	mux.Run(":8080")

	tracelog.Stop()
}
