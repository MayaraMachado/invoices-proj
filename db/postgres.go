package db

import (
    "os"
    "log"
    // "fmt"
    "database/sql"
    _ "github.com/lib/pq"
)

// type DatabaseConnection interface {
// 	CloseDbConnection()
// }

// type databaseConnection struct {
// 	db *sql.DB
// }

var db *sql.DB

const (
    dbhost = "DBHOST"
    dbport = "DBPORT"
    dbuser = "DBUSER"
    dbpass = "DBPASS"
    dbname = "DBNAME"
    dbtype = "postgres"
)

func NewDB() *sql.DB {
    var err error
    // var databaseInfo = "postgres://postgres:123@localhost/postgres?sslmode=disable"

    // config := dbConfig()
    // var err error
    // databaseInfo := fmt.Sprintf("host=%s port=%s user=%s "+
    //     "password=%s dbname=%s sslmode=disable",
    //     config[dbhost], config[dbport],
    //     config[dbuser], config[dbpass], config[dbname])

    db, err = sql.Open(dbtype, os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Panic(err)
    }

    if err = db.Ping(); err != nil {
        log.Panic(err)
    }

    return db
}

func dbConfig() map[string]string {
    conf := make(map[string]string)
    host, ok := os.LookupEnv(dbhost)
    if !ok {
        panic("DBHOST environment variable required but not set")
    }
    port, ok := os.LookupEnv(dbport)
    if !ok {
        panic("DBPORT environment variable required but not set")
    }
    user, ok := os.LookupEnv(dbuser)
    if !ok {
        panic("DBUSER environment variable required but not set")
    }
    password, ok := os.LookupEnv(dbpass)
    if !ok {
        panic("DBPASS environment variable required but not set")
    }
    name, ok := os.LookupEnv(dbname)
    if !ok {
        panic("DBNAME environment variable required but not set")
    }
    conf[dbhost] = host
    conf[dbport] = port
    conf[dbuser] = user
    conf[dbpass] = password
    conf[dbname] = name
    return conf
}