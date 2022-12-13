package db

import (
	"database/sql"
	"strconv"

	types "dic-morphy/types/struct"

	_ "github.com/go-sql-driver/mysql"
)

// type tpDB struct {}

func Connect(opt types.DBStruct) *sql.DB {
	dbn, err := sql.Open(opt.Driver, opt.Username+":"+opt.Password+"@tcp("+opt.Host+":"+strconv.Itoa(opt.Port)+")/"+opt.DBname)

	if err != nil {
		panic(err.Error())
	}

	return dbn
}

func Select(db sql.DB, query string, args ...any) *sql.Rows {
	rows, err := db.Query(query, args...)

	if err != nil {
		panic(err.Error())
	}

	return rows
}

func Begin(db sql.DB) (*sql.Tx, error) {
	return db.Begin()
}

func Close(db sql.DB) error {
	return db.Close()
}
