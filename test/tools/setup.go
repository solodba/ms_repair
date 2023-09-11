package tools

import (
	"github.com/solodba/mcube/apps"
	"github.com/solodba/mcube/logger"
	_ "github.com/solodba/ms_repair/apps/all"
	"github.com/solodba/ms_repair/conf"
)

func LoadConfig() {
	conf.Conf = conf.NewDefaultConfig()
	conf.Conf.MySQL.Username = "root"
	conf.Conf.MySQL.Password = "Root@123"
	conf.Conf.MySQL.Host = "118.195.139.53"
	conf.Conf.MySQL.Port = 3306
	conf.Conf.MySQL.DB = "mysql"
	conf.Conf.MySQL.MaxOpenConn = 50
	conf.Conf.MySQL.MaxIdleConn = 10
	conf.Conf.MySQL.MaxLifeTime = 600
	conf.Conf.MySQL.MaxIdleTime = 300
}

func DevelopmentSet() {
	LoadConfig()
	err := apps.InitInternalApps()
	if err != nil {
		logger.L().Panic().Msgf("initial object config error, err: %s", err.Error())
	}
}
