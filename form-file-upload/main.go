package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func FormCreate(w http.ResponseWriter, r *http.Request) {
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

func FormProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if err := r.ParseMultipartForm(2048); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var tmplt = template.Must(template.New("result").ParseFiles("view.htm"))

	var name = r.FormValue("name")

	uploadedFile, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()

	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filename := handler.Filename
	if name != "" {
		filename = fmt.Sprintf("%s%s", name, filepath.Ext(handler.Filename))
	}

	fileLocation := filepath.Join(dir, "files", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]string{"name": name, "message": filename}

	if err := tmplt.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", FormCreate)
	http.HandleFunc("/process", FormProcess)

	fmt.Println("Server started at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
