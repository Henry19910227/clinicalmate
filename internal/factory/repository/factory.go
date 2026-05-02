package repository

import (
	adminRepo "clinicalmate/internal/repository/admin"
	"gorm.io/gorm"
)

type factory struct {
	db              *gorm.DB
	adminRepository adminRepo.Repository
}

func New(db *gorm.DB) Factory {
	adminR := adminRepo.New(db)
	repoFactory := &factory{db: db, adminRepository: adminR}
	return repoFactory
}

func (f *factory) AdminRepository() adminRepo.Repository {
	return f.adminRepository
}
