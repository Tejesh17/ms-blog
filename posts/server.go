package main

import (
	"net/http"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/posts", ReturnPosts)
	mux.HandleFunc("/eventbus", RecieveEvent)
	handler := cors.Default().Handler(mux)

	http.ListenAndServe(":8080", handler)

}
