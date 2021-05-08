package main

import (
	"encoding/json"
	"net/http"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data interface{}
	var status int
	if r.Method != "GET" {
		data = struct {
			Error string `json:"error"`
		}{Error: "Method not allowed"}
		status = http.StatusMethodNotAllowed
	} else {
		data = []map[string]interface{}{{"id": 1, "name": "Jack"}}
		status = http.StatusOK
	}

	resp, _ := json.Marshal(data)
	w.WriteHeader(status)
	w.Write(resp)
}

func main() {
	http.HandleFunc("/users", getUsers)
	http.ListenAndServe(":8080", nil)
}
