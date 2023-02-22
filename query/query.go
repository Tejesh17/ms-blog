package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type GeneralEventBus struct {
	Type string      `json:"type"`
	Body interface{} `json:"body"`
}

type PostEvent struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type CommentEvent struct {
	PostID    int    `json:"postid"`
	Content   string `json:"content"`
	CommentID int    `json:"commentid"`
}

type PostsWithComments struct {
	ID       int            `json:"id"`
	Title    string         `json:"title"`
	Comments []CommentEvent `json:"comments"`
}

var allposts = make(map[int]PostsWithComments)

func recieveevent(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	var decodedEvent GeneralEventBus
	json.Unmarshal(body, &decodedEvent)
	if decodedEvent.Type == "PostCreated" {
		post := decodedEvent.Body.(map[string]interface{})
		newpost := PostsWithComments{
			ID:       int(post["id"].(float64)),
			Title:    post["title"].(string),
			Comments: []CommentEvent{},
		}
		allposts[int(post["id"].(float64))] = newpost
	}
	if decodedEvent.Type == "CommentCreated" {
		comment := decodedEvent.Body.(map[string]interface{})
		newComment := CommentEvent{
			CommentID: int(comment["commentid"].(float64)),
			Content:   comment["content"].(string),
			PostID:    int(comment["postid"].(float64)),
		}
		post := allposts[newComment.PostID]
		post.Comments = append(post.Comments, newComment)
		allposts[newComment.PostID] = post
	}
	fmt.Println(allposts)
}

func getposts(c *gin.Context) {
	c.JSON(200, allposts)
}
