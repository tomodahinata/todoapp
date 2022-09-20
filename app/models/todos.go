package models

import (
	"log"
	"time"
)

type Todo struct {
	ID        int
	Content   string
	Title     string
	Category  string
	Deadline  string
	Time      int
	UserID    int
	CreatedAt time.Time
}

func (u *User) CreateTodo(content string, title string, deadline string, category string) (err error) {
	cmd := `insert into todos(
	content,
	title,
	category,
	deadline,
	user_id,
	created_at) values(?,?,?,?,?,?)`

	_, err = Db.Exec(cmd, content, title, category, deadline, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetTodo(id int) (todo Todo, err error) {
	cmd := `
		SELECT 
			id,content,user_id,created_at 
		FROM 
			todos
		WHERE 
			id=?
			`
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
			&todo.Category,
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
			id, 
			content, 
			title,
			category, 
			deadline, 
			user_id, 
			created_at 
		FROM 
			todos
		WHERE
			user_id=?
		ORDER BY created_at DESC			
	`
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
			&todo.Category,
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
	cmd := `
			update todos set content =?,user_id=?,title =?,category=?,deadline=?
		WHERE 
			id=?
			`
	_, err = Db.Exec(cmd, t.Content, t.UserID, t.Title, t.Category, t.Deadline, t.ID /*title更新*/)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (t *Todo) DeleteTodo() error {
	cmd := `
			delete 
		FROM 
			todos 
		WHERE 
			id=?
	`
	_, err = Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
func (u *User) GetTodosSort(choice string) (todos []Todo, err error) {

	var cmd string
	if choice == "id昇順" {
		cmd = `
		SELECT 
			id, 
			content, 
			title,
			category, 
			deadline, 
			user_id, 
			created_at 
		FROM 
			todos
		WHERE
			user_id=?
		ORDER BY created_at ASC
			
	`
	} else if choice == "id降順" {
		cmd = `
		SELECT 
			id, 
			content, 
			title,
			category, 
			deadline, 
			user_id, 
			created_at 
		FROM 
			todos
		WHERE
			user_id=?
		ORDER BY created_at DESC
			
	`
	} else if choice == "期限昇順" {
		cmd = `
		SELECT 
			id, 
			content, 
			title,
			category, 
			deadline, 
			user_id, 
			created_at 
		FROM 
			todos
		WHERE
			user_id=?
		ORDER BY deadline ASC
			
	`
	} else if choice == "期限降順" {
		cmd = `
		SELECT 
			id, 
			content, 
			title,
			category, 
			deadline, 
			user_id, 
			created_at 
		FROM 
			todos
		WHERE
			user_id=?
		ORDER BY deadline DESC
			
	`

	}

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
			&todo.Category,
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
