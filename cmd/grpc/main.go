package main

import (
	"golectro-user/internal/command"
	"golectro-user/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	validate := config.NewValidator(viperConfig)

	if !command.NewCommandExecutor(db).Execute(log) {
		return
	}

	config.StartGRPC(viperConfig, db, validate, log)
}
