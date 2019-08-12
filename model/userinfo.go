package model

import (
	"context"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserInfo struct {
	ID       string
	MailAddr string
}

type SignUp struct {
	ID        string
	MailAddr  string
	PassHash  string
	CreatedAt time.Time
}

func NewSignUp(mailAddr, pw string) (*SignUp, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MaxCost)
	if err != nil {
		return nil, err
	}

	return &SignUp{
		ID:        xid.New().String(),
		MailAddr:  mailAddr,
		PassHash:  string(hash),
		CreatedAt: time.Now(),
	}, nil
}

type UserRepository interface {
	GetByEmail(ctx context.Context, mailAddr string) (*UserInfo, error)
	Insert(ctx context.Context, signUp *SignUp) error
	Delete(ctx context.Context, userID string) error
}
