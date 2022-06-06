package logger

import (
	"acc/internal/pkg/setting"
	"bytes"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func Setup() {
	var (
		writer   io.Writer = os.Stdout
		noColors           = false
	)

	if setting.ConfigSetting.Logger.Target == "file" {
		file := setting.ConfigSetting.Logger.RollingFile
		writer = &lumberjack.Logger{
			Filename:   file.FileName,
			MaxSize:    file.MaxSize,
			MaxBackups: file.MaxBackups,
			MaxAge:     file.MaxAge,
			Compress:   file.Compress,
		}
		noColors = true
	}

	formatter := nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		NoColors:        noColors,
		ShowFullLevel:   true,
		CallerFirst:     true,
		CustomCallerFormatter: func(frame *runtime.Frame) string {
			b := &bytes.Buffer{}
			b.WriteString(" ")
			b.WriteString(frame.File)
			b.WriteString(":")

			b.WriteString(strconv.Itoa(frame.Line))
			return b.String()
		},
	}

	logrus.SetOutput(writer)
	logrus.SetLevel(getLevel())
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&formatter)
}

func getLevel() logrus.Level {
	level := strings.ToLower(setting.ConfigSetting.Logger.Level)
	switch level {
	case "trace":
		return logrus.TraceLevel
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	default:
		return logrus.ErrorLevel
	}
}
