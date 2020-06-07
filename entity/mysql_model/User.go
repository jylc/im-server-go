package mysql_model

type User struct {
	UserId        int    `gorm:"column:user_id;PRIMARY_KEY" json:"user_id"`
	Telephone     string `gorm:"column:telephone" json:"telephone"`
	Password      string `gorm:"column:password" json:"password"`
	Nickname      string `gorm:"column:nickname" json:"nickname"`
	Sex           int    `gorm:"column:sex" json:"sex"`
	Avatar        string `gorm:"column:avatar" json:"avatar"`                   //头像
	Ctime         int64  `gorm:"column:c_time" json:"c_time"`                   //创建时间
	LastLoginTime int64  `gorm:"column:last_login_time" json:"last_login_time"` //最后登录时间
}
