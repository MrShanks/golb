package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
