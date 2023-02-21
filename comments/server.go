package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/comments/{id}", GetCommentsByPostID).Methods("GET")
	r.HandleFunc("/comments", GetAllComments).Methods("GET", "OPTIONS")
	r.HandleFunc("/comments", AddCommentToPost).Methods("POST")
	r.HandleFunc("/eventbus", RecieveEvent).Methods("POST")
	r.Use(mux.CORSMethodMiddleware(r))
	log.Fatal(http.ListenAndServe(":8081", r))
}
