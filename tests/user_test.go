package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/adred/wiki-player/app/models"
	"github.com/adred/wiki-player/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type register struct {
	Email     string `json:"email" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

func TestUserRegister(t *testing.T) {
	var jsonStr = []byte(`{ "email": "red.adaya@x.com", "username": "x", "first_name": "Red", "last_name": "Adaya", "password": "shadowfiend" }`)
	req, _ := http.NewRequest("POST", "/api/v1/users/register", bytes.NewBuffer(jsonStr))

	sSave := true
	sCreate := true
	w := httptest.NewRecorder()
	mux := gin.New()
	mux.POST("/api/v1/users/register", func(c *gin.Context) {
		var r register
		// Bind params
		c.Bind(&r)

		// Create DB connection
		dbHandle := utils.DbHandle()

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
		id, err := um.Create()
		if err != nil {
			c.Data(200, "application/json", []byte(`500`))
		} else {
			c.Data(200, "application/json", []byte(`200`))
		}

		// Set user ID to last inserted ID
		um.UserData.Id = id

		// Get new cookie store
		store := sessions.NewCookieStore([]byte(utils.ConfigEntry("SecretKey")))
		store.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   3600,
			HttpOnly: true,
		}

		// Set session
		session, err := store.Get(c.Request, "testSession")
		if err != nil {
			sCreate = false
		}

		// Set some session values
		session.Values["uid"] = um.UserData.Id
		session.Values["email"] = um.UserData.Email
		session.Values["username"] = um.UserData.Username
		session.Values["firstName"] = um.UserData.FirstName
		session.Values["lastName"] = um.UserData.LastName
		session.Values["accessLevel"] = um.UserData.AccessLevel

		// Save session
		err = session.Save(c.Request, c.Writer)
		if err != nil {
			sSave = false
		}
	})

	mux.ServeHTTP(w, req)

	if sCreate == false {
		t.Error("Test failed to get/create a session.")
	}

	if sSave == false {
		t.Error("Test failed to save session")
	}

	if w.Body.String() != "200" {
		t.Errorf("Status code should be 200, was: %s", w.Body.String())
	}

	if w.HeaderMap.Get("Content-Type") != "application/json" {
		t.Errorf("Content-Type should be application/json, was %s", w.HeaderMap.Get("Content-Type"))
	}
}
