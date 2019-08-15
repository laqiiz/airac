package model

import (
	"context"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// UserInfo represents OAuth returning info.
type UserInfo struct {
	ID       string
	MailAddr string
}

type Account struct {
	ID        string
	MailAddr  string
	PassHash  string
	CreatedAt time.Time
}

func NewSignUp(mailAddr, pw string) (*Account, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Account{
		ID:        xid.New().String(),
		MailAddr:  mailAddr,
		PassHash:  string(hash),
		CreatedAt: time.Now(),
	}, nil
}

type UserRepository interface {
	GetByEmail(ctx context.Context, mailAddr string) (*Account, error)
	Insert(ctx context.Context, signUp *Account) error
	Delete(ctx context.Context, userID string) error
}
