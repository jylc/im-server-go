package connections

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"im-server-go/config"
	"log"
)

var db *gorm.DB

func initMySql() {
	dsn := config.MysqlCfg.Config.FormatDSN()

	var err error
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("连接数据库失败：%s", err.Error())
	}

	db.DB().SetConnMaxLifetime(config.MysqlCfg.ConnMaxLifeTime)
	db.DB().SetMaxIdleConns(config.MysqlCfg.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.MysqlCfg.MaxOpenConns)

	err = db.DB().Ping()
	if err != nil {
		log.Fatalf("Ping数据库失败：%s", err.Error())
	}
}
