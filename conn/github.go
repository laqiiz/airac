package conn

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"os"
)

func GetGitHubConnect() *oauth2.Config {

	clientID := os.Getenv("GITHUB_CLIENT_ID")
	clientSecret := os.Getenv("GITHUB_CLIENT_SECRET")

	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     github.Endpoint,
		Scopes:       []string{"user:email", "repo"},
		RedirectURL:  "http://localhost:8000/github/callback",
	}

}
