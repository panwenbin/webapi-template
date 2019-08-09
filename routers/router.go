package routers

import (
	"app/actions"
	"github.com/gin-gonic/gin"
)

func Load() *gin.Engine {
	r := gin.Default()
	r.GET("/hello", actions.Hello)

	return r
}
