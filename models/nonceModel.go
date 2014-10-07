package models

import (
	"database/sql"
	"time"

	"github.com/adred/wiki-player/utils"
)

type (
	NonceModel struct {
		DbHandle *sql.DB
	}
)

func (nm *NonceModel) Create(uid string) (nonce string, err error) {
	stmt, err := nm.DbHandle.Prepare("INSERT INTO nonces VALUES ('', ?, ?, ?)")
	if err != nil {
		return "", err
	}

	nonce = utils.RandomString(16)
	created := time.Now().Local()

	_, err = stmt.Exec(uid, nonce, created)
	if err != nil {
		return "", err
	}

	return nonce, nil
}

func (nm *NonceModel) Verify(nonce string) (exists bool, err error) {
	stmt, err := nm.DbHandle.Prepare("SELECT id FROM nonces WHERE nonce = ?")
	if err != nil {
		return false, err
	}

	_, err = stmt.Query(nonce)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (nm *NonceModel) Delete(id string) (err error) {
	stmt, err := nm.DbHandle.Prepare("DELETE FROM nonces WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
