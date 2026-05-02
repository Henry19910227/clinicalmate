package admin

import (
	adminStore "clinicalmate/internal/store/admin"
)

type service struct {
	adminStore adminStore.Store
}

func New(adminStore adminStore.Store) Service {
	return &service{adminStore: adminStore}
}
