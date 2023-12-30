package db

import (
	cfg "blogsite/backend/pkg/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var dbUrl = cfg.DatabaseString

type Database struct {
	*sql.DB
}

func EstablishDBCon() (Database, error) {
	db, err := sql.Open("mysql", dbUrl)

	if err != nil {
		fmt.Print(err.Error())
		log.Fatalf("Failed to establish connection to: %s", dbUrl)
	}

	return Database{db}, nil
}
