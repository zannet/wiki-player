package models

import (
	"database/sql"
)

type (
	UserModel struct {
		DBHandle *sql.DB
	}

	user struct {
		Id       string
		First    string
		Last     string
		Username string
		Email    string
		Hash     string
		Joined   string
	}
)

func (sm *UserModel) GetAll() error {
	return nil
}

func (sm *UserModel) Get(id string) {

}
