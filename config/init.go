package config

import (
	"github.com/spf13/viper"
	"log"
)

var Cfg *viper.Viper

func init() {
	Cfg = viper.New()
	Cfg.SetConfigName("conf")
	Cfg.SetConfigType("json")
	Cfg.AddConfigPath("./config/")
	err := Cfg.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置文件失败：%s", err.Error())
	}

	initMySql()
}
