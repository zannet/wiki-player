package middlewares

import (
	"github.com/adred/wiki-player/utils"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/tracelog"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

func Session(store *sessions.CookieStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := store.Get(c.Request, utils.ConfigEntry("SessionName"))
		if err != nil {
			tracelog.CompletedError(emptySession, "Session", "Getting the session")
			c.Abort(500)
		}

		c.Set("session", session)
		defer context.Clear(c.Request)
	}
}
