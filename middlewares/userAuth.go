package middlewares

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/goinggo/tracelog"
	"github.com/gorilla/sessions"
)

var emptySession = errors.New("Empty session.")

func UserAuth(store *sessions.CookieStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := c.MustGet("session").(*sessions.Session)
		if session.Values["uid"] == nil {
			tracelog.CompletedError(emptySession, "UserAuth", "Checking of session uid value")
			c.Abort(401)
		}

		c.Next()
	}
}
