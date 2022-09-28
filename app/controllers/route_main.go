package controllers

import (
	"log"
	"net/http"
	"to-do-app/app/models"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		genereateHTML(w, "Hello", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}
func index(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		todos, _ := user.GetTodosByUser()
		user.Todos = todos
		genereateHTML(w, user, "layout", "private_navbar", "index")
	}
}
func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		genereateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}

func todoSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		title := r.PostFormValue("title")
		deadline := r.PostFormValue("deadline")
		category := r.PostFormValue("category")
		if err := user.CreateTodo(content, title, deadline, category); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/todos", http.StatusFound)

	}
}

func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		t, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}
		genereateHTML(w, t, "layout", "private_navbar", "todo_edit")
	}
}
func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		title := r.PostFormValue("title")
		category := r.PostFormValue("category")
		deadline := r.PostFormValue("deadline")
		t := &models.Todo{ID: id, Content: content, UserID: user.ID, Deadline: deadline, Title: title, Category: category}
		if err := t.UpdateTodo(); err != nil {
			log.Println(err)
		}
		// タイトル更新

		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}
func todoDelete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		t, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}
		if err := t.DeleteTodo(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}

func todoSort(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {

		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		choice := r.PostFormValue("choice")
		todos, err := user.GetTodosSort(choice)
		if err != nil {
			log.Println(err)
		}

		user.Todos = todos
		genereateHTML(w, user, "layout", "private_navbar", "index")
	}
}

func todoChatNew(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		chats, _ := user.GetChatsByUser()
		user.Chats = chats
		genereateHTML(w, user, "layout", "private_navbar", "chat_new")
	}
}

func todoChatSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		chat := r.PostFormValue("chat")
		if err := user.ChatTodo(chat); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos/chat_new", http.StatusFound)
	}
}

func todoChatPick(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		picks, _ := user.GetPicksByUser()
		user.Picks = picks
		genereateHTML(w, user, "layout", "private_navbar", "chat_pick")
	}
}

func todoChatCreate(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		pick := r.PostFormValue("pick")
		group_name := r.PostFormValue("group_name")
		if err := user.PickTodo(pick, group_name); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos/chat_grouplist", http.StatusFound)
	}
}

func todoChatGroupList(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		picks, _ := user.GetPicksByUser()
		user.Picks = picks
		genereateHTML(w, user, "layout", "private_navbar", "chat_group")
	}
}

func todoChatGroup(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		p, err := models.GetPick(id)
		if err != nil {
			log.Println(err)
		}
		log.Println(id)
		// パスに含まれたID のグループとチャット一覧を取得
		genereateHTML(w, p, "layout", "private_navbar", "chat_group")
	}
}
func todoChatGroupBack(w http.ResponseWriter, r *http.Request, id int) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/todos/chat_group", http.StatusFound)
	}
}
