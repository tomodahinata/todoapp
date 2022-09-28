package models

import (
	"log"
	"time"
)

type Chat struct {
	ID        int
	Chat      string
	UserID    int
	CreatedAt time.Time
	Name      string
}

func (u *User) ChatTodo(chat string) (err error) {
	cmd := `insert into chats(
	chat,
	user_id,
	created_at) values(?,?,?)`

	_, err = Db.Exec(cmd, chat, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
func GetChat(id int) (chat Chat, err error) {
	cmd := `
		SELECT 
			id,chat,user_id,created_at 
		FROM 
			chats
		WHERE 
			id=?
			`
	chat = Chat{}

	err = Db.QueryRow(cmd, id).Scan(
		&chat.ID,
		&chat.Chat,
		&chat.UserID,
		&chat.CreatedAt)

	return chat, err
}

func GetChats() (chats []Chat, err error) {
	cmd := `SELECT id,chat,user_id,created_at FROM chats`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var chat Chat
		err = rows.Scan(&chat.ID,
			&chat.Chat,
			&chat.UserID,
			&chat.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		chats = append(chats, chat)
	}
	rows.Close()

	return chats, err
}
func (u *User) GetChatsByUser() (chats []Chat, err error) {
	cmd := `
		SELECT 
			id, 
			chat,  	
			user_id, 
			created_at 
		FROM 
			chats
		WHERE
			user_id=?
		ORDER BY created_at DESC			
	`
	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var chat Chat
		err = rows.Scan(
			&chat.ID,
			&chat.Chat,
			&chat.UserID,
			&chat.CreatedAt)

		if err != nil {
			log.Fatalln(err)
		}
		chats = append(chats, chat)
	}
	rows.Close()

	return chats, err
}
