package admin

import "clinicalmate/internal/factory/controller"

type Router interface {
	Set(factory controller.Factory)
}
