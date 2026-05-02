package controller

import (
	adminCtrl "clinicalmate/internal/controller/admin"
	"clinicalmate/internal/factory/service"
)

type factory struct {
	serviceFactory service.Factory
}

func New(serviceFactory service.Factory) Factory {
	return &factory{serviceFactory: serviceFactory}
}

func (f *factory) AdminController() adminCtrl.Controller {
	return adminCtrl.New(f.serviceFactory.AdminService())
}
