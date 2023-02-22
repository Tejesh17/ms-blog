package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/eventbus", recieveevent)
	r.GET("/posts", getposts)
	r.Run(":8082") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
