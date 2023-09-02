package graph

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var DB *sql.DB

func Database() {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_URL")) // Corrected syntax
	if err != nil {
		panic(err)
	}

	DB = db
}
