package main

import (
	"github.com/tanakornwry/mars-exploration-project/config"
	"github.com/tanakornwry/mars-exploration-project/infrastructure/router"
)

var Config config.Config

func main() {
	Config = config.LoadConfiguration("./config/config.json")

	r := router.SetupRouter()
	r.Run(":" + Config.Port)
}
