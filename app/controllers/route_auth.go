package controllers

import "net/http"

func signup(w http.ResponseWriter, r *http.Request) {
	genereateHTML(w, nil, "layout", "public_navbar", "signup")
}
