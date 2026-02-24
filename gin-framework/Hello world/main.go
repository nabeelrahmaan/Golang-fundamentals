package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Greet struct {
	message string `json:"message"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		resp := Greet{
			message: "Hello world.",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func main() {
	http.HandleFunc("/api/hello", Handler)

	fmt.Println("server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
