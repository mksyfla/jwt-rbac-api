package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", os.Getenv("dbUser"), os.Getenv("dbPass"), os.Getenv("dbName")),
	)

	if err != nil {
		panic(err)
	}

	return db
}
