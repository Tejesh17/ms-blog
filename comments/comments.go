package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Comment struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

type AllComments struct {
	PostID   int       `json:"postid"`
	Comments []Comment `json:"comments"`
}

type PostComment struct {
	PostID  int    `json:"postid"`
	Comment string `json:"comment"`
}

type EventBusComment struct {
	PostID    int    `json:"postid"`
	Content   string `json:"content"`
	CommentID int    `json:"commentid"`
}

type EventBusBody struct {
	Type string          `json:"type"`
	Body EventBusComment `json:"body"`
}

type GeneralEventBus struct {
	Type string      `json:"type"`
	Body interface{} `json:"body"`
}

var allcomments []AllComments

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

}

func GetCommentsByPostID(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	PostID := vars["id"]
	IntPostID, err := strconv.Atoi(PostID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Post ID sent is not an integer"))
		return
	}
	comments := AllComments{}
	for _, post := range allcomments {
		if post.PostID == IntPostID {
			comments = post
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(comments); err != nil {
		panic(err)
	}
}

func GetAllComments(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(allcomments); err != nil {
		panic(err)
	}
}

func AddCommentToPost(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var comment PostComment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, fmt.Sprintf("error decoding JSON: %v", err), http.StatusBadRequest)
		return
	}
	result, commentid := AddComment(comment.PostID, comment.Comment, &allcomments)

	commentevent := EventBusBody{
		Type: "CommentCreated",
		Body: EventBusComment{
			PostID:    comment.PostID,
			Content:   comment.Comment,
			CommentID: commentid,
		},
	}
	data, err := json.Marshal(commentevent)
	if err != nil {
		log.Fatal(err)
	}
	http.Post("http://localhost:8085/event", "application/json", bytes.NewBuffer(data))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}

func AddComment(postID int, content string, allComments *[]AllComments) (AllComments, int) {
	for i, c := range *allComments {
		if c.PostID == postID {
			newComment := Comment{ID: len(c.Comments) + 1, Content: content}
			(*allComments)[i].Comments = append(c.Comments, newComment)
			return (*allComments)[i], len(c.Comments) + 1
		}
	}
	newComment := Comment{ID: 1, Content: content}
	newPost := AllComments{PostID: postID, Comments: []Comment{newComment}}
	*allComments = append(*allComments, newPost)
	return newPost, 1
}

func RecieveEvent(w http.ResponseWriter, r *http.Request) {
	a, _ := ioutil.ReadAll(r.Body)
	var EventBody GeneralEventBus
	_ = json.Unmarshal(a, &EventBody)
}
