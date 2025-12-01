package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
)

var tpl = template.Must(template.ParseGlob("templates/*.html"))

var posts = []Post{}

type Post struct {
	Name string
	Date string
}

type ArticleData struct {
	ArticleContent template.HTML
	Date           string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	pathNotFound(w, r, "/")

	fileName := "./static/posts/how_to_approach_a_programming_project.md"

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

func BlogHanlder(w http.ResponseWriter, r *http.Request) {
	pathNotFound(w, r, "/blog")

	if err := tpl.ExecuteTemplate(w, "blog.html", posts); err != nil {
		fmt.Printf("Could not execute blog template: %v\n", err)
	}
}

func pathNotFound(w http.ResponseWriter, r *http.Request, path string) {
	if r.URL.Path != path {
		http.NotFound(w, r)
		fmt.Printf("Page not found: %v", r.URL.Path)
		return
	}
}

func readPosts() {
	entries, err := os.ReadDir("./static/posts/")
	if err != nil {
		fmt.Printf("Could not read post folder: %v", err)
		return
	}

	for _, file := range entries {
		fi, err := file.Info()
		if err != nil {
			fmt.Printf("Error reading file info: %v", err)
			return
		}

		posts = append(posts, Post{polishName(file.Name()), fi.ModTime().Format("2006-01-02 15:04:05")})
	}
}

func polishName(filename string) string {
	noSpaces := strings.TrimSpace(filename)
	noUnderlines := strings.ReplaceAll(noSpaces, "_", " ")
	firstCapital := CapitalizeASCIIString(noUnderlines)
	noExtension := strings.TrimSuffix(firstCapital, ".md")
	return noExtension
}

func CapitalizeASCIIString(s string) string {
	if len(s) == 0 {
		return ""
	}
	// Convert first byte to string, uppercase it, add the rest
	return strings.ToUpper(s[:1]) + s[1:]
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	readPosts()

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/blog", BlogHanlder)
	http.HandleFunc("/.well-known/appspecific/com.chrome.devtools.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
