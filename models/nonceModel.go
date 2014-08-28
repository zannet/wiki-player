package models

import (
	"math/rand"
	"time"

	"bitbucket.org/adred/wiki-player/utils"
)

type (
	NonceModel struct {
		DB *utils.DB
	}
)

func (nm *NonceModel) Create() (string, error) {
	stmt, err := nm.DB.Handle.Prepare("INSERT INTO nonces VALUES ('', ?, ?, ?)")
	if err != nil {
		return "", err
	}

	appId := 1
	nonce := randString(10)
	created := time.Now().Local()

	_, err = stmt.Exec(appId, nonce, created)
	if err != nil {
		return "", err
	}

	return nonce, nil
}

func (nm *NonceModel) Verify(nonce string) (bool, error) {
	stmt, err := nm.DB.Handle.Prepare("SELECT id FROM nonces WHERE key = ?")
	if err != nil {
		return false, err
	}

	_, err = stmt.Query(nonce)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (nm *NonceModel) Delete(id string) error {
	stmt, err := nm.DB.Handle.Prepare("DELETE FROM nonces WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func randString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
