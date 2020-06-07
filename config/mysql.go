package config

import (
	"github.com/go-sql-driver/mysql"
	"time"
)

type mysqlConfig struct {
	ConnMaxLifeTime time.Duration
	MaxIdleConns    int
	MaxOpenConns    int
	Config          *mysql.Config
}

var MysqlCfg = &mysqlConfig{}

func initMySql() {
	cfg := Cfg.GetStringMap("mysql")

	MysqlCfg.Config = &mysql.Config{
		User:                 cfg["user"].(string),
		Passwd:               cfg["passwd"].(string),
		Net:                  cfg["net"].(string),
		Addr:                 cfg["addr"].(string),
		DBName:               cfg["dbname"].(string),
		Collation:            cfg["collation"].(string),
		AllowNativePasswords: cfg["allow_native_passwords"].(bool),
	}
	MysqlCfg.ConnMaxLifeTime, _ = time.ParseDuration(cfg["conn_max_life_time"].(string))
	MysqlCfg.MaxIdleConns = int(cfg["max_idle_conns"].(float64))
	MysqlCfg.MaxOpenConns = int(cfg["max_open_conns"].(float64))
}
