package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type StaticController struct {
	Store *sessions.CookieStore
}

func (sc *StaticController) Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func (sc *StaticController) About(c *gin.Context) {
}

func (sc *StaticController) Tos(c *gin.Context) {
}

func (sc *StaticController) PrivacyPolicy(c *gin.Context) {
}

func (sc *StaticController) Credits(c *gin.Context) {
}