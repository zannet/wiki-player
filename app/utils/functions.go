package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var alpha = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"

// RandomString generates a random string of fixed size
func RandomString(size int) string {
	buf := make([]byte, size)
	for i := 0; i < size; i++ {
		buf[i] = alpha[rand.Intn(len(alpha))]
	}
	return string(buf)
}

// ComputeHmac256 creates an HMAC hash
func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// LoadHTMLFiles loads all html files in a given directory
func LoadHTMLFiles(root string, mux *gin.Engine) error {
	err := filepath.Walk(root, func(path string, info os.FileInfo, e error) error {
		if info.IsDir() {
			return nil
		}
		// Load html file
		mux.LoadHTMLFiles(path)
		return e
	})

	return err
}
