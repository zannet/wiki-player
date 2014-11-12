package interfaces

import (
	"github.com/gin-gonic/gin"
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
