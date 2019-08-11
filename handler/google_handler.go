package handler

import (
	"errors"
	"github.com/laqiiz/airac/conn"
	"google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
	"html/template"
	"net/http"
)

type GoogleHandler struct {
}

type GoogleCallbackRequest struct {
	Code  string
	State string
}

func (c *GoogleHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	url := conn.GetGoogleConnect().AuthCodeURL("")
	http.Redirect(w, r, url, 302)
}

func (c *GoogleHandler) GetCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request = GoogleCallbackRequest{
		Code:  r.URL.Query().Get("code"),
		State: r.URL.Query().Get("state"),
	}

	connect := conn.GetGoogleConnect()

	tok, err := connect.Exchange(ctx, request.Code)
	if err != nil {
		panic(err)
	}

	if tok.Valid() == false {
		panic(errors.New("valid token"))
	}

	service, err := oauth2.NewService(ctx, option.WithHTTPClient(connect.Client(ctx, tok)))
	if err != nil {
		panic(err)
	}

	tokenInfo, err := service.Tokeninfo().AccessToken(tok.AccessToken).Context(ctx).Do()
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{}
	data["ID"] = tokenInfo.UserId
	data["Email"] = tokenInfo.Email

	tpl := template.Must(template.ParseFiles("view/google/callback.tpl"))

	if err := tpl.Execute(w, data); err != nil {
		_, _ = w.Write([]byte(err.Error()))
	}

}
