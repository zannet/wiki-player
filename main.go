package main

import (
	// "fmt"
	// "html/template"

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
	s := &controllers.StaticController{Store: store}

	// Init Gin
	mux := gin.Default()
	// Load templates
	mux.LoadHTMLFiles("static/*")

	// Routes for static pages
	static := mux.Group("/")
	{
		// Routes for static pages
		static.GET("/", s.Index)
		static.GET("/about", s.About)
		static.GET("/tos", s.Tos)
		static.GET("/privacy-policy", s.PrivacyPolicy)
		static.GET("/credits", s.Credits)
	}

	// Routes that don't authorization
	basic := mux.Group("/api/v1")
	basic.Use(middlewares.Session(store))
	{
		basic.POST("/users/login", uc.Login)
		basic.POST("/users/register", uc.Register)
	}

	// Routes that need authorization
	auth := mux.Group("/api/v1")
	auth.Use(middlewares.Session(store), middlewares.UserAuth(store))
	{
		auth.GET("/songs", sc.Index)
		auth.GET("/users/delete/:nonce", uc.ConfirmDelete)
		auth.POST("/users/delete", uc.Delete)
		auth.POST("/users/update", uc.Update)
		auth.POST("/users/logout", uc.Logout)
	}

	// nonce := mux.Group("/api/v1")
	// nonce.Use(middlewares.Nonce(nm))

	// Listen and serve on 0.0.0.0:8080
	mux.Run(":8080")

	tracelog.Stop()
}
