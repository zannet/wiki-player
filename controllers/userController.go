package controllers

import (
	"time"

	"bitbucket.org/adred/wiki-player/models"
	"bitbucket.org/adred/wiki-player/utils"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/tracelog"
	"github.com/gorilla/sessions"
)

type (
	UserController struct {
		UM    *models.UserModel
		Store *sessions.CookieStore
	}

	Login struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	Register struct {
		Email     string `json:"email" binding:"required"`
		Username  string `json:"username" binding:"required"`
		Password  string `json:"password" binding:"required"`
		FirstName string `json:"first_name" binding:"required"`
		LastName  string `json:"last_name" binding:"required"`
	}
)

func (uc *UserController) Login(c *gin.Context) {
	// Bind params
	var g Login
	c.Bind(&g)

	// Check if user exists and get instance if it does
	ud, err := uc.UM.Get("email", g.Email)
	if err != nil {
		tracelog.CompletedError(err, "UserController", "uc.UM.NewUserModel")
		c.JSON(401, gin.H{"message": "Invalid email.", "status": 401})
	}

	// Compare hashes
	hash := utils.ComputeHmac256(g.Password, utils.ConfigEntry("Salt"))
	if hash != ud.Hash {
		tracelog.CompletedError(err, "UserController", "Hashes comparison")
		c.JSON(401, gin.H{"message": "Invalid password.", "status": 401})
	}

	// Set session
	uc.UM.UserData = ud
	err = uc.setSession(c)
	if err != nil {
		tracelog.CompletedError(err, "UserController", "uc.setSession")
		c.JSON(500, gin.H{"message": "Something went wrong.", "status": 500})
	}
}

func (uc *UserController) Register(c *gin.Context) {
	// Bind params
	var r Register
	c.Bind(&r)

	// Set user data
	uc.UM.UserData.Email = r.Email
	uc.UM.UserData.Username = r.Username
	uc.UM.UserData.FirstName = r.FirstName
	uc.UM.UserData.LastName = r.LastName
	uc.UM.UserData.Hash = utils.ComputeHmac256(r.Password, utils.ConfigEntry("Salt"))
	uc.UM.UserData.AccessLevel = 10 // Figure out how to set this properly
	uc.UM.UserData.Joined = time.Now().Local()

	// Save user
	id, err := uc.UM.Save()
	if err != nil {
		tracelog.CompletedError(err, "UserController", "uc.UM.Save")
		c.JSON(500, gin.H{"message": "Something went wrong.", "status": 500})
	}

	// Set session
	err = uc.setSession(c)
	if err != nil {
		tracelog.CompletedError(err, "UserController", "uc.setSession")
		c.JSON(500, gin.H{"message": "Something went wrong.", "status": 500})
	}

	// Redirect user to home page

}

func (uc *UserController) setSession(c *gin.Context) (err error) {
	// Store in session variable
	session, _ := uc.Store.Get(c.Request, "session")

	// Set some session values
	session.Values["uid"] = uc.UM.Userdata.Id
	session.Values["email"] = uc.UM.Userdata.Email
	session.Values["username"] = uc.UM.Userdata.Username
	session.Values["firstName"] = uc.UM.Userdata.FirstName
	session.Values["lastName"] = uc.UM.Userdata.LastName
	session.Values["accessLevel"] = uc.UM.Userdata.AccessLevel

	// Save session
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		return err
	}

	return nil
}
