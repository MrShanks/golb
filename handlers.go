package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	mdContent, err := os.ReadFile(homePageArticle)
	if err != nil {
		http.Error(w, "Could not load article", http.StatusInternalServerError)
		log.Printf("Could not load article: %v", err)
		return
	}

	fileInfo, err := os.Stat(homePageArticle)
	if err != nil {
		http.Error(w, "Error reading file's stats", http.StatusInternalServerError)
		log.Printf("Could not read file stats: %v", err)
		return
	}

	var buf bytes.Buffer
	if err := md.Convert(mdContent, &buf); err != nil {
		http.Error(w, "Could not parse markdown", http.StatusInternalServerError)
		log.Printf("Could not parse markdown: %v", err)
		return
	}

	data := ArticleData{
		ArticleContent: template.HTML(buf.String()),
		Date:           fileInfo.ModTime().Format("2006-01-02 15:04:05"),
	}

	if err := tpl.ExecuteTemplate(w, "home.html", data); err != nil {
		http.Error(w, "Could not execute home template", http.StatusInternalServerError)
		return
	}
}

func BlogHanlder(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/blog" {
		http.NotFound(w, r)
		return
	}

	ordered := []Post{}

	for _, p := range posts {
		ordered = append(ordered, p)
	}

	sort.Slice(ordered, func(i, j int) bool {
		return ordered[i].Name < ordered[j].Name
	})

	if err := tpl.ExecuteTemplate(w, "blog.html", ordered); err != nil {
		log.Printf("Could not execute blog template: %v\n", err)
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	articleSlug := r.PathValue("slug")
	post, ok := posts[articleSlug]
	if !ok {
		http.NotFound(w, r)
		return
	}

	if err := tpl.ExecuteTemplate(w, "post.html", post); err != nil {
		log.Printf("Could not execute blog template: %v\n", err)
	}
}
