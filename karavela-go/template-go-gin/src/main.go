package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "~> ${{ values.apiName | lower }} It's live.")
	})

	r.GET("/healthcheck", func(c *gin.Context) {
		c.String(http.StatusOK, "~> ${{ values.apiName | lower }} is healthy.")
	})

	//r.Run(":${{ values.port | lower }}")
	r.Run(":3000")
}
