package connections

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"im-server-go/config"
	"im-server-go/entity/mysql_model"
	"log"
	"math/rand"
	"strings"
	"time"
)

type MySqlDB struct {
	db *gorm.DB
}

var mysqlDb *MySqlDB

func InitMySql() IUser {
	dsn := config.MysqlCfg.Config.FormatDSN()
	mysqlDb = &MySqlDB{db: nil}
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("连接数据库失败：%s", err.Error())
	}
	mysqlDb.db = db
	mysqlDb.db.DB().SetConnMaxLifetime(config.MysqlCfg.ConnMaxLifeTime)
	mysqlDb.db.DB().SetMaxIdleConns(config.MysqlCfg.MaxIdleConns)
	mysqlDb.db.DB().SetMaxOpenConns(config.MysqlCfg.MaxOpenConns)

	// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	mysqlDb.db.SingularTable(true)
	err = mysqlDb.db.DB().Ping()
	if err != nil {
		log.Fatalf("Ping数据库失败：%s", err.Error())
		return nil
	}
	return mysqlDb
}

//用户注册
func (md *MySqlDB) UserSignUp(telephone, pwd string) bool {
	var alpha = []rune("abcdefghijklmnopqrstuvwxyz")
	defaultNick := make([]rune, 10)
	for i := range defaultNick {
		defaultNick[i] = alpha[rand.Intn(len(alpha))]
	}

	user := mysql_model.User{
		Telephone:     telephone,
		Password:      pwd,
		Nickname:      string(defaultNick),
		Sex:           0,
		Avatar:        "",
		Status:        1,
		Ctime:         time.Now().Unix(),
		LastLoginTime: time.Now().Unix(),
	}
	//判断是否已注册
	//NewRecord主键为空返回`true`,创建`user`后返回`false`
	if isNew := md.db.NewRecord(user); isNew == false {
		fmt.Println("[mysql] This telephone number has been registered!!!")
		return false
	}
	//创建新记录
	md.db.Create(&user)
	//判断是否创建成功,创建失败为true
	if isNew := md.db.NewRecord(user); isNew == true {
		fmt.Println("[mysql] This telephone number registered error!!!")
		return false
	}
	return true
}

//用户登录
func (md *MySqlDB) UserSignIn(telephone, pwd string) bool {
	user := &mysql_model.User{}
	err := md.db.Where("telephone = ?", telephone).First(user).Error
	if err != nil {
		fmt.Println("[mysql] Can't find user!!! err", err)
		return false
	}
	if strings.Compare(user.Password, pwd) == 0 {
		return true
	}
	return false
}

//获取用户信息
func (md *MySqlDB) GetUserInfo(telephone string) (mysql_model.User, error) {
	user := mysql_model.User{}
	err := md.db.Where("telephone = ?", telephone).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
