package store

import (
	adminStore "clinicalmate/internal/store/admin"
)

type Factory interface {
	AdminStore() adminStore.Store
}
