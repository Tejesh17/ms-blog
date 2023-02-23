package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

type Posts struct {
	Title string `json:"title"`
	ID    int    `json:"id"`
}

type GeneralEventBus struct {
	Type string      `json:"type"`
	Body interface{} `json:"body"`
}

var posts []Posts

func ReturnPosts(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
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
		num := rand.Intn(99999)
		NewPost := Posts{
			Title: title["title"],
			ID:    num,
		}
		posts = append(posts, NewPost)
		newpost, _ := json.Marshal(NewPost)
		eventbusbody := GeneralEventBus{
			Type: "PostCreated",
			Body: NewPost,
		}
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(eventbusbody)
		if err != nil {
			log.Fatal(err)
		}
		http.Post("http://event-bus-srv:8085/event", "application/json", &buf)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(newpost)
	}
}

func RecieveEvent(w http.ResponseWriter, r *http.Request) {
	a, _ := ioutil.ReadAll(r.Body)
	var EventBody GeneralEventBus
	json.Unmarshal(a, &EventBody)
}
