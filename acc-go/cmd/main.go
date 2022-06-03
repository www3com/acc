package main

import (
	"acc/internal/pkg/db"
	"acc/internal/pkg/logger"
	"acc/internal/pkg/setting"
	"acc/internal/router"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func main() {
	var config string
	flag.StringVar(&config, "config", "./configs/config.yml", "Configuration file")
	flag.Parse()

	setting.Setup(config)
	logger.Setup()
	db.Setup()

	serverConf := setting.ConfigSetting.Server
	addr := fmt.Sprintf(":%d", serverConf.Port)
	server := &http.Server{
		Addr:           addr,
		Handler:        routers.InitRouter(),
		ReadTimeout:    serverConf.ReadTimeout * time.Second,
		WriteTimeout:   serverConf.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	logrus.Infof("Start http server listening %s", addr)
	err := server.ListenAndServe()
	logrus.Errorf("Start http server error, detail: %v", err)
}
