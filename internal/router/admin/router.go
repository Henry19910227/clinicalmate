package admin

import (
	"clinicalmate/internal/factory/controller"
	"github.com/gin-gonic/gin"
)

type router struct {
	group *gin.RouterGroup
}

func New(group *gin.RouterGroup) Router {
	r := &router{group: group}
	return r
}

func (r *router) Set(factory controller.Factory) {
	r.group.GET("/admin", func(c *gin.Context) {})
}
