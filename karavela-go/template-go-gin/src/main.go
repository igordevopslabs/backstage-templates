package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

var log = logrus.New()

func init() {
	log.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "~> ${{ values.apiName | lower }} It's live.")
	})

	r.GET("/healthcheck", func(c *gin.Context) {
		c.String(http.StatusOK, "~> ${{ values.apiName | lower }} is healthy.")
	})

	r.Run(":${{ values.port | lower }}")
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		duration := time.Since(start)

		log.WithFields(logrus.Fields{
			"method":      r.Method,
			"url":         r.URL.Path,
			"remote_addr": r.RemoteAddr,
			"duration":    duration,
			"statusCode":  rw.statusCode,
		}).Info("Request completed")
	})
}
