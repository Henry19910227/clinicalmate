package config

import (
	appConfig "clinicalmate/internal/config/app"
	mysqlConfig "clinicalmate/internal/config/mysql"
)

type Factory interface {
	MysqlConfig() mysqlConfig.Config
	AppConfig() appConfig.Config
}
