package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ViewData map[string]interface{}

type Contact struct {
	Phone string
	Email string
}

type Student struct {
	Name    string
	Age     int
	Hobbies []string
	Contact Contact
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]interface{}{
			"main": ViewData{
				"title": "GO WEB",
				"name":  "Haidar Aziz Habibulloh",
			},
			"student": Student{
				Name:    "Aziz",
				Age:     10,
				Hobbies: []string{"Football", "Drawing", "Cycling"},
				Contact: Contact{
					Phone: "089518566669",
					Email: "azizzz@gmail.com",
				},
			},
		}

		var tmplt = template.Must(template.ParseFiles(
			"views/pages/index.htm",
			"views/components/header.htm",
		))

		var err = tmplt.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	fmt.Println("Server started at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
