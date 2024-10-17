package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"udemy-todo-app/config"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var Db *sql.DB

var err error

const (
	tableNameUser    = "users"
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
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

	cmdU := fmt.Sprintf(`create table if not exists %s(
        id serial primary key,
        uuid text not null unique,
        name text,
        email text,
        password text,
        created_at timestamp)`, tableNameUser)
	_, err = Db.Exec(cmdU)
	if err != nil {
		log.Fatalln("Error creating users table:", err)
	}

	cmdT := fmt.Sprintf(`create table if not exists %s(
        id serial primary key,
        content text,
        user_id integer,
        created_at timestamp,
        constraint fk_user_id foreign key(user_id) references users (id)
        )`, tableNameTodo)
	_, err = Db.Exec(cmdT)
	if err != nil {
		log.Fatalln("Error creating todos table:", err)
	}
	cmdTI := fmt.Sprintf(`create index if not exists user_todos on %s(user_id)`, tableNameTodo)
	_, err = Db.Exec(cmdTI)
	if err != nil {
		log.Fatalln("Error creating index on todos table:", err)
	}

	cmdS := fmt.Sprintf(`create table if not exists %s(
        id serial primary key,
        uuid text not null unique,
        user_id integer,
        email text,
        created_at timestamp,
        constraint fk_user_id foreign key(user_id) references users (id)
        )`, tableNameSession)
	Db.Exec(cmdS)
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
