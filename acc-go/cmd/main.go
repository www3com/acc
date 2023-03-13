package main

import (
	"acc/internal/pkg/conf"
	routers "acc/internal/router"
	"flag"
	"github.com/upbos/go-saber/db"
	"github.com/upbos/go-saber/log"
	"net/http"
)

func main() {
	var confPath string
	flag.StringVar(&confPath, "conf", "./configs/config.yml", "Configuration file")
	flag.Parse()

	config := conf.Parse(confPath)
	log.Init(config.Log)
	db.Init(config.DataSource)

	serverConf := config.Server
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
