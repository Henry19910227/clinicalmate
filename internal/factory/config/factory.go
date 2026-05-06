package config

import (
	appConfig "clinicalmate/internal/config/app"
	mysqlConfig "clinicalmate/internal/config/mysql"
	model "clinicalmate/internal/model/config"
)

type factory struct {
	cfg model.Config
}

func New(cfg model.Config) Factory {
	return &factory{cfg: cfg}
}

func (f *factory) MysqlConfig() mysqlConfig.Config {
	return mysqlConfig.New(&f.cfg.Database)
}

func (f *factory) AppConfig() appConfig.Config {
	return appConfig.New(&f.cfg.App)
}
