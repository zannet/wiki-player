package middlewares

import (
	"github.com/adred/wiki-player/models"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/tracelog"
)

type (
	N struct {
		Nonce string `json:"nonce" binding:"required"`
	}
)

func Nonce(nonce *models.NonceModel) gin.HandlerFunc {
	return func(c *gin.Context) {
		var n N
		// Bind params
		c.Bind(&n)

		_, err := nonce.Verify(n.Nonce)
		if err != nil {
			tracelog.CompletedError(err, "Nonce", "nonce.Verify")
			c.Abort(401)
		}

		c.Next()
	}
}
