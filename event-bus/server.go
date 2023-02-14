package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/event", events)
	r.Run(":8085") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
