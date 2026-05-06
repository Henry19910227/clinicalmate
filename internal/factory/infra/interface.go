package infra

import mysqlInfra "clinicalmate/internal/infra/mysql"

type Factory interface {
	MysqlInfra() mysqlInfra.Infra
}
