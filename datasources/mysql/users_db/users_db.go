package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	user   = "USERS_USERNAME"
	pass   = "USERS_PASSWD"
	host   = "USERS_HOST"
	scheme = "USERS_SCHERME"
)

var (
	Db *sql.DB
)

func init() {
	var err error

	log.Println(os.Getenv(user), os.Getenv(pass), os.Getenv(host), os.Getenv(scheme))

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", 
	os.Getenv(user), os.Getenv(pass), os.Getenv(host), os.Getenv(scheme))

	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := Db.Ping(); err != nil {
		panic(err)
	}

	log.Println("database sucessfully connected")
}
