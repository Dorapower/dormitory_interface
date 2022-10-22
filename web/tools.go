package web

import (
	"fmt"
	"net/http"
)

func CheckSession(r *http.Request) bool {
	cookie, err := r.Cookie("session")
	if err == http.ErrNoCookie {
		return false
	} else if err != nil {
		panic(err)
	}
	session := cookie.Value
	if _, ok := Sessions[session]; !ok {
		fmt.Println("Session not found, " + session)
		for k := range Sessions {
			fmt.Println(k)
		}
		return false
	}
	return true
}
