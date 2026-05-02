package service

import (
	"clinicalmate/internal/factory/store"
)

type factory struct {
	storeFactory store.Factory
}

func New(storeFactory store.Factory) Factory {
	serviceFactory := &factory{storeFactory: storeFactory}
	return serviceFactory
}
