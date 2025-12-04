package main

import "html/template"

type Post struct {
	URL            string
	Name           string
	Date           string
	Path           string
	ArticleContent template.HTML
}

type ArticleData struct {
	ArticleContent template.HTML
	Date           string
}
