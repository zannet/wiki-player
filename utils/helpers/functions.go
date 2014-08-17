package helpers

import (
	"errors"

	"github.com/zenazn/goji/web"
)

var ErrTypeNotPresent = errors.New("Expected type not present in the request context.")

// Context returns the type stored in the context
func Context(c web.C, key string) (interface{}, error) {
	val, ok := c.Env[key]
	if !ok {
		return val, ErrTypeNotPresent
	}

	return val, nil
}
