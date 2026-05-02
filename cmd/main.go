package main

import (
	appConfig "clinicalmate/config/app"
	"clinicalmate/internal/database"
	controllerFactory "clinicalmate/internal/factory/controller"
	repoFactory "clinicalmate/internal/factory/repository"
	serviceFactory "clinicalmate/internal/factory/service"
	storeFactory "clinicalmate/internal/factory/store"
	adminRouter "clinicalmate/internal/router/admin"
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	configFile := flag.String("config", "config.yaml", "YAML configuration file name")
	flag.Parse()

	cfg := appConfig.New(*configFile).Config()

	rdb := database.New(&cfg.Database)
	repoF := repoFactory.New(rdb.Connect())
	storeF := storeFactory.New(repoF)
	serviceF := serviceFactory.New(storeF)
	controllerF := controllerFactory.New(serviceF)

	g := gin.Default()
	adminR := adminRouter.New(g.Group("/api/v1"))
	adminR.Set(controllerF)

	_ = g.Run(fmt.Sprintf("%s:%d", cfg.App.Ip, cfg.App.Port))
}
