package models

import (
	"log"
	"time"
)

type Todo struct {
	ID      int
	Content string
	Title   string
	// 期日更新
	Deadline  string
	Time      int
	UserID    int
	CreatedAt time.Time
}

func (u *User) CreateTodo(content string, title string, deadline string) (err error) {
	cmd := `insert into todos(
	content,
	title,
	deadline,
	user_id,
	created_at) values(?,?,?,?,?)`

	_, err = Db.Exec(cmd, content, title, u.ID, deadline, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetTodo(id int) (todo Todo, err error) {
	cmd := `SELECT id,content,user_id,created_at FROM todos
	WHERE id=?`
	todo = Todo{}

	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt)

	return todo, err
}

func GetTodos() (todos []Todo, err error) {
	cmd := `SELECT id,content,user_id,created_at FROM todos`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID,
			&todo.Content,
			&todo.Title,
			// 期日更新
			&todo.Deadline,
			&todo.UserID,
			&todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}
func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `
		SELECT 
			id, content, title, deadline, user_id, created_at 
		FROM 
			todos
		WHERE 
			user_id=?
	`
	log.Println(u.ID)
	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.Title,
			&todo.Deadline,
			&todo.UserID,
			&todo.CreatedAt)

		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

func (t *Todo) UpdateTodo() error {
	cmd := `update todos set content =?,user_id=?,title =?
	WHERE id=?`
	_, err = Db.Exec(cmd, t.Content, t.UserID, t.Title, t.ID /*title更新*/)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (t *Todo) DeleteTodo() error {
	cmd := `delete FROM todos WHERE id=?`
	_, err = Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
