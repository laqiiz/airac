package main

import (
	"fmt"
	"github.com/laqiiz/airac/handler"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	log.Println("airac start")

	// set the port this service will be run on
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000" // default port
	} else {
		_, err := strconv.Atoi(port)
		if err != nil {
			log.Fatal("env PORT is invalid")
		}
	}

	// index
	index := handler.IndexHandler{}
	http.HandleFunc("/", index.Index)

	// HealthCheck
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "ok")
	})

	// Google
	http.HandleFunc("/google/oauth2", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello World")
	})
	http.HandleFunc("/google/callback", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello World")
	})

	// Twitter
	http.HandleFunc("/twitter/oauth", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello World")
	})
	http.HandleFunc("/twitter/callback", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello World")
	})
	http.HandleFunc("/twitter/post", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello World")
	})

	// Facebook
	http.HandleFunc("/facebook/oauth2", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello World")
	})
	http.HandleFunc("/facebook/callback", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello World")
	})

	// GitHub
	http.HandleFunc("/github/oauth2", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello World")
	})
	http.HandleFunc("/github/callback", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello World")
	})

	log.Println("airac start in :" + port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
