package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"

	"bitbucket.org/adred/wiki-player/models"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/tracelog"
)

// Binding from JSON
type ClientJSON struct {
	Nonce  string `json:"nonce" binding:"required"`
	ApiKey string `json:"apiKey" binding:"required"`
	// Hash   string `json:"hash" binding:"required"`
}

func computeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func ClientAuth(dbHandle *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var json ClientJSON

		c.Bind(&json)

		// Check validity of Nonce
		nm := &models.NonceModel{DBHandle: dbHandle}
		ok, err := nm.Verify(json.Nonce)
		if err != nil {
			tracelog.CompletedError(err, "ClientAuth", "nm.Verify")
			c.JSON(401, gin.H{"message": "Invalid nonce.", "status": 401})
			c.Abort(401)
		}

		// Check validity of ApiKey
		cm := &models.ClientModel{DBHandle: dbHandle}
		ok, err = cm.Verify(json.ApiKey)
		if err != nil {
			tracelog.CompletedError(err, "ClientAuth", "cm.Verify")
			c.JSON(401, gin.H{"message": "Invalid api key.", "status": 401})
			c.Abort(401)
		}

		// Check validity of Hash
		//hash := computeHmac256(c.Request.URL)
		pkey, _ := cm.PrivateKey(json.ApiKey)
		fmt.Println(pkey)
		fmt.Println(ok)
	}
}
