package main

import (
	"github.com/adred/wiki-player/app/controllers"
	"github.com/adred/wiki-player/app/factories"
	"github.com/adred/wiki-player/app/middlewares"
	"github.com/adred/wiki-player/app/models"
	"github.com/adred/wiki-player/lib"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/tracelog"
	"github.com/gorilla/sessions"
)

// Main go routine
func main() {
	// Start logger
	tracelog.StartFile(1, lib.ConfigEntry("LogDir"), 1)

	// Get new cookie store
	store := sessions.NewCookieStore([]byte(lib.ConfigEntry("SecretKey")))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
	}

	// Create DB connection
	dbHandle := lib.DbHandle()
	// Close DB
	defer dbHandle.Close()

	// Init Models
	sm := &models.Song{DbHandle: dbHandle}
	nm := &models.Nonce{DbHandle: dbHandle}
	um, err := factories.NewUserModel(dbHandle, lib.EnvConfigEntry("Mode"))
	if err != nil {
		tracelog.CompletedError(err, "main", "Model creation failed")
		panic(err.Error())
	}

	// Init Controllers
	sc := &controllers.Song{SM: sm}
	nc := &controllers.Nonce{NM: nm}
	vc := &controllers.View{Store: store}
	uc, err := factories.NewUserController(um, store, lib.EnvConfigEntry("Mode"))
	if err != nil {
		tracelog.CompletedError(err, "main", "Controller creation failed")
		panic(err.Error())
	}

	// Init Gin
	mux := gin.Default()

	// Load templates
	mux.LoadHTMLGlob("app/views/*")

	// Serve static files
	mux.Static("/public", lib.ConfigEntry("StaticDir"))

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
	mux.Run(":9000")

	tracelog.Stop()
}
