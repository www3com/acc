package main

import (
	"acc/internal/pkg/conf"
	routers "acc/internal/router"
	"flag"
	"github.com/upbos/go-base/db"
	"github.com/upbos/go-base/log"
	"net/http"
)

func main() {
	var confPath string
	flag.StringVar(&confPath, "conf", "./configs/config.yml", "Configuration file")
	flag.Parse()

	conf.Init(confPath)
	log.Init(conf.Info.Log)
	db.Init(conf.Info.DataSource)

	serverConf := conf.Info.Server
	server := &http.Server{
		Addr:           serverConf.Addr,
		Handler:        routers.InitRouter(),
		ReadTimeout:    serverConf.ReadTimeout,
		WriteTimeout:   serverConf.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	log.Infof("Start http server listening %s", serverConf.Addr)
	err := server.ListenAndServe()
	log.Errorf(err, "Start http server error, detail: %v")
}
