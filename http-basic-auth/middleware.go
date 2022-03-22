package main

import "net/http"

const USERNAME = "emiya"
const PASSWORD = "1password1"

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()

	if !ok {
		w.Write([]byte(`Something went wrong!`))
		return false
	}

	isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		w.Write([]byte(`Something went wrong!`))
		return false
	}

	return true
}

func RequestIsGet(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte(`Only GET method allowed`))
		return false
	}

	return true
}
