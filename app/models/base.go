package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"

	"log"
	"to-do-app/config"

	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser    = "users"
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
	tableNameChat    = "chats"
	tableChatPick    = "picks"
)

func init() {
	fmt.Println(config.Config.SQLDriver, config.Config.DbName)
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	uuid STRING NOT NULL UNIQUE,
	name STRING,
	email STRING,
	password STRING,
	created_at DATETIME)`, tableNameUser)
	Db.Exec(cmdU)

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		title TEXT,
		deadline TEXT,
		category TEXT,
		user_id INTEGER,
		created_at DATETIME)`, tableNameTodo)
	Db.Exec(cmdT)

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STRING,
		user_id INTEGER,
		created_at DATETIME)`, tableNameSession)
	Db.Exec(cmdS)

	cmdO := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		chat STRING,
		user_id INTEGER,
		created_at DATETIME)`, tableNameChat)
	Db.Exec(cmdO)

	cmdP := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		group_name INTEGER,
		user1_id INTEGER,
		user2_id INTEGER,
		created_at DATETIME)`, tableChatPick)
	Db.Exec(cmdP)
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
