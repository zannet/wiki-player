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
	// Temporarily set mode to "test" so we can use custom templates dir
	gin.SetMode("test")

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
	nm := &models.NonceModel{DbHandle: dbHandle}
	um := &models.UserModel{DbHandle: dbHandle, UserData: &models.UserData{}}

	// Init Controllers
	sc := &controllers.SongController{SM: sm}
	nc := &controllers.NonceController{NM: nm}
	uc := &controllers.UserController{UM: um, Store: store}
	vc := &controllers.ViewController{Store: store}

	// Init Gin
	mux := gin.Default()

	// Load templates
	utils.LoadHTMLFiles("views", mux)
	// Serve static files
	mux.Static("/static", utils.ConfigEntry("StaticDir"))

	// Routes for static pages
	static := mux.Group("/")
	{
		static.GET("/", vc.Index)
		static.GET("/about", vc.About)
		static.GET("/tos", vc.Tos)
		static.GET("/privacy-policy", vc.PrivacyPolicy)
		static.GET("/credits", vc.Credits)
	}

	// Routes that don't need authorization
	basic := mux.Group("/api/v1")
	basic.Use(middlewares.Session(store))
	{
		basic.GET("/nonce", nc.Create)
		basic.POST("/users/login", uc.Login)
		basic.POST("/users/register", uc.Register)
	}

	// Routes that need authorization
	auth := mux.Group("/api/v1")
	auth.Use(middlewares.Session(store), middlewares.UserAuth(store), middlewares.Nonce(nm))
	{
		auth.GET("/songs", sc.Index)
		auth.GET("/users/delete/:nonce", uc.ConfirmDelete)
		auth.POST("/users/delete", uc.Delete)
		auth.POST("/users/update", uc.Update)
		auth.POST("/users/logout", uc.Logout)
	}

	// Listen and serve on 0.0.0.0:8080
	mux.Run(":8080")

	tracelog.Stop()
}
