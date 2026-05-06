package service

import (
	storeFactory "clinicalmate/internal/factory/store"
	adminService "clinicalmate/internal/service/admin"
)

type factory struct {
	storeFac storeFactory.Factory
	adminSvc adminService.Service
}

func New(storeFactory storeFactory.Factory) Factory {
	adminSvc := adminService.New(storeFactory)
	return &factory{storeFac: storeFactory, adminSvc: adminSvc}
}

func (f *factory) AdminService() adminService.Service {
	return f.adminSvc
}
