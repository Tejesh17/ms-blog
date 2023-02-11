package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/posts", ReturnPosts)
	log.Fatal((http.ListenAndServe(":8080", nil)))
}
