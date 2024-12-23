package util

import (
	"basic_management/learn-go-management/model"
	"flag"
	"os"

	"github.com/sirupsen/logrus"
)

var Logger logrus.Logger

func init() {

	Logger = *logrus.New()

	Logger.Out = os.Stdout

	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

}

func SetLogger() {

	loglevel := flag.String(model.LogLevel, model.LogLevelInfo, "log-level (debug, error, info, warning)")

	flag.Parse()
	switch *loglevel {
	case model.LogLevelDebug:
		Logger.SetLevel(logrus.DebugLevel)
	case model.LogLevelError:
		Logger.SetLevel(logrus.ErrorLevel)
	case model.LogLevelWarning:
		Logger.SetLevel(logrus.WarnLevel)
	default:
		Logger.SetLevel(logrus.InfoLevel)
	}

}

func Log(loglevel, packageLevel, functionName string, message, parameter interface{}) {
	switch loglevel {
	case model.LogLevel:
		if parameter != nil {
			Logger.Debugf("packageLevel: %s, functionName: %s, message: %v, parameter: %v\n", packageLevel, functionName, message, parameter)
		} else {
			Logger.Debugf("packageLevel: %s, functionName: %s, message: %v\n", packageLevel, functionName, message)
		}
	case model.LogLevelError:
		if parameter != nil {
			Logger.Errorf("packageLevel: %s, functionName: %s, message: %v, parameter: %v\n", packageLevel, functionName, message, parameter)
		} else {
			Logger.Errorf("packageLevel: %s, functionName: %s, message: %v\n", packageLevel, functionName, message)
		}
	case model.LogLevelWarning:
		if parameter != nil {
			Logger.Warnf("packageLevel: %s, functionName: %s, message: %v, parameter: %v\n", packageLevel, functionName, message, parameter)
		} else {
			Logger.Warnf("packageLevel: %s, functionName: %s, message: %v\n", packageLevel, functionName, message)
		}
	default:
		if parameter != nil {
			Logger.Infof("packageLevel: %s, functionName: %s, message: %v, parameter: %v\n", packageLevel, functionName, message, parameter)
		} else {
			Logger.Infof("packageLevel: %s, functionName: %s, message: %v\n", packageLevel, functionName, message)
		}
	}
}
