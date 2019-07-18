package main

import _ "app/init"

import (
	"app/actions"
	"app/databases"
	"github.com/gin-gonic/gin"
)

func main() {
	databases.AutoMigrate()

	r := gin.Default()
	r.GET("/hello", actions.Hello)

	r.Run()
}
