package middlewares

import (
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/adred/wiki-player/models"
	"github.com/adred/wiki-player/utils"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/tracelog"
)

type (
	ClientParams struct {
		Nonce  string `json:"nonce" binding:"required"`
		ApiKey string `json:"apiKey" binding:"required"`
		Hash   string `json:"hash" binding:"required"`
	}

	ClientHashableParams struct {
		Nonce  string
		ApiKey string
	}
)

func ClientAuth(dbHandle *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p ClientParams

		c.Bind(&p)

		// Check validity of Nonce
		nm := &models.NonceModel{DbHandle: dbHandle}
		_, err := nm.Verify(p.Nonce)
		if err != nil {
			tracelog.CompletedError(err, "ClientAuth", "nm.Verify")
			c.JSON(401, gin.H{"message": "Invalid nonce.", "status": 401})
			c.Abort(401)
		}

		// Check validity of ApiKey
		cm := &models.ClientModel{DbHandle: dbHandle}
		_, err = cm.Verify(p.ApiKey)
		if err != nil {
			tracelog.CompletedError(err, "ClientAuth", "cm.Verify")
			c.JSON(401, gin.H{"message": "Invalid api key.", "status": 401})
			c.Abort(401)
		}

		// Get private key
		privateKey, err := cm.PrivateKey(p.ApiKey)
		if err != nil {
			tracelog.CompletedError(err, "ClientAuth", "cm.PrivateKey")
			c.JSON(401, gin.H{"message": "Couldn't retrieve the private key.", "status": 401})
			c.Abort(401)
		}

		// Hash request body
		hp := ClientHashableParams{Nonce: p.Nonce, ApiKey: p.ApiKey}
		payload, err := json.Marshal(&hp)
		if err != nil {
			tracelog.CompletedError(err, "ClientAuth", "json.Marshal")
			c.JSON(401, gin.H{"message": "Couldn't marshal request body.", "status": 401})
			c.Abort(401)
		}

		// Check validity of Hash
		// Client also must hash a JSON {Nonce:"abc123",ApiKey:"abc123"} with a private key
		var hashMismatch = errors.New("Hashes do not match.")
		hash := utils.ComputeHmac256(string(payload), privateKey)
		if p.Hash != hash {
			tracelog.CompletedError(hashMismatch, "ClientAuth", "Hashes comparison")
			c.JSON(401, gin.H{"message": "Hashes do not match.", "status": 401})
			c.Abort(401)
		}
	}
}
