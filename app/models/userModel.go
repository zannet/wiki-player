package models

import (
	"database/sql"
	"errors"
	"strconv"
	"time"
)

// UserModelInterface is the Interface for User models
type UserModelInterface interface {
	User(field, value string) (*UserData, error)
	Update() error
	Create() (string, error)
	Delete(nonce string) error
}

var errInvalidMode = errors.New("Invalid App Mode")

// NewUserModel returns instance of User models
func NewUserModel(dbHandle *sql.DB, mode string) (UserModelInterface, error) {
	if mode == "mock" {
		return &MockUserModel{DbHandle: dbHandle, UserData: &UserData{}}, nil
	} else if mode == "real" {
		return &UserModel{DbHandle: dbHandle, UserData: &UserData{}}, nil
	} else {
		return nil, errInvalidMode
	}
}

type (
	// UserModel is type of this class
	UserModel struct {
		DbHandle *sql.DB
		UserData *UserData
	}

	// UserData defines the fields of the users table
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

// User returns UserData instance
func (um *UserModel) User(field, value string) (*UserData, error) {
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

// Update updates the user
func (um *UserModel) Update() error {
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
func (um *UserModel) Create() (string, error) {
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
func (um *UserModel) Delete(nonce string) error {
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
