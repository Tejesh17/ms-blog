package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

type Posts struct {
	Title string `json:"title"`
	ID    int    `json:"id"`
}

type EventBusBody struct {
	Type string                 `json:"type"`
	Body map[string]interface{} `json:"body"`
}

var posts []Posts

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func ReturnPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	enableCors(&w)
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

func RecieveEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
	var EventBody EventBusBody
	if err := json.NewDecoder(r.Body).Decode(&EventBody); err != nil {
		http.Error(w, fmt.Sprintf("error decoding JSON: %v", err), http.StatusBadRequest)
		return
	}
	fmt.Println(EventBody)
}
