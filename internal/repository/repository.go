package repository

import "go-templ/internal/domain"

type Repositories struct {
	User  domain.UserRepository
	Cache domain.CacheRepository
}
