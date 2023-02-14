package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func events(c *gin.Context) {

	body, _ := ioutil.ReadAll(c.Request.Body)

	http.Post("http://localhost:8080/eventbus", "application/json", bytes.NewReader(body))
	http.Post("http://localhost:8081/eventbus", "application/json", bytes.NewReader(body))
	http.Post("http://localhost:8082/eventbus", "application/json", bytes.NewReader(body))

}
