package mysql

import (
	"database/sql"
	"net/url"
	"os"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

type Connection struct {
    DB *sql.DB
}

func NewConnection() *Connection {
    username := os.Getenv("MYSQL_USERNAME")
    password, err := url.QueryUnescape(os.Getenv("MYSQL_PASSWORD"))
    host := os.Getenv("MYSQL_HOST")
    port := os.Getenv("MYSQL_PORT")
    database := os.Getenv("MYSQL_DATABASE")

    if err != nil {
        panic("Invalid msqyl password provided in dot files.")
    }

    url := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database
    db, err := sql.Open("mysql", url)

    if err != nil {
        log.Fatal("Could not connect to database.")
        panic(err)
    }

    connection := Connection{
        DB: db,
    }

    return &connection
}
