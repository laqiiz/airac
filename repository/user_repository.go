package repository

import (
	"context"
	"github.com/laqiiz/airac/model"
	"github.com/patrickmn/go-cache"
	"time"
)

func NewMemUserRepository() model.UserRepository {
	return &MemUseRepository{
		cache: cache.New(30*time.Minute, 60*time.Minute), // TODO
	}
}

type MemUseRepository struct {
	cache *cache.Cache
}

func (r *MemUseRepository) GetByEmail(ctx context.Context, mailAddr string) (*model.UserInfo, error) {
	info, found := r.cache.Get(mailAddr)
	if !found {
		return nil, model.NotFound
	}

	userInfo := info.(*model.UserInfo)
	return userInfo, nil
}

func (r *MemUseRepository) Insert(ctx context.Context, signUp *model.SignUp) error {
	r.cache.Set(signUp.ID, signUp, cache.DefaultExpiration)
	return nil
}

func (r *MemUseRepository) Delete(ctx context.Context, userID string) error {
	r.cache.Delete(userID)
	return nil
}
