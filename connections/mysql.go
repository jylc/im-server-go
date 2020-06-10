package connections

import (
	"errors"
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
func (md *MySqlDB) UserSignUp(telephone, pwd, token string) error {
	var alpha = []rune("abcdefghijklmnopqrstuvwxyz")
	defaultNick := make([]rune, 10)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := range defaultNick {
		defaultNick[i] = alpha[r.Intn(len(alpha))]
	}

	user := mysql_model.User{
		Telephone:     telephone,
		Password:      pwd,
		Nickname:      string(defaultNick),
		Sex:           0,
		Avatar:        "",
		Status:        1,
		AccessToken:   token,
		Ctime:         time.Now().Unix(),
		LastLoginTime: time.Now().Unix(),
	}
	//判断是否已注册
	if !md.db.Where("telephone = ?", telephone).Find(&mysql_model.User{}).RecordNotFound() {
		return errors.New("this telephone number has been registered")
	}
	//创建新记录
	if err := md.db.Save(&user).Error; err != nil {
		return err
	}
	//判断是否创建成功,创建失败为true
	isNew := md.db.NewRecord(user)
	if isNew == true {
		return errors.New("this telephone number registered failed")
	}
	return nil
}

//用户登录
func (md *MySqlDB) UserSignIn(telephone, pwd, token string) bool {
	user := &mysql_model.User{}
	err := md.db.Where("telephone = ?", telephone).First(user).Error
	if err != nil {
		fmt.Println("[mysql] Can't find user!!! err", err)
		return false
	}
	if strings.Compare(user.Password, pwd) == 0 {
		//登录成功，更新token
		md.UpdateToken(user, token)
		return true
	}
	return false
}

//获取通过手机号查询用户信息
func (md *MySqlDB) GetUserInfo(telephone string) (mysql_model.User, error) {
	user := mysql_model.User{}
	err := md.db.Where("telephone = ?", telephone).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

//TODO 更新用户token
func (md *MySqlDB) UpdateToken(user *mysql_model.User, token string) bool {
	err := md.db.Model(user).Update("access_token", token).Error
	if err != nil {
		log.Println("create token error")
		return false
	}
	return true
}
