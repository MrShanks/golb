package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-highlighting/v2"
)

var tpl = template.Must(template.ParseGlob("templates/*.html"))
var posts = make(map[string]Post)
var md = goldmark.New(
	goldmark.WithExtensions(
		highlighting.NewHighlighting(
			highlighting.WithStyle("dracula"))))

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	loadPosts()

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/blog", BlogHanlder)
	http.HandleFunc("/post/{slug}", PostHandler)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      nil, // uses DefaultServeMux
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Println("Server running on http://localhost:8080")
	log.Fatal(srv.ListenAndServe())
}
