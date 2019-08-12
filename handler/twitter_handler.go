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

func (h *twitterHandler) AuthRedirect(w http.ResponseWriter, r *http.Request) {
	con := conn.GetTwitterConnect()

	rt, err := con.RequestTemporaryCredentials(nil, "http://localhost:8000/twitter/callback", nil)
	if err != nil {
		panic(err)
	}

	h.session.Put(r.Context(), "request_token", rt.Token)
	h.session.Put(r.Context(), "request_token_secret", rt.Secret)

	http.Redirect(w, r, con.AuthorizationURL(rt, nil), 302)
}

func (h *twitterHandler) Callback(w http.ResponseWriter, r *http.Request) {
	request := TwitterCallbackRequest{
		Token:    r.URL.Query().Get("oauth_token"),
		Verifier: r.URL.Query().Get("oauth_verifier"),
	}

	at, err := conn.GetAccessToken(
		&oauth.Credentials{
			Token:  h.session.Get(r.Context(), "request_token").(string),
			Secret: h.session.Get(r.Context(), "request_token_secret").(string),
		},
		request.Verifier,
	)
	if err != nil {
		panic(err)
	}

	h.session.Put(r.Context(), "oauth_secret", at.Secret)
	h.session.Put(r.Context(), "oauth_token", at.Token)

	var a conn.Account
	if err = conn.GetMe(at, &a); err != nil {
		panic(err)
	}

	data := map[string]interface{}{}
	data["ID"] = a.ID
	data["ProfileImageURL"] = a.ProfileImageURL
	data["ScreenName"] = a.ScreenName
	data["Email"] = a.Email

	tpl := template.Must(template.ParseFiles("view/twitter/callback.tpl"))
	if err := tpl.Execute(w, data); err != nil {
		_, _ = w.Write([]byte(err.Error()))
	}

}
