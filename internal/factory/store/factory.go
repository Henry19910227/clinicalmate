package store

import (
	"clinicalmate/internal/factory/repository"
	"clinicalmate/internal/store/admin"
)

type factory struct {
	repoFactory repository.Factory
}

func New(repoFactory repository.Factory) Factory {
	return &factory{repoFactory: repoFactory}
}

func (f *factory) AdminStore() admin.Store {
	return admin.New(f.repoFactory.AdminRepository())
}
