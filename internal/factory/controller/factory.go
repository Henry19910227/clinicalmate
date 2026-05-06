package controller

import (
	adminController "clinicalmate/internal/controller/admin"
	serviceFactory "clinicalmate/internal/factory/service"
)

type factory struct {
	serviceFac serviceFactory.Factory
	adminCtrl  adminController.Controller
}

func New(serviceFac serviceFactory.Factory) Factory {
	adminCtrl := adminController.New(serviceFac)
	return &factory{serviceFac: serviceFac, adminCtrl: adminCtrl}
}

func (f *factory) AdminController() adminController.Controller {
	return f.adminCtrl
}
