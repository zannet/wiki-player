package utils

import (
	"database/sql"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/goinggo/tracelog"
)

type (
	dbSingleton struct {
		once  sync.Once
		value *sql.DB
	}
)

var handle dbSingleton

func MustLoadDB() {
	handle.once.Do(func() {
		conn, err := sql.Open(ConfigEntry("Driver"),
			ConfigEntry("Username")+":"+ConfigEntry("Password")+"@/"+ConfigEntry("Database"))
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

func DBHandle() *sql.DB {
	MustLoadDB()
	return handle.value
}
