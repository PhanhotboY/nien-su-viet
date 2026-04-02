package defaultLogger

import (
	"os"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/config/settings"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger/zap"
)

var l logger.Logger

func initLogger() {
	logType := os.Getenv("LogConfig_LogType")

	switch logType {
	case "Zap", "":
	default:
		l = zap.NewZapLogger(
			settings.LoadConfig(),
		)
		break
		// case "Logrus":
		// 	l = logrous.NewLogrusLogger(
		// 		&config.LogOptions{LogType: models.Logrus, CallerEnabled: false},
		// 		constants.Dev,
		// 	)
		// 	break
	}
}

func GetLogger() logger.Logger {
	if l == nil {
		initLogger()
	}

	return l
}
