package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"os"
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

	if err := tpl.ExecuteTemplate(w, "blog.html", posts); err != nil {
		log.Printf("Could not execute blog template: %v\n", err)
	}
}
