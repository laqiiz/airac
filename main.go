package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/laqiiz/airac/handler"
	"github.com/laqiiz/airac/middleware"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
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

	// Session
	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour // TODO Magic Number

	mux := http.NewServeMux()

	// index
	index := handler.IndexHandler{}
	mux.HandleFunc("/", middleware.Entry(index.Index))

	// HealthCheck
	mux.HandleFunc("/health", middleware.Entry(handler.Health))

	// Google
	google := handler.GoogleHandler{}
	mux.HandleFunc("/google/oauth2", middleware.Entry(google.Redirect))
	mux.HandleFunc("/google/callback", middleware.Entry(google.GetCallback))

	// Twitter
	twitterHandler := handler.NewTwitterHandler(sessionManager)
	mux.HandleFunc("/twitter/oauth", twitterHandler.Redirect)
	mux.HandleFunc("/twitter/callback", twitterHandler.GetCallback)

	// Facebook
	fb := handler.FacebookHandler{}
	mux.HandleFunc("/facebook/oauth2", middleware.Entry(fb.Redirect))
	mux.HandleFunc("/facebook/callback", middleware.Entry(fb.GetCallback))

	// GitHub
	github := handler.GitHubHandler{}
	mux.HandleFunc("/github/oauth2", github.Redirect)
	mux.HandleFunc("/github/callback", github.GetCallback)

	log.Println("airac start in :" + port)

	log.Fatal(http.ListenAndServe(":"+port, sessionManager.LoadAndSave(mux)))
}
