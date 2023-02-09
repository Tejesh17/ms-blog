package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
)

type Posts struct {
	Title string `json:"title"`
	ID    int    `json:"id"`
}

var posts []Posts

func ReturnPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == "GET" {
		PostsJson, err := json.Marshal(posts)
		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(PostsJson)
		}
		return
	}
	if r.Method == "POST" {
		fmt.Println("POST")
		body, _ := io.ReadAll((r.Body))

		var title map[string]string
		json.Unmarshal(body, &title)
		fmt.Println(title["title"])
		num := rand.Intn(99999)
		NewPost := Posts{
			Title: title["title"],
			ID:    num,
		}
		posts = append(posts, NewPost)
		newpost, _ := json.Marshal(NewPost)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(newpost)
	}
}

func main() {
	http.HandleFunc("/posts", ReturnPosts)
	log.Fatal((http.ListenAndServe(":8080", nil)))
}
