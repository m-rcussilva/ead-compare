package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/m-rcussilva/go-private/tree/main/2024/projects/02-ead-compare/ead-compare-backend/src/config"
)

func ConnectionDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.StringConnectDB)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
