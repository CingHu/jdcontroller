package config

import (
	"testing"
)

func TestParseConfig(t *testing.T) {
	pigeonConfig := GetConfig()
	t.Logf(pigeonConfig.MysqlIp)
	t.Logf(pigeonConfig.MysqlPort)
	t.Logf(pigeonConfig.MysqlUser)
	t.Logf(pigeonConfig.MysqlPasswd)
	t.Logf(pigeonConfig.MysqlDb)
	t.Logf(pigeonConfig.ApiPort)
	t.Logf(pigeonConfig.ControllerPort)
}
