package mysql

import (
	mysqlCfg "clinicalmate/internal/config/mysql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type infra struct {
	db *gorm.DB
}

func New(cfg mysqlCfg.Config) Infra {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.Username(), cfg.Password(), cfg.Host(), cfg.Port(), cfg.Database())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB: %v", err)
	}
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns())
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns())
	return &infra{db: db.Debug()}
}

func (i *infra) GORM() *gorm.DB {
	return i.db
}
