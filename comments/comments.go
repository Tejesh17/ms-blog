package main

import (
	"encoding/json"
	"fmt"
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

var allcomments []AllComments

func GetCommentsByPostID(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(allcomments); err != nil {
		panic(err)
	}
}

func AddCommentToPost(w http.ResponseWriter, r *http.Request) {
	var comment PostComment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, fmt.Sprintf("error decoding JSON: %v", err), http.StatusBadRequest)
		return
	}
	result := AddComment(comment.PostID, comment.Comment, &allcomments)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}

func AddComment(postID int, content string, allComments *[]AllComments) AllComments {
	for i, c := range *allComments {
		if c.PostID == postID {
			newComment := Comment{ID: len(c.Comments) + 1, Content: content}
			(*allComments)[i].Comments = append(c.Comments, newComment)
			return (*allComments)[i]
		}
	}
	newComment := Comment{ID: 1, Content: content}
	newPost := AllComments{PostID: postID, Comments: []Comment{newComment}}
	*allComments = append(*allComments, newPost)
	return newPost
}
