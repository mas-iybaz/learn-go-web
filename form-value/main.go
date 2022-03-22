package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func messageCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var tmplt = template.Must(template.New("form").ParseFiles("view.htm"))
		var err = tmplt.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func messageProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tmplt = template.Must(template.New("result").ParseFiles("view.htm"))

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var name = r.FormValue("name")
		var msg = r.FormValue("message")
		var data = map[string]string{"name": name, "message": msg}

		if err := tmplt.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/", messageCreate)
	http.HandleFunc("/process", messageProcess)

	fmt.Println("Server started at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
