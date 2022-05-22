package setting

import (
	"fmt"
	"strings"
	"time"
)

// Logger 系统日志配置
type Logger struct {
	Daemon bool   // 是否是后台进程
	Level  string `yaml:"level"`  // 日志级别
	Output string `yaml:"output"` // 输出对象，stdout,file
	File   file   `yaml:"file"`   // 文件保存配置
}

// 系统日志滚动策略
type file struct {
	FileName   string `yaml:"fileName"`   // 文件路径
	MaxSize    int    `yaml:"maxSize"`    // 每个日志文件保存的最大尺寸 单位：M
	MaxBackups int    `yaml:"maxBackups"` // 日志文件最多保存多少个备份
	MaxAge     int    `yaml:"maxAge"`     // 文件最多保存多少天
	Compress   bool   `yaml:"compress"`   // 是否压缩
}

var (
	LOGGER_STDOUT = "stdout"
	LOGGER_FILE   = "file"
)

var LoggerSetting = &Logger{
	Daemon: false,
	Level:  "debug",
	Output: LOGGER_STDOUT,
	File: file{
		FileName:   "./logs/qiankun.log",
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     15,
		Compress:   true,
	},
}

// Server Web Server 配置
type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{
	RunMode:      "release",
	HttpPort:     8989,
	ReadTimeout:  60,
	WriteTimeout: 60,
}

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	Db       string
}

var DatabaseSetting = &Database{
	Host:     "192.168.3.3",
	Port:     5432,
	User:     "acc",
	Password: "acc@2022",
	Db:       "acc",
}

type Email struct {
	Host     string
	Port     int
	UserName string
	Password string
}

var EmailSetting = &Email{
	Host:     "smtp.126.com",
	Port:     25,
	UserName: "wjzchina2008@126.com",
	Password: "EJOSPQEVVBIDEIHJ",
}

func Setup() {
	fmt.Println("读取参数文件成功")
}

func (s *Logger) IsDebug() bool {
	if strings.ToLower(s.Level) == "debug" {
		return true
	}
	return false
}
