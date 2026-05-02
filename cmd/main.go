package main

import (
	"clinicalmate/internal/database"
	controllerFactory "clinicalmate/internal/factory/controller"
	repoFactory "clinicalmate/internal/factory/repository"
	serviceFactory "clinicalmate/internal/factory/service"
	storeFactory "clinicalmate/internal/factory/store"
	adminRouter "clinicalmate/internal/router/admin"
	"github.com/gin-gonic/gin"
)

func main() {
	rdb := database.New()
	repoF := repoFactory.New(rdb.Connect())
	storeF := storeFactory.New(repoF)
	serviceF := serviceFactory.New(storeF)
	controllerF := controllerFactory.New(serviceF)

	g := gin.Default()
	adminR := adminRouter.New(g.Group("/api/v1"))
	adminR.Set(controllerF)

	_ = g.Run(":8080")
}
