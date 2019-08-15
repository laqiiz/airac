package repository

import (
	"context"
	"github.com/laqiiz/airac/model"
	"github.com/patrickmn/go-cache"
	"log"
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

func (r *MemUseRepository) GetByEmail(ctx context.Context, mailAddr string) (*model.Account, error) {
	log.Println("start")
	info, found := r.cache.Get(mailAddr)
	if !found {
		log.Println("not found")
		return nil, model.NotFound
	}
	log.Println("found")

	a := info.(*model.Account)
	return a, nil
}

func (r *MemUseRepository) Insert(ctx context.Context, a *model.Account) error {
	r.cache.Set(a.MailAddr, a, cache.DefaultExpiration)
	return nil
}

func (r *MemUseRepository) Delete(ctx context.Context, userID string) error {
	r.cache.Delete(userID)
	return nil
}
