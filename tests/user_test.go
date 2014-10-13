package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
    "github.com/gin-gonic/gin"
    "github.com/adred/wiki-player/app/models"
)

register struct {
    Email     string `json:"email" binding:"required"`
    Username  string `json:"username" binding:"required"`
    Password  string `json:"password" binding:"required"`
    FirstName string `json:"first_name" binding:"required"`
    LastName  string `json:"last_name" binding:"required"`
}

func TestUserRegister(t *testing.T) {
    var jsonStr = []byte(`{ "email": "red.adaya@gmail.com", "username": "shadowfiend", "first_name": "Red", "last_name": "Adaya", "password": "shadowfiend" }`)
    req, err := http.NewRequest("POST", "/api/v1/users/register", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    name := ""
    mux := gin.New()
    mux.POST("/api/v1/users/register", func(c *Context) {
        var r register
        // Bind params
        c.Bind(&r)

        // Set user data
        um := &models.UserModel{DbHandle: dbHandle, UserData: &models.UserData{}}
        um.UserData.Email = r.Email
        um.UserData.Username = r.Username
        um.UserData.FirstName = r.FirstName
        um.UserData.LastName = r.LastName
        um.UserData.Hash = utils.ComputeHmac256(r.Password, utils.ConfigEntry("Salt"))
        um.UserData.AccessLevel = 10 // Figure out how to set this properly
        um.UserData.Joined = time.Now().Local()

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
    })
    mux.ServeHTTP(w, req)

	Convey("Given some integer with a starting value", t, func() {
		x := 1

		Convey("When the integer is incremented", func() {
			x++

			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 2)
			})
		})
	})
}
