package main

import (
	"net/http"

	"bitbucket.org/adred/wiki-player/controllers"
	// "bitbucket.org/adred/wiki-player/middlewares"
	"bitbucket.org/adred/wiki-player/models"
	"bitbucket.org/adred/wiki-player/utils"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/tracelog"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

// Our main go routine
func main() {
	// Start logger
	tracelog.StartFile(1, utils.ConfigEntry("LogDir"), 1)

	store := sessions.NewCookieStore([]byte(utils.ConfigEntry("SecretKey")))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 8, // 8 hours
		HttpOnly: true,
	}

	// Create DB connection
	dbHandle := utils.DbHandle()
	// Close DB
	defer dbHandle.Close()

	// Init Models
	sm := &models.SongModel{DbHandle: dbHandle}
	// cm := &models.ClientModel{DbHandle: dbHandle}
	// nm := &models.NonceModel{DbHandle: dbHandle}
	um := &models.UserModel{DbHandle: dbHandle, UserData: &models.UserData{}}

	// Init Controllers
	sc := &controllers.SongController{SM: sm}
	// cc := &controllers.ClientController{CM: cm}
	// nc := &controllers.NonceController{NM: nm}
	uc := &controllers.UserController{UM: um, Store: store}

	// Init Gin
	mux := gin.Default()

	// Middlewares
	// mux.Use(middlewares.ClientAuth(dbHandle))
	// mux.Use(middlewares.UserAuth(dbHandle))

	// Routes
	mux.GET("/", sc.Index)
	// mux.POST("/client", cc.Index) // mux.GET("/nonce", nc.Create)
	mux.POST("/users/login", uc.Login)
	mux.POST("/users/logout", uc.Logout)
	mux.POST("/users/register", uc.Register)
	mux.GET("/users/session", uc.CheckSessionStatus)

	// Listen and serve on 0.0.0.0:8080
	http.ListenAndServe(":8080", context.ClearHandler(mux))

	tracelog.Stop()
}
