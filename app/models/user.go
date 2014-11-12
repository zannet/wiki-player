package models

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/adred/wiki-player/app/interfaces"
)

type (
	// User is type of this class
	User struct {
		DbHandle *sql.DB
		UserData *userData
	}

	// userData defines the fields of the users table
	userData struct {
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

// User returns UserData instance
func (um *User) User(field, value string) (interfaces.UserModel, error) {
	query := "SELECT id, email, username, first_name, last_name, hash, access_level, joined FROM users WHERE "
	query += field
	query += " = ?"

	stmt, err := um.DbHandle.Prepare(query)
	if err != nil {
		return nil, err
	}

	var accessLevel int
	var joined time.Time
	var id, email, username, firstName, lastName, hash string

	err = stmt.QueryRow(value).Scan(&id, &email, &username, &firstName, &lastName, &hash, &accessLevel, &joined)
	if err != nil {
		return nil, err
	}

	return &User{
		UserData: &userData{
			Id:          id,
			Email:       email,
			Username:    username,
			FirstName:   firstName,
			LastName:    lastName,
			Hash:        hash,
			AccessLevel: accessLevel,
			Joined:      joined,
		},
	}, nil
}

// Update updates the user
func (um *User) Update() error {
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

// Create creates a user
func (um *User) Create() (string, error) {
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

// Delete deletes a user
func (um *User) Delete(nonce string) error {
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
