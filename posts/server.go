package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/posts", ReturnPosts)
	http.HandleFunc("/eventbus", RecieveEvent)
	log.Fatal((http.ListenAndServe(":8080", nil)))
}
