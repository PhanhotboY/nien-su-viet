package initialize

import (
	"net/http"

	"github.com/phanhotboy/nien-su-viet/apps/cms/global"
)

func Run() (http.Handler, int) {
	LoadConfig()

	db, err := InitPostgreSQL()
	if err != nil {
		panic(err)
	}
	InitRabbitMQ()

	logLevel := "debug"
	if global.Config.Server.Env == "production" {
		logLevel = "info"
	}

	r := InitRouter(db, logLevel)
	return r, global.Config.Server.Port

}
