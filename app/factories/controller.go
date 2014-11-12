package factories

import (
	"errors"

	"github.com/adred/wiki-player/app/controllers"
	"github.com/adred/wiki-player/app/interfaces"
	"github.com/adred/wiki-player/app/mockControllers"
	"github.com/adred/wiki-player/app/mockModels"
	"github.com/adred/wiki-player/app/models"
	"github.com/gorilla/sessions"
)

var errInvalidMode = errors.New("Invalid App Mode")

// NewUserController returns instance of User controller
func NewUserController(um interfaces.UserModel, store *sessions.CookieStore, mode string) (interfaces.UserController, error) {
	if mode == "mock" {
		return &mockControllers.User{UM: um.(*mockModels.User), Store: store}, nil
	} else if mode == "real" {
		return &controllers.User{UM: um.(*models.User), Store: store}, nil
	} else {
		return nil, errInvalidMode
	}
}
