package userTests

import (
	"net/http"
	"strings"
	"testing"
)

func Register() (*http.Response, error) {
	url := "http://localhost:9000/api/v1/users/register"
	jsonStr := `{ "email": "red.adaya@x.com", "username": "x", "first_name": "Red", "last_name": "Adaya", "password": "shadowfiend" }`

	rdr := strings.NewReader(jsonStr)
	req, err := http.NewRequest("POST", url, rdr)
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
		t.Errorf("Success expected: %d", resp.StatusCode)
	}

	// TODO: Validate the format returned -- it should be in JSON
}
