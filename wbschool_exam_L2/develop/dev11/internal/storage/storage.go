package storage

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type Storage struct {
	*cache.Cache
}

func NewCache() *Storage {
	return &Storage{
		cache.New(5*time.Minute, 10*time.Minute),
	}
}
