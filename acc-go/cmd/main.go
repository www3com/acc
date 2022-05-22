package main

import (
	"accounting-service/internal/pkg/db"
	"accounting-service/internal/pkg/i18n"
	"accounting-service/internal/pkg/logger"
	"accounting-service/internal/pkg/setting"
	"accounting-service/internal/pkg/translator"
	"accounting-service/internal/router"
	"fmt"
	"net/http"
	"time"
)

func init() {
	setting.Setup()
	logger.Setup()
	db.Setup()
	translator.Setup()
	i18n.Setup()
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
