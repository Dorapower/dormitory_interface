package main

import (
	"dormitory_interface/web"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", web.HelloHandler)
	http.HandleFunc("/hello", web.HelloHandler)
	http.HandleFunc("/login", web.LoginHandler)
	http.HandleFunc("/password", web.PasswordHandler)
	http.HandleFunc("/student", web.StudentInfoHandler)
	http.HandleFunc("/building", web.BuildingListHandler)
	http.HandleFunc("/available", web.AvaliableCountHandler)
	fmt.Println("Listening on port 80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}
