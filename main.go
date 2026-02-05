package main

import (
	"golang_crud/src/plugin"
	"golang_crud/src/plugin/dependencies"
	"golang_crud/src/plugin/logger"
)

func main() {
	logger.Init()
	dependencies := dependencies.Init()
	plugin.RegisterRoutes(&dependencies.UserService)
}
