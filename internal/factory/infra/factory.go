package infra

import (
	configFactory "clinicalmate/internal/factory/config"
	mysqlInfra "clinicalmate/internal/infra/mysql"
)

type factory struct {
	configFac configFactory.Factory
	mysqlInf  mysqlInfra.Infra
}

func New(configFac configFactory.Factory) Factory {
	mysqlInf := mysqlInfra.New(configFac.MysqlConfig())
	return &factory{configFac: configFac, mysqlInf: mysqlInf}
}

func (f *factory) MysqlInfra() mysqlInfra.Infra {
	return f.mysqlInf
}
