package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}
var newid int = 1
var Users []User

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Users)

	case http.MethodPost:
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err!= nil{
			http.Error(w,"Invalid JSON", http.StatusBadRequest)
			return
		}
		user.Id = newid
		newid++
		Users = append(Users, user)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func main(){
	http.HandleFunc("/api/users", UserHandler)

	fmt.Println("server is running on port 4000...")
	http.ListenAndServe(":4000", nil)
}