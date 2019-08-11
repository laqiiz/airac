package handler

import (
	"errors"
	"github.com/huandu/facebook"
	"github.com/laqiiz/airac/conn"
	"html/template"
	"net/http"
	"strconv"
)

type FacebookOauth2Handler struct {
}

func (c *FacebookOauth2Handler) Get(w http.ResponseWriter, r *http.Request) {
	url := conn.GetFacebookConnect().AuthCodeURL("")
	http.Redirect(w, r, url, 302)
}

type CallbackRequest struct {
	Code  string `form:"code"`
	State int    `form:"state"`
}

func (c *FacebookOauth2Handler) GetCallback(w http.ResponseWriter, r *http.Request) {
	state, err := strconv.Atoi(r.Header.Get("state"))
	if err != nil {
		panic(err)
	}
	var request = CallbackRequest{
		Code:  r.Header.Get("code"),
		State: state,
	}

	config := conn.GetFacebookConnect()

	tok, err := config.Exchange(r.Context(), request.Code)
	if err != nil {
		panic(err)
	}

	if tok.Valid() == false {
		panic(errors.New("valid token"))
	}

	session := &facebook.Session{
		Version:    "v2.8",
		HttpClient: config.Client(r.Context(), tok),
	}

	res, err := session.Get("/me?fields=id,name,email", nil)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{}
	data["ID"] = res["id"]
	data["Name"] = res["name"]
	data["Email"] = res["email"]

	tpl := template.Must(template.ParseFiles("view/facebook/callback.tpl"))

	if err := tpl.Execute(w, map[string]string{}); err != nil {
		_, _ = w.Write([]byte(err.Error()))
	}
}
