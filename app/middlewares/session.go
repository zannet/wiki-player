package middlewares

import (
	"github.com/adred/wiki-player/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/tracelog"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

// Session attaches session to gin context
func Session(store *sessions.CookieStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := store.Get(c.Request, utils.ConfigEntry("SessionName"))
		if err != nil {
			tracelog.CompletedError(err, "Session", "Getting the session")
			c.Error(err, "Failed to create session")
			c.AbortWithStatus(500)
		}

		c.Set("session", session)
		defer context.Clear(c.Request)
	}
}
