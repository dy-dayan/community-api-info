package main

import (
	"github.com/dy-dayan/community-api-info/router"
	"github.com/dy-gopkg/kit/web"
)

func main() {
	// Create service
	web.Init()

	r := router.Init()

	// Register Handler
	web.DefaultService.Handle("/", r)

	// Run server
	web.Run()
}
