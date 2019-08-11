package conn

import (
	"golang.org/x/oauth2"
	"os"
)

const (
	fbAuthorizeEndpoint = "https://www.facebook.com/dialog/oauth"
	fbTokenEndpoint     = "https://graph.facebook.com/oauth/access_token"
)

func GetFacebookConnect() *oauth2.Config {

	clientID := os.Getenv("FACEBOOK_CLIENT_ID")
	clientSecret := os.Getenv("FACEBOOK_CLIENT_SECRET")

	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  fbAuthorizeEndpoint,
			TokenURL: fbTokenEndpoint,
		},
		Scopes:      []string{"email"},
		RedirectURL: "http://localhost:8000/facebook/callback",
	}
}
