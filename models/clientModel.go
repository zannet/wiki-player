package models

import (
	"database/sql"
	"time"

	"bitbucket.org/adred/wiki-player/utils"
)

type (
	ClientModel struct {
		DbHandle *sql.DB
	}
	// Table structure
	client struct {
		Id         string
		UserId     string
		clientName string
		PrivateKey string
		ApiKey     string
		Registered string
	}
)

func (cm *ClientModel) Register(userId string, clientName string) (map[string]string, error) {
	stmt, err := cm.DbHandle.Prepare("INSERT INTO clients VALUES ('', ?, ?, ?, ?, ?)")
	if err != nil {
		return map[string]string{}, err
	}

	apiKey := utils.RandomString(32)
	privateKey := utils.RandomString(32)
	registered := time.Now().Local()

	_, err = stmt.Exec(userId, clientName, apiKey, privateKey, registered)
	if err != nil {
		return map[string]string{}, err
	}

	return map[string]string{"apiKey": apiKey, "privateKey": privateKey}, nil
}

func (cm *ClientModel) Verify(apiKey string) (bool, error) {
	stmt, err := cm.DbHandle.Prepare("SELECT id FROM clients WHERE apiKey = ?")
	if err != nil {
		return false, err
	}

	_, err = stmt.Query(apiKey)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (cm *ClientModel) PrivateKey(apiKey string) (string, error) {
	stmt, err := cm.DbHandle.Prepare("SELECT privateKey FROM clients WHERE apiKey = ?")
	if err != nil {
		return "", err
	}

	var privateKey string
	err = stmt.QueryRow(apiKey).Scan(&privateKey)
	if err != nil {
		return "", err
	}

	return privateKey, nil
}

func (cm *ClientModel) Delete(id string) error {
	stmt, err := cm.DbHandle.Prepare("DELETE FROM clients WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
