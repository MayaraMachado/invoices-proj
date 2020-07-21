package db

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

func InitPostgresDB(dataSourceName string) *sql.DB {
    var err error
    db, err := sql.Open("postgres", dataSourceName)
    if err != nil {
        log.Panic(err)
    }

    if err = db.Ping(); err != nil {
        log.Panic(err)
    }

    return db
}
