package connections

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"im-server-go/config"
	"log"
)

type mySqlDB struct {
	db *gorm.DB
}

var mysqlDb *mySqlDB

func initMySql() {
	dsn := config.MySQLConfig.Config.FormatDSN()
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("连接数据库失败：%s", err.Error())
	}
	mysqlDb = &mySqlDB{}
	mysqlDb.db = db
	mysqlDb.db.DB().SetConnMaxLifetime(config.MySQLConfig.ConnMaxLifeTime)
	mysqlDb.db.DB().SetMaxIdleConns(config.MySQLConfig.MaxIdleConns)
	mysqlDb.db.DB().SetMaxOpenConns(config.MySQLConfig.MaxOpenConns)

	// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `auth`
	mysqlDb.db.SingularTable(true)
	err = mysqlDb.db.DB().Ping()
	if err != nil {
		log.Fatalf("Ping数据库失败：%s", err.Error())
	}
}

func GetInstance() *gorm.DB {
	return mysqlDb.db
}
