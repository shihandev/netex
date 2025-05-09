package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type user struct {
	Name string `json:"Name"`
	Sex  bool   `json:"Sex"`
	Age  uint   `json:"Age"`
}

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
	}
}

func updateUsers(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	_, err = w.Write([]byte(fmt.Sprintf("hello, %s", string(body))))
	if err != nil {
		return
	}
}
