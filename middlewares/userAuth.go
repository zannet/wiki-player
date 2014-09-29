package middlewares

import (
	"database/sql"
	// "encoding/json"
	// "errors"

	// "bitbucket.org/adred/wiki-player/models"
	// "bitbucket.org/adred/wiki-player/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	// "github.com/goinggo/tracelog"
)

type (
	UserParams struct {
		token string `json:"token" binding:"required"`
	}
)

func UserAuth(store *sessions.CookieStore) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
