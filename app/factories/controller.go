package factories

import (
	"errors"

	"github.com/adred/wiki-player/app/controllers"
	"github.com/adred/wiki-player/app/interfaces"
	"github.com/adred/wiki-player/app/models"
	"github.com/gorilla/sessions"
)

var errInvalidMode = errors.New("Invalid App Mode")

// NewUserController returns instance of User controller
func NewUserController(um interfaces.UserModelInterface, store *sessions.CookieStore, mode string) (interfaces.UserControllerInterface, error) {
	if mode == "mock" {
		return &controllers.MockUserController{UM: um.(*models.MockUserModel), Store: store}, nil
	} else if mode == "real" {
		return &controllers.UserController{UM: um.(*models.UserModel), Store: store}, nil
	} else {
		return nil, errInvalidMode
	}
}
