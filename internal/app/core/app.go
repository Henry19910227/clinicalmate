package core

import (
	adminControler "clinicalmate/internal/controller/admin"
	"github.com/gin-gonic/gin"
)

type app struct {
	engine    *gin.Engine
	adminCtrl adminControler.Controller
}

func New(engine *gin.Engine, adminCtrl adminControler.Controller) App {
	return &app{engine: engine, adminCtrl: adminCtrl}
}

func (a *app) Run() error {
	return a.engine.Run()
}
