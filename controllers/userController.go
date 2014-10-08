package controllers

import (
	"time"

	"github.com/adred/wiki-player/models"
	"github.com/adred/wiki-player/utils"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/tracelog"
	"github.com/gorilla/sessions"
)

type (
	UserController struct {
		UM    *models.UserModel
		Store *sessions.CookieStore
	}

	// Login struct is used for login payload binding
	Login struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Register struct is used for register payload binding
	Register struct {
		Email     string `json:"email" binding:"required"`
		Username  string `json:"username" binding:"required"`
		Password  string `json:"password" binding:"required"`
		FirstName string `json:"first_name" binding:"required"`
		LastName  string `json:"last_name" binding:"required"`
	}

	// Update struct is used for update payload binding
	Update struct {
		Email     string `json:"email" binding:"required"`
		Password  string `json:"password" binding:"required"`
		FirstName string `json:"first_name" binding:"required"`
		LastName  string `json:"last_name" binding:"required"`
	}
)

// Login logs the user in
func (uc *UserController) Login(c *gin.Context) {
	var g Login
	// Bind params
	c.Bind(&g)

	// Check if user exists and get UserData instance if it does
	ud, err := uc.UM.Get("email", g.Username)
	if err != nil {
		// Mybe the user provided the username instead of email
		ud, err = uc.UM.Get("username", g.Username)
		if err != nil {
			tracelog.CompletedError(err, "UserController", "uc.UM.NewUserModel")
			c.JSON(401, gin.H{"message": "Invalid Username.", "status": 401})
			return
		}
	}

	// Compare hashes
	hash := utils.ComputeHmac256(g.Password, utils.ConfigEntry("Salt"))
	if hash != ud.Hash {
		tracelog.CompletedError(err, "UserController", "Hashes comparison")
		c.JSON(401, gin.H{"message": "Invalid password.", "status": 401})
		return
	}

	// Set session
	uc.UM.UserData = ud
	err = uc.setSession(c)
	if err != nil {
		tracelog.CompletedError(err, "UserController", "uc.setSession")
		c.JSON(500, gin.H{"message": "Something went wrong.", "status": 500})
		return
	}

	c.JSON(200, gin.H{"message": "Logged in successfully.", "status": 200})
}

// Logout logs the user out
func (uc *UserController) Logout(c *gin.Context) {
	uc.clearSession(c)

	c.JSON(200, gin.H{"message": "Logged out successfully.", "status": 200})
}

// Register registers the user
func (uc *UserController) Register(c *gin.Context) {
	var r Register
	// Bind params
	c.Bind(&r)

	// Set user data
	uc.UM.UserData.Email = r.Email
	uc.UM.UserData.Username = r.Username
	uc.UM.UserData.FirstName = r.FirstName
	uc.UM.UserData.LastName = r.LastName
	uc.UM.UserData.Hash = utils.ComputeHmac256(r.Password, utils.ConfigEntry("Salt"))
	uc.UM.UserData.AccessLevel = 10 // Figure out how to set this properly
	uc.UM.UserData.Joined = time.Now().Local()

	// Create user
	id, err := uc.UM.Create()
	if err != nil {
		tracelog.CompletedError(err, "UserController", "uc.UM.Save")
		c.JSON(500, gin.H{"message": "Something went wrong.", "status": 500})
		return
	}

	// Set user ID to last inserted ID
	uc.UM.UserData.Id = id

	// Set session
	err = uc.setSession(c)
	if err != nil {
		tracelog.CompletedError(err, "UserController", "uc.setSession")
		c.JSON(500, gin.H{"message": "Something went wrong.", "status": 500})
		return
	}

	c.JSON(200, gin.H{"message": "Registered successfully.", "status": 200})
}

// Update udpates the user
func (uc *UserController) Update(c *gin.Context) {
	var u Update
	// Bind params
	c.Bind(&u)

	// Set user data
	uc.UM.UserData.Email = u.Email
	uc.UM.UserData.FirstName = u.FirstName
	uc.UM.UserData.LastName = u.LastName
	uc.UM.UserData.Hash = utils.ComputeHmac256(u.Password, utils.ConfigEntry("Salt"))

	// Update user
	err := uc.UM.Update()
	if err != nil {
		tracelog.CompletedError(err, "UserController", "uc.UM.Update")
		c.JSON(500, gin.H{"message": "Something went wrong.", "status": 500})
		return
	}

	c.JSON(200, gin.H{"message": "User updated successfully.", "status": 200})
}

// Delete sends delete confirmation email to the user
func (uc *UserController) Delete(c *gin.Context) {
	//
	// Update user
	// err := uc.UM.Update()
	// if err != nil {
	// 	tracelog.CompletedError(err, "UserController", "uc.UM.Update")
	// 	c.JSON(500, gin.H{"message": "Something went wrong.", "status": 500})
	// 	return
	// }

	// c.JSON(200, gin.H{"message": "User updated successfully.", "status": 200})
}

// Delete deletes the user
func (uc *UserController) ConfirmDelete(c *gin.Context) {
	// Delete user
	err := uc.UM.Delete(c.Params.ByName("nonce"))
	if err != nil {
		tracelog.CompletedError(err, "UserController", "uc.UM.Delete")
		c.JSON(500, gin.H{"message": "Something went wrong.", "status": 500})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted successfully.", "status": 200})
}

// setSession sets the session
func (uc *UserController) setSession(c *gin.Context) (err error) {
	// Get session
	session := c.MustGet("session").(*sessions.Session)

	// Set some session values
	session.Values["uid"] = uc.UM.UserData.Id
	session.Values["email"] = uc.UM.UserData.Email
	session.Values["username"] = uc.UM.UserData.Username
	session.Values["firstName"] = uc.UM.UserData.FirstName
	session.Values["lastName"] = uc.UM.UserData.LastName
	session.Values["accessLevel"] = uc.UM.UserData.AccessLevel

	// Save session
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		return err
	}

	return nil
}

// clearSession destroys the session
func (uc *UserController) clearSession(c *gin.Context) (err error) {
	// Get session
	session := c.MustGet("session").(*sessions.Session)
	session.Options.MaxAge = -3600

	// Save session
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		return err
	}

	return nil
}
