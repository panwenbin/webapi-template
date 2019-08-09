package routers

import (
	"app/actions"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func Load() *gin.Engine {
	r.GET("/hello", actions.Hello)

	return r
}

func init() {
	r = gin.Default()
}