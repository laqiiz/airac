package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/gorilla/mux"
	"github.com/laqiiz/airac/handler"
	"github.com/laqiiz/airac/middleware"
	"github.com/laqiiz/airac/repository"
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

	// gorilla router
	r := mux.NewRouter()

	// index
	index := handler.IndexHandler{}
	r.HandleFunc("/", middleware.Entry(index.Index)).Methods(http.MethodGet)

	// HealthCheck
	r.HandleFunc("/health", middleware.Entry(handler.Health)).Methods(http.MethodGet)

	// Google
	google := handler.GoogleHandler{}
	r.HandleFunc("/google/oauth2", middleware.Entry(google.AuthRedirect)).Methods(http.MethodGet)
	r.HandleFunc("/google/callback", middleware.Entry(google.Callback)).Methods(http.MethodGet)

	// Twitter
	twitterHandler := handler.NewTwitterHandler(sessionManager)
	r.HandleFunc("/twitter/oauth", twitterHandler.AuthRedirect).Methods(http.MethodGet)
	r.HandleFunc("/twitter/callback", twitterHandler.Callback).Methods(http.MethodGet)

	// Facebook
	fb := handler.FacebookHandler{}
	r.HandleFunc("/facebook/oauth2", middleware.Entry(fb.AuthRedirect)).Methods(http.MethodGet)
	r.HandleFunc("/facebook/callback", middleware.Entry(fb.Callback)).Methods(http.MethodGet)

	// GitHub
	github := handler.GitHubHandler{}
	r.HandleFunc("/github/oauth2", github.AuthRedirect).Methods(http.MethodGet)
	r.HandleFunc("/github/callback", github.Callback).Methods(http.MethodGet)

	// SignUp,SignIn,SignOut,DeleteAccount
	ur := repository.NewMemUserRepository()
	signupHandler := handler.NewSignupHandler(sessionManager, ur)
	r.HandleFunc("/signup", signupHandler.SignUp).Methods(http.MethodPost)

	log.Println("airac start in :" + port)

	log.Fatal(http.ListenAndServe(":"+port, sessionManager.LoadAndSave(r)))
}
