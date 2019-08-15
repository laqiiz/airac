package handler

import (
	"encoding/json"
	"github.com/alexedwards/scs/v2"
	"github.com/laqiiz/airac/jwt"
	"github.com/laqiiz/airac/model"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
)

type SignHandler struct {
	r       model.UserRepository
	session *scs.SessionManager
}

func NewSignHandler(session *scs.SessionManager, r model.UserRepository) SignHandler {
	return SignHandler{
		r:       r,
		session: session,
	}
}

type SignUp struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"min=8,max=100"`
}

type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *SignHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var signUp SignUp
	if err := json.NewDecoder(r.Body).Decode(&signUp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	// validate at only sign-up because do not save database for illegal variable
	if err := validator.New().Struct(signUp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	ctx := r.Context()

	_, err := h.r.GetByEmail(ctx, signUp.Email)
	if err != nil && err != model.NotFound {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}
	if err == nil {
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

		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(body)
		return
	}

	h.session.Put(ctx, "user_id", up.ID)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(up)
}

func (h *SignHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var signIn SignIn
	if err := json.NewDecoder(r.Body).Decode(&signIn); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	ctx := r.Context()
	userInfo, err := h.r.GetByEmail(ctx, signIn.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	if err == model.NotFound {
		body := model.ProblemError{
			Title: "mail addr is not found",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(body)
		return
	}

	h.session.Put(ctx, "user_id", userInfo.ID)

	jwtToken, err := jwt.GenerateToken(userInfo.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	// Cookie
	http.SetCookie(w, &http.Cookie{  //TODO 他のオプションも検討
		Name:  "jwtToken",
		Value: jwtToken,
	})

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(userInfo)
}
