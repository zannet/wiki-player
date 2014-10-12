package middlewares

import (
	"errors"

	"github.com/adred/wiki-player/models"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/tracelog"
)

type (
	// N struct is used for nonce payload binding
	N struct {
		Nonce string `json:"nonce"`
	}
)

var errNonce = errors.New("Empty or no nonce sent.")

// Nonce verifies the nonce attached to a request
func Nonce(nm *models.NonceModel) gin.HandlerFunc {
	return func(c *gin.Context) {
		var n N
		// Bind params
		c.Bind(&n)

		var err error
		var id string
		if n.Nonce == "" {
			val, ok := c.Request.URL.Query()["nonce"]
			if ok && val[0] != "" {
				id, err = nm.Verify(val[0])
			} else {
				tracelog.CompletedError(errNonce, "Nonce", "Empty or no nonce sent")
				c.Abort(401)
			}
		} else {
			id, err = nm.Verify(n.Nonce)
		}

		if err != nil {
			tracelog.CompletedError(err, "Nonce", "nm.Verify")
			c.Abort(401)
		}

		// Delete nonce once verified
		err = nm.Delete(id)
		if err != nil {
			tracelog.CompletedError(err, "Nonce", "nm.Delete")
			c.Abort(401)
		}
	}
}
