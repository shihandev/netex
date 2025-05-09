package main

import (
	"encoding/json"
	"github.com/gorilla/schema"
	"net/http"
)

var users = []user{
	{Name: "alice", Sex: false, Age: 17},
	{Name: "arina", Sex: false, Age: 22},
	{Name: "alex", Sex: true, Age: 20},
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
	var decoder = schema.NewDecoder()
	err := decoder.Decode(&users, r.PostForm)
	if err != nil {
		return
	}
	resp, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	_, err = w.Write([]byte(resp))
	if err != nil {
		return
	}
}
