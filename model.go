package main

import "html/template"

type Post struct {
	Name string
	Date string
}

type ArticleData struct {
	ArticleContent template.HTML
	Date           string
}
