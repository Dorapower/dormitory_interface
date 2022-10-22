package web

import (
	"dormitory_interface/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func StudentInfoHandler(w http.ResponseWriter, r *http.Request) {
	if !CheckSession(r) {
		_, err := w.Write([]byte("Please login first"))
		if err != nil {
			panic(err)
		}
		return
	}
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	username := r.Form.Get("username")
	if !testValid(username) {
		_, err := w.Write([]byte("Invalid username"))
		fmt.Println(username)
		if err != nil {
			panic(err)
		}
		return
	}
	data := sql.QueryStudent(username)
	fmt.Println(data)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}

func BuildingListHandler(w http.ResponseWriter, r *http.Request) {
	if !CheckSession(r) {
		_, err := w.Write([]byte("Please login first"))
		if err != nil {
			panic(err)
		}
		return
	}
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	username := r.Form.Get("username")
	if !testValid(username) {
		_, err := w.Write([]byte("Invalid username"))
		fmt.Println(username)
		if err != nil {
			panic(err)
		}
		return
	}
	data := sql.QueryBuildingList(username)
	fmt.Println(data)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}

func AvaliableCountHandler(w http.ResponseWriter, r *http.Request) {
	if !CheckSession(r) {
		_, err := w.Write([]byte("Please login first"))
		if err != nil {
			panic(err)
		}
		return
	}
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	buildingNo := r.Form.Get("building")

	data := sql.QueryAvaliableCount(buildingNo)
	fmt.Println(data)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}
