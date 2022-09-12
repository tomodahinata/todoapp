package controllers

import (
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	genereateHTML(w, "Hello", "layout", "public_navbar", "top")
}
