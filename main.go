package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/", getHeaders)

	r.Run(":8080")
}

func getHeaders(c *gin.Context) {
	c.JSON(200, c.Request.Header)
}
