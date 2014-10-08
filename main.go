package main

import (
	"github.com/adred/wiki-player/controllers"
	"github.com/adred/wiki-player/middlewares"
	"github.com/adred/wiki-player/models"
	"github.com/adred/wiki-player/utils"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/tracelog"
	"github.com/gorilla/sessions"
)

// Main go routine
func main() {
	// Start logger
	tracelog.StartFile(1, utils.ConfigEntry("LogDir"), 1)

	// Get new cookie store
	store := sessions.NewCookieStore([]byte(utils.ConfigEntry("SecretKey")))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
	}

	// Create DB connection
	dbHandle := utils.DbHandle()
	// Close DB
	defer dbHandle.Close()

	// Init Models
	sm := &models.SongModel{DbHandle: dbHandle}
	// nm := &models.NonceModel{DbHandle: dbHandle}
	um := &models.UserModel{DbHandle: dbHandle, UserData: &models.UserData{}}

	// Init Controllers
	sc := &controllers.SongController{SM: sm}
	// nc := &controllers.NonceController{NM: nm}
	uc := &controllers.UserController{UM: um, Store: store}

	// Init Gin
	mux := gin.Default()
	// Create private routes group
	private := mux.Group("/")

	// Public middlewares
	mux.Use(middlewares.Session(store))

	// Private middlewares
	private.Use(middlewares.Session(store), middlewares.UserAuth(store))

	// Public routes
	mux.POST("/users/login", uc.Login)
	mux.POST("/users/register", uc.Register)

	// Private routes
	private.GET("/", sc.Index)
	private.GET("/users/delete/:nonce", uc.ConfirmDelete)
	private.POST("/users/delete", uc.Delete)
	private.POST("/users/update", uc.Update)
	private.POST("/users/logout", uc.Logout)

	// Listen and serve on 0.0.0.0:8080
	mux.Run(":8080")

	tracelog.Stop()
}
