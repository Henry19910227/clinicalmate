package admin

import adminService "clinicalmate/internal/service/admin"

type controller struct {
	adminService adminService.Service
}

func New(adminService adminService.Service) Controller {
	return &controller{adminService: adminService}
}
