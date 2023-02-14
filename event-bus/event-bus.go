package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func events(c *gin.Context) {
	http.Post("http://localhost:8080/eventbus", "application/json", c.Request.Body)
	http.Post("http://localhost:8081/eventbus", "application/json", c.Request.Body)
	http.Post("http://localhost:8082/eventbus", "application/json", c.Request.Body)
}
