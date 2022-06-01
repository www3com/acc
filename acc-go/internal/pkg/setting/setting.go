package setting

import (
	"acc/internal/pkg/util"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"time"
)

type Server struct {
	RunMode      string        `yaml:"run-mode"`
	Port         int           `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read-timeout"`
	WriteTimeout time.Duration `yaml:"write-timeout"`
}

type DataSource struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Email struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// Logger 系统日志配置
type Logger struct {
	Level  string `yaml:"level"`  // 日志级别
	Output string `yaml:"output"` // 输出对象，stdout,file
	File   file   `yaml:"file"`   // 文件保存配置
}

// 系统日志滚动策略
type file struct {
	FileName   string `yaml:"file-name"`   // 文件路径
	MaxSize    int    `yaml:"max-size"`    // 每个日志文件保存的最大尺寸 单位：M
	MaxBackups int    `yaml:"max-backups"` // 日志文件最多保存多少个备份
	MaxAge     int    `yaml:"max-age"`     // 文件最多保存多少天
	Compress   bool   `yaml:"compress"`    // 是否压缩
}

type Config struct {
	Server     Server     `yaml:"server"`
	DataSource DataSource `yaml:"data-source"`
	Logger     Logger     `yaml:"logger"`
	Email      Email      `yaml:"email"`
}

var (
	LOGGER_STDOUT = "stdout"
	LOGGER_FILE   = "file"
)

var ConfigSetting = &Config{
	Server: Server{
		RunMode:      "release",
		Port:         8989,
		ReadTimeout:  60,
		WriteTimeout: 60,
	},
	DataSource: DataSource{
		Host:     "localhost",
		Port:     5432,
		User:     "acc",
		Password: "acc",
		Database: "acc",
	},
	Email: Email{
		Host:     "smtp.126.com",
		Port:     25,
		User:     "",
		Password: "",
	},
	Logger: Logger{
		Level:  "debug",
		Output: LOGGER_STDOUT,
		File: file{
			FileName:   "./logs/acc.log",
			MaxSize:    100,
			MaxBackups: 3,
			MaxAge:     15,
			Compress:   true,
		},
	},
}

func Setup(configPath string) {
	if configPath == "" {
		configPath = "./configs/config.yaml"
	}
	path := util.GetAbsPath(util.GetCwd(), configPath)
	config, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Read configuration file error")
		os.Exit(1)
	}
	// 替换环境变量
	config = []byte(os.ExpandEnv(string(config)))
	err = yaml.Unmarshal(config, ConfigSetting)
	if err != nil {
		fmt.Println("Parsing configuration file error")
		os.Exit(1)
	}
}
