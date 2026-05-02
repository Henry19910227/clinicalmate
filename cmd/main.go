package main

import (
	"clinicalmate/config"
	"clinicalmate/internal/database"
	"clinicalmate/internal/router/admin"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg.DBDSN)

	r := admin.New(db)
	r.Run(":" + cfg.Port)
}
