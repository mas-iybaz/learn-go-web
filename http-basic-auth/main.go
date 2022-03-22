package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/phone", handlePhone)

	server := new(http.Server)
	server.Addr = ":8080"

	fmt.Println("Server started at Localhost:8080")
	server.ListenAndServe()
}

func handlePhone(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r) {
		return
	}
	if !RequestIsGet(w, r) {
		return
	}

	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectPhone(id))
		return
	}

	OutputJSON(w, GetPhones())
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
