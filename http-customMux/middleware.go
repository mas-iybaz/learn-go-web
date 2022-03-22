package main

import "net/http"

const USERNAME = "emiya"
const PASSWORD = "1password1"

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		if !ok {
			w.Write([]byte(`Something went wrong!`))
			return
		}

		isValid := (username == USERNAME) && (password == PASSWORD)
		if !isValid {
			w.Write([]byte(`Something went wrong!`))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func MiddlewareRequestIsGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.Write([]byte(`Only GET method allowed`))
			return
		}

		next.ServeHTTP(w, r)
	})
}
