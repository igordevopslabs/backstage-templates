package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	defaultPort    = "{{ values.port | lower }}"
	defaultEnvName = "{{ values.apiEnv | lower }}"
)

type fixedResponse string

func (s fixedResponse) ServeHTTP(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, s) }

func main() {
	http.Handle("/healthcheck", fixedResponse("~> {{ values.apiName | lower }} is healthy..."))
	http.Handle("/", fixedResponse(fmt.Sprintf("It's live!! env: %s\n", getEnv("ENV_NAME", defaultEnvName))))

	addr := fmt.Sprintf(":%s", getEnv("PORT", defaultPort))

	log.Printf("listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
