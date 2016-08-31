package config

import (
	"strconv"
	"fmt"
)

const (
	LOG_LEVEL = "DEBUG"
	CONFIG_FILE = "/config.ini"
	LOG_CONFIG_FILE = "/log.properties"
)

type PigeonConfig struct {
	MysqlIp string
	MysqlPort string
	MysqlUser string
	MysqlPasswd string
	MysqlDb string
	ApiPort string
	ControllerPort string
	CacheFlushInterval int
	FlowmodIdleTimeout int
}

var pigeonConfig *PigeonConfig

func GetConfig() *PigeonConfig {
	return pigeonConfig
}


func InitConfig(configDir string) *PigeonConfig {

	pigeonConfig = new(PigeonConfig)
	cfg, err := ReadDefault(configDir + CONFIG_FILE)
	if err != nil {
	 	fmt.Println("Can't find config file", configDir)
	}
	//-h -u -p -D
	if cfg.HasSection("default") {
		configItem, _ := cfg.String("default", "mysql_ip")
		pigeonConfig.MysqlIp = configItem
		configItem, _ = cfg.String("default", "mysql_port")
		pigeonConfig.MysqlPort = configItem
		configItem, _ = cfg.String("default", "mysql_user")
		pigeonConfig.MysqlUser = configItem
		configItem, _ = cfg.String("default", "mysql_passwd")
		pigeonConfig.MysqlPasswd = configItem
		configItem, _ = cfg.String("default", "mysql_db")
		pigeonConfig.MysqlDb = configItem
		configItem, _ = cfg.String("default", "api_port")
		pigeonConfig.ApiPort = configItem
		configItem, _ = cfg.String("default", "controller_port")
		pigeonConfig.ControllerPort = configItem
		configItem, _ = cfg.String("default", "cache_flush_interval")
		cacheFlushInterval, err := strconv.Atoi(configItem)
		if err != nil {
			cacheFlushInterval = 600
		}
		pigeonConfig.CacheFlushInterval = cacheFlushInterval
		configItem, _ = cfg.String("default", "flowmod_idle_timeout")
		flowmodIdleTimeout, err := strconv.Atoi(configItem)
		if err != nil {
			flowmodIdleTimeout = 600
		}
		pigeonConfig.FlowmodIdleTimeout = flowmodIdleTimeout
	}
	return pigeonConfig
}
