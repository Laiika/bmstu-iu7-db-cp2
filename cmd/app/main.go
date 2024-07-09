package main

import (
	"db_cp_6/config"
	"db_cp_6/internal/app"
	"db_cp_6/pkg/logger"
)

func main() {
	log := logger.GetLogger()
	cfg := config.GetConfig(log)

	app.Run(cfg, log)
}
