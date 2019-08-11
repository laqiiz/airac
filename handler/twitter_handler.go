package handler

import (
	"github.com/alexedwards/scs/v2"
	"github.com/garyburd/go-oauth/oauth"
	"github.com/laqiiz/airac/conn"
	"html/template"
	"net/http"
)

type TwitterCallbackRequest struct {
	Token    string
	Verifier string
}

type twitterHandler struct {
	session *scs.SessionManager
}

func NewTwitterHandler(session *scs.SessionManager) *twitterHandler {
	return &twitterHandler{
		session: session,
	}
}

func (c *twitterHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	con := conn.GetTwitterConnect()

	// ここのURLはTwitterのDeveloperツールで設定するURLに一致する必要があるので、ポートが固定になっちゃう
	rt, err := con.RequestTemporaryCredentials(nil, "http://localhost:8000/twitter/callback", nil)
	if err != nil {
		panic(err)
	}

	c.session.Put(r.Context(), "request_token", rt.Token)
	c.session.Put(r.Context(), "request_token_secret", rt.Secret)

	http.Redirect(w, r, con.AuthorizationURL(rt, nil), 302)
}

func (c *twitterHandler) GetCallback(w http.ResponseWriter, r *http.Request) {
	request := TwitterCallbackRequest{
		Token:    r.URL.Query().Get("oauth_token"),
		Verifier: r.URL.Query().Get("oauth_verifier"),
	}

	at, err := conn.GetAccessToken(
		&oauth.Credentials{
			Token:  c.session.Get(r.Context(), "request_token").(string),
			Secret: c.session.Get(r.Context(), "request_token_secret").(string),
		},
		request.Verifier,
	)
	if err != nil {
		panic(err)
	}

	c.session.Put(r.Context(), "oauth_secret", at.Secret)
	c.session.Put(r.Context(), "oauth_token", at.Token)

	var account conn.Account
	if err = conn.GetMe(at, &account); err != nil {
		panic(err)
	}

	data := map[string]interface{}{}
	data["ID"] = account.ID
	data["ProfileImageURL"] = account.ProfileImageURL
	data["ScreenName"] = account.ScreenName
	data["Email"] = account.Email

	tpl := template.Must(template.ParseFiles("view/twitter/callback.tpl"))
	if err := tpl.Execute(w, data); err != nil {
		_, _ = w.Write([]byte(err.Error()))
	}

}

func (c *twitterHandler) Tweet(w http.ResponseWriter, r *http.Request) {
	at := oauth.Credentials{
		Secret: c.session.Get(r.Context(), "oauth_secret").(string),
		Token:  c.session.Get(r.Context(), "oauth_token").(string),
	}

	if err := conn.PostTweet(&at); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/?message=投稿しました", 302)
}
