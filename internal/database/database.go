package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr/v2"
)

type Database struct {
	sess     *dbr.Session
	readSess *dbr.Session
}

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("root:root@tcp(localhost:%v)/%v", os.Getenv("DB_PORT"), os.Getenv("DB_NANE")))
	if err != nil {
		return db, fmt.Errorf("db connection error:%w", err)
	}
	return db, nil
}
