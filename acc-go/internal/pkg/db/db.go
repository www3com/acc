package db

import (
	"acc/internal/pkg/logger"
	"acc/internal/pkg/setting"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
	"os"
)

var DB *gorm.DB

func Setup() {
	ds := setting.ConfigSetting.DataSource
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		ds.Host, ds.Port, ds.User, ds.Database, ds.Password)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger2.Default.LogMode(logger2.Info),
	})
	if err != nil {
		logger.Error("Database connection error", err.Error())
		os.Exit(1)
	}

	logger.Info("Connected to the database successfully")
}
