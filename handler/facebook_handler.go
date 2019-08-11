package handler

import (
	"errors"
	"github.com/huandu/facebook"
	"github.com/laqiiz/airac/conn"
	"html/template"
	"net/http"
	"strconv"
)

type FacebookHandler struct {
}

type FacebookCallbackRequest struct {
	Code  string
	State int
}

func (h *FacebookHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, conn.GetFacebookConnect().AuthCodeURL(""), 302)
}

func (h *FacebookHandler) GetCallback(w http.ResponseWriter, r *http.Request) {
	state, err := strconv.Atoi(r.URL.Query().Get("state"))
	if err != nil {
		state = 0 // TODO
	}
	var request = FacebookCallbackRequest{
		Code:  r.URL.Query().Get("code"),
		State: state,
	}

	con := conn.GetFacebookConnect()

	tok, err := con.Exchange(r.Context(), request.Code)
	if err != nil {
		panic(err)
	}

	if tok.Valid() == false {
		panic(errors.New("valid token"))
	}

	session := &facebook.Session{
		Version:    "v2.8",
		HttpClient: con.Client(r.Context(), tok),
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

	if err := tpl.Execute(w, data); err != nil {
		_, _ = w.Write([]byte(err.Error()))
	}
}
