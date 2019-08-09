package main

import _ "app/init"

import (
	"app/databases"
	"app/routers"
)

func main() {
	databases.AutoMigrate()

	r := routers.Load()
	r.Run()
}
