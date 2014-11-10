package models

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

type (
	// MockUserModel is type of this class
	MockUserModel struct {
		DbHandle *sql.DB
		UserData map[string]string
	}
)

// User returns UserData instance
func (um *MockUserModel) User(field, value string) (map[string]string, error) {
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

	return map[string]string{
		"Id":          id,
		"Email":       email,
		"Username":    username,
		"FirstName":   firstName,
		"LastName":    lastName,
		"Hash":        hash,
		"AccessLevel": strconv.Itoa(accessLevel),
		"Joined":      fmt.Sprint(joined),
	}, nil
}

// Update updates the user
func (um *MockUserModel) Update() error {
	stmt, err := um.DbHandle.Prepare("UPDATE users SET email = ?, first_name = ?, last_name = ?, hash = ? WHERE id = ?")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(um.UserData["Email"], um.UserData["FirstName"], um.UserData["LastName"], um.UserData["Hash"], um.UserData["Id"])
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
func (um *MockUserModel) Create() (string, error) {
	stmt, err := um.DbHandle.Prepare("INSERT INTO users VALUES ('', ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return "", err
	}

	joined, err := time.Parse(time.RFC3339, um.UserData["Joined"])
	if err != nil {
		return "", err
	}

	accessLevel, err := strconv.Atoi(um.UserData["AccessLevel"])
	if err != nil {
		return "", err
	}

	res, err := stmt.Exec(um.UserData["Email"], um.UserData["Username"], um.UserData["FirstName"], um.UserData["LastName"],
		um.UserData["Hash"], accessLevel, joined)
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
func (um *MockUserModel) Delete(nonce string) error {
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
