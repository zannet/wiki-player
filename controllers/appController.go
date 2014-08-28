package controllers

import (
	"bitbucket.org/adred/wiki-player/models"
	"github.com/gin-gonic/gin"
)

type AppController struct {
	AM *models.AppModel
}

func (ac *AppController) Index(c *gin.Context) {

}

func (ac *AppController) Get(c *gin.Context) {
}
