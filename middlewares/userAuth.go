package middlewares

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/goinggo/tracelog"
	"github.com/gorilla/sessions"
)

var errSession = errors.New("Empty session.")

// UserAuth checks if a session exists
func UserAuth(store *sessions.CookieStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := c.MustGet("session").(*sessions.Session)
		if session.Values["uid"] == nil {
			tracelog.CompletedError(errSession, "UserAuth", "Checking of session uid value")
			c.Abort(401)
		}
	}
}
