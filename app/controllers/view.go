package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type View struct {
	Store *sessions.CookieStore
}

func (vc *View) Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func (vc *View) About(c *gin.Context) {
}

func (vc *View) Tos(c *gin.Context) {
}

func (vc *View) PrivacyPolicy(c *gin.Context) {
}

func (vc *View) Credits(c *gin.Context) {
}
