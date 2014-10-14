package main

import (
	"net/http"
	"strings"
	"testing"
)

func TestUserRegister(t *testing.T) {
	url := "http://localhost:9000/api/v1/users/register"
	jsonStr := `{ "email": "red.adaya@x.com", "username": "x", "first_name": "Red", "last_name": "Adaya", "password": "shadowfiend" }`

	rdr := strings.NewReader(jsonStr)
	req, err := http.NewRequest("POST", url, rdr)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	// TODO: Validate the format returned -- it should be in JSON
}
