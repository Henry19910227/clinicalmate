package app

import (
	coreApp "clinicalmate/internal/app/core"
	controllerFactory "clinicalmate/internal/factory/controller"
	"github.com/gin-gonic/gin"
)

type factory struct {
	engine        *gin.Engine
	controllerFac controllerFactory.Factory
	coreApp       coreApp.App
}

func New(engine *gin.Engine, controllerFac controllerFactory.Factory) Factory {
	coreApp := coreApp.New(engine, controllerFac)
	return &factory{engine: engine, controllerFac: controllerFac, coreApp: coreApp}
}

func (f *factory) CoreApp() coreApp.App {
	return f.coreApp
}
