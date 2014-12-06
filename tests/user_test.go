package userTests

import (
	"net/http"
	"strings"
	"testing"
)

var (
	testUsername string = "x"
	testEmail string = "red.adaya@x.com"
)

func CheckUsername() (*http.Response, error) {
	url := "http://localhost:9000/api/v1/users/check-username/" + testUsername
	req, err := http.NewRequest("GET", url, nil)
	resp, err := http.DefaultClient.Do(req)

	return resp, err
}

func CheckEmail() (*http.Response, error) {
	url := "http://localhost:9000/api/v1/users/check-email/" + testEmail
	req, err := http.NewRequest("GET", url, nil)
	resp, err := http.DefaultClient.Do(req)

	return resp, err
}


func Register() (*http.Response, error) {
	url := "http://localhost:9000/api/v1/users/register"
	userData := `{ "email": "`+ testEmail +`", "username": "`+ testUsername +`", "first_name": "Red", "last_name": "Adaya", "password": "shadowfiend" }`

	rdr := strings.NewReader(userData)
	req, err := http.NewRequest("POST", url, rdr)
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)

	return resp, err
}


// Test functions

func TestCheckUsername(t *testing.T) {
	resp, err := CheckUsername()

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code is 200 but returned %d", resp.StatusCode)
	}

	// TODO: Validate the format returned -- it should be in JSON
}

func TestCheckEmail(t *testing.T) {
	resp, err := CheckEmail()

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code is 200 but returned %d", resp.StatusCode)
	}

	// TODO: Validate the format returned -- it should be in JSON
}

func TestUserRegister(t *testing.T) {
	resp, err := Register()

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code is 200 but returned %d", resp.StatusCode)
	}

	// TODO: Validate the format returned -- it should be in JSON
}
