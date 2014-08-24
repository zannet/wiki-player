package utils

import (
	"database/sql"

	"bitbucket.org/adred/cloud-music-player/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goinggo/tracelog"
)

type DB struct {
	Handle *sql.DB
}

func NewDB() (db *DB, err error) {
	handle, err := sql.Open(config.Entry("Driver"),
		config.Entry("Username")+":"+config.Entry("Password")+"@/"+config.Entry("Database"))

	if err != nil {
		tracelog.CompletedError(err, "DB", "NewDB")
		return &DB{Handle: nil}, err
	}
	err = handle.Ping()
	if err != nil {
		tracelog.CompletedError(err, "DB", "handle.Ping")
		return &DB{Handle: nil}, err
	}

	return &DB{Handle: handle}, nil
}
