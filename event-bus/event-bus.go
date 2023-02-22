package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GeneralEventBus struct {
	Type string      `json:"type"`
	Body interface{} `json:"body"`
}

func events(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	fmt.Println(bytes.NewBuffer(body))
	var eventbody GeneralEventBus
	json.Unmarshal(body, &eventbody)
	http.Post("http://posts-srv:8080/eventbus", "application/json", bytes.NewBuffer(body))
	http.Post("http://comments-srv:8081/eventbus", "application/json", bytes.NewBuffer(body))
	http.Post("http://query-srv:8082/eventbus", "application/json", bytes.NewBuffer(body))
}
