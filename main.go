package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var tpl = template.Must(template.ParseGlob("templates/*.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if err := tpl.Execute(w, "home.html"); err != nil {
		fmt.Println(err)
	}
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", HomeHandler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
