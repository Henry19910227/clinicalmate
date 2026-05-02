package service

import adminSvc "clinicalmate/internal/service/admin"

type Factory interface {
	AdminService() adminSvc.Service
}
