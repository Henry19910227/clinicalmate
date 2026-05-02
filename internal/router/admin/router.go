package admin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	_ = v1

	return r
}
