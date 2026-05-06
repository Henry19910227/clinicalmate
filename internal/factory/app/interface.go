package app

import "clinicalmate/internal/app/core"

type Factory interface {
	CoreApp() core.App
}
