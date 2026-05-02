package database

import "gorm.io/gorm"

type RDB interface {
	Connect() *gorm.DB
}
