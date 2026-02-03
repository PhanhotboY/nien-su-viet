package initialize

import (
	"fmt"
	"net/http"

	"github.com/phanhotboy/nien-su-viet/apps/cms/global"
)

func Run() (http.Handler, string) {
	LoadConfig()
	fmt.Print(global.Config)

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
