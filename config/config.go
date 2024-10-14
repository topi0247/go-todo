package config

import (
    "log"
	"gopkg.in/ini.v1"
    "udemy-todo-app/utils"
)

type ConfigList struct {
    Port string
    SQLDriver string
    DbName string
    DbUser string
    DbPassword string
    DbHost string
    DbPort string
    LogFile string
}

var Config ConfigList

func init() {
    LoadConfig()
}

func init() {
    LoadConfig()
    utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
    cfg, err := ini.Load("config.ini")
    if err != nil {
        log.Fatalln(err)
    }
    Config = ConfigList{
        Port: cfg.Section("web").Key("port").MustString("8080"),
        SQLDriver: cfg.Section("db").Key("driver").String(),
        DbName: cfg.Section("db").Key("name").String(),
        DbUser: cfg.Section("db").Key("user").String(),
        DbPassword: cfg.Section("db").Key("password").String(),
        DbHost: cfg.Section("db").Key("host").String(),
        DbPort: cfg.Section("db").Key("port").String(),
        LogFile: cfg.Section("web").Key("logfile").String(),
    }
}
