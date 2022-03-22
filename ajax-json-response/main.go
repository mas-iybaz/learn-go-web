package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func jsonData(w http.ResponseWriter, r *http.Request) {
	data := []struct {
		Name string
		Age  int
	}{
		{"Alpha", 22},
		{"Beta", 21},
		{"Charlie", 20},
		{"Delta", 23},
	}

	w.Header().Set("Content-Type", "application/json")

	// Using json.Marshal
	/*
		jsonBytes, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(jsonBytes)
	*/

	// Using json encoder -> json.NewEncoder().Encode()
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/json-data", jsonData)

	fmt.Println("Server started at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
