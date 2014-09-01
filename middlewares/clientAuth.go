package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"

	"bitbucket.org/adred/wiki-player/models"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/tracelog"
)

// Binding from JSON
type ClientJSON struct {
	Nonce  string `json:"nonce" binding:"required"`
	ApiKey string `json:"apiKey" binding:"required"`
	Hash   string `json:"hash" binding:"required"`
}

type Hashable struct {
	Nonce  string
	ApiKey string
}

func computeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func ClientAuth(dbHandle *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var j ClientJSON

		c.Bind(&j)

		// Check validity of Nonce
		nm := &models.NonceModel{DBHandle: dbHandle}
		_, err := nm.Verify(j.Nonce)
		if err != nil {
			tracelog.CompletedError(err, "ClientAuth", "nm.Verify")
			c.JSON(401, gin.H{"message": "Invalid nonce.", "status": 401})
			c.Abort(401)
		}

		// Check validity of ApiKey
		cm := &models.ClientModel{DBHandle: dbHandle}
		_, err = cm.Verify(j.ApiKey)
		if err != nil {
			tracelog.CompletedError(err, "ClientAuth", "cm.Verify")
			c.JSON(401, gin.H{"message": "Invalid api key.", "status": 401})
			c.Abort(401)
		}

		// Get private key
		privateKey, err := cm.PrivateKey(j.ApiKey)
		if err != nil {
			tracelog.CompletedError(err, "ClientAuth", "cm.PrivateKey")
			c.JSON(401, gin.H{"message": "Couldn't retrieve the private key.", "status": 401})
			c.Abort(401)
		}

		// Hash request body
		h := Hashable{Nonce: j.Nonce, ApiKey: j.ApiKey}
		payload, err := json.Marshal(h)
		if err != nil {
			tracelog.CompletedError(err, "ClientAuth", "json.Marshal")
			c.JSON(401, gin.H{"message": "Couldn't marshal request body.", "status": 401})
			c.Abort(401)
		}

		// Check validity of Hash
		// Client also must hash a JSON {Nonce:"abc123",ApiKey:"abc123"} with a private key
		var hashMismatch = errors.New("Hashes do not match.")
		hash := computeHmac256(string(payload), privateKey)

		if j.Hash != hash {
			tracelog.CompletedError(hashMismatch, "ClientAuth", "Hashes comparison")
			c.JSON(401, gin.H{"message": "Hashes do not match.", "status": 401})
			c.Abort(401)
		}
	}
}
