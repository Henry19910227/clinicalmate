package repository

import adminRepo "clinicalmate/internal/repository/admin"

type Factory interface {
	AdminRepository() adminRepo.Repository
}
