package db

import (
	"acc/internal/pkg/setting"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strings"
	"time"
)

var DB *gorm.DB

type dbWriter struct{}

func (w *dbWriter) Printf(format string, v ...interface{}) {
	logrus.Debugf(format, v...)
}

func Setup() {
	ds := setting.ConfigSetting.DataSource
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		ds.Host, ds.Port, ds.User, ds.Database, ds.Password)

	slowLogger := logger.New(
		&dbWriter{},
		logger.Config{
			Colorful:                  false,
			SlowThreshold:             1 * time.Second,
			LogLevel:                  getLevel(),
			IgnoreRecordNotFoundError: true,
		},
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: slowLogger,
	})

	if err != nil {
		logrus.Error("Database connection error", err.Error())
		os.Exit(1)
	}

	logrus.Info("Connected to the database successfully")
}

func getLevel() logger.LogLevel {
	level := strings.ToLower(setting.ConfigSetting.Logger.Level)
	switch level {
	case "trace", "debug", "info":
		return logger.Info
	case "warn":
		return logger.Warn
	default:
		return logger.Error
	}
}
