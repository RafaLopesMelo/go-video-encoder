package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

type Connection struct {
    db *sql.DB
}

func NewConnection() *Connection {
    db, err := sql.Open("mysql", "curseduca:%#curseduca/video_encoder")

    if err != nil {
        log.Fatal("Could not connect to database.")
        panic(err)
    }

    connection := Connection{
        db: db,
    }

    return &connection
}
