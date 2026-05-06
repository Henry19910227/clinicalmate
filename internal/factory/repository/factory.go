package repository

import (
	infraFactory "clinicalmate/internal/factory/infra"
	adminRepository "clinicalmate/internal/repository/admin"
)

type factory struct {
	adminRepo adminRepository.Repository
}

func New(infraFac infraFactory.Factory) Factory {
	adminR := adminRepository.New(infraFac.MysqlInfra().GORM())
	repoFactory := &factory{adminRepo: adminR}
	return repoFactory
}

func (f *factory) AdminRepository() adminRepository.Repository {
	return f.adminRepo
}
