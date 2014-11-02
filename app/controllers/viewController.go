package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type ViewController struct {
	Store *sessions.CookieStore
}

func (vc *ViewController) Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func (vc *ViewController) About(c *gin.Context) {
}

func (vc *ViewController) Tos(c *gin.Context) {
}

func (vc *ViewController) PrivacyPolicy(c *gin.Context) {
}

func (vc *ViewController) Credits(c *gin.Context) {
}
