package userTests

import (
	"net/http"
	"strings"
	"testing"
	"fmt"
)

var (
	registerUrl string = "http://localhost:9000/api/v1/users/register"
	testUsername string = "x"
	userData string = `{ "email": "red.adaya@x.com", "username": "`+ testUsername +`", "first_name": "Red", "last_name": "Adaya", "password": "shadowfiend" }`
)

func Register() (*http.Response, error) {
	rdr := strings.NewReader(userData)
	req, err := http.NewRequest("POST", registerUrl, rdr)
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)

	return resp, err
}

func TestUserRegister(t *testing.T) {
	resp, err := Register()

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code is 200 but returned: %d", resp.StatusCode)
	}

	// TODO: Validate the format returned -- it should be in JSON
}
