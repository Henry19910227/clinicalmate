package admin

import adminRepo "clinicalmate/internal/repository/admin"

type store struct {
	adminRepo adminRepo.Repository
}

func New(adminRepo adminRepo.Repository) Store {
	return &store{adminRepo: adminRepo}
}
