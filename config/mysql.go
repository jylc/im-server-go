package config

import (
	"github.com/go-sql-driver/mysql"
	"im-server-go/utils"
	"time"
)

type mysqlConfig struct {
	ConnMaxLifeTime time.Duration
	MaxIdleConns    int
	MaxOpenConns    int
	Config          *mysql.Config
}

var MySQLConfig = &mysqlConfig{}

func initMySql() {
	cfg := Cfg.GetStringMap("mysql")

	MySQLConfig.Config = &mysql.Config{
		User:                 utils.AsString(cfg["auth"]),
		Passwd:               utils.AsString(cfg["passwd"]),
		Net:                  utils.AsString(cfg["net"]),
		Addr:                 utils.AsString(cfg["addr"]),
		DBName:               utils.AsString(cfg["dbname"]),
		Collation:            utils.AsString(cfg["collation"]),
		AllowNativePasswords: (cfg["allow_native_passwords"]).(bool),
	}
	MySQLConfig.ConnMaxLifeTime, _ = time.ParseDuration(utils.AsString(cfg["conn_max_life_time"]))
	MySQLConfig.MaxIdleConns = utils.AsInt(cfg["max_idle_conns"])
	MySQLConfig.MaxOpenConns = utils.AsInt(cfg["max_open_conns"])
}
