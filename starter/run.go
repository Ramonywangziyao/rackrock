package starter

import (
	"rackrock/config"
	"rackrock/logger"
	"rackrock/starter/component"
)

func RunServer() {
	//  conf
	config.Init()

	// log
	logger.Init()

	//  init db
	component.InitDatabase()

	// do run
	serverRun()
}
