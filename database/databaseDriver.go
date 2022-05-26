package database

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var Db *sql.DB

func ConnectToDatabase() {
	// Capture connection properties.
	err2 := os.Setenv("DBUSER", "root")
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	err2 = os.Setenv("DBPASS", "")
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "go_database",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	Db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := Db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
