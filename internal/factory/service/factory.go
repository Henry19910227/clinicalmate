package service

import (
	"clinicalmate/internal/factory/store"
	adminSvc "clinicalmate/internal/service/admin"
)

type factory struct {
	storeFactory store.Factory
}

func New(storeFactory store.Factory) Factory {
	return &factory{storeFactory: storeFactory}
}

func (f *factory) AdminService() adminSvc.Service {
	return adminSvc.New(f.storeFactory.AdminStore())
}
