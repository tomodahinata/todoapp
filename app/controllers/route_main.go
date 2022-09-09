package controllers

import (
	"fmt"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	fmt.Println("テスト")
	genereateHTML(w, "Hello", "layout", "top")
}
