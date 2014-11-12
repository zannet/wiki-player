package factories

import (
	"database/sql"

	"github.com/adred/wiki-player/app/interfaces"
	"github.com/adred/wiki-player/app/models"
)

// NewUserModel returns instance of User models
func NewUserModel(dbHandle *sql.DB, mode string) (interfaces.UserModelInterface, error) {
	if mode == "mock" {
		return &models.MockUserModel{DbHandle: dbHandle}, nil
	} else if mode == "real" {
		return &models.UserModel{DbHandle: dbHandle}, nil
	} else {
		return nil, errInvalidMode
	}
}
