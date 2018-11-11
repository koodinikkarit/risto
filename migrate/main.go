package main

import (
	"github.com/koodinikkarit/database"
	"github.com/koodinikkarit/risto/config"
)

func main() {
	config.ValidateConfig()

	database.Migrate(
		config.MysqlUsername,
		config.MysqlPassword,
		config.MysqlHost,
		config.MysqlPort,
		config.MysqlDatabase,
	)
}
