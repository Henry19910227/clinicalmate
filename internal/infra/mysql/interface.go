package mysql

import "gorm.io/gorm"

type Infra interface {
	GORM() *gorm.DB
}
