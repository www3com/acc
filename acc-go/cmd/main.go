package main

import (
	"acc/internal/pkg/db"
	"acc/internal/pkg/logger"
	"acc/internal/pkg/setting"
	"acc/internal/router"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// 保证顺序
	setting.Setup("./configs/config-dev.yaml")
	logger.Setup()
	db.Setup()

	endPoint := fmt.Sprintf(":%d", setting.ConfigSetting.Server.Port)
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routers.InitRouter(),
		ReadTimeout:    setting.ConfigSetting.Server.ReadTimeout * time.Second,
		WriteTimeout:   setting.ConfigSetting.Server.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	logger.Infof("Start http server listening %s", endPoint)

	server.ListenAndServe()
}
