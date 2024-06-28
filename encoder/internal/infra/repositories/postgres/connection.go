package postgres

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type Connection struct {
    DB *sql.DB
}

func NewConnection() *Connection {
    username := os.Getenv("POSTGRES_USERNAME")
    password, err := url.QueryUnescape(os.Getenv("POSTGRES_PASSWORD"))
    host := os.Getenv("POSTGRES_HOST")
    port := os.Getenv("POSTGRES_PORT")
    database := os.Getenv("POSTGRES_DATABASE")

    if err != nil {
        panic("Invalid postgres password provided in dot files.")
    }

    url := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        host,
        port,
        username,
        password,
        database,
    )

    db, err := sql.Open("postgres", url)

    if err != nil {
        log.Fatal("Could not connect to database.")
        panic(err)
    }

    connection := Connection{
        DB: db,
    }

    return &connection
}
