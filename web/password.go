package web

import (
	"dormitory_interface/sql"
	"io"
	"net/http"
)

func changePassword(username string, password string) bool {
	if !testValid(password) {
		return false
	}
	sql.UpdatePassword(username, password)
	return true
}

func PasswordHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	if !CheckSession(r) {
		_, err := io.WriteString(w, "Please login first")
		if err != nil {
			panic(err)
		}
		return
	}
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	if changePassword(username, password) {
		_, err := io.WriteString(w, "Change password success")
		if err != nil {
			panic(err)
		}
	} else {
		_, err := io.WriteString(w, "Change password failed")
		if err != nil {
			panic(err)
		}
	}
}
