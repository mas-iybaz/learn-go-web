package main

import (
	"fmt"
	"net/http"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Welcome"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello World"
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)
	http.HandleFunc("/closure", func(w http.ResponseWriter, r *http.Request) {
		text := "Text"
		w.Write([]byte(text))
	})

	var address = "localhost:8080"
	fmt.Printf("server started at %s \n", address)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
}
