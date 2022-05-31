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

func init() {
	setting.Setup()
	logger.Setup()
	db.Setup()
}

func main() {

	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routers.InitRouter(),
		ReadTimeout:    setting.ServerSetting.ReadTimeout * time.Second,
		WriteTimeout:   setting.ServerSetting.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	logger.Infof("start http server listening %s", endPoint)

	server.ListenAndServe()
}
