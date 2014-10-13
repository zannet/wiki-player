package utils

import (
	"database/sql"
	"sync"

	_ "github.com/go-sql-driver/mysql" // Automatically load mysql driver
	"github.com/goinggo/tracelog"
)

type (
	dbSingleton struct {
		once  sync.Once
		value *sql.DB
	}
)

var handle dbSingleton

// MustLoadDB loads the database
func MustLoadDB() {
	handle.once.Do(func() {
		conn, err := sql.Open(ConfigEntry("Driver"),
			ConfigEntry("Username")+":"+ConfigEntry("Password")+"@/"+ConfigEntry("Database")+"?parseTime=true")
		if err != nil {
			tracelog.CompletedError(err, "MustLoadDB", "sql.Open")
			panic(err.Error())
		}

		handle.value = conn
		err = handle.value.Ping()
		if err != nil {
			tracelog.CompletedError(err, "MustLoadDB", "handle.value.Ping")
			panic(err.Error())
		}
	})
}

// DbHandle retuns a handle of the database connection
func DbHandle() *sql.DB {
	MustLoadDB()
	return handle.value
}
