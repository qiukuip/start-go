package main

import (
	"database/sql"
	"fmt"
	// "os"
  "log"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
  log.SetPrefix("db-access: ")
	log.SetFlags(0)

	cfg := mysql.NewConfig()
	// cfg.User = os.Getenv("DBUSER")
	// cfg.Passwd = os.Getenv("DBPASS")
	cfg.User = "root"
	cfg.Passwd = "ABab1234!"
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "recordings"

	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")
}
