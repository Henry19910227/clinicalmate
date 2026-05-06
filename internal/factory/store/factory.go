package store

import (
	repositoryFactory "clinicalmate/internal/factory/repository"
	adminStore "clinicalmate/internal/store/admin"
)

type factory struct {
	repositoryFac repositoryFactory.Factory
	adminSto      adminStore.Store
}

func New(repositoryFac repositoryFactory.Factory) Factory {
	adminSto := adminStore.New(repositoryFac)
	return &factory{repositoryFac: repositoryFac, adminSto: adminSto}
}

func (f *factory) AdminStore() adminStore.Store {
	return f.adminSto
}
