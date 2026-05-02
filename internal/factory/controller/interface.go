package controller

import adminCtrl "clinicalmate/internal/controller/admin"

type Factory interface {
	AdminController() adminCtrl.Controller
}
