package main

import (
	"fmt"
	"github.com/laqiiz/airac/handler"
	"github.com/laqiiz/airac/middleware"
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
	http.HandleFunc("/", middleware.Entry(index.Index))

	// HealthCheck
	http.HandleFunc("/health",middleware.Entry(handler.Health))

	// Google
	google := handler.GoogleOAuthController{}
	http.HandleFunc("/google/oauth2", middleware.Entry(google.Redirect))
	http.HandleFunc("/google/callback", middleware.Entry(google.GetCallback))

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
	fb := handler.FacebookOauthHandler{}
	http.HandleFunc("/facebook/oauth2", middleware.Entry(fb.Redirect))
	http.HandleFunc("/facebook/callback", middleware.Entry(fb.GetCallback))

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
