package models

import (
	"bitbucket.org/adred/wiki-player/utils"
)

type (
	UserModel struct {
		DB *utils.DB
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
