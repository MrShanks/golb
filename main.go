package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
)

var tpl = template.Must(template.ParseGlob("templates/*.html"))

type ArticleData struct {
	ArticleContent template.HTML
	Date           string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	fileName := "./static/posts/go_projects_to_try_out.md"

	mdContent, err := os.ReadFile(fileName)
	if err != nil {
		http.Error(w, "Could not load article", http.StatusInternalServerError)
		fmt.Printf("Error reading file: %v", err)
		return
	}

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Printf("Error reading file's stats: %v", err)
	}

	md := goldmark.New(
		goldmark.WithExtensions(
			highlighting.NewHighlighting(
				highlighting.WithStyle("dracula"))))

	var buf bytes.Buffer
	if err := md.Convert(mdContent, &buf); err != nil {
		http.Error(w, "Could not parse markdown", http.StatusInternalServerError)
		fmt.Printf("Error parsing markdown: %v", err)
		return
	}

	data := ArticleData{
		ArticleContent: template.HTML(buf.String()),
		Date:           fileInfo.ModTime().Format("2006-01-02 15:04:05"),
	}

	if err := tpl.ExecuteTemplate(w, "home.html", data); err != nil {
		fmt.Printf("Could not execute home template: %v\n", err)
	}
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", HomeHandler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
