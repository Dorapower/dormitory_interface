package web

import (
	"crypto/sha1"
	"dormitory_interface/sql"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
)

var Sessions = make(map[string]struct{})

func testValid(str string) bool {
	l := len(str)
	if l < 6 || l > 32 {
		return false
	}
	for _, c := range str {
		if !(c >= '0' && c <= '9') && !(c >= 'a' && c <= 'z') && !(c >= 'A' && c <= 'Z') {
			return false
		}
	}
	return true
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	fmt.Printf("username: %s, password: %s\n", username, password)
	if !testValid(username) || !testValid(password) {
		_, err := io.WriteString(w, "Invalid username or password")
		if err != nil {
			panic(err)
		}
		return
	}
	if sql.QueryLogin(username, password) {
		h := sha1.New()
		h.Write([]byte(username + password))
		sha1Hash := hex.EncodeToString(h.Sum(nil))
		cookie := http.Cookie{
			Name:  "session",
			Value: sha1Hash,
		}
		http.SetCookie(w, &cookie)
		Sessions[sha1Hash] = struct{}{}
		_, err := io.WriteString(w, "Login success")
		if err != nil {
			panic(err)
		}
	} else {
		_, err := io.WriteString(w, "Login failed (username or password incorrect)")
		if err != nil {
			panic(err)
		}
	}
}
