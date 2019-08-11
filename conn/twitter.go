package conn

import (
	"encoding/json"
	"errors"
	"github.com/garyburd/go-oauth/oauth"
	"io"
	"log"
	"net/url"
	"os"
)

var (
	tempCredKey  string
	tokenCredKey string
)

type Account struct {
	ID              string `json:"id_str"`
	ScreenName      string `json:"screen_name"`
	ProfileImageURL string `json:"profile_image_url"`
	Email           string `json:"email"`
}

func init() {
	tempCredKey = os.Getenv("twitter_Consumer_Key")
	tokenCredKey = os.Getenv("twitter_Consumer_Secret")
}

func GetConnect() *oauth.Client {
	return &oauth.Client{
		TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
		ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authorize",
		TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
		Credentials: oauth.Credentials{
			Token:  tempCredKey,
			Secret: tokenCredKey,
		},
	}
}

func GetAccessToken(rt *oauth.Credentials, oauthVerifier string) (*oauth.Credentials, error) {
	at, _, err := GetConnect().RequestToken(nil, rt, oauthVerifier)
	return at, err
}

func GetMe(at *oauth.Credentials, user *Account) error {
	v := url.Values{}
	v.Set("include_email", "true")

	resp, err := GetConnect().Get(nil, at, "https://api.twitter.com/1.1/account/verify_credentials.json", v)
	if err != nil {
		return err
	}
	defer dclose(resp.Body)

	if resp.StatusCode >= 500 {
		return errors.New("twitter is unavailable")
	}

	if resp.StatusCode >= 400 {
		return errors.New("twitter request is invalid")
	}

	return json.NewDecoder(resp.Body).Decode(user)

}

func PostTweet(at *oauth.Credentials) error {
	v := url.Values{}
	v.Set("status", "test post by sample api\npost from：https://github.com/laqiiz/airac")

	resp, err := GetConnect().Post(nil, at, "https://api.twitter.com/1.1/statuses/update.json", v)
	if err != nil {
		return err
	}
	defer dclose(resp.Body)

	return nil
}

func dclose(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}
