package handler

import (
	"github.com/google/go-github/github"
	"github.com/laqiiz/airac/conn"
	"golang.org/x/oauth2"
	"html/template"
	"net/http"
	"strconv"
)

type GitHubHandler struct {
}

type CallbackRequest struct {
	Code  string
	State int
}

func (h *GitHubHandler) AuthRedirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, conn.GetGitHubConnect().AuthCodeURL(""), 302)
}

func (h *GitHubHandler) Callback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	state, err := strconv.Atoi(r.URL.Query().Get("state"))
	if err != nil {
		state = 0 // TODO
	}

	request := CallbackRequest{
		Code:  r.URL.Query().Get("code"),
		State: state,
	}

	tok, err := conn.GetGitHubConnect().Exchange(ctx, request.Code)
	if err != nil {
		panic(err)
	}

	tc := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: tok.AccessToken},
	))

	user, _, err := github.NewClient(tc).Users.Get(ctx, "")
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{}
	data["ID"] = user.GetID()
	data["AvatarURL"] = user.GetAvatarURL()
	data["Name"] = user.GetName()

	tpl := template.Must(template.ParseFiles("view/github/callback.tpl"))

	if err := tpl.Execute(w, data); err != nil {
		_, _ = w.Write([]byte(err.Error()))
	}

}
