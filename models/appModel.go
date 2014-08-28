package models

import (
	"bitbucket.org/adred/wiki-player/utils"
)

type (
	AppModel struct {
		DB *utils.DB
	}

	app struct {
		Id         string
		Name       string
		Username   string
		Key        string
		Registered string
	}
)

func (sm *AppModel) GetAll() error {
	return nil
}

func (sm *AppModel) Get(id string) {

}
