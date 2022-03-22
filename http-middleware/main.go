package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	mux := http.DefaultServeMux

	mux.HandleFunc("/phone", handlePhone)

	var handler http.Handler = mux

	handler = MiddlewareAuth(handler)
	handler = MiddlewareRequestIsGet(handler)

	server := new(http.Server)
	server.Addr = ":8080"
	server.Handler = handler

	fmt.Println("Server started at Localhost:8080")
	server.ListenAndServe()
}

func handlePhone(w http.ResponseWriter, r *http.Request) {
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
