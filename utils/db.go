package utils

import (
	"database/sql"

	"bitbucket.org/adred/cloud-music-player/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goinggo/tracelog"
)

type DB struct {
	Conn *sql.DB
}

func (db *DB) Open() (err error) {
	db.Conn, err = sql.Open(config.Entry("Driver"),
		config.Entry("Username")+":"+config.Entry("Password")+"@/"+config.Entry("Database"))

	if err != nil {
		tracelog.CompletedError(err, "DB", "Open")
		return err
	}
	err = db.Conn.Ping()
	if err != nil {
		tracelog.CompletedError(err, "DB", "db.Conn.Ping")
		return err
	}

	return nil
}

func (db *DB) Close() {
	db.Conn.Close()
}
