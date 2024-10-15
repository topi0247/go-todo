package models

import (
    "crypto/sha1"
    "fmt"
    "log"
    "database/sql"
    "udemy-todo-app/config"

    "github.com/google/uuid"
    _ "github.com/lib/pq"
)

var Db *sql.DB

var err error

const (
    tableNameUser = "users"
)

func init() {
    var err error
    connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
        config.Config.DbUser,
        config.Config.DbPassword,
        config.Config.DbName,
        config.Config.DbHost,
        config.Config.DbPort)
    Db, err = sql.Open(config.Config.SQLDriver, connStr)
    if err != nil {
        log.Fatalln(err)
    }

    cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
        id SERIAL PRIMARY KEY,
        uuid TEXT NOT NULL UNIQUE,
        name TEXT,
        email TEXT,
        password TEXT,
        created_at TIMESTAMP)`, tableNameUser)
    _, err = Db.Exec(cmdU)
    if err != nil {
        log.Fatalln(err)
    }
}

func createUUID() (uuidobj uuid.UUID) {
    uuidobj, _ = uuid.NewUUID()
    return uuidobj
}

func  Encrypt(plaintext string) (cryptext string) {
    cryptext = fmt.Sprintf("%x", sha1.Sum([] byte(plaintext)))
    return cryptext
}