package setting

import (
	"acc/internal/consts"
	"acc/internal/pkg/file"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"time"
)

type Server struct {
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
	Level       string      `yaml:"level"` // 日志级别
	Target      string      `yaml:"target"`
	RollingFile RollingFile `yaml:"file"` // 文件保存配置
}

type RollingFile struct {
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

var ConfigSetting = &Config{
	Server: Server{
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
		Target: "console",
		RollingFile: RollingFile{
			FileName:   "./logs/acc.log",
			MaxSize:    100,
			MaxBackups: 3,
			MaxAge:     15,
			Compress:   true,
		},
	},
}

func Setup(configPath string) {
	path := file.GetAbsPath(file.GetCwd(), configPath)
	config, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf(consts.RedBold+"Read configuration error, detail: %s\n", err)
		os.Exit(1)
	}
	// 替换环境变量
	config = []byte(os.ExpandEnv(string(config)))
	err = yaml.Unmarshal(config, ConfigSetting)
	if err != nil {
		fmt.Printf(consts.RedBold+"Parsing configuration error, detail: %s\n", err)
		os.Exit(1)
	}
}
