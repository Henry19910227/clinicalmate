package store

import (
	"clinicalmate/internal/factory/repository"
	"clinicalmate/internal/store/admin"
)

type factory struct {
	repoFactory repository.Factory
}

func (f *factory) AdminService() admin.Store {
	adminRepo := f.repoFactory.AdminRepository()
	return admin.New(adminRepo)
}
