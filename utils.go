package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
)

func loadPosts() {
	entries, err := os.ReadDir("./static/posts/")
	if err != nil {
		log.Printf("Could not read post folder: %v", err)
		return
	}

	for _, file := range entries {
		fi, err := file.Info()
		if err != nil {
			log.Printf("Error reading file info: %v", err)
			return
		}

		mdContent, err := os.ReadFile(fmt.Sprintf("./static/posts/%s", fi.Name()))
		if err != nil {
			log.Printf("Could not load article: %v", err)
			return
		}

		var buf bytes.Buffer
		if err := md.Convert(mdContent, &buf); err != nil {
			log.Printf("Could not parse markdown: %v", err)
			return
		}
		post := Post{URL: strings.TrimSuffix(file.Name(), ".md"), Name: polishName(file.Name()), Date: fi.ModTime().Format("2006-01-02 15:04:05"), Path: fmt.Sprintf("./static/posts/%s", file.Name()), ArticleContent: template.HTML(buf.String())}
		posts[post.URL] = post
	}
}

func polishName(filename string) string {
	noSpaces := strings.TrimSpace(filename)
	noUnderlines := strings.ReplaceAll(noSpaces, "_", " ")
	firstCapital := capitalizeASCIIString(noUnderlines)
	noExtension := strings.TrimSuffix(firstCapital, ".md")
	return noExtension
}

func capitalizeASCIIString(s string) string {
	if len(s) == 0 {
		return ""
	}
	// Convert first byte to string, uppercase it, add the rest
	return strings.ToUpper(s[:1]) + s[1:]
}
