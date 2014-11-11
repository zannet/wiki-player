package factory

import (
    "errors"
    "database/sql"

    "github.com/adred/wiki-player/models"
    "github.com/adred/wiki-player/controllers"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/sessions"
)

// UserControllerInterface is the Interface for User controllers
type UserControllerInterface interface {
    Login(c *gin.Context)
    Logout(c *gin.Context)
    Register(c *gin.Context)
    Update(c *gin.Context)
    Delete(c *gin.Context)
    ConfirmDelete(c *gin.Context)
}

var errInvalidMode = errors.New("Invalid App Mode")

// NewUserController returns instance of User controller
func NewUserController(um UserModelInterface, store *sessions.CookieStore, mode string) (UserControllerInterface, error) {
    if mode == "mock" {
        return &controllers.MockUserController{UM: um.(*models.MockUserModel), Store: store}, nil
    } else if mode == "real" {
        return &controllers.UserController{UM: um.(*models.UserModel), Store: store}, nil
    } else {
        return nil, errInvalidMode
    }
}

// UserModelInterface is the Interface for User models
type UserModelInterface interface {
    User(field, value string) (*models.UserData, error)
    Update() error
    Create() (string, error)
    Delete(nonce string) error
}

// NewUserModel returns instance of User models
func NewUserModel(dbHandle *sql.DB, mode string) (UserModelInterface, error) {
    if mode == "mock" {
        return &models.MockUserModel{DbHandle: dbHandle, UserData: &models.UserData{}}, nil
    } else if mode == "real" {
        return &models.UserModel{DbHandle: dbHandle, UserData: &models.UserData{}}, nil
    } else {
        return nil, errInvalidMode
    }
}
