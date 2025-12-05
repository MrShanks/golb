package main

import (
	"log"
	"net/http"
	"sort"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if err := tpl.ExecuteTemplate(w, "home.html", posts["how_to_approach_a_programming_project"]); err != nil {
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
