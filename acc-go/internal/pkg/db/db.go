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
	dbSetting := setting.DatabaseSetting
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", dbSetting.Host, dbSetting.Port, dbSetting.User, dbSetting.Db, dbSetting.Password)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger2.Default.LogMode(logger2.Info),
	})
	if err != nil {
		logger.Error("初始化数据库连接错误", err.Error())
		os.Exit(1)
	}

	logger.Info("初始化数据库连接成功")
}
