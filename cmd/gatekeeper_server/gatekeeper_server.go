package main

import (
	"github.com/spf13/viper"
	"is.having.coffee/gatekeeper/gatekeeper-server/config"
	"is.having.coffee/gatekeeper/gatekeeper-server/database"
)

func main() {
	if err := config.InitConfig(); err != nil {
		panic(err)
	}

	if err := database.InitDatabase(); err != nil {
		panic(err)
	}

	defer database.Db.Close()

	if router, err := initAPI(); err != nil {
		panic(err)
	} else {
		go router.Run(viper.GetString(config.ConfigListenSocket))
	}

	if viper.GetBool(config.ConfigConsole) {
		shell := initShell()
		shell.Run()
	}
}