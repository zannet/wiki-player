package controllers

import (
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
	u, err := uc.UM.Get("email", g.Email)
	if err != nil {
		tracelog.CompletedError(err, "UserController", "uc.UM.NewUserModel")
		c.JSON(401, gin.H{"message": "Invalid email.", "status": 401})
	}

	// Compare hashes
	hash := utils.ComputeHmac256(g.Password, utils.ConfigEntry("Salt"))
	if hash != u.Hash {
		tracelog.CompletedError(err, "UserController", "Hashes comparison")
		c.JSON(401, gin.H{"message": "Invalid password.", "status": 401})
	}

	// Set session
	ud := map[string]string{"id": u.Id, "email": u.Email, "username": u.Username, "firstName": u.FirstName,
		"lastName": u.LastName, "accessLevel": 10}
	err = uc.setSession(c, ud)
	if err != nil {
		tracelog.CompletedError(err, "UserController", "uc.setSession")
		c.JSON(500, gin.H{"message": "Something went wrong.", "status": 500})
	}
}

func (sc *UserController) Register(c *gin.Context) {
	// Bind params
	var r Register
	c.Bind(&r)

	// Create password hash
	hash := utils.ComputeHmac256(r.Password, utils.ConfigEntry("Salt"))

	// Set user data
	sc.UM.UserData.Email = r.Email
	sc.UM.UserData.Username = r.Username
	sc.UM.UserData.FirstName = r.FirstName
	sc.UM.UserData.LastName = r.LastName
	sc.UM.UserData.Hash = hash
	sc.UM.UserData.AccessLevel = 10 // Figure out how to set this properly
	sc.UM.UserData.Joined = time.Now().Local()

	// Save user
	id, err := sc.UM.Save()
	if err != nil {
		tracelog.CompletedError(err, "UserController", "sc.UM.Save")
		c.JSON(500, gin.H{"message": "Something went wrong.", "status": 500})
	}

	// Set session
	ud := map[string]string{"id": id, "email": r.Email, "username": r.Username, "firstName": r.FirstName,
		"lastName": r.LastName, "accessLevel": 10}
	err = uc.setSession(c, ud)
	if err != nil {
		tracelog.CompletedError(err, "UserController", "uc.setSession")
		c.JSON(500, gin.H{"message": "Something went wrong.", "status": 500})
	}

	// Redirect user to home page

}

func (uc *UserController) setSession(c *gin.Context, ud map[string]string) (err error) {
	// Store in session variable
	session, _ := uc.Store.Get(c.Request, "session")

	// Set some session values
	session.Values["uid"] = ud["id"]
	session.Values["email"] = ud["email"]
	session.Values["username"] = ud["username"]
	session.Values["firstName"] = ud["firstName"]
	session.Values["lastName"] = ud["lastName"]
	session.Values["accessLevel"] = ud["accessLevel"]

	// Save session
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		return err
	}

	return nil
}
