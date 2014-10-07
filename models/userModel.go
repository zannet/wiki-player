package models

import (
	"database/sql"
	"strconv"
	"time"
)

type (
	UserModel struct {
		DbHandle *sql.DB
		UserData *UserData
	}

	UserData struct {
		Id          string
		Email       string
		Username    string
		FirstName   string
		LastName    string
		Hash        string
		AccessLevel int
		Joined      time.Time
	}
)

func (um *UserModel) Get(field, value string) (ud *UserData, err error) {
	query := "SELECT id, email, username, first_name, last_name, hash, access_level, joined FROM users WHERE "
	query += field
	query += " = ?"

	stmt, err := um.DbHandle.Prepare(query)
	if err != nil {
		return &UserData{}, err
	}

	var accessLevel int
	var joined time.Time
	var id, email, username, firstName, lastName, hash string

	err = stmt.QueryRow(value).Scan(&id, &email, &username, &firstName, &lastName, &hash, &accessLevel, &joined)
	if err != nil {
		return &UserData{}, err
	}

	return &UserData{
		Id:          id,
		Email:       email,
		Username:    username,
		FirstName:   firstName,
		LastName:    lastName,
		Hash:        hash,
		AccessLevel: accessLevel,
		Joined:      joined,
	}, nil
}

func (um *UserModel) Update() (err error) {
	stmt, err := um.DbHandle.Prepare("UPDATE users SET email = ?, first_name = ?, last_name = ?, hash = ? WHERE id = ?")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(um.UserData.Email, um.UserData.FirstName, um.UserData.LastName, um.UserData.Hash, um.UserData.Id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (um *UserModel) Create() (id string, err error) {
	stmt, err := um.DbHandle.Prepare("INSERT INTO users VALUES ('', ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return "", err
	}

	res, err := stmt.Exec(um.UserData.Email, um.UserData.Username, um.UserData.FirstName, um.UserData.LastName,
		um.UserData.Hash, um.UserData.AccessLevel, um.UserData.Joined)
	if err != nil {
		return "", err
	}

	lId, err := res.LastInsertId()
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(lId, 10), nil
}

func (um *UserModel) Delete(nonce string) (err error) {
	stmt, err := um.DbHandle.Prepare("DELETE FROM users WHERE nonce = ?")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(nonce)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
