package main

import (
	configFactory "clinicalmate/internal/factory/config"
	controllerFactory "clinicalmate/internal/factory/controller"
	infraFactory "clinicalmate/internal/factory/infra"
	repoFactory "clinicalmate/internal/factory/repository"
	serviceFactory "clinicalmate/internal/factory/service"
	storeFactory "clinicalmate/internal/factory/store"
	model "clinicalmate/internal/model/config"
	adminRouter "clinicalmate/internal/router/admin"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"os"
)

func main() {
	configFile := flag.String("config", "config.yaml", "YAML configuration file name")
	flag.Parse()

	data, err := os.ReadFile(*configFile)
	if err != nil {
		panic(err)
	}

	var cfg model.Config
	if err = yaml.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	// 創建 config 層
	configFac := configFactory.New(cfg)

	// 創建 infra 層
	infraFac := infraFactory.New(configFac)

	// 創建 repo 層
	repoF := repoFactory.New(infraFac)

	// 創建 store 層
	storeF := storeFactory.New(repoF)

	// 創建 service 層
	serviceF := serviceFactory.New(storeF)

	// 創建 controller 層
	controllerF := controllerFactory.New(serviceF)

	// 創建 app 層

	g := gin.Default()
	adminR := adminRouter.New(g.Group("/api/v1"))
	adminR.Set(controllerF)

	_ = g.Run(fmt.Sprintf("%s:%d", configFac.AppConfig().Ip(), configFac.AppConfig().Port()))
}
