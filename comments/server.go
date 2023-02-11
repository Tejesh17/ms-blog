package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/comments/{id}", GetCommentsByPostID).Methods("GET")
	r.HandleFunc("/comments", GetAllComments).Methods("GET")
	r.HandleFunc("/comments", AddCommentToPost).Methods("POST")
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8081", r))

}
