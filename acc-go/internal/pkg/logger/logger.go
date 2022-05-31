package logger

import (
	"acc/internal/pkg/setting"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var log = logrus.New()

func Setup() {

	log.SetLevel(logrus.DebugLevel)

	noColors := false
	if setting.ConfigSetting.Logger.Output == setting.LOGGER_STDOUT {
		log.SetOutput(os.Stdout)
	} else {
		noColors = true
		output := &lumberjack.Logger{
			Filename:   setting.ConfigSetting.Logger.File.FileName,
			MaxSize:    setting.ConfigSetting.Logger.File.MaxSize,
			MaxBackups: setting.ConfigSetting.Logger.File.MaxBackups,
			MaxAge:     setting.ConfigSetting.Logger.File.MaxAge,
			Compress:   setting.ConfigSetting.Logger.File.Compress,
		}
		log.SetOutput(output)
	}

	log.SetFormatter(&nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		NoColors:        noColors,
		ShowFullLevel:   true,
	})
}

func GetLog() *logrus.Logger {
	return log
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}
