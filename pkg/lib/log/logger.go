package log

import (
	"fmt"
	"github.com/tallongsun/go-scaffold/pkg/lib/config"
	"os"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func Init() {
	Logger = logrus.New()
	Logger.Formatter = new(logrus.JSONFormatter)
	Logger.SetReportCaller(true)
	mode := config.Config.Get("mode")

	if mode == "alpha" || mode == "beta" || mode == "release" {
		if mode == "alpha" || mode == "beta" {
			Logger.SetLevel(logrus.DebugLevel)
		} else {
			Logger.SetLevel(logrus.InfoLevel)
		}
		logf, err := rotatelogs.New(
			"logs/app.%Y%m%d.log",
			rotatelogs.WithMaxAge(24*time.Hour),
			rotatelogs.WithRotationTime(24*time.Hour),
		)
		if err != nil {
			fmt.Errorf("failed to create rotatelogs: %s", err)
			return
		}
		Logger.Out = logf
	} else {
		Logger.Out = os.Stdout
	}
}
