package main

import (
	"encoding/json"
	"github.com/gorilla/schema"
	"net/http"
)

var users = []user{
	{Name: "alice", Sex: false, Age: 17},
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appliation/json")
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func updateUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var temp user
	decoder := schema.NewDecoder()
	if err := decoder.Decode(&temp, r.Form); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	users = append(users, temp)
}
