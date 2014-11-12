package factories

import (
	"database/sql"

	"github.com/adred/wiki-player/app/interfaces"
	"github.com/adred/wiki-player/app/mockModels"
	"github.com/adred/wiki-player/app/models"
)

// NewUserModel returns instance of User models
func NewUserModel(dbHandle *sql.DB, mode string) (interfaces.UserModel, error) {
	if mode == "mock" {
		return &mockModels.User{DbHandle: dbHandle}, nil
	} else if mode == "real" {
		return &models.User{DbHandle: dbHandle}, nil
	} else {
		return nil, errInvalidMode
	}
}
