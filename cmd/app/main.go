package main

import (
	"db_cp_6/config"
	"db_cp_6/internal/app"
	"db_cp_6/pkg/logger"
)

//	@title			DB course project API
//	@version		1.0
//	@description	This is db course project backend API.

//	@contact.name	API Support
//	@contact.email	evgeniazavojskih@gmail.com

// @host		localhost:8080
// @BasePath	/api/v1
// @Schemes	http
func main() {
	log := logger.GetLogger()
	cfg := config.GetConfig(log)

	app.Run(cfg, log)
}
