package models

import (
	"log"
	"time"
)

type Pick struct {
	ID         int
	Group_name string
	User1ID    int
	User2ID    int
	Chats      []Chat
	CreatedAt  time.Time
}

func (u *User) PickTodo(pick string, group_name string) (err error) {
	cmd := `insert into picks(
	group_name,	
	user1_id,
	user2_id,
	created_at) values(?,?,?,?)`

	_, err = Db.Exec(cmd, group_name, pick, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
func GetPick(id int) (pick Pick, err error) {
	cmd := `
		SELECT 
			id,group_name,user1_id,user2_id,created_at 
		FROM 
			picks
		WHERE 
			id=?
			`
	pick = Pick{}

	err = Db.QueryRow(cmd, id).Scan(
		&pick.ID,
		&pick.Group_name,
		&pick.User1ID,
		&pick.User2ID,
		&pick.CreatedAt)

	return pick, err
}

func GetPicks() (picks []Pick, err error) {
	cmd := `SELECT id,group_name,user1_id,user2_id,created_at FROM picks`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var pick Pick
		err = rows.Scan(&pick.ID,
			&pick.Group_name,
			&pick.User1ID,
			&pick.User2ID,
			&pick.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		picks = append(picks, pick)
	}
	rows.Close()

	return picks, err
}
func (u *User) GetPicksByUser() (picks []Pick, err error) {
	cmd := `
		SELECT 
			id, 
			group_name,	
			user1_id,  	
			user2_id, 
			created_at 
		FROM 
			picks
		WHERE
			user1_id=?	
		ORDER BY created_at DESC			
	`
	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var pick Pick
		err = rows.Scan(
			&pick.ID,
			&pick.Group_name,
			&pick.User1ID,
			&pick.User2ID,
			&pick.CreatedAt)

		if err != nil {
			log.Fatalln(err)
		}
		picks = append(picks, pick)
	}
	rows.Close()

	return picks, err
}
