package main

import (
	"fmt"
	"net/http"
	"time"
)

type M map[string]interface{}

var cookieName = "cookieData"

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/delete", handleDelete)

	fmt.Println("Server started at Localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	cookieName := "cookieData"

	c := &http.Cookie{}

	if storedCookie, _ := r.Cookie(cookieName); storedCookie != nil {
		c = storedCookie
	}

	if c.Value == "" {
		c = &http.Cookie{}
		c.Name = cookieName
		c.Value = "Hjs87ytE440OpLytr6729n7gTe23Kli1"
		c.Expires = time.Now().Add(5 * time.Minute)
		http.SetCookie(w, c)
	}

	w.Write([]byte(c.Value))
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{}
	c.Name = cookieName
	c.Expires = time.Unix(0, 0)
	c.MaxAge = -1
	http.SetCookie(w, c)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
