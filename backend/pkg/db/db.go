package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var dbUrl = "root:123@tcp(127.0.0.1:3306)/users"

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
