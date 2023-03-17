package conf

import (
	"fmt"
	"github.com/upbos/go-saber/db"
	"github.com/upbos/go-saber/file"
	"github.com/upbos/go-saber/log"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"time"
)

type Server struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Config struct {
	Server     Server         `yaml:"server"`
	DataSource *db.DataSource `yaml:"data-source"`
	Log        *log.Logger    `yaml:"logger"`
}

var conf = &Config{
	Server: Server{
		Addr:         ":8989",
		ReadTimeout:  60_000_000_000, // 60s
		WriteTimeout: 60_000_000_000, // 60s
	},
	DataSource: &db.DataSource{
		Host:     "localhost",
		Port:     5432,
		User:     "acc",
		Password: "acc",
		Database: "acc",
	},
	Log: &log.Logger{
		Level:   "debug",
		Console: true,
		File: &log.RollingFile{
			FileName:   "./logs/acc.log",
			MaxSize:    100,
			MaxBackups: 3,
			MaxAge:     15,
			Compress:   true,
		},
	},
}

const redBold = "\033[31;1m"

func Parse(configPath string) *Config {
	path := file.GetAbsPath(file.GetCwd(), configPath)
	config, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf(redBold+"Read configuration error, detail: %s\n", err)
		os.Exit(1)
	}
	// 替换环境变量
	config = []byte(os.ExpandEnv(string(config)))
	err = yaml.Unmarshal(config, conf)
	if err != nil {
		fmt.Printf(redBold+"Parsing configuration error, detail: %s\n", err)
		os.Exit(1)
	}
	return conf
}
