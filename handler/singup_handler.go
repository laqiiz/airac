package handler

import (
	"encoding/json"
	"github.com/alexedwards/scs/v2"
	"github.com/laqiiz/airac/model"
	"gopkg.in/go-playground/validator.v9"
	"io"
	"net/http"
	"strconv"
)

func NewSignupHandler(session *scs.SessionManager, r model.UserRepository) SignUpHandler {
	return SignUpHandler{
		r:       r,
		session: session,
	}
}

type SignUpHandler struct {
	r       model.UserRepository
	session *scs.SessionManager
}

type SignUp struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"min=8,max=100"`
}

type AccountCreated struct {
	Warning bool   `json:"warning"`
	Message string `json:"message"`
	UserID  uint   `json:"userID"`
}

func (h *SignUpHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var signUp SignUp
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	if err := json.Unmarshal(body, &signUp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	validate := validator.New()
	if err := validate.Struct(signUp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	ctx := r.Context()

	if _, err := h.r.GetByEmail(ctx, signUp.Email); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	if err != model.NotFound {
		body := model.ProblemError{
			Title: "mail addr already exists",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		_ = json.NewEncoder(w).Encode(body)
		return
	}

	up, err := model.NewSignUp(signUp.Email, signUp.Password)
	if err != nil {
		panic(err) //TODO Temporary impl
	}

	if err := h.r.Insert(ctx, up); err != nil {
		body := model.ProblemError{
			Title: "insert error: " + err.Error(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(body)
		return
	}

	h.session.Put(ctx, "user_id", up.ID)

	_ = json.NewEncoder(w).Encode(up)
}
