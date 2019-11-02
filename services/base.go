package services

import "nomal-gin/cache"

type BaseService struct {
	cache cache.ICache
}

func NewBaseService(cache cache.ICache) *BaseService{
	return &BaseService{
		cache,
	}
}