package controllers

import (
	"bitbucket.org/adred/wiki-player/models"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UM *models.UserModel
}

func (sc *UserController) Index(c *gin.Context) {
}

func (sc *UserController) Get(c *gin.Context) {
}
